package provider

import (
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/authn"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	datahubclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client"
	datalakeclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client"
	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/resources/datahub"
	"github.com/cloudera/terraform-provider-cdp/resources/datalake"
	"github.com/cloudera/terraform-provider-cdp/resources/environments"
	"github.com/cloudera/terraform-provider-cdp/utils"
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

	profileField = "profile"

	// TODO: add endpoint URLs
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema:        providerSchema(),
		ResourcesMap:  resourcesMap(),
		ConfigureFunc: configureProvider,
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
	config := getCdpConfig(d)
	transport, err := authn.GetAPIKeyAuthTransport(config.CdpApiEndpointUrl, config.BaseAPIPath)
	if err != nil {
		return nil, err
	}

	cdpClients := utils.CdpClients{
		Environments: environmentsclient.New(transport, nil),
		Datalake:     datalakeclient.New(transport, nil),
		Datahub:      datahubclient.New(transport, nil),
	}

	return &cdpClients, nil
}

func getCdpConfig(d *schema.ResourceData) *cdp.Config {
	accessKeyId := d.Get(cdpPrivateKeyField).(string)
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
		"cdp_environments_credential":  environments.ResourceCredential(),
		"cdp_environments_environment": environments.ResourceEnvironment(),
		"cdp_datalake_datalake":        datalake.ResourceDatalake(),
		"cdp_datahub_cluster":          datahub.ResourceCluster(),
	}
}
