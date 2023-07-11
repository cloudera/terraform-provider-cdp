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
	"os"
	"runtime"
	"strconv"

	"github.com/cloudera/terraform-provider-cdp/resources/datahub"
	"github.com/cloudera/terraform-provider-cdp/resources/iam"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/resources/datalake"
	"github.com/cloudera/terraform-provider-cdp/resources/environments"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &CdpProvider{}
)

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CdpProvider{
			version: version,
		}
	}
}

type CdpProvider struct {
	version string
}

type CdpProviderModel struct {
	CdpAccessKeyId           types.String `tfsdk:"cdp_access_key_id"`
	CdpPrivateKey            types.String `tfsdk:"cdp_private_key"`
	Profile                  types.String `tfsdk:"cdp_profile"`
	CdpRegion                types.String `tfsdk:"cdp_region"`
	AltusEndpointUrl         types.String `tfsdk:"endpoint_url"`
	CdpEndpointUrl           types.String `tfsdk:"cdp_endpoint_url"`
	CdpConfigFile            types.String `tfsdk:"cdp_config_file"`
	CdpSharedCredentialsFile types.String `tfsdk:"cdp_shared_credentials_file"`
	LocalEnvironment         types.Bool   `tfsdk:"local_environment"`
}

// Metadata returns the provider type name.
func (p *CdpProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cdp"
	resp.Version = p.version
}

func (p *CdpProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The Cloudera Data Platform (CDP) provider manages the lifecycle of resources supported by CDP like Credentials, Environment, Datalake, Datahub and other data services.",
		Attributes: map[string]schema.Attribute{
			"cdp_access_key_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "CDP access key id to authenticate the requests. It can be provided in the provider config (not recommended!), or it can be sourced from the `CDP_ACCESS_KEY_ID` environment variable, or via a shared credentials file. If `cdp_profile` is specified credentials for the specific profile will be used.",
			},
			"cdp_private_key": schema.StringAttribute{
				Optional:            true,
				Sensitive:           true,
				MarkdownDescription: "CDP private key associated with the given access key. It can be provided in the provider config(not recommended!), or it can also be sourced from the `CDP_PRIVATE_KEY` environment variable, or via a shared credentials file. If `cdp_profile` is specified credentials for the specific profile will be used.",
			},
			"cdp_profile": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "CDP Profile to use for the configuration in shared credentials file (`~/.cdp/credentials`). It can also be sourced from the `CDP_PROFILE` environment variable.",
			},
			"cdp_region": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "CDP Control Plane Region to send the API requests to, not to be confused by the cloud workload region (AWS, Azure or GCP). Defaults to `us-west-1` and can also be sourced from `CDP_REGION` environment variable. More details can be found [here](https://docs.cloudera.com/cdp-public-cloud/cloud/cp-regions/topics/cdp-control-plane-regions.html).",
			},
			"cdp_config_file": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "CDP configuration file. Defaults to `~/.cdp/config`.",
			},
			"cdp_shared_credentials_file": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "CDP shared credentials file. Defaults to `~/.cdp/credentials`.",
			},
			"endpoint_url": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Endpoint URL to use. Customize the endpoint URL format for connecting to alternate endpoints for IAM and Workload Management services. See the Custom [Service Endpoints Guide](guides/custom-service-endpoints.md) for more information about connecting to alternate CDP endpoints.",
			},
			"cdp_endpoint_url": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "CDP Endpoint URL to use. Customize the endpoint URL format for connecting to alternate endpoints for CDP services. See the Custom [Service Endpoints Guide](guides/custom-service-endpoints.md) for more information about connecting to alternate CDP endpoints.",
			},
			"local_environment": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Defines whether CDP Control Plane runs locally. Defaults to `false`. This is an advanced configuration, not intended to be used by end users.",
			},
		},
	}
}

func (p *CdpProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring CDP client")
	var data CdpProviderModel
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	// Create a new CDP client using the configuration values
	client, err := cdp.NewClient(getCdpConfig(ctx, &data, p.version, req.TerraformVersion))

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create CDP API Client",
			"An unexpected error occurred when creating the CDP API client. "+
				"If the error is not clear, please contact Cloudera.\n\n"+
				"CDP API Client Error: "+err.Error(),
		)
		return
	}

	// Make the CDP client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client
}

// getOrDefaultFromEnv returns the string value if it is non-empty. Otherwise returns the environment
// variable value from the operating system.
func getOrDefaultFromEnv(val basetypes.StringValue, envVars ...string) string {
	if !val.IsNull() {
		return val.ValueString()
	}

	for _, envVar := range envVars {
		env, ok := os.LookupEnv(envVar)
		if ok {
			return env
		}
	}
	return ""
}

func getOrDefaultBoolFromEnv(ctx context.Context, val basetypes.BoolValue, envVars ...string) bool {
	if !val.IsNull() {
		return val.ValueBool()
	}

	for _, envVar := range envVars {
		env, ok := os.LookupEnv(envVar)
		if ok {
			boolVal, err := strconv.ParseBool(env)
			if err != nil {
				tflog.Warn(ctx, fmt.Sprintf("Error parsing boolean: %v", env))
				return false
			}
			return boolVal
		}
	}
	return false
}

