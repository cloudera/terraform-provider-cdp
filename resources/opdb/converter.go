// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package opdb

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	opdbmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/models"
)

func fromModelToDatabaseRequest(model databaseResourceModel, ctx context.Context) *opdbmodels.CreateDatabaseRequest {
	tflog.Debug(ctx, "Conversion from databaseResourceModel to CreateDatabaseRequest started.")
	req := opdbmodels.CreateDatabaseRequest{}
	req.DatabaseName = model.DatabaseName.ValueStringPointer()
	req.EnvironmentName = model.Environment.ValueStringPointer()
	req.ScaleType = opdbmodels.ScaleType(model.ScaleType.ValueString())
	req.StorageType = opdbmodels.StorageType(model.StorageType.ValueString())
	req.DisableExternalDB = model.DisableExternalDB.ValueBool()

	req.DisableMultiAz = model.DisableMultiAz.ValueBool()
	req.SubnetID = model.SubnetID.ValueString()

	req.JavaVersion = int64To32(model.JavaVersion)
	req.NumEdgeNodes = int64To32(model.NumEdgeNodes)

	tflog.Debug(ctx, fmt.Sprintf("Conversion from databaseResourceModel to CreateDatabaseRequest has finished with request: %+v.", req))
	return &req
}

func int64To32(in types.Int64) int32 {
	n64 := in.ValueInt64()
	return int32(n64)
}
