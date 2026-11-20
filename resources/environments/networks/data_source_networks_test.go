// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package networks

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
	"github.com/cloudera/terraform-provider-cdp/mocks"
)

func createRawNetworksDataSource(credential string) tftypes.Value {
	cloudNetType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"id":         tftypes.String,
			"name":       tftypes.String,
			"properties": tftypes.Map{ElementType: tftypes.String},
			"subnets":    tftypes.Map{ElementType: tftypes.String},
		},
	}
	platNetType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"platform": tftypes.String,
			"networks": tftypes.List{ElementType: cloudNetType},
		},
	}
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"credential":        tftypes.String,
				"region":            tftypes.String,
				"availability_zone": tftypes.String,
				"platform_variant":  tftypes.String,
				"network_id":        tftypes.String,
				"shared_project_id": tftypes.String,
				"subnet_ids":        tftypes.List{ElementType: tftypes.String},
				"networks":          tftypes.List{ElementType: platNetType},
			},
		},
		map[string]tftypes.Value{
			"credential":        tftypes.NewValue(tftypes.String, credential),
			"region":            tftypes.NewValue(tftypes.String, "us-west-2"),
			"availability_zone": tftypes.NewValue(tftypes.String, nil),
			"platform_variant":  tftypes.NewValue(tftypes.String, nil),
			"network_id":        tftypes.NewValue(tftypes.String, nil),
			"shared_project_id": tftypes.NewValue(tftypes.String, nil),
			"subnet_ids":        tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, nil),
			"networks":          tftypes.NewValue(tftypes.List{ElementType: platNetType}, nil),
		},
	)
}

func TestReadNetworks(t *testing.T) {
	testCases := map[string]struct {
		response       any
		errorResponse  any
		expectedError  bool
		expectedDetail string
	}{
		"OK": {
			response: &operations.ListNetworksOK{
				Payload: &models.ListNetworksResponse{
					Networks: []*models.PlatformNetwork{
						{
							Platform: "AWS",
							Networks: []*models.CloudNetwork{
								{
									ID:   "vpc-123",
									Name: "my-vpc",
								},
							},
						},
					},
				},
			},
			expectedError: false,
		},
		"BadRequest": {
			response: nil,
			errorResponse: &operations.ListNetworksDefault{
				Payload: &models.Error{
					Code:    "BAD_REQUEST",
					Message: "Invalid credential",
				},
			},
			expectedError:  true,
			expectedDetail: "Invalid credential",
		},
		"TransportError": {
			response:       nil,
			errorResponse:  errors.New("connection refused"),
			expectedError:  true,
			expectedDetail: "connection refused",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()

			mockClient := new(mocks.MockEnvironmentClientService)
			mockClient.On("ListNetworksContext", mock.Anything, mock.Anything).Return(tc.response, tc.errorResponse)

			ds := &networksDataSource{
				client: &cdp.Client{Environments: newMockEnvironments(mockClient)},
			}

			req := datasource.ReadRequest{
				Config: tfsdk.Config{
					Raw:    createRawNetworksDataSource("my-cred"),
					Schema: networksSchema,
				},
			}
			resp := &datasource.ReadResponse{
				State:       tfsdk.State{Schema: networksSchema},
				Diagnostics: make(diag.Diagnostics, 0),
			}

			ds.Read(ctx, req, resp)

			assert.Equal(t, tc.expectedError, resp.Diagnostics.HasError())
			if tc.expectedError {
				assert.True(t, strings.Contains(resp.Diagnostics.Errors()[0].Detail(), tc.expectedDetail))
			} else {
				var state networksModel
				resp.State.Get(ctx, &state)
				assert.Equal(t, 1, len(state.Networks))
				assert.Equal(t, "AWS", state.Networks[0].Platform.ValueString())
			}

			mockClient.AssertExpectations(t)
		})
	}
}
