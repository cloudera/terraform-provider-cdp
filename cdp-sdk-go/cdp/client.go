package cdp

import (
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/authn"
	datahubclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client"
	datalakeclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client"
	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	iamclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/client"
	mlclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/client"
)

type Client struct {
	Environments *environmentsclient.Environments
	Datalake     *datalakeclient.Datalake
	Datahub      *datahubclient.Datahub
	Iam          *iamclient.Iam
	Ml           *mlclient.Ml
}

func NewClient(config *Config) (*Client, error) {
	environmentsClient, err := NewEnvironmentsClient(config)
	if err != nil {
		return nil, err
	}
	datalakeClient, err := NewDatalakeClient(config)
	if err != nil {
		return nil, err
	}

	datahubClient, err := NewDatahubClient(config)
	if err != nil {
		return nil, err
	}

	iamClient, err := NewIamClient(config)
	if err != nil {
		return nil, err
	}

	mlClient, err := NewMlClient(config)
	if err != nil {
		return nil, err
	}

	return &Client{
		Environments: environmentsClient,
		Datalake:     datalakeClient,
		Datahub:      datahubClient,
		Iam:          iamClient,
		Ml:           mlClient,
	}, nil
}

func NewIamClient(config *Config) (*iamclient.Iam, error) {
	transport, err := authn.GetAPIKeyAuthTransport(config.toInternalConfig(), "iam", true)
	if err != nil {
		return nil, err
	}
	return iamclient.New(transport, nil), nil
}

func NewEnvironmentsClient(config *Config) (*environmentsclient.Environments, error) {
	transport, err := authn.GetAPIKeyAuthTransport(config.toInternalConfig(), "environments", false)
	if err != nil {
		return nil, err
	}
	return environmentsclient.New(transport, nil), nil
}

func NewDatalakeClient(config *Config) (*datalakeclient.Datalake, error) {
	transport, err := authn.GetAPIKeyAuthTransport(config.toInternalConfig(), "datalake", false)
	if err != nil {
		return nil, err
	}
	return datalakeclient.New(transport, nil), nil
}

func NewDatahubClient(config *Config) (*datahubclient.Datahub, error) {
	transport, err := authn.GetAPIKeyAuthTransport(config.toInternalConfig(), "datahub", false)
	if err != nil {
		return nil, err
	}
	return datahubclient.New(transport, nil), nil
}

func NewMlClient(config *Config) (*mlclient.Ml, error) {
	transport, err := authn.GetAPIKeyAuthTransport(config.toInternalConfig(), "ml", false)
	if err != nil {
		return nil, err
	}
	return mlclient.New(transport, nil), nil
}
