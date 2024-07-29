// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package cdp

import (
	"os"
	"testing"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/common"
)

func unsetEnvs(keys ...string) {
	for _, key := range keys {
		os.Unsetenv(key)
	}
}

func TestGetCdpProfileFromConfig(t *testing.T) {
	os.Setenv(CdpProfileEnvVar, "foo")
	os.Setenv(cdpDefaultProfileEnvVar, "bar")
	defer unsetEnvs(CdpProfileEnvVar, cdpDefaultProfileEnvVar)
	config := Config{
		Profile: "baz",
	}
	profile, err := config.GetCdpProfile()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	common.AssertEquals(t, "baz", profile)
}

func TestGetCdpProfileFromEnv(t *testing.T) {
	os.Setenv(CdpProfileEnvVar, "foo")
	os.Setenv(cdpDefaultProfileEnvVar, "bar")
	defer unsetEnvs(CdpProfileEnvVar, cdpDefaultProfileEnvVar)
	config := Config{
		Profile: "",
	}
	profile, err := config.GetCdpProfile()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	common.AssertEquals(t, "bar", profile)
}

func TestGetCdpProfileFromEnv2(t *testing.T) {
	os.Setenv(CdpProfileEnvVar, "")
	os.Setenv(cdpDefaultProfileEnvVar, "bar")
	defer unsetEnvs(CdpProfileEnvVar, cdpDefaultProfileEnvVar)
	config := Config{}
	profile, err := config.GetCdpProfile()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	common.AssertEquals(t, "bar", profile)
}

func TestGetCdpProfileFromDefault(t *testing.T) {
	os.Setenv(CdpProfileEnvVar, "")
	os.Setenv(cdpDefaultProfileEnvVar, "")
	defer unsetEnvs(CdpProfileEnvVar, cdpDefaultProfileEnvVar)
	config := Config{}
	profile, err := config.GetCdpProfile()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	common.AssertEquals(t, cdpDefaultProfile, profile)
}

func TestGetCredentialsNotFound(t *testing.T) {
	// empty env vars.
	os.Setenv(CdpAccessKeyIdEnvVar, "")
	os.Setenv(CdpPrivateKeyEnvVar, "")
	os.Setenv(CdpProfileEnvVar, "")
	os.Setenv(cdpDefaultProfileEnvVar, "")
	defer unsetEnvs(CdpAccessKeyIdEnvVar, CdpPrivateKeyEnvVar, CdpProfileEnvVar, cdpDefaultProfileEnvVar)

	path := "testdata/test-credentials"
	profile := "profile_non_existing"

	emptyConfig := Config{
		Profile:         profile,
		CredentialsFile: path,
	}
	err := emptyConfig.LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
	_, err = emptyConfig.GetCredentials()
	if err == nil {
		t.Fatal(err)
	}
}

func TestGetCdpCredentials(t *testing.T) {
	// empty env vars.
	os.Setenv(CdpAccessKeyIdEnvVar, "value-from-env")
	os.Setenv(CdpPrivateKeyEnvVar, "value-from-env")
	os.Setenv(CdpProfileEnvVar, "")
	os.Setenv(cdpDefaultProfileEnvVar, "")
	defer unsetEnvs(CdpAccessKeyIdEnvVar, CdpPrivateKeyEnvVar, CdpProfileEnvVar, cdpDefaultProfileEnvVar)

	path := "testdata/test-credentials"
	profile := "file_cdp_credentials_provider_profile"

	config := Config{
		Profile:         profile,
		CredentialsFile: path,
		Credentials: &Credentials{
			AccessKeyId: "value-from-config",
			PrivateKey:  "value-from-config"},
	}
	err := config.LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
	credentials, err := config.GetCredentials()
	if err != nil {
		t.Fatal(err)
	}

	if credentials.AccessKeyId != "value-from-config" && credentials.PrivateKey != "value-from-config" {
		if err != nil {
			t.Errorf("Expected the provider chain to prioritize credentials from config")
		}
	}

	config.Credentials = nil
	credentials, err = config.GetCredentials()
	if err != nil {
		t.Fatal(err)
	}
	if credentials.AccessKeyId != "value-from-env" && credentials.PrivateKey != "value-from-env" {
		if err != nil {
			t.Errorf("Expected the provider chain to prioritize credentials from env after config")
		}
	}

	os.Setenv(CdpAccessKeyIdEnvVar, "")
	os.Setenv(CdpPrivateKeyEnvVar, "")
	credentials, err = config.GetCredentials()
	if err != nil {
		t.Fatal(err)
	}
	if credentials.AccessKeyId != "value-from-file" && credentials.PrivateKey != "value-from-file" {
		if err != nil {
			t.Errorf("Expected the provider chain to use credentials from file last")
		}
	}
}

