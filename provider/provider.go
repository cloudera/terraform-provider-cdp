package provider

import (
	"context"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/authn"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/resources/environments"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
	AltusEndpointUrl         types.String `tfsdk:"endpoint_url"`
	CdpEndpointUrl           types.String `tfsdk:"cdp_endpoint_url"`
	CdpConfigFile            types.String `tfsdk:"cdp_config_file"`
	CdpSharedCredentialsFile types.String `tfsdk:"cdp_shared_credentials_file"`
}

// Metadata returns the provider type name.
func (p *CdpProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cdp"
	resp.Version = p.version
}

func (p *CdpProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"cdp_access_key_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "CDP access key id",
			},
			"cdp_private_key": schema.StringAttribute{
				Optional:            true,
				Sensitive:           true,
				MarkdownDescription: "CDP private key associated with the given access key",
			},
			"cdp_profile": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "CDP Profile to use for the configuration in ~/.cdp/",
			},
			"endpoint_url": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Altus Endpoint URL Format",
			},
			"cdp_endpoint_url": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "CDP Endpoint URL",
			},
			"cdp_config_file": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "CDP configuration file. Defaults to ~/.cdp/config",
			},
			"cdp_shared_credentials_file": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "CDP shared credentials file. Defaults to ~/.cdp/credentials",
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
	client, err := cdp.NewClient(getCdpConfig(ctx, data))

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

func getCdpConfig(ctx context.Context, data CdpProviderModel) *cdp.Config {
	tflog.Info(ctx, "Setting up CDP config")

	accessKeyId := getOrDefaultFromEnv(data.CdpAccessKeyId, "CDP_ACCESS_KEY_ID")
	privateKey := getOrDefaultFromEnv(data.CdpPrivateKey, "CDP_PRIVATE_KEY")
	cdpProfile := getOrDefaultFromEnv(data.Profile, "CDP_PROFILE", "CDP_DEFAULT_PROFILE")
	altusEndpointUrl := getOrDefaultFromEnv(data.AltusEndpointUrl, "ENDPOINT_URL")
	cdpEndpointUrl := getOrDefaultFromEnv(data.CdpEndpointUrl, "CDP_ENDPOINT_URL")
	cdpConfigFile := getOrDefaultFromEnv(data.CdpConfigFile, "CDP_CONFIG_FILE")
	cdpSharedCredentialsFile := getOrDefaultFromEnv(data.CdpSharedCredentialsFile, "CDP_SHARED_CREDENTIALS_FILE")

	config := cdp.Config{}
	config.WithProfile(cdpProfile)
	config.WithAltusApiEndpointUrl(altusEndpointUrl)
	config.WithCdpApiEndpointUrl(cdpEndpointUrl)
	config.WithCredentials(&authn.Credentials{
		AccessKeyId: accessKeyId,
		PrivateKey:  privateKey,
	})
	config.WithConfigFile(cdpConfigFile)
	config.WithCredentialsFile(cdpSharedCredentialsFile)

	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "privateKey")
	ctx = tflog.SetField(ctx, "accessKeyId", accessKeyId)
	ctx = tflog.SetField(ctx, "privateKey", privateKey)
	ctx = tflog.SetField(ctx, "cdpProfile", cdpProfile)
	ctx = tflog.SetField(ctx, "altusEndpointUrl", altusEndpointUrl)
	ctx = tflog.SetField(ctx, "cdpEndpointUrl", cdpEndpointUrl)
	ctx = tflog.SetField(ctx, "cdpConfigFile", cdpConfigFile)
	ctx = tflog.SetField(ctx, "cdpSharedCredentialsFile", cdpSharedCredentialsFile)

	tflog.Info(ctx, "CDP config set up. Creating client")
	return &config
}

func (p *CdpProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		environments.NewAwsCredentialResource,
	}
}

func (p *CdpProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		environments.NewAWSCredentialPrerequisitesDataSource,
	}
	/*	return map[string]*schema.Resource{
			"cdp_iam_group": iamresources.DataSourceGroup(),
			"cdp_environments_aws_credential_prerequisites": environmentsresources.DataSourceAWSCredentialPrerequisites(),
		}
	*/

}
