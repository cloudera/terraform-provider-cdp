// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package provider

import (
	"context"
	"fmt"
	"github.com/cloudera/terraform-provider-cdp/resources/dw/virtualwarehouse/impala"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/resources/datahub"
	"github.com/cloudera/terraform-provider-cdp/resources/datalake"
	"github.com/cloudera/terraform-provider-cdp/resources/de"
	dwaws "github.com/cloudera/terraform-provider-cdp/resources/dw/cluster/aws"
	dwdatabasecatalog "github.com/cloudera/terraform-provider-cdp/resources/dw/databasecatalog"
	"github.com/cloudera/terraform-provider-cdp/resources/dw/virtualwarehouse/hive"
	"github.com/cloudera/terraform-provider-cdp/resources/environments"
	"github.com/cloudera/terraform-provider-cdp/resources/iam"
	"github.com/cloudera/terraform-provider-cdp/resources/ml"
	"github.com/cloudera/terraform-provider-cdp/resources/opdb"
	testUtil "github.com/cloudera/terraform-provider-cdp/utils/test"
)

const (
	testVersion          = "0.1.0"
	testTerraformVersion = "v1.4.2"
	testEnvVariable      = "someEnvVariable"
)

type cdpConfigTestScenarioBase struct {
	name     string
	envVar   string
	envValue string
}

type cdpConfigTestScenarioForString struct {
	cdpConfigTestScenarioBase
	modelValue     types.String
	expectedResult string
}

type cdpConfigTestScenarioForBool struct {
	cdpConfigTestScenarioBase
	modelValue     types.Bool
	expectedResult bool
}

func createCdpProviderModel() *CdpProviderModel {
	return &CdpProviderModel{
		CdpAccessKeyId:           types.StringValue("cdp-access-key"),
		CdpPrivateKey:            types.StringValue("cdp-private-key"),
		Profile:                  types.StringValue("profile"),
		CdpRegion:                types.StringValue("region"),
		AltusEndpointUrl:         types.StringValue("altus-endpoint-url"),
		CdpEndpointUrl:           types.StringValue("cdp-endpoint-url"),
		CdpConfigFile:            types.StringValue("cdp-config-file"),
		CdpSharedCredentialsFile: types.StringValue("cdp-shared-credentials-file"),
		LocalEnvironment:         types.BoolValue(false),
	}
}

func TestCdpProvider_ProviderOverridesUserAgent(t *testing.T) {
	model := createCdpProviderModel()

	config := getCdpConfig(context.Background(), model, testVersion, testTerraformVersion)
	userAgent := config.GetUserAgentOrDefault()

	r, _ := regexp.Compile(`^CDPTFPROVIDER/.+ Terraform/.+ Go/.+ .+_.+$`)
	if !r.MatchString(userAgent) {
		t.Fatalf("Failed to match the User-Agent regex: %v", userAgent)
	}
}

func TestCdpProvider_ProviderClientApplicationName(t *testing.T) {
	model := createCdpProviderModel()

	config := getCdpConfig(context.Background(), model, testVersion, testTerraformVersion)
	clientApplicationName := config.ClientApplicationName

	if clientApplicationName != "terraform-provider-cdp" {
		t.Fatalf("Terraform provider should have set client application name. Got: %v", clientApplicationName)
	}
}

func TestCdpProvider_ProviderCdpRegion(t *testing.T) {
	model := createCdpProviderModel()

	config := getCdpConfig(context.Background(), model, testVersion, testTerraformVersion)

	cdpRegion, err := config.GetCdpRegion()
	if err != nil {
		t.Fatalf("Error getting cdp region: %v", err)
	}
	if cdpRegion != "region" {
		t.Fatalf("Terraform provider should have set cdp region. Got: %v", cdpRegion)
	}
}

