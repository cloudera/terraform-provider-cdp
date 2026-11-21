// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package vmtypes

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

func createRawVMTypesDataSource(credName string) tftypes.Value {
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"credential_name":   tftypes.String,
				"credential_crn":    tftypes.String,
				"region":            tftypes.String,
				"availability_zone": tftypes.String,
				"architecture":      tftypes.String,
				"platform_variant":  tftypes.String,
				"resource_type":     tftypes.String,
				"vm_types":          tftypes.List{ElementType: tftypes.Object{AttributeTypes: vmTypesObjectType()}},
			},
		},
		map[string]tftypes.Value{
			"credential_name":   tftypes.NewValue(tftypes.String, credName),
			"credential_crn":    tftypes.NewValue(tftypes.String, nil),
			"region":            tftypes.NewValue(tftypes.String, "us-west-2"),
			"availability_zone": tftypes.NewValue(tftypes.String, nil),
			"architecture":      tftypes.NewValue(tftypes.String, nil),
			"platform_variant":  tftypes.NewValue(tftypes.String, nil),
			"resource_type":     tftypes.NewValue(tftypes.String, nil),
			"vm_types":          tftypes.NewValue(tftypes.List{ElementType: tftypes.Object{AttributeTypes: vmTypesObjectType()}}, nil),
		},
	)
}

func vmTypesObjectType() map[string]tftypes.Type {
	return map[string]tftypes.Type{
		"platform":                tftypes.String,
		"default_virtual_machine": tftypes.Object{AttributeTypes: vmTypeObjectType()},
		"virtual_machines":        tftypes.List{ElementType: tftypes.Object{AttributeTypes: vmTypeObjectType()}},
	}
}

func vmTypeObjectType() map[string]tftypes.Type {
	return map[string]tftypes.Type{
		"name": tftypes.String,
		"vm_type_metadata": tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"configs": tftypes.List{ElementType: tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"maximum_number":        tftypes.Number,
						"maximum_size":          tftypes.Number,
						"minimum_number":        tftypes.Number,
						"minimum_size":          tftypes.Number,
						"volume_parameter_type": tftypes.String,
					},
				}},
				"properties": tftypes.Map{ElementType: tftypes.String},
			},
		},
	}
}

func TestVMTypesConfigValidatorRejectsNoCredential(t *testing.T) {
	ctx := context.TODO()
	ds := &vmTypesDataSource{}

	raw := tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"credential_name":   tftypes.String,
				"credential_crn":    tftypes.String,
				"region":            tftypes.String,
				"availability_zone": tftypes.String,
				"architecture":      tftypes.String,
				"platform_variant":  tftypes.String,
				"resource_type":     tftypes.String,
				"vm_types":          tftypes.List{ElementType: tftypes.Object{AttributeTypes: vmTypesObjectType()}},
			},
		},
		map[string]tftypes.Value{
			"credential_name":   tftypes.NewValue(tftypes.String, nil),
			"credential_crn":    tftypes.NewValue(tftypes.String, nil),
			"region":            tftypes.NewValue(tftypes.String, nil),
			"availability_zone": tftypes.NewValue(tftypes.String, nil),
			"architecture":      tftypes.NewValue(tftypes.String, nil),
			"platform_variant":  tftypes.NewValue(tftypes.String, nil),
			"resource_type":     tftypes.NewValue(tftypes.String, nil),
			"vm_types":          tftypes.NewValue(tftypes.List{ElementType: tftypes.Object{AttributeTypes: vmTypesObjectType()}}, nil),
		},
	)

	req := datasource.ValidateConfigRequest{
		Config: tfsdk.Config{Raw: raw, Schema: vmTypesSchema},
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

func TestVMTypesConfigValidatorAcceptsCredentialName(t *testing.T) {
	ctx := context.TODO()
	ds := &vmTypesDataSource{}

	req := datasource.ValidateConfigRequest{
		Config: tfsdk.Config{Raw: createRawVMTypesDataSource("my-cred"), Schema: vmTypesSchema},
	}

	var diagnostics diag.Diagnostics
	for _, v := range ds.ConfigValidators(ctx) {
		resp := &datasource.ValidateConfigResponse{}
		v.ValidateDataSource(ctx, req, resp)
		diagnostics.Append(resp.Diagnostics...)
	}

	assert.False(t, diagnostics.HasError())
}

func TestReadVMTypes(t *testing.T) {
	testCases := map[string]struct {
		response       any
		errorResponse  any
		expectedError  bool
		expectedDetail string
	}{
		"OK": {
			response: &operations.ListVMTypesOK{
				Payload: &models.ListVMTypesResponse{
					VMTypes: []*models.PlatformVMTypes{
						{
							Platform: "AWS",
							VirtualMachines: []*models.VMType{
								{Name: "m5.xlarge"},
							},
						},
					},
				},
			},
			expectedError: false,
		},
		"BadRequest": {
			response: nil,
			errorResponse: &operations.ListVMTypesDefault{
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
			mockClient.On("ListVMTypesContext", mock.Anything, mock.Anything).Return(tc.response, tc.errorResponse)

			ds := &vmTypesDataSource{
				client: &cdp.Client{Environments: newMockEnvironments(mockClient)},
			}

			req := datasource.ReadRequest{
				Config: tfsdk.Config{
					Raw:    createRawVMTypesDataSource("my-cred"),
					Schema: vmTypesSchema,
				},
			}
			resp := &datasource.ReadResponse{
				State:       tfsdk.State{Schema: vmTypesSchema},
				Diagnostics: make(diag.Diagnostics, 0),
			}

			ds.Read(ctx, req, resp)

			assert.Equal(t, tc.expectedError, resp.Diagnostics.HasError())
			if tc.expectedError {
				assert.True(t, strings.Contains(resp.Diagnostics.Errors()[0].Detail(), tc.expectedDetail))
			} else {
				var state vmTypesModel
				resp.State.Get(ctx, &state)
				assert.Equal(t, 1, len(state.VMTypes))
				assert.Equal(t, "AWS", state.VMTypes[0].Platform.ValueString())
			}

			mockClient.AssertExpectations(t)
		})
	}
}
