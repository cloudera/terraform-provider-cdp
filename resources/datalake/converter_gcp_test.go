// Copyright 2023 Cloudera. All Rights Reserved.
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
	"github.com/go-openapi/strfmt"
	"testing"

	datalakemodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
)

func TestToGcpDatalakeResourceModel(t *testing.T) {
	dlCrn := "datalakeCrn"
	name := "dlName"
	creationDate := strfmt.NewDateTime()
	input := &datalakemodels.CreateGCPDatalakeResponse{
		Datalake: &datalakemodels.Datalake{
			CertificateExpirationState: "someState",
			CreationDate:               creationDate,
			Crn:                        &dlCrn,
			DatalakeName:               &name,
			EnableRangerRaz:            false,
			EnvironmentCrn:             "envCrn",
			MultiAz:                    false,
			Status:                     "some cool status",
			StatusReason:               "some more cole reason",
		},
	}
	toModify := &gcpDatalakeResourceModel{}
	toGcpDatalakeResourceModel(input, toModify)
	if toModify.Crn.ValueString() != dlCrn {
		t.Errorf("The CRN (%s) is not the expected: %s", toModify.Crn.ValueString(), dlCrn)
	}
	if toModify.DatalakeName.ValueString() != name {
		t.Errorf("The Datalake name (%s) is not the expected: %s", toModify.DatalakeName.ValueString(), name)
	}
	if toModify.CreationDate.ValueString() != creationDate.String() {
		t.Errorf("The creation date (%s) is not the expected: %s", toModify.CreationDate.ValueString(), creationDate.String())
	}
	if toModify.EnableRangerRaz.ValueBool() != input.Datalake.EnableRangerRaz {
		t.Errorf("The EnableRangerRaz (%t) is not the expected: %t", toModify.EnableRangerRaz.ValueBool(), input.Datalake.EnableRangerRaz)
	}
	if toModify.EnvironmentCrn.ValueString() != input.Datalake.EnvironmentCrn {
		t.Errorf("The CRN (%s) is not the expected: %s", toModify.EnvironmentCrn.ValueString(), input.Datalake.EnvironmentCrn)
	}
	if toModify.MultiAz.ValueBool() != input.Datalake.MultiAz {
		t.Errorf("The MultiAz (%t) is not the expected: %t", toModify.MultiAz.ValueBool(), input.Datalake.MultiAz)
	}
	if toModify.Status.ValueString() != input.Datalake.Status {
		t.Errorf("The Status (%s) is not the expected: %s", toModify.Status.ValueString(), input.Datalake.Status)
	}
	if toModify.StatusReason.ValueString() != input.Datalake.StatusReason {
		t.Errorf("The StatusReason (%s) is not the expected: %s", toModify.StatusReason.ValueString(), input.Datalake.StatusReason)
	}
}
