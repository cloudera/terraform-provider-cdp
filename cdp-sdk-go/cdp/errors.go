package cdp

import (
	datahubmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
	datalakemodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	iammodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
	mlmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/models"
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

func IsMlError(err *mlmodels.Error, code string, message string) bool {
	return err.Code == code && strings.Contains(err.Message, message)
}
