// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datalake

import (
	"context"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	datalakemodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
)

func TestResolveUnknownToNull(t *testing.T) {
	t.Run("unknown becomes null", func(t *testing.T) {
		v := types.StringUnknown()
		resolveUnknownToNull(&v)
		if !v.IsNull() {
			t.Errorf("expected null, got %s", v.String())
		}
	})
	t.Run("null stays null", func(t *testing.T) {
		v := types.StringNull()
		resolveUnknownToNull(&v)
		if !v.IsNull() {
			t.Errorf("expected null, got %s", v.String())
		}
	})
	t.Run("value unchanged", func(t *testing.T) {
		v := types.StringValue("7.2.18")
		resolveUnknownToNull(&v)
		if v.ValueString() != "7.2.18" {
			t.Errorf("expected 7.2.18, got %s", v.ValueString())
		}
	})
}

func TestUnknownRuntimeResolvedOnCreate(t *testing.T) {
	crn := "crn:datalake"
	name := "dl"
	creationDate := strfmt.NewDateTime()

	t.Run("GCP creation polling converter", func(t *testing.T) {
		model := &gcpDatalakeResourceModel{Runtime: types.StringUnknown()}
		datalakeDetailsToGcpDatalakeResourceModelFromCreation(&datalakemodels.DatalakeDetails{
			Crn: &crn, DatalakeName: &name, CreationDate: creationDate,
		}, model)
		if !model.Runtime.IsNull() {
			t.Errorf("expected null, got %s", model.Runtime.String())
		}
	})

	t.Run("GCP initial create response", func(t *testing.T) {
		model := &gcpDatalakeResourceModel{Runtime: types.StringUnknown()}
		toGcpDatalakeResourceModel(&datalakemodels.CreateGCPDatalakeResponse{
			Datalake: &datalakemodels.Datalake{
				Crn: &crn, DatalakeName: &name, CreationDate: creationDate,
			},
		}, model)
		if !model.Runtime.IsNull() {
			t.Errorf("expected null, got %s", model.Runtime.String())
		}
	})

	t.Run("AWS initial create response", func(t *testing.T) {
		model := &awsDatalakeResourceModel{Runtime: types.StringUnknown()}
		toAwsDatalakeResourceModel(&datalakemodels.CreateAWSDatalakeResponse{
			Datalake: &datalakemodels.Datalake{
				Crn: &crn, DatalakeName: &name, CreationDate: creationDate,
			},
		}, model)
		if !model.Runtime.IsNull() {
			t.Errorf("expected null, got %s", model.Runtime.String())
		}
	})

	t.Run("Azure initial create response", func(t *testing.T) {
		model := &azureDatalakeResourceModel{Runtime: types.StringUnknown()}
		toAzureDatalakeResourceModel(&datalakemodels.CreateAzureDatalakeResponse{
			Datalake: &datalakemodels.Datalake{
				Crn: &crn, DatalakeName: &name, CreationDate: creationDate,
			},
		}, model)
		if !model.Runtime.IsNull() {
			t.Errorf("expected null, got %s", model.Runtime.String())
		}
	})
}

func TestGcpReadConverterAfterImport(t *testing.T) {
	crn := "crn:datalake"
	name := "dl"
	email := "sa@project.iam.gserviceaccount.com"
	location := "gs://bucket/path"

	t.Run("nil model config is initialized from response", func(t *testing.T) {
		var diags diag.Diagnostics
		model := &gcpDatalakeResourceModel{DatalakeName: types.StringValue(name)}
		datalakeDetailsToGcpDatalakeResourceModel(context.Background(), &datalakemodels.DatalakeDetails{
			Crn: &crn, DatalakeName: &name, CreationDate: strfmt.NewDateTime(),
			GcpConfiguration:         &datalakemodels.GCPConfiguration{ServiceAccountEmail: email},
			CloudStorageBaseLocation: location,
		}, model, &diags)
		if diags.HasError() {
			t.Fatalf("unexpected diagnostics: %v", diags)
		}
		if model.CloudProviderConfiguration == nil {
			t.Fatal("expected CloudProviderConfiguration to be initialized")
		}
		if model.CloudProviderConfiguration.ServiceAccountEmail.ValueString() != email {
			t.Errorf("expected %s, got %s", email, model.CloudProviderConfiguration.ServiceAccountEmail.ValueString())
		}
		if model.CloudProviderConfiguration.StorageLocation.ValueString() != location {
			t.Errorf("expected %s, got %s", location, model.CloudProviderConfiguration.StorageLocation.ValueString())
		}
	})

	t.Run("nil response GcpConfiguration skips service account", func(t *testing.T) {
		var diags diag.Diagnostics
		model := &gcpDatalakeResourceModel{DatalakeName: types.StringValue(name)}
		datalakeDetailsToGcpDatalakeResourceModel(context.Background(), &datalakemodels.DatalakeDetails{
			Crn: &crn, DatalakeName: &name, CreationDate: strfmt.NewDateTime(),
			CloudStorageBaseLocation: location,
		}, model, &diags)
		if diags.HasError() {
			t.Fatalf("unexpected diagnostics: %v", diags)
		}
		if model.CloudProviderConfiguration == nil {
			t.Fatal("expected CloudProviderConfiguration to be initialized from storage location")
		}
		if model.CloudProviderConfiguration.StorageLocation.ValueString() != location {
			t.Errorf("expected %s, got %s", location, model.CloudProviderConfiguration.StorageLocation.ValueString())
		}
	})
}
