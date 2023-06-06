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
	transport, err := buildClientTransportWithDefaultHttpTransport(config, config.GetEndpoint("iam", true))
	if err != nil {
		return nil, err
	}
	return iamclient.New(transport, nil), nil
}

func NewEnvironmentsClient(config *Config) (*environmentsclient.Environments, error) {
	transport, err := buildClientTransportWithDefaultHttpTransport(config, config.GetEndpoint("environments", false))
	if err != nil {
		return nil, err
	}
	return environmentsclient.New(transport, nil), nil
}

func NewDatalakeClient(config *Config) (*datalakeclient.Datalake, error) {
	transport, err := buildClientTransportWithDefaultHttpTransport(config, config.GetEndpoint("datalake", false))
	if err != nil {
		return nil, err
	}
	return datalakeclient.New(transport, nil), nil
}

func NewDatahubClient(config *Config) (*datahubclient.Datahub, error) {
	transport, err := buildClientTransportWithDefaultHttpTransport(config, config.GetEndpoint("datahub", false))
	if err != nil {
		return nil, err
	}
	return datahubclient.New(transport, nil), nil
}

func NewMlClient(config *Config) (*mlclient.Ml, error) {
	transport, err := buildClientTransportWithDefaultHttpTransport(config, config.GetEndpoint("ml", false))
	if err != nil {
		return nil, err
	}
	return mlclient.New(transport, nil), nil
}

func NewDwClient(config *Config) (*dwclient.Dw, error) {
	transport, err := buildClientTransportWithDefaultHttpTransport(config, config.GetEndpoint("dw", false))
	if err != nil {
		return nil, err
	}
	return dwclient.New(transport, nil), nil
}
