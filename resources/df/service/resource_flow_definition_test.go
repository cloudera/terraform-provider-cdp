// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package service

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"

	dfmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/df/models"
)

func TestLatestVersionCrn(t *testing.T) {
	tests := map[string]struct {
		versions []struct {
			Crn     string `json:"crn"`
			Version int32  `json:"version"`
		}
		expected string
	}{
		"empty": {
			versions: nil,
			expected: "",
		},
		"single": {
			versions: []struct {
				Crn     string `json:"crn"`
				Version int32  `json:"version"`
			}{
				{Crn: "crn:v1", Version: 1},
			},
			expected: "crn:v1",
		},
		"multiple_returns_latest": {
			versions: []struct {
				Crn     string `json:"crn"`
				Version int32  `json:"version"`
			}{
				{Crn: "crn:v1", Version: 1},
				{Crn: "crn:v3", Version: 3},
				{Crn: "crn:v2", Version: 2},
			},
			expected: "crn:v3",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := latestVersionCrn(tc.versions)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestLatestFlowVersionCrn(t *testing.T) {
	crn1 := "crn:v1"
	crn2 := "crn:v2"
	crn3 := "crn:v3"
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3

	tests := map[string]struct {
		versions []*dfmodels.FlowVersion
		expected string
	}{
		"empty": {
			versions: nil,
			expected: "",
		},
		"single": {
			versions: []*dfmodels.FlowVersion{
				{Crn: &crn1, Version: v1},
			},
			expected: "crn:v1",
		},
		"multiple_returns_latest": {
			versions: []*dfmodels.FlowVersion{
				{Crn: &crn1, Version: v1},
				{Crn: &crn3, Version: v3},
				{Crn: &crn2, Version: v2},
			},
			expected: "crn:v3",
		},
		"nil_crn": {
			versions: []*dfmodels.FlowVersion{
				{Crn: nil, Version: v1},
			},
			expected: "",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := latestFlowVersionCrn(tc.versions)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestComputeFlowFileSha(t *testing.T) {
	tests := map[string]struct {
		file        types.String
		expectedSha string
	}{
		"null": {
			file:        types.StringNull(),
			expectedSha: "",
		},
		"empty": {
			file:        types.StringValue(""),
			expectedSha: "",
		},
		"content": {
			file:        types.StringValue(`{"flowContents":{}}`),
			expectedSha: fmt.Sprintf("%x", sha256.Sum256([]byte(`{"flowContents":{}}`))),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			m := &flowDefinitionModel{File: tc.file}
			computeFlowFileSha(m)
			assert.Equal(t, tc.expectedSha, m.FileSha.ValueString())
		})
	}
}

func TestComputeFlowFileSha_DifferentContent(t *testing.T) {
	m1 := &flowDefinitionModel{File: types.StringValue(`{"version":1}`)}
	m2 := &flowDefinitionModel{File: types.StringValue(`{"version":2}`)}

	computeFlowFileSha(m1)
	computeFlowFileSha(m2)

	assert.NotEqual(t, m1.FileSha.ValueString(), m2.FileSha.ValueString())
}
