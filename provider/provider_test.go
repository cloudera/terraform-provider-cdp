// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"regexp"
	"testing"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"cdp": providerserver.NewProtocol6WithError(New("test")()),
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
}

func TestProviderOverridesUserAgent(t *testing.T) {
	model := CdpProviderModel{
		CdpAccessKeyId:           types.StringValue("cdp-access-key"),
		CdpPrivateKey:            types.StringValue("cdp-private-key"),
		Profile:                  types.StringValue("profile"),
		AltusEndpointUrl:         types.StringValue("altus-endpoint-url"),
		CdpEndpointUrl:           types.StringValue("cdp-endpoint-url"),
		CdpConfigFile:            types.StringValue("cdp-config-file"),
		CdpSharedCredentialsFile: types.StringValue("cdp-shared-credentials-file"),
		LocalEnvironment:         types.BoolValue(false),
	}

	config := getCdpConfig(context.Background(), model, "0.1.0", "v1.4.2")
	userAgent := config.GetUserAgentOrDefault()

	r, _ := regexp.Compile(`^CDPTFPROVIDER/.+ Terraform/.+ Go/.+ .+_.+$`)
	if !r.MatchString(userAgent) {
		t.Fatalf("Failed to match the User-Agent regex: %v", userAgent)
	}
}
