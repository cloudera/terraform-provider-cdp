// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package auditcredential

import (
	"context"
	"errors"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/mocks"
)

func TestCreateAwsGovCloudAuditCredential(t *testing.T) {
	ctx := context.TODO()
	mockClient := new(mocks.MockEnvironmentClientService)

	cred := testAuditCredential()
	cred.AwsCredentialProperties.RoleArn = "arn:aws-us-gov:iam::123456789012:role/audit-role"

	createMatcher := func(params *operations.SetAWSGovCloudAuditCredentialParams) bool {
		return *params.Input.RoleArn == "arn:aws-us-gov:iam::123456789012:role/audit-role"
	}
	mockClient.On("SetAWSGovCloudAuditCredentialContext", mock.Anything, mock.MatchedBy(createMatcher)).Return(
		&operations.SetAWSGovCloudAuditCredentialOK{Payload: &models.SetAWSGovCloudAuditCredentialResponse{Credential: cred}}, nil)

	r := &awsGovCloudAuditCredentialResource{client: &cdp.Client{Environments: NewMockEnvironments(mockClient)}}
	schemaResp := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResp)

	req := resource.CreateRequest{Plan: tfsdk.Plan{Raw: createRawAwsGovCloudAuditCredentialResource(""), Schema: schemaResp.Schema}}
	resp := &resource.CreateResponse{State: tfsdk.State{Schema: schemaResp.Schema}}

	r.Create(ctx, req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	var state awsAuditCredentialResourceModel
	resp.State.Get(ctx, &state)
	assert.Equal(t, "test-audit-cred", state.ID.ValueString())
	mockClient.AssertExpectations(t)
}

func TestCreateAwsGovCloudAuditCredentialError(t *testing.T) {
	ctx := context.TODO()
	mockClient := new(mocks.MockEnvironmentClientService)

	mockClient.On("SetAWSGovCloudAuditCredentialContext", mock.Anything, mock.Anything).Return(
		nil, errors.New("connection timeout"))

	r := &awsGovCloudAuditCredentialResource{client: &cdp.Client{Environments: NewMockEnvironments(mockClient)}}
	schemaResp := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResp)

	req := resource.CreateRequest{Plan: tfsdk.Plan{Raw: createRawAwsGovCloudAuditCredentialResource(""), Schema: schemaResp.Schema}}
	resp := &resource.CreateResponse{State: tfsdk.State{Schema: schemaResp.Schema}}

	r.Create(ctx, req, resp)

	assert.True(t, resp.Diagnostics.HasError())
	assert.Equal(t, "Create Aws Govcloud Audit Credential", resp.Diagnostics.Errors()[0].Summary())
	mockClient.AssertExpectations(t)
}
