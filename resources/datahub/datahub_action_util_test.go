// Copyright 2026 Cloudera. All Rights Reserved.
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
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	datahubclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client/operations"
	datahubmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
)

type dummyClientTransport struct {
	submit func(operation *runtime.ClientOperation) (interface{}, error)
}

func (f *dummyClientTransport) Submit(operation *runtime.ClientOperation) (interface{}, error) {
	return f.submit(operation)
}

func TestDescribeDatahubWithDiagnosticHandle_Success(t *testing.T) {
	t.Parallel()

	clusterName := "test-datahub"
	expectedCluster := &datahubmodels.Cluster{
		ClusterName: &clusterName,
		Crn:         new("crn:cdp:datahub:us-west-1:tenant:cluster:test-datahub"),
		Status:      "AVAILABLE",
	}

	transport := &dummyClientTransport{
		submit: func(operation *runtime.ClientOperation) (interface{}, error) {
			if operation.ID != "describeCluster" {
				t.Fatalf("unexpected operation ID: %s", operation.ID)
			}

			params, ok := operation.Params.(*operations.DescribeClusterParams)
			if !ok {
				t.Fatalf("unexpected params type: %T", operation.Params)
			}

			if params.Input == nil || params.Input.ClusterName == nil {
				t.Fatal("expected describe cluster input with cluster name")
			}

			if got := *params.Input.ClusterName; got != clusterName {
				t.Fatalf("unexpected cluster name: got %q, want %q", got, clusterName)
			}

			return &operations.DescribeClusterOK{
				Payload: &datahubmodels.DescribeClusterResponse{
					Cluster: expectedCluster,
				},
			}, nil
		},
	}

	client := &cdp.Client{
		Datahub: datahubclient.New(transport, nil),
	}

	var diags diag.Diagnostics
	var state tfsdk.State

	result, err := describeDatahubWithDiagnosticHandle(
		clusterName,
		"test-id",
		context.Background(),
		client,
		&diags,
		&state,
	)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if result == nil {
		t.Fatal("expected cluster result, got nil")
		return
	}

	if result.ClusterName != expectedCluster.ClusterName {
		t.Fatalf("unexpected cluster name: got %s, want %s", *result.ClusterName, *expectedCluster.ClusterName)
		return
	}

	if diags.HasError() {
		t.Fatalf("expected no diagnostics errors, got: %+v", diags)
	}

	if len(diags) != 0 {
		t.Fatalf("expected no diagnostics entries, got: %+v", diags)
	}
}

func TestDescribeDatahubWithDiagnosticHandle_NotFound(t *testing.T) {
	t.Parallel()

	clusterName := "missing-datahub"

	notFoundErr := &operations.DescribeClusterDefault{
		Payload: &datahubmodels.Error{
			Code:    "NOT_FOUND",
			Message: "cluster not found",
		},
	}

	transport := &dummyClientTransport{
		submit: func(operation *runtime.ClientOperation) (interface{}, error) {
			if operation.ID != "describeCluster" {
				t.Fatalf("unexpected operation ID: %s", operation.ID)
			}
			return nil, notFoundErr
		},
	}

	client := &cdp.Client{
		Datahub: datahubclient.New(transport, nil),
	}

	var diags diag.Diagnostics
	var state tfsdk.State

	defer func() {
		recovered := recover()
		if recovered == nil {
			t.Fatal("expected panic caused by zero-value tfsdk.State.RemoveResource, but got none")
		}

		if diags.HasError() {
			t.Fatalf("expected warning only, got error diagnostics: %+v", diags)
		}

		if len(diags) != 1 {
			t.Fatalf("expected exactly 1 diagnostic entry, got %d", len(diags))
		}

		if diags[0].Severity() != diag.SeverityWarning {
			t.Fatalf("expected warning diagnostic, got severity: %v", diags[0].Severity())
		}

		if got := diags[0].Summary(); got != "Resource not found on provider" {
			t.Fatalf("unexpected warning summary: got %q", got)
		}

		if got := diags[0].Detail(); !strings.Contains(got, "Datahub not found, removing from state.") {
			t.Fatalf("unexpected warning detail: got %q", got)
		}
	}()

	_, _ = describeDatahubWithDiagnosticHandle(
		clusterName,
		"test-id",
		context.Background(),
		client,
		&diags,
		&state,
	)
}

func TestDescribeDatahubWithDiagnosticHandle_GenericError(t *testing.T) {
	t.Parallel()

	clusterName := "test-datahub"

	apiErr := &operations.DescribeClusterDefault{
		Payload: &datahubmodels.Error{
			Code:    "BAD_REQUEST",
			Message: "bad request from API",
		},
	}

	transport := &dummyClientTransport{
		submit: func(operation *runtime.ClientOperation) (interface{}, error) {
			if operation.ID != "describeCluster" {
				t.Fatalf("unexpected operation ID: %s", operation.ID)
			}
			return nil, apiErr
		},
	}

	client := &cdp.Client{
		Datahub: datahubclient.New(transport, nil),
	}

	var diags diag.Diagnostics
	var state tfsdk.State

	result, err := describeDatahubWithDiagnosticHandle(
		clusterName,
		"test-id",
		context.Background(),
		client,
		&diags,
		&state,
	)

	if result != nil {
		t.Fatalf("expected nil result, got: %+v", result)
	}

	if !errors.Is(err, apiErr) {
		t.Fatalf("expected returned error to be the API error, got: %v", err)
	}

	if !diags.HasError() {
		t.Fatalf("expected diagnostics error, got: %+v", diags)
	}

	if len(diags) != 1 {
		t.Fatalf("expected exactly 1 diagnostic entry, got %d", len(diags))
	}

	if diags[0].Severity() != diag.SeverityError {
		t.Fatalf("expected error diagnostic, got severity: %v", diags[0].Severity())
	}

	if got := diags[0].Summary(); got != "Read Datahub" {
		t.Fatalf("unexpected error summary: got %q", got)
	}

	if got := diags[0].Detail(); !strings.Contains(got, "Failed to read Datahub, unexpected error: bad request from API") {
		t.Fatalf("unexpected error detail: got %q", got)
	}
}
