package cdp

import (
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/authn"
	datahubclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client"
	datalakeclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client"
	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
)

type Client struct {
	Environments *environmentsclient.Environments
	Datalake     *datalakeclient.Datalake
	Datahub      *datahubclient.Datahub
}

func NewClient(config *Config) Client {
	return Client{
		Environments: NewEnvironmentsClient(config),
		Datalake:     NewDatalakeClient(config),
		Datahub:      NewDatahubClient(config),
	}

}

func NewEnvironmentsClient(config *Config) *environmentsclient.Environments {
	transport, err := authn.GetAPIKeyAuthTransport(config.CdpApiEndpointUrl, "")
	if err != nil {
		panic(err)
	}
	return environmentsclient.New(transport, nil)
}

func NewDatalakeClient(config *Config) *datalakeclient.Datalake {
	transport, err := authn.GetAPIKeyAuthTransport(config.CdpApiEndpointUrl, "")
	if err != nil {
		panic(err)
	}
	return datalakeclient.New(transport, nil)
}

func NewDatahubClient(config *Config) *datahubclient.Datahub {
	transport, err := authn.GetAPIKeyAuthTransport(config.CdpApiEndpointUrl, "")
	if err != nil {
		panic(err)
	}
	return datahubclient.New(transport, nil)
}
