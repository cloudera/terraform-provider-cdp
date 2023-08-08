// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package cdpacctest

import (
	"fmt"
	"os"
	"regexp"
	"sync"
	"testing"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pkg/errors"
)

const (
	ResourcePrefix = "tf-acc-test"
)

var TestAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"cdp": providerserver.NewProtocol6WithError(provider.New("test")()),
}

var (
	preCheckOnce sync.Once
	crnPattern   = regexp.MustCompile(`^crn:([\w\-]+):(\w+):([\w\-]+):([\w\-]+):(\w+):(\S+)$`)

	AwsExternalProvider = map[string]resource.ExternalProvider{
		"aws": {
			Source:            "hashicorp/aws",
			VersionConstraint: "~> 5.0",
		},
	}
	HttpExternalProvider = map[string]resource.ExternalProvider{
		"http": {
			Source:            "hashicorp/http",
			VersionConstraint: "~> 3.4",
		},
	}

	cdpClientOnce sync.Once
	cdpClient     *cdp.Client
)

//nolint:all
func PreCheck(t *testing.T) {
	preCheckOnce.Do(func() {
		if os.Getenv(cdp.CdpProfileEnvVar) == "" && os.Getenv(cdp.CdpAccessKeyIdEnvVar) == "" {
			t.Fatalf("Terraform acceptance testing requires either %s or %s environment variables to be set", cdp.CdpProfileEnvVar, cdp.CdpAccessKeyIdEnvVar)
		}

		if os.Getenv(cdp.CdpAccessKeyIdEnvVar) != "" {
			if _, ok := os.LookupEnv(cdp.CdpPrivateKeyEnvVar); !ok {
				t.Fatalf("Environment variable %s should be set together with %s", cdp.CdpPrivateKeyEnvVar, cdp.CdpAccessKeyIdEnvVar)
			}
		}

		// TODO: check whether AWS, Azure or GCP is configured.
	})
}

func ConcatExternalProviders(providerMaps ...map[string]resource.ExternalProvider) map[string]resource.ExternalProvider {
	result := make(map[string]resource.ExternalProvider)

	for _, p := range providerMaps {
		for k, v := range p {
			result[k] = v
		}
	}

	return result
}

func TestAccCdpProviderConfig() string {
	return `
provider "cdp" {
}
`
}

func TestAccAwsProviderConfig() string {
	return `
provider "aws" {
}
`
}

// CheckCrn Checks whether the value is set and is a properly formatted CRN
func CheckCrn(value string) error {
	if value == "" {
		return errors.New("Expected CRN to be set")
	}

	if !crnPattern.MatchString(value) {
		return errors.New(fmt.Sprintf("Provided value does not match expected CRN format: %s", value))
	}

	return nil
}

func GetCdpClientForAccTest() *cdp.Client {
	cdpClientOnce.Do(func() {
		config := cdp.NewConfig()
		var err error
		cdpClient, err = cdp.NewClient(config)
		if err != nil {
			panic(fmt.Sprintf("error while creating CDP Client: %v", err))
		}
	})
	return cdpClient
}
