package cdp

import (
	datahubmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
	datalakemodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"strings"
)

// These functions should be generated in the appropriate gen package in an errors module.

func IsEnvironmentsError(err *environmentsmodels.Error, code string, message string) bool {
	return err.Code == code && strings.Contains(err.Message, message)
}

func IsDatalakeError(err *datalakemodels.Error, code string, message string) bool {
	return err.Code == code && strings.Contains(err.Message, message)
}

func IsDatahubError(err *datahubmodels.Error, code string, message string) bool {
	return err.Code == code && strings.Contains(err.Message, message)
}
