// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datahub

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"testing"
)

func TestForceDeleteRequestedForAws(t *testing.T) {
	tests := []struct {
		name           string
		model          *awsDatahubResourceModel
		expectedResult bool
	}{
		{
			name:           "when DestroyOptions nil",
			model:          &awsDatahubResourceModel{DestroyOptions: nil},
			expectedResult: false,
		},
		{
			name:           "when DestroyOptions not nil but ForceDeleteCluster is",
			model:          &awsDatahubResourceModel{DestroyOptions: &DestroyOptions{ForceDeleteCluster: types.BoolNull()}},
			expectedResult: false,
		},
		{
			name:           "when neither DestroyOptions or ForceDeleteCluster are nil but ForceDeleteCluster is false",
			model:          &awsDatahubResourceModel{DestroyOptions: &DestroyOptions{ForceDeleteCluster: types.BoolValue(false)}},
			expectedResult: false,
		},
		{
			name:           "when ForceDeleteCluster is true",
			model:          &awsDatahubResourceModel{DestroyOptions: &DestroyOptions{ForceDeleteCluster: types.BoolValue(true)}},
			expectedResult: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.model.forceDeleteRequested() != test.expectedResult {
				t.Errorf("Did not get the expected output! Expected: %t, got: %t", test.expectedResult, test.model.forceDeleteRequested())
			}
		})
	}
}

func TestForceDeleteRequestedForAzure(t *testing.T) {
	tests := []struct {
		name           string
		model          *azureDatahubResourceModel
		expectedResult bool
	}{
		{
			name:           "when DestroyOptions nil",
			model:          &azureDatahubResourceModel{DestroyOptions: nil},
			expectedResult: false,
		},
		{
			name:           "when DestroyOptions not nil but ForceDeleteCluster is",
			model:          &azureDatahubResourceModel{DestroyOptions: &DestroyOptions{ForceDeleteCluster: types.BoolNull()}},
			expectedResult: false,
		},
		{
			name:           "when neither DestroyOptions or ForceDeleteCluster are nil but ForceDeleteCluster is false",
			model:          &azureDatahubResourceModel{DestroyOptions: &DestroyOptions{ForceDeleteCluster: types.BoolValue(false)}},
			expectedResult: false,
		},
		{
			name:           "when ForceDeleteCluster is true",
			model:          &azureDatahubResourceModel{DestroyOptions: &DestroyOptions{ForceDeleteCluster: types.BoolValue(true)}},
			expectedResult: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.model.forceDeleteRequested() != test.expectedResult {
				t.Errorf("Did not get the expected output! Expected: %t, got: %t", test.expectedResult, test.model.forceDeleteRequested())
			}
		})
	}
}