func TestCdpProvider_GetCdpConfigCdpAccessKeyId(t *testing.T) {
	tests := []cdpConfigTestScenarioForString{
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when CdpAccessKeyId is not empty",
				envVar:   "CDP_ACCESS_KEY_ID",
				envValue: "",
			},
			modelValue:     types.StringValue("some CdpAccessKeyId value"),
			expectedResult: "some CdpAccessKeyId value",
		},
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when CdpAccessKeyId is empty",
				envVar:   "CDP_ACCESS_KEY_ID",
				envValue: "some CdpAccessKeyId env variable value",
			},
			modelValue:     types.StringNull(),
			expectedResult: "some CdpAccessKeyId env variable value",
		},
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when both CdpAccessKeyId and CDP_ACCESS_KEY_ID contains no value",
				envVar:   "CDP_ACCESS_KEY_ID",
				envValue: "",
			},
			modelValue:     types.StringNull(),
			expectedResult: "",
		},
	}
	for _, test := range tests {
		t.Setenv(test.envVar, test.envValue)
		t.Run(test.name, func(t *testing.T) {
			defer t.Setenv(test.envVar, "")
			model := createCdpProviderModel()
			model.CdpAccessKeyId = test.modelValue

			config := getCdpConfig(context.Background(), model, testVersion, testTerraformVersion)

			cred := config.Credentials
			if cred.AccessKeyId != test.expectedResult {
				t.Errorf("AccessKeyId diverges from the expected! Expected: %s, got: %s\n",
					test.expectedResult, cred.AccessKeyId)
			}
		})
	}
}

func TestCdpProvider_GetCdpConfigCdpPrivateKey(t *testing.T) {
	tests := []cdpConfigTestScenarioForString{
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when CdpPrivateKey is not empty",
				envVar:   "CDP_PRIVATE_KEY",
				envValue: "",
			},
			modelValue:     types.StringValue("some CdpPrivateKey value"),
			expectedResult: "some CdpPrivateKey value",
		},
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when CdpPrivateKey is empty",
				envVar:   "CDP_PRIVATE_KEY",
				envValue: "some CdpPrivateKey env variable value",
			},
			modelValue:     types.StringNull(),
			expectedResult: "some CdpPrivateKey env variable value",
		},
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when both CdpAccessKeyId and CDP_PRIVATE_KEY contains no value",
				envVar:   "CDP_PRIVATE_KEY",
				envValue: "",
			},
			modelValue:     types.StringNull(),
			expectedResult: "",
		},
	}
	for _, test := range tests {
		t.Setenv(test.envVar, test.envValue)
		t.Run(test.name, func(t *testing.T) {
			defer t.Setenv(test.envVar, "")
			model := createCdpProviderModel()
			model.CdpPrivateKey = test.modelValue

			config := getCdpConfig(context.Background(), model, testVersion, testTerraformVersion)

			cred := config.Credentials
			if cred.PrivateKey != test.expectedResult {
				t.Errorf("AccessKeyId diverges from the expected! Expected: %s, got: %s\n",
					test.expectedResult, cred.PrivateKey)
			}
		})
	}
}

func TestCdpProvider_GetCdpConfigProfile(t *testing.T) {
	tests := []cdpConfigTestScenarioForString{
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when CDP_PROFILE is not empty",
				envVar:   "CDP_DEFAULT_PROFILE",
				envValue: "",
			},
			modelValue:     types.StringValue("some CDP_PROFILE value"),
			expectedResult: "some CDP_PROFILE value",
		},
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when CDP_PROFILE is empty",
				envVar:   "CDP_DEFAULT_PROFILE",
				envValue: "some CDP_PROFILE env variable value",
			},
			modelValue:     types.StringNull(),
			expectedResult: "some CDP_PROFILE env variable value",
		},
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when both CDP_PROFILE and CDP_DEFAULT_PROFILE contains no value",
				envVar:   "CDP_DEFAULT_PROFILE",
				envValue: "",
			},
			modelValue:     types.StringNull(),
			expectedResult: "",
		},
	}
	for _, test := range tests {
		t.Setenv(test.envVar, test.envValue)
		t.Run(test.name, func(t *testing.T) {
			defer t.Setenv(test.envVar, "")
			model := createCdpProviderModel()
			model.Profile = test.modelValue

			config := getCdpConfig(context.Background(), model, testVersion, testTerraformVersion)

			if config.Profile != test.expectedResult {
				t.Errorf("Profile diverges from the expected! Expected: %s, got: %s\n",
					test.expectedResult, config.Profile)
			}
		})
	}
}

