// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package cdp

import (
	datahubmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
	datalakemodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	iammodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
	mlmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/models"
	opdbmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/models"
	"strings"
)

// These functions should be generated in the appropriate gen package in an errors module.

func IsIamError(err *iammodels.Error, code string, message string) bool {
	return err.Code == code && strings.Contains(err.Message, message)
}

func IsEnvironmentsError(err *environmentsmodels.Error, code string, message string) bool {
	return err.Code == code && strings.Contains(err.Message, message)
}

func IsDatalakeError(err *datalakemodels.Error, code string, message string) bool {
	return err.Code == code && strings.Contains(err.Message, message)
}

func IsDatahubError(err *datahubmodels.Error, code string, message string) bool {
	return err.Code == code && strings.Contains(err.Message, message)
}

func IsDatabaseError(err *opdbmodels.Error, code string, message string) bool {
	return err.Code == code && strings.Contains(err.Message, message)
}

func IsMlError(err *mlmodels.Error, code string, message string) bool {
	return err.Code == code && strings.Contains(err.Message, message)
}