func TestLoadConfigFileNotFound(t *testing.T) {
	os.Setenv("CDP_CONFIG_FILE", "")
	defer unsetEnvs("CDP_CONFIG_FILE")
	config := Config{
		ConfigFile: "testdata/non-existent-file",
	}
	// should silently ignore non existent files
	err := config.LoadConfig()
	if err != nil {
		t.Errorf("Unexpected error from LoadConfig()")
	}
}

func TestLoadConfigFile(t *testing.T) {
	os.Setenv("CDP_CONFIG_FILE", "")
	defer unsetEnvs("CDP_CONFIG_FILE")
	config := Config{
		ConfigFile: "testdata/test-config",
	}
	err := config.LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
	endpoint, err := config.GetCdpApiEndpoint()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	common.AssertEquals(t, "value1", endpoint)
}

func TestLoadConfigFileFromEnv(t *testing.T) {
	os.Setenv("CDP_CONFIG_FILE", "testdata/test-config")
	defer unsetEnvs("CDP_CONFIG_FILE")
	config := Config{
		ConfigFile: "testdata/test-config",
	}
	err := config.LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
	endpoint, err := config.GetCdpApiEndpoint()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	common.AssertEquals(t, "value1", endpoint)
}

func TestGetCdpProfileCaseSensitivity(t *testing.T) {
	// We support case sensitive section names, but case insensitive option names
	config := Config{
		ConfigFile: "testdata/test-config",
		Profile:    "UPPER_CASE_PROFILE",
	}
	err := config.LoadConfig()
	if err != nil {
		t.Fatal(err)
	}

	endpoint, err := config.GetCdpApiEndpoint()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	common.AssertEquals(t, "value6", endpoint)
}

func TestGetCdpRegionFromConfig(t *testing.T) {
	config := Config{
		CdpRegion: "foo",
	}
	os.Setenv("CDP_REGION", "bar")
	defer unsetEnvs("CDP_REGION")
	region, err := config.GetCdpRegion()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	common.AssertEquals(t, "foo", region)
}

func TestGetCdpRegionFromEnv(t *testing.T) {
	config := Config{}
	os.Setenv("CDP_REGION", "foo")
	defer unsetEnvs("CDP_REGION")
	region, err := config.GetCdpRegion()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	common.AssertEquals(t, "foo", region)
	cdpEndpoint, err := config.GetCdpApiEndpoint()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	common.AssertEquals(t, "https://api.foo.cdp.cloudera.com/", cdpEndpoint)
}

func TestGetCdpRegionFromConfigFile(t *testing.T) {
	config := Config{
		Profile:    "foo",
		ConfigFile: "testdata/test-config",
	}
	err := config.LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
	region, err := config.GetCdpRegion()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	common.AssertEquals(t, "value5", region)
}

func TestGetCdpRegionFromDefault(t *testing.T) {
	config := Config{}
	region, err := config.GetCdpRegion()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	common.AssertEquals(t, defaultCdpRegion, region)
}

func TestGetEndpointsWithProfile(t *testing.T) {
	config := Config{
		Profile:    "foo",
		ConfigFile: "testdata/test-config",
	}
	err := config.LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
	cdpEndpoint, err := config.GetCdpApiEndpoint()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	altusEndpoint, err := config.GetAltusApiEndpoint()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	iamEndpoint, err := config.GetEndpoint("iam", true)
	if err != nil {
		t.Fatalf("Error getting the endpoint: %v", err)
	}
	common.AssertEquals(t, "value3", cdpEndpoint)
	common.AssertEquals(t, "value4%s", altusEndpoint)
	common.AssertEquals(t, "value4iam", iamEndpoint)
}

func TestGetEndpointsWithRegionInProfile(t *testing.T) {
	config := Config{
		Profile:    "bar",
		ConfigFile: "testdata/test-config",
	}
	err := config.LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
	cdpEndpoint, err := config.GetCdpApiEndpoint()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	altusEndpoint, err := config.GetAltusApiEndpoint()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	iamEndpoint, err := config.GetEndpoint("iam", true)
	if err != nil {
		t.Fatalf("Error getting the endpoint: %v", err)
	}
	common.AssertEquals(t, "https://api.value6.cdp.cloudera.com/", cdpEndpoint)
	common.AssertEquals(t, "https://api.value6.cdp.cloudera.com/", altusEndpoint)
	common.AssertEquals(t, "https://api.value6.cdp.cloudera.com/", iamEndpoint)
}