func TestCdpProvider_GetCdpConfigCdpRegion(t *testing.T) {
	tests := []cdpConfigTestScenarioForString{
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when CdpRegion is not empty",
				envVar:   "CDP_REGION",
				envValue: "",
			},
			modelValue:     types.StringValue("some CdpRegion value"),
			expectedResult: "some CdpRegion value",
		},
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when CdpRegion is empty",
				envVar:   "CDP_REGION",
				envValue: "some CdpRegion env variable value",
			},
			modelValue:     types.StringNull(),
			expectedResult: "some CdpRegion env variable value",
		},
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when both CdpRegion and CDP_REGION contains no value",
				envVar:   "CDP_REGION",
				envValue: "",
			},
			modelValue:     types.StringNull(),
			expectedResult: "",
		},
	}
	for _, test := range tests {
		t.Setenv(test.envVar, test.envValue)
		t.Run(test.name, func(t *testing.T) {
			defer t.Setenv(test.envVar, "")
			model := createCdpProviderModel()
			model.CdpRegion = test.modelValue

			config := getCdpConfig(context.Background(), model, testVersion, testTerraformVersion)

			if config.CdpRegion != test.expectedResult {
				t.Errorf("CdpRegion diverges from the expected! Expected: %s, got: %s\n",
					test.expectedResult, config.CdpRegion)
			}
		})
	}
}

func TestCdpProvider_GetCdpConfigAltusEndpointUrl(t *testing.T) {
	tests := []cdpConfigTestScenarioForString{
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when AltusEndpointUrl is not empty",
				envVar:   "ENDPOINT_URL",
				envValue: "",
			},
			modelValue:     types.StringValue("some AltusEndpointUrl value"),
			expectedResult: "some AltusEndpointUrl value",
		},
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when AltusEndpointUrl is empty",
				envVar:   "ENDPOINT_URL",
				envValue: "some AltusEndpointUrl env variable value",
			},
			modelValue:     types.StringNull(),
			expectedResult: "some AltusEndpointUrl env variable value",
		},
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when both AltusEndpointUrl and ENDPOINT_URL contains no value",
				envVar:   "ENDPOINT_URL",
				envValue: "",
			},
			modelValue:     types.StringNull(),
			expectedResult: "",
		},
	}
	for _, test := range tests {
		t.Setenv(test.envVar, test.envValue)
		t.Run(test.name, func(t *testing.T) {
			defer t.Setenv(test.envVar, "")
			model := createCdpProviderModel()
			model.AltusEndpointUrl = test.modelValue

			config := getCdpConfig(context.Background(), model, testVersion, testTerraformVersion)

			if config.AltusApiEndpointUrl != test.expectedResult {
				t.Errorf("AltusApiEndpointUrl diverges from the expected! Expected: %s, got: %s\n",
					test.expectedResult, config.CdpRegion)
			}
		})
	}
}

func TestCdpProvider_GetCdpConfigCdpEndpointUrl(t *testing.T) {
	tests := []cdpConfigTestScenarioForString{
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when CdpEndpointUrl is not empty",
				envVar:   "CDP_ENDPOINT_URL",
				envValue: "",
			},
			modelValue:     types.StringValue("some CdpEndpointUrl value"),
			expectedResult: "some CdpEndpointUrl value",
		},
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when CdpEndpointUrl is empty",
				envVar:   "CDP_ENDPOINT_URL",
				envValue: "some CdpEndpointUrl env variable value",
			},
			modelValue:     types.StringNull(),
			expectedResult: "some CdpEndpointUrl env variable value",
		},
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when both CdpEndpointUrl and CDP_ENDPOINT_URL contains no value",
				envVar:   "CDP_ENDPOINT_URL",
				envValue: "",
			},
			modelValue:     types.StringNull(),
			expectedResult: "",
		},
	}
	for _, test := range tests {
		t.Setenv(test.envVar, test.envValue)
		t.Run(test.name, func(t *testing.T) {
			defer t.Setenv(test.envVar, "")
			model := createCdpProviderModel()
			model.CdpEndpointUrl = test.modelValue

			config := getCdpConfig(context.Background(), model, testVersion, testTerraformVersion)

			if config.CdpApiEndpointUrl != test.expectedResult {
				t.Errorf("CdpEndpointUrl diverges from the expected! Expected: %s, got: %s\n",
					test.expectedResult, config.CdpRegion)
			}
		})
	}
}

