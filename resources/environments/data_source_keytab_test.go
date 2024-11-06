// Copyright 2024 Cloudera. All Rights Reserved.
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
	"errors"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	mocks "github.com/cloudera/terraform-provider-cdp/mocks/github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
)

func createRawKeytabDataSource(environment string) tftypes.Value {
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"environment": tftypes.String,
				"actor_crn":   tftypes.String,
				"keytab":      tftypes.String,
			},
		},
		map[string]tftypes.Value{
			"environment": tftypes.NewValue(tftypes.String, environment),
			"actor_crn":   tftypes.NewValue(tftypes.String, "test-actor"),
			"keytab":      tftypes.NewValue(tftypes.String, "test-description"),
		},
	)
}

func TestFetchKeytabs(t *testing.T) {
	testCases := map[string]struct {
		expectedResponse      interface{}
		expectedErrorResponse interface{}
		expectedError         bool
		expectedSummary       string
		expectedDetail        string
		expectedEnvironment   string
		expectedActorCrn      string
		expectedKeytab        string
	}{
		"OK": {
			expectedResponse: &operations.GetKeytabOK{
				Payload: &models.GetKeytabResponse{
					Contents: "ZG9uJ3QgYmUgdG9vIGN1cmlvdXMgOik=",
				},
			},
			expectedErrorResponse: nil,
			expectedError:         false,
			expectedDetail:        "",
			expectedEnvironment:   "test-environment",
			expectedActorCrn:      "test-actor",
			expectedKeytab:        "ZG9uJ3QgYmUgdG9vIGN1cmlvdXMgOik=",
		},
		"BadRequest": {
			expectedResponse: nil,
			expectedErrorResponse: &operations.GetKeytabDefault{
				Payload: &models.Error{
					Code:    "BAD_REQUEST",
					Message: "Missing environment name field",
				},
			},
			expectedError:       true,
			expectedDetail:      "Missing environment name field",
			expectedEnvironment: "",
			expectedActorCrn:    "",
			expectedKeytab:      "",
		},
		"TransportError": {
			expectedResponse:      nil,
			expectedErrorResponse: errors.New("request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)"),
			expectedError:         true,
			expectedDetail:        "request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)",
			expectedEnvironment:   "",
			expectedActorCrn:      "",
			expectedKeytab:        "",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()

			mockClient := new(mocks.MockEnvironmentClientService)

			createMatcher := func(params *operations.GetKeytabParams) bool {
				match := *params.Input.EnvironmentName == "test-environment"
				match = match && params.Input.ActorCrn == "test-actor"
				return match
			}
			mockClient.On("GetKeytab", mock.MatchedBy(createMatcher)).Return(testCase.expectedResponse, testCase.expectedErrorResponse)

			pcResource := &keytabDataSource{
				client: &cdp.Client{Environments: NewMockEnvironments(mockClient)},
			}

			req := datasource.ReadRequest{
				Config: tfsdk.Config{
					Raw:    createRawKeytabDataSource("test-environment"),
					Schema: KeytabSchema,
				},
				ProviderMeta: tfsdk.Config{},
			}

			resp := &datasource.ReadResponse{
				State: tfsdk.State{
					Schema: KeytabSchema,
				},
				Diagnostics: make(diag.Diagnostics, 0),
			}

			pcResource.Read(ctx, req, resp)

			assert.Equal(t, testCase.expectedError, resp.Diagnostics.HasError())
			if testCase.expectedError {
				assert.True(t, strings.Contains(resp.Diagnostics.Errors()[0].Detail(), testCase.expectedDetail))
			}

			var state KeytabModel
			resp.State.Get(ctx, &state)

			assert.Equal(t, testCase.expectedEnvironment, state.Environment.ValueString())
			assert.Equal(t, testCase.expectedActorCrn, state.ActorCrn.ValueString())
			assert.Equal(t, testCase.expectedKeytab, state.Keytab.ValueString())

			mockClient.AssertExpectations(t)
		})
	}

}
