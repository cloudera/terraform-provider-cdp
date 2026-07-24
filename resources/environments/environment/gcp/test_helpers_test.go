// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package gcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/mocks"
	"github.com/cloudera/terraform-provider-cdp/resources/environments/freeipa"
)

const (
	testEnvName            = "test-env"
	testNewKey             = "ssh-rsa NEW_KEY"
	testOldKey             = "ssh-rsa OLD_KEY"
	testSameKey            = "ssh-rsa SAME_KEY"
	testServiceUnavailable = "service unavailable"

	testOldProxyConfigName = "proxy1"
	testNewProxyConfigName = "proxy2"

	testOldDockerRegistryCrn = "crn:cdp:docker:us-west-1:old-registry"
	testNewDockerRegistryCrn = "crn:cdp:docker:us-west-1:new-registry"

	testOldCredentialName  = "old-credential"
	testNewCredentialName  = "new-credential"
	testSameCredentialName = "same-credential"

	testGatewaySchemePublic  = "PUBLIC"
	testGatewaySchemePrivate = "PRIVATE"

	testOldCatalogURL  = "https://old-catalog.example.com"
	testNewCatalogURL  = "https://new-catalog.example.com"
	testSameCatalogURL = "https://same-catalog.example.com"

	testOldDefaultSG  = "sg-old-default"
	testNewDefaultSG  = "sg-new-default"
	testOldKnoxSG     = "sg-old-knox"
	testNewKnoxSG     = "sg-new-knox"
	testSameDefaultSG = "sg-same-default"
	testSameKnoxSG    = "sg-same-knox"
)

func NewMockEnvironments(mockClient *mocks.MockEnvironmentClientService) *environmentsclient.Environments {
	return &environmentsclient.Environments{
		Operations: mockClient,
	}
}

func newFreeIpaObject(catalog string) types.Object {
	instances, _ := types.SetValueFrom(context.TODO(), freeipa.FreeIpaInstanceType, []freeipa.FreeIpaInstance{})
	recipes, _ := types.SetValueFrom(context.TODO(), types.StringType, []string{})
	obj, _ := basetypes.NewObjectValueFrom(context.TODO(), freeipa.FreeIpaDetailsType.AttrTypes, &freeipa.FreeIpaDetails{
		Catalog:              types.StringValue(catalog),
		ImageID:              types.StringValue("img-1"),
		Os:                   types.StringValue("centos7"),
		InstanceCountByGroup: types.Int32Value(1),
		InstanceType:         types.StringValue("m5.xlarge"),
		Instances:            instances,
		MultiAz:              types.BoolValue(false),
		Recipes:              recipes,
		Architecture:         types.StringValue("X86_64"),
	})
	return obj
}