func TestCdpProvider_GetCdpConfigCdpConfigFile(t *testing.T) {
	tests := []cdpConfigTestScenarioForString{
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when CdpConfigFile is not empty",
				envVar:   "CDP_CONFIG_FILE",
				envValue: "",
			},
			modelValue:     types.StringValue("some CdpConfigFile value"),
			expectedResult: "some CdpConfigFile value",
		},
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when CdpConfigFile is empty",
				envVar:   "CDP_CONFIG_FILE",
				envValue: "some CdpEndpointUrl env variable value",
			},
			modelValue:     types.StringNull(),
			expectedResult: "some CdpEndpointUrl env variable value",
		},
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when both CdpConfigFile and CDP_CONFIG_FILE contains no value",
				envVar:   "CDP_CONFIG_FILE",
				envValue: "",
			},
			modelValue:     types.StringNull(),
			expectedResult: "",
		},
	}
	for _, test := range tests {
		t.Setenv(test.envVar, test.envValue)
		t.Run(test.name, func(t *testing.T) {
			defer t.Setenv(test.envVar, "")
			model := createCdpProviderModel()
			model.CdpConfigFile = test.modelValue

			config := getCdpConfig(context.Background(), model, testVersion, testTerraformVersion)

			if config.ConfigFile != test.expectedResult {
				t.Errorf("CdpConfigFile diverges from the expected! Expected: %s, got: %s\n",
					test.expectedResult, config.CdpRegion)
			}
		})
	}
}

func TestCdpProvider_GetCdpConfigCdpSharedCredentialsFile(t *testing.T) {
	tests := []cdpConfigTestScenarioForString{
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when CdpSharedCredentialsFile is not empty",
				envVar:   "CDP_SHARED_CREDENTIALS_FILE",
				envValue: "",
			},
			modelValue:     types.StringValue("some CdpSharedCredentialsFile value"),
			expectedResult: "some CdpSharedCredentialsFile value",
		},
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when CdpSharedCredentialsFile is empty",
				envVar:   "CDP_SHARED_CREDENTIALS_FILE",
				envValue: "some CdpSharedCredentialsFile env variable value",
			},
			modelValue:     types.StringNull(),
			expectedResult: "some CdpSharedCredentialsFile env variable value",
		},
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when both CdpSharedCredentialsFile and CDP_SHARED_CREDENTIALS_FILE contains no value",
				envVar:   "CDP_SHARED_CREDENTIALS_FILE",
				envValue: "",
			},
			modelValue:     types.StringNull(),
			expectedResult: "",
		},
	}
	for _, test := range tests {
		t.Setenv(test.envVar, test.envValue)
		t.Run(test.name, func(t *testing.T) {
			defer t.Setenv(test.envVar, "")
			model := createCdpProviderModel()
			model.CdpSharedCredentialsFile = test.modelValue

			config := getCdpConfig(context.Background(), model, testVersion, testTerraformVersion)

			if config.CredentialsFile != test.expectedResult {
				t.Errorf("CdpSharedCredentialsFile diverges from the expected! Expected: %s, got: %s\n",
					test.expectedResult, config.CdpRegion)
			}
		})
	}
}

func TestCdpProvider_GetCdpConfigLocalEnvironment(t *testing.T) {
	tests := []cdpConfigTestScenarioForBool{
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when LocalEnvironment is true",
				envVar:   "LOCAL_ENVIRONMENT",
				envValue: "",
			},
			modelValue:     types.BoolValue(true),
			expectedResult: true,
		},
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when LocalEnvironment is false",
				envVar:   "LOCAL_ENVIRONMENT",
				envValue: "",
			},
			modelValue:     types.BoolValue(false),
			expectedResult: false,
		},
		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when LocalEnvironment is null and env var value is false",
				envVar:   "LOCAL_ENVIRONMENT",
				envValue: "false",
			},
			modelValue:     types.BoolNull(),
			expectedResult: false,
		},

		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when LocalEnvironment is null and env var value is true",
				envVar:   "LOCAL_ENVIRONMENT",
				envValue: "true",
			},
			modelValue:     types.BoolNull(),
			expectedResult: true,
		},

		{
			cdpConfigTestScenarioBase: cdpConfigTestScenarioBase{
				name:     "test when LocalEnvironment is null and env var value is empty",
				envVar:   "LOCAL_ENVIRONMENT",
				envValue: "",
			},
			modelValue:     types.BoolNull(),
			expectedResult: false,
		},
	}
	for _, test := range tests {
		t.Setenv(test.envVar, test.envValue)
		t.Run(test.name, func(t *testing.T) {
			defer t.Setenv(test.envVar, "")
			model := createCdpProviderModel()
			model.LocalEnvironment = test.modelValue

			config := getCdpConfig(context.Background(), model, testVersion, testTerraformVersion)

			if config.LocalEnvironment != test.expectedResult {
				t.Errorf("LocalEnvironment diverges from the expected! Expected: %t, got: %t\n",
					test.expectedResult, config.LocalEnvironment)
			}
		})
	}
}

