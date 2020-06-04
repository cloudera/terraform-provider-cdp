package utils

import (
	datahubclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client"
	datalakeclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client"
	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
)

type CdpClients struct {
	Environments *environmentsclient.Environments
	Datalake     *datalakeclient.Datalake
	Datahub      *datahubclient.Datahub
}
