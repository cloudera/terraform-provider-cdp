// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package environments

import (
	"context"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/stretchr/testify/assert"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

func TestCloudPlatformIsAwsSetsAwsEnvironment(t *testing.T) {
	ctx := context.TODO()
	awsCloudPlatform := "AWS"
	env := createEnvironmentTestModel(&awsCloudPlatform)
	resp := &datasource.ReadResponse{}
	data := &EnvironmentConfigModel{}

	fillEnvironmentPlatformValid(ctx, env, resp, data)

	assert.NotNil(t, data.Aws)
	assert.Nil(t, data.Azure)
	assert.Nil(t, data.Gcp)
	assert.False(t, resp.Diagnostics.HasError())
}

func TestCloudPlatformIsAzureSetsAzureEnvironment(t *testing.T) {
	ctx := context.TODO()
	azureCloudPlatform := "AZURE"
	env := createEnvironmentTestModel(&azureCloudPlatform)
	resp := &datasource.ReadResponse{}
	data := &EnvironmentConfigModel{}

	fillEnvironmentPlatformValid(ctx, env, resp, data)

	assert.Nil(t, data.Aws)
	assert.NotNil(t, data.Azure)
	assert.Nil(t, data.Gcp)
	assert.False(t, resp.Diagnostics.HasError())
}

func TestCloudPlatformIsGcpSetsGcpEnvironment(t *testing.T) {
	ctx := context.TODO()
	gcpCloudPlatform := "GCP"
	env := createEnvironmentTestModel(&gcpCloudPlatform)
	resp := &datasource.ReadResponse{}
	data := &EnvironmentConfigModel{}

	fillEnvironmentPlatformValid(ctx, env, resp, data)

	assert.Nil(t, data.Aws)
	assert.Nil(t, data.Azure)
	assert.NotNil(t, data.Gcp)
	assert.False(t, resp.Diagnostics.HasError())
}

func TestCloudPlatformIsUnknownAddsError(t *testing.T) {
	ctx := context.TODO()
	unknownCloudPlatform := "somethingCompletelyDifferent"
	env := createEnvironmentTestModel(&unknownCloudPlatform)
	resp := &datasource.ReadResponse{}
	data := &EnvironmentConfigModel{}

	fillEnvironmentPlatformValid(ctx, env, resp, data)

	assert.Nil(t, data.Aws)
	assert.Nil(t, data.Azure)
	assert.Nil(t, data.Gcp)
	assert.True(t, resp.Diagnostics.HasError())
	assert.Equal(t, "Unknown cloud platform", resp.Diagnostics.Errors()[0].Summary())
}

func TestCloudPlatformIsNilAddsError(t *testing.T) {
	ctx := context.TODO()
	env := createEnvironmentTestModel(nil)
	resp := &datasource.ReadResponse{}
	data := &EnvironmentConfigModel{}

	fillEnvironmentPlatformValid(ctx, env, resp, data)

	assert.Nil(t, data.Aws)
	assert.Nil(t, data.Azure)
	assert.Nil(t, data.Gcp)
	assert.True(t, resp.Diagnostics.HasError())
	assert.Equal(t, "Cloud platform not set", resp.Diagnostics.Errors()[0].Summary())
}

func createEnvironmentTestModel(cloudPlatform *string) *models.Environment {
	status := "ACTIVE"
	env := &models.Environment{
		Authentication:                   nil,
		AwsComputeClusterConfiguration:   nil,
		AwsDetails:                       nil,
		AzureComputeClusterConfiguration: nil,
		AzureDetails:                     nil,
		BackupStorage:                    nil,
		CloudPlatform:                    cloudPlatform,
		ComputeClusterEnabled:            false,
		Created:                          strfmt.DateTime{},
		Creator:                          "",
		CredentialName:                   nil,
		Crn:                              nil,
		CustomDockerRegistry:             nil,
		DataServices:                     nil,
		Description:                      "",
		EnvironmentName:                  nil,
		Freeipa:                          nil,
		GcpDetails:                       nil,
		LogStorage:                       nil,
		Network:                          nil,
		ProxyConfig:                      nil,
		Region:                           nil,
		ReportDeploymentLogs:             false,
		SecurityAccess:                   nil,
		Status:                           &status,
		StatusReason:                     "",
		Tags:                             nil,
		TunnelEnabled:                    false,
		TunnelType:                       "",
		WorkloadAnalytics:                false,
	}
	return env
}