func TestGetOrDefaultFromEnvWithValue(t *testing.T) {
	val := types.StringValue("expected")
	if got := getOrDefaultFromEnv(val, testEnvVariable); got != val.ValueString() {
		t.Errorf("Assertion faled! Expected: %s, got: %s\n", val.ValueString(), got)
	}
}

func TestGetOrDefaultFromEnvWhenValuePresentInEnv(t *testing.T) {
	defer t.Setenv(testEnvVariable, "")
	val := "expected"
	t.Setenv(testEnvVariable, val)
	if got := getOrDefaultFromEnv(types.StringNull(), testEnvVariable); got != val {
		t.Errorf("Assertion faled! Expected: %s, got: %s\n", val, got)
	}
}

func TestGetOrDefaultFromEnvWhenValuePresentInEnvButNotFirst(t *testing.T) {
	defer t.Setenv(testEnvVariable, "")
	val := "expected"
	t.Setenv(testEnvVariable, val)
	if got := getOrDefaultFromEnv(types.StringNull(), "customOtherEnvVariable", testEnvVariable); got != val {
		t.Errorf("Assertion faled! Expected: %s, got: %s\n", val, got)
	}
}

func TestGetOrDefaultBoolFromEnvWithValue(t *testing.T) {
	val := types.BoolValue(true)
	if got := getOrDefaultBoolFromEnv(context.Background(), val, testEnvVariable); got != val.ValueBool() {
		t.Errorf("Assertion faled! Expected: %t, got: %t\n", val.ValueBool(), got)
	}
}

func TestGetOrDefaultBoolFromEnvWhenValuePresentInEnv(t *testing.T) {
	defer t.Setenv(testEnvVariable, "")
	val := "true"
	t.Setenv(testEnvVariable, val)
	if got := getOrDefaultBoolFromEnv(context.Background(), types.BoolNull(), testEnvVariable); !got {
		t.Errorf("Assertion faled! Expected: %t, got: %t\n", true, got)
	}
}

func TestGetOrDefaultBoolFromEnvWhenValuePresentInEnvButNotFirst(t *testing.T) {
	defer t.Setenv(testEnvVariable, "")
	val := "true"
	t.Setenv(testEnvVariable, val)
	if got := getOrDefaultBoolFromEnv(context.Background(), types.BoolNull(), "customOtherEnvVariable", testEnvVariable); !got {
		t.Errorf("Assertion faled! Expected: %t, got: %t\n", true, got)
	}
}

func TestCdpProvider_Resources(t *testing.T) {
	expectedResources := []func() resource.Resource{
		environments.NewAwsCredentialResource,
		environments.NewAwsEnvironmentResource,
		environments.NewIDBrokerMappingsResource,
		environments.NewUserSyncResource,
		environments.NewAzureCredentialResource,
		environments.NewAzureEnvironmentResource,
		environments.NewGcpEnvironmentResource,
		environments.NewGcpCredentialResource,
		environments.NewProxyConfigurationResource,
		environments.NewAzureImageTermsResource,
		datalake.NewAwsDatalakeResource,
		datalake.NewAzureDatalakeResource,
		datalake.NewGcpDatalakeResource,
		iam.NewGroupResource,
		iam.NewMachineUserResource,
		iam.NewMachineUserGroupAssignmentResource,
		iam.NewMachineUserRoleAssignmentResource,
		iam.NewMachineUserResourceRoleAssignmentResource,
		datahub.NewAwsDatahubResource,
		datahub.NewAzureDatahubResource,
		datahub.NewGcpDatahubResource,
		opdb.NewDatabaseResource,
		ml.NewWorkspaceResource,
		de.NewServiceResource,
		hive.NewHiveResource,
		impala.NewImpalaResource,
		dwaws.NewDwClusterResource,
		dwdatabasecatalog.NewDwDatabaseCatalogResource,
	}

	provider := CdpProvider{testVersion}
	resources := provider.Resources(context.Background())

	unexpectedResources, missingResources := compareResources(resources, expectedResources)

	if len(unexpectedResources) > 0 {
		t.Errorf("The following unexpected resource(s) got introduced: %s", strings.Join(unexpectedResources, ","))
	}
	if len(missingResources) > 0 {
		t.Errorf("The following resource(s) got removed: %s", strings.Join(missingResources, ","))
	}
}

