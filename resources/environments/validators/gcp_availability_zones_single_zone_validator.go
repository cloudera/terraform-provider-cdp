// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package validators

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/cloudera/terraform-provider-cdp/utils"
)

type gcpAvailabilityZonesSingleZoneValidator struct{}

func (v gcpAvailabilityZonesSingleZoneValidator) Description(_ context.Context) string {
	return "availability_zones must contain at most one zone"
}

func (v gcpAvailabilityZonesSingleZoneValidator) MarkdownDescription(_ context.Context) string {
	return "`availability_zones` must contain at most one zone"
}

func (v gcpAvailabilityZonesSingleZoneValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	elementCount := len(req.ConfigValue.Elements())
	multipleAvailabilityZonesGiven := elementCount > 1

	validationLogJSON := v.prepareValidationLog(req, elementCount, multipleAvailabilityZonesGiven)
	utils.LogSilently(ctx, "Validating GCP availability zones configuration: ", validationLogJSON)

	if multipleAvailabilityZonesGiven {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid availability zones configuration",
			"`availability_zones` can contain only one zone until multi-zone support is added for GCP.",
		)
	}
}

func (v gcpAvailabilityZonesSingleZoneValidator) prepareValidationLog(req validator.SetRequest, elementCount int, multipleAvailabilityZonesGiven bool) []byte {
	validationLog := map[string]interface{}{
		"path": req.Path.String(),
		"values": map[string]interface{}{
			"availability_zones": req.ConfigValue.String(),
		},
		"checks": map[string]interface{}{
			"availability_zones_count":                elementCount,
			"multiple_availability_zones_given":       multipleAvailabilityZonesGiven,
			"availability_zones_single_zone_required": true,
		},
	}

	validationLogJSON, _ := json.Marshal(validationLog)
	return validationLogJSON
}

func GCPAvailabilityZonesSingleZoneValidator() validator.Set {
	return gcpAvailabilityZonesSingleZoneValidator{}
}
