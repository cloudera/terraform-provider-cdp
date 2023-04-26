package cdp

import (
	datahubclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client"
	datalakeclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client"
	dwclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/client"
	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	iamclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/client"
	mlclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/client"
)

type Client struct {
	config       *Config
	Environments *environmentsclient.Environments
	Datalake     *datalakeclient.Datalake
	Datahub      *datahubclient.Datahub
	Iam          *iamclient.Iam
	Ml           *mlclient.Ml
	Dw           *dwclient.Dw
}

func NewClient(config *Config) (*Client, error) {
	if err := config.initConfig(); err != nil {
		return nil, err
	}

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

	dwClient, err := NewDwClient(config)
	if err != nil {
		return nil, err
	}

	return &Client{
		config:       config,
		Environments: environmentsClient,
		Datalake:     datalakeClient,
		Datahub:      datahubClient,
		Iam:          iamClient,
		Ml:           mlClient,
		Dw:           dwClient,
	}, nil
}

func NewIamClient(config *Config) (*iamclient.Iam, error) {
	credentials, err := config.GetCredentials()
	if err != nil {
		return nil, err
	}
	transport, err := GetAPIKeyAuthTransport(config.Context, config.Logger, credentials, config.GetEndpoint("iam", true), config.BaseApiPath, config.GetLocalEnvironment())
	if err != nil {
		return nil, err
	}
	return iamclient.New(transport, nil), nil
}

func NewEnvironmentsClient(config *Config) (*environmentsclient.Environments, error) {
	credentials, err := config.GetCredentials()
	if err != nil {
		return nil, err
	}
	transport, err := GetAPIKeyAuthTransport(config.Context, config.Logger, credentials, config.GetEndpoint("environments", false), config.BaseApiPath, config.GetLocalEnvironment())
	if err != nil {
		return nil, err
	}
	return environmentsclient.New(transport, nil), nil
}

func NewDatalakeClient(config *Config) (*datalakeclient.Datalake, error) {
	credentials, err := config.GetCredentials()
	if err != nil {
		return nil, err
	}
	transport, err := GetAPIKeyAuthTransport(config.Context, config.Logger, credentials, config.GetEndpoint("datalake", false), config.BaseApiPath, config.GetLocalEnvironment())
	if err != nil {
		return nil, err
	}
	return datalakeclient.New(transport, nil), nil
}

func NewDatahubClient(config *Config) (*datahubclient.Datahub, error) {
	credentials, err := config.GetCredentials()
	if err != nil {
		return nil, err
	}
	transport, err := GetAPIKeyAuthTransport(config.Context, config.Logger, credentials, config.GetEndpoint("datahub", false), config.BaseApiPath, config.GetLocalEnvironment())
	if err != nil {
		return nil, err
	}
	return datahubclient.New(transport, nil), nil
}

func NewMlClient(config *Config) (*mlclient.Ml, error) {
	credentials, err := config.GetCredentials()
	if err != nil {
		return nil, err
	}
	transport, err := GetAPIKeyAuthTransport(config.Context, config.Logger, credentials, config.GetEndpoint("ml", false), config.BaseApiPath, config.GetLocalEnvironment())
	if err != nil {
		return nil, err
	}
	return mlclient.New(transport, nil), nil
}

func NewDwClient(config *Config) (*dwclient.Dw, error) {
	credentials, err := config.GetCredentials()
	if err != nil {
		return nil, err
	}
	transport, err := GetAPIKeyAuthTransport(config.Context, config.Logger, credentials, config.GetEndpoint("dw", false), config.BaseApiPath, config.GetLocalEnvironment())
	if err != nil {
		return nil, err
	}
	return dwclient.New(transport, nil), nil
}
