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
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"math/big"
	"os"
	"regexp"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	environmentoperations "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/provider"
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
	TimeExternalProvider = map[string]resource.ExternalProvider{
		"time": {
			Source: "hashicorp/time",
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

func RandomShortWithPrefix(name string) string {
	upperLimit := big.NewInt(1000)
	nBig, err := rand.Int(rand.Reader, upperLimit)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s-%d", name, nBig.Int64())
}

func TestAccCdpProviderConfig() string {
	return `
provider "cdp" {
}
`
}

type AwsProvider struct {
	accessKeyID     string
	secretAccessKey string
	region          string
}

func NewAwsProvider(accessKeyID, secretAccessKey, region string) *AwsProvider {
	return &AwsProvider{
		accessKeyID:     accessKeyID,
		secretAccessKey: secretAccessKey,
		region:          region,
	}
}

func TestAccAwsProviderConfig(p *AwsProvider) string {
	return fmt.Sprintf(`
		provider "aws" {
		  access_key = %[1]q
		  secret_key = %[2]q
		  region  = %[3]q
		}
`, p.accessKeyID, p.secretAccessKey, p.region)
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

type AwsAccountCredentials struct {
	name          string
	accountID     string
	externalID    string
	defaultPolicy string
}

func NewAwsAccountCredentials(name string) *AwsAccountCredentials {
	return &AwsAccountCredentials{
		name: name,
	}
}

func getEnvironmentPrerequisites(t *testing.T, cloudPlatform string) *environmentsmodels.GetCredentialPrerequisitesResponse {
	client := GetCdpClientForAccTest()
	response, err := client.Environments.
		Operations.
		GetCredentialPrerequisites(
			environmentoperations.NewGetCredentialPrerequisitesParams().
				WithInput(&environmentsmodels.GetCredentialPrerequisitesRequest{
					CloudPlatform: &cloudPlatform,
				}),
		)
	assert.Nil(t, err)
	payload := response.GetPayload()
	assert.NotNil(t, payload)
	return payload
}

func (a *AwsAccountCredentials) WithPolicy(t *testing.T) *AwsAccountCredentials {
	payload := getEnvironmentPrerequisites(t, "AWS")
	assert.NotNil(t, payload)
	decodedBytes, err := base64.StdEncoding.DecodeString(*payload.Aws.PolicyJSON)
	assert.Nil(t, err)
	a.defaultPolicy = string(decodedBytes)
	return a
}

func (a *AwsAccountCredentials) WithExternalID(t *testing.T) *AwsAccountCredentials {
	payload := getEnvironmentPrerequisites(t, "AWS")
	assert.NotNil(t, payload)
	a.externalID = *payload.Aws.ExternalID
	return a
}

func (a *AwsAccountCredentials) WithAccountID(t *testing.T) *AwsAccountCredentials {
	payload := getEnvironmentPrerequisites(t, "AWS")
	assert.NotNil(t, payload)
	a.accountID = payload.AccountID
	return a
}

func CreateDefaultRoleAndPolicy(p *AwsAccountCredentials) string {
	return fmt.Sprintf(`
		resource "aws_iam_role" "cdp_test_role" {
		  name = "%[1]s-role"
		
		  assume_role_policy = <<EOF
		{
			"Version": "2012-10-17",
			"Statement": [
				{
					"Sid": "Statement1",
					"Effect": "Allow",
					"Principal": {
						"AWS": "arn:aws:iam::%[2]s:root"
					},
					"Action": "sts:AssumeRole",
					"Condition": {
						"StringEquals": {
							"sts:ExternalId": %[3]q
						}
					}
				}
			]
		}
		EOF
		
		  tags = {
			owner = "cdw-terraform-test@cloudera.com"
		  }
		}

		resource "aws_iam_policy" "cdp_test_policy" {
		  name        = "%[1]s-policy"
		  description = "DefaultCBPolicy for CDP, replace the static file with a CLI call"
		
		  policy = <<EOF
		  %[4]s
		EOF
		}

		resource "aws_iam_policy_attachment" "test-attach" {
		  name       = "test_attachment"
		  roles      = [aws_iam_role.cdp_test_role.name]
		  policy_arn = aws_iam_policy.cdp_test_policy.arn
		}
		`, p.name, p.accountID, p.externalID, p.defaultPolicy)
}
