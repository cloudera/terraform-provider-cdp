package provider

import (
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/authn"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	datahubresources "github.com/cloudera/terraform-provider-cdp/resources/datahub"
	datalakeresources "github.com/cloudera/terraform-provider-cdp/resources/datalake"
	environmentsresources "github.com/cloudera/terraform-provider-cdp/resources/environments"
	iamresources "github.com/cloudera/terraform-provider-cdp/resources/iam"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	// Name of environment variable holding the users CDP access key ID.
	cdpAccessKeyIdEnvVar = "CDP_ACCESS_KEY_ID"

	// Name of environment variable holding the users CDP private key.
	cdpPrivateKeyEnvVar = "CDP_PRIVATE_KEY"

	// Name of system environment variable holding the name of profile to use
	// when reading the credentials file. Overrides cdpDefaultProfile.
	cdpDefaultProfileEnvVar = "CDP_DEFAULT_PROFILE"

	// TODO: is this CDP_PROFILE or CDP_DEFAULT_PROFILE? Both are respected for now.
	cdpProfileEnvVar = "CDP_PROFILE"

	//==== Below are fields for the provider =====

	// Provider key for configuring CDP access key id
	cdpAccessKeyIdField = "cdp_access_key_id"

	// Provider key for configuring CDP private key
	cdpPrivateKeyField = "cdp_private_key"

	profileField = "cdp_profile"

	// TODO: add endpoint URLs
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema:         providerSchema(),
		ResourcesMap:   resourcesMap(),
		DataSourcesMap: dataSourcesMap(),
		ConfigureFunc:  configureProvider,
	}
}

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		profileField: {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{cdpProfileEnvVar, cdpDefaultProfileEnvVar}, nil),
			Description: "CDP Profile to use for the configuration in ~/.cdp/",
		},
		cdpAccessKeyIdField: {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc(cdpAccessKeyIdEnvVar, nil),
			Description: "CDP access key id",
		},
		cdpPrivateKeyField: {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc(cdpPrivateKeyEnvVar, nil),
			Description: "CDP private key associated with the given access key",
		},
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	return cdp.NewClient(getCdpConfig(d))
}

func getCdpConfig(d *schema.ResourceData) *cdp.Config {
	accessKeyId := d.Get(cdpAccessKeyIdField).(string)
	privateKey := d.Get(cdpPrivateKeyField).(string)
	profile := d.Get(profileField).(string)

	config := cdp.NewConfig()
	config.WithProfile(profile)
	config.WithCredentials(&authn.Credentials{
		AccessKeyId: accessKeyId,
		PrivateKey:  privateKey,
	})
	return config
}

func resourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"cdp_iam_group":                       iamresources.ResourceGroup(),
		"cdp_environments_aws_credential":     environmentsresources.ResourceAWSCredential(),
		"cdp_environments_aws_environment":    environmentsresources.ResourceAWSEnvironment(),
		"cdp_environments_id_broker_mappings": environmentsresources.ResourceIDBrokerMappings(),
		"cdp_datalake_aws_datalake":           datalakeresources.ResourceAWSDatalake(),
		"cdp_datahub_aws_cluster":             datahubresources.ResourceAWSCluster(),
	}
}

func dataSourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"cdp_iam_group": iamresources.DataSourceGroup(),
		"cdp_environments_aws_credential_prerequisites": environmentsresources.DataSourceAWSCredentialPrerequisites(),
	}
}