func getCdpConfig(ctx context.Context, data *CdpProviderModel, version string, terraformVersion string) *cdp.Config {
	tflog.Info(ctx, "Setting up CDP config")

	accessKeyId := getOrDefaultFromEnv(data.CdpAccessKeyId, "CDP_ACCESS_KEY_ID")
	privateKey := getOrDefaultFromEnv(data.CdpPrivateKey, "CDP_PRIVATE_KEY")
	cdpProfile := getOrDefaultFromEnv(data.Profile, "CDP_PROFILE", "CDP_DEFAULT_PROFILE")
	cdpRegion := getOrDefaultFromEnv(data.CdpRegion, "CDP_REGION")
	altusEndpointUrl := getOrDefaultFromEnv(data.AltusEndpointUrl, "ENDPOINT_URL")
	cdpEndpointUrl := getOrDefaultFromEnv(data.CdpEndpointUrl, "CDP_ENDPOINT_URL")
	cdpConfigFile := getOrDefaultFromEnv(data.CdpConfigFile, "CDP_CONFIG_FILE")
	cdpSharedCredentialsFile := getOrDefaultFromEnv(data.CdpSharedCredentialsFile, "CDP_SHARED_CREDENTIALS_FILE")
	localEnvironment := getOrDefaultBoolFromEnv(ctx, data.LocalEnvironment, "LOCAL_ENVIRONMENT")

	config := cdp.NewConfig()
	config.WithContext(ctx)
	config.WithProfile(cdpProfile)
	config.WithCdpRegion(cdpRegion)
	config.WithAltusApiEndpointUrl(altusEndpointUrl)
	config.WithCdpApiEndpointUrl(cdpEndpointUrl)
	config.WithCredentials(&cdp.Credentials{
		AccessKeyId: accessKeyId,
		PrivateKey:  privateKey,
	})
	config.WithConfigFile(cdpConfigFile)
	config.WithCredentialsFile(cdpSharedCredentialsFile)
	config.WithLocalEnvironment(localEnvironment)
	config.WithLogger(new(TFLoggerAdaptor))
	config.WithUserAgent(getUserAgent(version, terraformVersion))
	config.WithClientApplicationName("terraform-provider-cdp")
	config.WithVersion(version)

	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "privateKey")
	ctx = tflog.SetField(ctx, "accessKeyId", accessKeyId)
	ctx = tflog.SetField(ctx, "privateKey", privateKey)
	ctx = tflog.SetField(ctx, "cdpProfile", cdpProfile)
	ctx = tflog.SetField(ctx, "cdpRegion", cdpRegion)
	ctx = tflog.SetField(ctx, "altusEndpointUrl", altusEndpointUrl)
	ctx = tflog.SetField(ctx, "cdpEndpointUrl", cdpEndpointUrl)
	ctx = tflog.SetField(ctx, "cdpConfigFile", cdpConfigFile)
	ctx = tflog.SetField(ctx, "cdpSharedCredentialsFile", cdpSharedCredentialsFile)

	tflog.Info(ctx, "CDP config set up. Creating client")
	return config
}

func (p *CdpProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		environments.NewAwsCredentialResource,
		environments.NewAwsEnvironmentResource,
		environments.NewIDBrokerMappingsResource,
		environments.NewAzureCredentialResource,
		environments.NewAzureEnvironmentResource,
		environments.NewGcpEnvironmentResource,
		environments.NewGcpCredentialResource,
		datalake.NewAwsDatalakeResource,
		datalake.NewAzureDatalakeResource,
		iam.NewGroupResource,
		datahub.NewAwsDatahubResource,
		datahub.NewAzureDatahubResource,
	}
}

func (p *CdpProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		environments.NewAWSCredentialPrerequisitesDataSource,
		iam.NewGroupDataSource,
	}
}

// getUserAgent returns a string to be set for the User-Agent header in HTTP requests. We follow the same format
// with the python based CDP CLI and Java based CDP SDK. However, there is no easy way to detect the OS version without
// running uname, so we don't do that. Can be added later if needed.
func getUserAgent(version string, terraformVersion string) string {
	return fmt.Sprintf("CDPTFPROVIDER/%s Terraform/%s Go/%s %s_%s", version, terraformVersion, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}

// TFLoggerAdaptor implements cdp.Logger to send CDP SDK logs to tflog
type TFLoggerAdaptor struct {
}

func (l *TFLoggerAdaptor) Errorf(ctx context.Context, format string, args ...any) {
	tflog.Error(ctx, fmt.Sprintf(format, args...))
}

func (l *TFLoggerAdaptor) Warnf(ctx context.Context, format string, args ...any) {
	tflog.Warn(ctx, fmt.Sprintf(format, args...))
}

func (l *TFLoggerAdaptor) Infof(ctx context.Context, format string, args ...any) {
	tflog.Info(ctx, fmt.Sprintf(format, args...))
}

func (l *TFLoggerAdaptor) Debugf(ctx context.Context, format string, args ...any) {
	tflog.Debug(ctx, fmt.Sprintf(format, args...))
}
