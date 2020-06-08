package cdp

import (
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/authn"
	datahubclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client"
	datalakeclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client"
	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	iamclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/client"
)

type Client struct {
	Environments *environmentsclient.Environments
	Datalake     *datalakeclient.Datalake
	Datahub      *datahubclient.Datahub
	IAM          *iamclient.Iam
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

	return &Client{
		Environments: environmentsClient,
		Datalake:     datalakeClient,
		Datahub:      datahubClient,
		IAM:          iamClient,
	}, nil
}

func NewIamClient(config *Config) (*iamclient.Iam, error) {
	transport, err := authn.GetAPIKeyAuthTransport(config.toInternalConfig(), true)
	if err != nil {
		return nil, err
	}
	return iamclient.New(transport, nil), nil
}

func NewEnvironmentsClient(config *Config) (*environmentsclient.Environments, error) {
	transport, err := authn.GetAPIKeyAuthTransport(config.toInternalConfig(), false)
	if err != nil {
		return nil, err
	}
	return environmentsclient.New(transport, nil), nil
}

func NewDatalakeClient(config *Config) (*datalakeclient.Datalake, error) {
	transport, err := authn.GetAPIKeyAuthTransport(config.toInternalConfig(), false)
	if err != nil {
		return nil, err
	}
	return datalakeclient.New(transport, nil), nil
}

func NewDatahubClient(config *Config) (*datahubclient.Datahub, error) {
	transport, err := authn.GetAPIKeyAuthTransport(config.toInternalConfig(), false)
	if err != nil {
		return nil, err
	}
	return datahubclient.New(transport, nil), nil
}
