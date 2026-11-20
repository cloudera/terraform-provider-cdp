// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package regions

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

func createRawRegionsDataSource(credential string) tftypes.Value {
	azType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"region":             tftypes.String,
			"availability_zones": tftypes.List{ElementType: tftypes.String},
		},
	}
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"credential":                tftypes.String,
				"environment_crn":           tftypes.String,
				"region":                    tftypes.String,
				"availability_zone":         tftypes.String,
				"availability_zones_needed": tftypes.Bool,
				"platform_variant":          tftypes.String,
				"names":                     tftypes.List{ElementType: tftypes.String},
				"default_region":            tftypes.String,
				"display_names":             tftypes.Map{ElementType: tftypes.String},
				"availability_zones_result": tftypes.List{ElementType: azType},
			},
		},
		map[string]tftypes.Value{
			"credential":                tftypes.NewValue(tftypes.String, credential),
			"environment_crn":           tftypes.NewValue(tftypes.String, nil),
			"region":                    tftypes.NewValue(tftypes.String, nil),
			"availability_zone":         tftypes.NewValue(tftypes.String, nil),
			"availability_zones_needed": tftypes.NewValue(tftypes.Bool, nil),
			"platform_variant":          tftypes.NewValue(tftypes.String, nil),
			"names":                     tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, nil),
			"default_region":            tftypes.NewValue(tftypes.String, nil),
			"display_names":             tftypes.NewValue(tftypes.Map{ElementType: tftypes.String}, nil),
			"availability_zones_result": tftypes.NewValue(tftypes.List{ElementType: azType}, nil),
		},
	)
}

func createRawRegionsDataSourceBothNil() tftypes.Value {
	azType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"region":             tftypes.String,
			"availability_zones": tftypes.List{ElementType: tftypes.String},
		},
	}
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"credential":                tftypes.String,
				"environment_crn":           tftypes.String,
				"region":                    tftypes.String,
				"availability_zone":         tftypes.String,
				"availability_zones_needed": tftypes.Bool,
				"platform_variant":          tftypes.String,
				"names":                     tftypes.List{ElementType: tftypes.String},
				"default_region":            tftypes.String,
				"display_names":             tftypes.Map{ElementType: tftypes.String},
				"availability_zones_result": tftypes.List{ElementType: azType},
			},
		},
		map[string]tftypes.Value{
			"credential":                tftypes.NewValue(tftypes.String, nil),
			"environment_crn":           tftypes.NewValue(tftypes.String, nil),
			"region":                    tftypes.NewValue(tftypes.String, nil),
			"availability_zone":         tftypes.NewValue(tftypes.String, nil),
			"availability_zones_needed": tftypes.NewValue(tftypes.Bool, nil),
			"platform_variant":          tftypes.NewValue(tftypes.String, nil),
			"names":                     tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, nil),
			"default_region":            tftypes.NewValue(tftypes.String, nil),
			"display_names":             tftypes.NewValue(tftypes.Map{ElementType: tftypes.String}, nil),
			"availability_zones_result": tftypes.NewValue(tftypes.List{ElementType: azType}, nil),
		},
	)
}

func TestRegionsConfigValidatorRejectsNoCredentialOrCrn(t *testing.T) {
	ctx := context.TODO()
	ds := &regionsDataSource{}

	req := datasource.ValidateConfigRequest{
		Config: tfsdk.Config{Raw: createRawRegionsDataSourceBothNil(), Schema: regionsSchema},
	}

	var diagnostics diag.Diagnostics
	for _, v := range ds.ConfigValidators(ctx) {
		resp := &datasource.ValidateConfigResponse{}
		v.ValidateDataSource(ctx, req, resp)
		diagnostics.Append(resp.Diagnostics...)
	}

	assert.True(t, diagnostics.HasError())
	assert.Contains(t, diagnostics.Errors()[0].Summary(), "Missing Attribute Configuration")
}

func TestRegionsConfigValidatorAcceptsCredential(t *testing.T) {
	ctx := context.TODO()
	ds := &regionsDataSource{}

	req := datasource.ValidateConfigRequest{
		Config: tfsdk.Config{Raw: createRawRegionsDataSource("my-cred"), Schema: regionsSchema},
	}

	var diagnostics diag.Diagnostics
	for _, v := range ds.ConfigValidators(ctx) {
		resp := &datasource.ValidateConfigResponse{}
		v.ValidateDataSource(ctx, req, resp)
		diagnostics.Append(resp.Diagnostics...)
	}

	assert.False(t, diagnostics.HasError())
}

func TestReadRegions(t *testing.T) {
	testCases := map[string]struct {
		response       any
		errorResponse  any
		expectedError  bool
		expectedDetail string
	}{
		"OK": {
			response: &operations.ListRegionsOK{
				Payload: &models.ListRegionsResponse{
					Names:         []string{"us-west-2", "eu-west-1"},
					DefaultRegion: "us-west-2",
				},
			},
			expectedError: false,
		},
		"BadRequest": {
			response: nil,
			errorResponse: &operations.ListRegionsDefault{
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
			mockClient.On("ListRegionsContext", mock.Anything, mock.Anything).Return(tc.response, tc.errorResponse)

			ds := &regionsDataSource{
				client: &cdp.Client{Environments: newMockEnvironments(mockClient)},
			}

			req := datasource.ReadRequest{
				Config: tfsdk.Config{
					Raw:    createRawRegionsDataSource("my-cred"),
					Schema: regionsSchema,
				},
			}
			resp := &datasource.ReadResponse{
				State:       tfsdk.State{Schema: regionsSchema},
				Diagnostics: make(diag.Diagnostics, 0),
			}

			ds.Read(ctx, req, resp)

			assert.Equal(t, tc.expectedError, resp.Diagnostics.HasError())
			if tc.expectedError {
				assert.True(t, strings.Contains(resp.Diagnostics.Errors()[0].Detail(), tc.expectedDetail))
			} else {
				var state regionsModel
				resp.State.Get(ctx, &state)
				assert.Equal(t, 2, len(state.Names.Elements()))
				assert.Equal(t, "us-west-2", state.DefaultRegion.ValueString())
			}

			mockClient.AssertExpectations(t)
		})
	}
}
