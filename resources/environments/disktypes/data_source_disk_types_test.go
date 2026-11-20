// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package disktypes

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

func createRawDiskTypesDataSource() tftypes.Value {
	entryType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"platform":          tftypes.String,
			"default_disk_type": tftypes.String,
			"disk_types":        tftypes.List{ElementType: tftypes.String},
			"display_names":     tftypes.Map{ElementType: tftypes.String},
			"disk_mappings":     tftypes.Map{ElementType: tftypes.String},
		},
	}
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"platform_disk_types": tftypes.List{ElementType: entryType},
			},
		},
		map[string]tftypes.Value{
			"platform_disk_types": tftypes.NewValue(tftypes.List{ElementType: entryType}, nil),
		},
	)
}

func TestReadDiskTypes(t *testing.T) {
	testCases := map[string]struct {
		response       any
		errorResponse  any
		expectedError  bool
		expectedDetail string
	}{
		"OK": {
			response: &operations.ListDiskTypesOK{
				Payload: &models.ListDiskTypesResponse{
					PlatformDiskTypes: []*models.PlatformDiskTypesEntry{
						{
							Platform:        "AWS",
							DefaultDiskType: "gp2",
							DiskTypes:       []string{"gp2", "gp3"},
						},
					},
				},
			},
			expectedError: false,
		},
		"BadRequest": {
			response: nil,
			errorResponse: &operations.ListDiskTypesDefault{
				Payload: &models.Error{
					Code:    "BAD_REQUEST",
					Message: "Something went wrong",
				},
			},
			expectedError:  true,
			expectedDetail: "Something went wrong",
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
			mockClient.On("ListDiskTypesContext", mock.Anything, mock.Anything).Return(tc.response, tc.errorResponse)

			ds := &diskTypesDataSource{
				client: &cdp.Client{Environments: newMockEnvironments(mockClient)},
			}

			req := datasource.ReadRequest{
				Config: tfsdk.Config{
					Raw:    createRawDiskTypesDataSource(),
					Schema: diskTypesSchema,
				},
			}
			resp := &datasource.ReadResponse{
				State:       tfsdk.State{Schema: diskTypesSchema},
				Diagnostics: make(diag.Diagnostics, 0),
			}

			ds.Read(ctx, req, resp)

			assert.Equal(t, tc.expectedError, resp.Diagnostics.HasError())
			if tc.expectedError {
				assert.True(t, strings.Contains(resp.Diagnostics.Errors()[0].Detail(), tc.expectedDetail))
			} else {
				var state diskTypesModel
				resp.State.Get(ctx, &state)
				assert.Equal(t, 1, len(state.PlatformDiskTypes))
				assert.Equal(t, "AWS", state.PlatformDiskTypes[0].Platform.ValueString())
			}

			mockClient.AssertExpectations(t)
		})
	}
}