func TestDefaultEndpoints(t *testing.T) {
	config := Config{
		ConfigFile: "testdata/test-config",
		Profile:    "non-existing",
	}
	err := config.LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
	cdpEndpoint, err := config.GetCdpApiEndpoint()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	altusEndpoint, err := config.GetAltusApiEndpoint()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	iamEndpoint, err := config.GetEndpoint("iam", true)
	if err != nil {
		t.Fatalf("Error getting the endpoint: %v", err)
	}
	common.AssertEquals(t, "https://api.us-west-1.cdp.cloudera.com/", cdpEndpoint)
	common.AssertEquals(t, "https://%sapi.us-west-1.altus.cloudera.com/", altusEndpoint)
	common.AssertEquals(t, "https://iamapi.us-west-1.altus.cloudera.com/", iamEndpoint)
}

func TestGetEndpointsWithRegionUsWest1(t *testing.T) {
	config := Config{
		CdpRegion: RegionUsWest1,
	}
	cdpEndpoint, err := config.GetCdpApiEndpoint()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	altusEndpoint, err := config.GetAltusApiEndpoint()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	iamEndpoint, err := config.GetEndpoint("iam", true)
	if err != nil {
		t.Fatalf("Error getting the endpoint: %v", err)
	}

	common.AssertEquals(t, "https://api.us-west-1.cdp.cloudera.com/", cdpEndpoint)
	common.AssertEquals(t, "https://%sapi.us-west-1.altus.cloudera.com/", altusEndpoint)
	common.AssertEquals(t, "https://iamapi.us-west-1.altus.cloudera.com/", iamEndpoint)
}

func TestGetEndpointsWithRegionEu1(t *testing.T) {
	config := Config{
		CdpRegion: RegionEu1,
	}
	cdpEndpoint, err := config.GetCdpApiEndpoint()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	altusEndpoint, err := config.GetAltusApiEndpoint()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	iamEndpoint, err := config.GetEndpoint("iam", true)
	if err != nil {
		t.Fatalf("Error getting the endpoint: %v", err)
	}

	common.AssertEquals(t, "https://api.eu-1.cdp.cloudera.com/", cdpEndpoint)
	common.AssertEquals(t, "https://api.eu-1.cdp.cloudera.com/", altusEndpoint)
	common.AssertEquals(t, "https://api.eu-1.cdp.cloudera.com/", iamEndpoint)
}

func TestGetEndpointsWithRegionAp1(t *testing.T) {
	config := Config{
		CdpRegion: RegionAp1,
	}
	cdpEndpoint, err := config.GetCdpApiEndpoint()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	altusEndpoint, err := config.GetAltusApiEndpoint()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	iamEndpoint, err := config.GetEndpoint("iam", true)
	if err != nil {
		t.Fatalf("Error getting the endpoint: %v", err)
	}

	common.AssertEquals(t, "https://api.ap-1.cdp.cloudera.com/", cdpEndpoint)
	common.AssertEquals(t, "https://api.ap-1.cdp.cloudera.com/", altusEndpoint)
	common.AssertEquals(t, "https://api.ap-1.cdp.cloudera.com/", iamEndpoint)
}

func TestGetEndpointsWithRegionUsg1(t *testing.T) {
	config := Config{
		CdpRegion: RegionUsg1,
	}
	cdpEndpoint, err := config.GetCdpApiEndpoint()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	altusEndpoint, err := config.GetAltusApiEndpoint()
	if err != nil {
		t.Fatalf("Error getting the config value: %v", err)
	}
	iamEndpoint, err := config.GetEndpoint("iam", true)
	if err != nil {
		t.Fatalf("Error getting the endpoint: %v", err)
	}

	common.AssertEquals(t, "https://api.usg-1.cdp.clouderagovt.com/", cdpEndpoint)
	common.AssertEquals(t, "https://api.usg-1.cdp.clouderagovt.com/", altusEndpoint)
	common.AssertEquals(t, "https://api.usg-1.cdp.clouderagovt.com/", iamEndpoint)
}

func TestConfig_WithBaseApiPath(t *testing.T) {
	var config Config
	config.WithBaseApiPath("/foo")
	common.AssertEquals(t, "/foo", config.GetBaseApiPath())
}
