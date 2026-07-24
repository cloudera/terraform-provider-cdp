// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package aws

import (
	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/mocks"
)

const (
	testEnvName            = "test-env"
	testNewKey             = "ssh-rsa NEW_KEY"
	testOldKey             = "ssh-rsa OLD_KEY"
	testSameKey            = "ssh-rsa SAME_KEY"
	testServiceUnavailable = "service unavailable"
	testClusterInitFailed  = "cluster init failed"
	testOldDefaultSG       = "sg-old-default"
	testNewDefaultSG       = "sg-new-default"
	testOldKnoxSG          = "sg-old-knox"
	testNewKnoxSG          = "sg-new-knox"
	testSameDefaultSG      = "sg-same-default"
	testSameKnoxSG         = "sg-same-knox"
)

func NewMockEnvironments(mockClient *mocks.MockEnvironmentClientService) *environmentsclient.Environments {
	return &environmentsclient.Environments{
		Operations: mockClient,
	}
}