func TestCdpProvider_GetUserAgent(t *testing.T) {
	userAgentContent := getUserAgent(testVersion, testTerraformVersion)
	expected := fmt.Sprintf("CDPTFPROVIDER/%s Terraform/%s Go/%s %s_%s",
		testVersion, testTerraformVersion, runtime.Version(), runtime.GOOS, runtime.GOARCH)

	if expected != userAgentContent {
		t.Errorf("Assertion failed! Expected: %s, got: %s\n", expected, userAgentContent)
	}
}

func TestCdpProvider_Schema(t *testing.T) {
	expectedSchemaFields := []string{
		"cdp_shared_credentials_file",
		"cdp_access_key_id",
		"local_environment",
		"cdp_endpoint_url",
		"cdp_config_file",
		"cdp_private_key",
		"cdp_profile",
		"endpoint_url",
		"cdp_region",
	}

	provider := CdpProvider{testVersion}
	sresp := &fwprovider.SchemaResponse{
		Schema:      schema.Schema{},
		Diagnostics: nil,
	}

	provider.Schema(context.Background(), fwprovider.SchemaRequest{}, sresp)

	var sresps []string
	for attribute := range sresp.Schema.GetAttributes() {
		sresps = append(sresps, attribute)
	}

	unexpectedFields, missingFields := testUtil.CompareStringSlices(sresps, expectedSchemaFields)

	if len(unexpectedFields) > 0 {
		t.Errorf("The following unexpected field(s) got introduced: %s", strings.Join(unexpectedFields, ","))
	}
	if len(missingFields) > 0 {
		t.Errorf("The following field(s) got removed: %s", strings.Join(missingFields, ","))
	}
}

func TestCdpProvider_Metadata(t *testing.T) {
	provider := CdpProvider{testVersion}
	mResp := &fwprovider.MetadataResponse{}

	provider.Metadata(context.Background(), fwprovider.MetadataRequest{}, mResp)

	if mResp.TypeName != "cdp" {
		t.Errorf("The MetadataRequest's Type does not match with the expected ('cdp'): %s\n", mResp.TypeName)
	}
	if mResp.Version != testVersion {
		t.Errorf("The MetadataRequest's version does not match with the expected ('%s'): %s\n", testVersion, mResp.Version)
	}
}

func TestCdpProvider_DataSources(t *testing.T) {
	expectedDataSources := []func() datasource.DataSource{
		environments.NewAWSCredentialPrerequisitesDataSource,
		environments.NewKeytabDataSource,
		iam.NewGroupDataSource,
	}

	provider := CdpProvider{testVersion}
	resources := provider.DataSources(context.Background())

	unexpectedResources, missingResources := compareDataSources(resources, expectedDataSources)

	if len(unexpectedResources) > 0 {
		t.Errorf("The following unexpected DataSource(s) got introduced: %s", strings.Join(unexpectedResources, ","))
	}
	if len(missingResources) > 0 {
		t.Errorf("The following DataSource(s) got removed: %s", strings.Join(missingResources, ","))
	}
}

func compareResources(actual, expected []func() resource.Resource) (unexpected, missing []string) {
	actualNames := testUtil.ToStringSliceFunc(actual, getPackageAndNameOfResource)
	expectedNames := testUtil.ToStringSliceFunc(expected, getPackageAndNameOfResource)
	return testUtil.CompareStringSlices(actualNames, expectedNames)
}

func compareDataSources(actual, expected []func() datasource.DataSource) (unexpected, missing []string) {
	actualNames := testUtil.ToStringSliceFunc(actual, getPackageAndNameOfDataSource)
	expectedNames := testUtil.ToStringSliceFunc(expected, getPackageAndNameOfDataSource)
	return testUtil.CompareStringSlices(actualNames, expectedNames)
}

func getPackageAndNameOfResource(f func() resource.Resource) string {
	separated := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), string(os.PathSeparator))
	return separated[len(separated)-1]
}

func getPackageAndNameOfDataSource(f func() datasource.DataSource) string {
	separated := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), string(os.PathSeparator))
	return separated[len(separated)-1]
}
