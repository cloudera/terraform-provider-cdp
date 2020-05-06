package cdp

import (
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/authn"
	environmentclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
)

type Client struct {
	Environments *environmentclient.Environments
	IAM  *iam.IAM
}

func NewEnvironmentsClient(config *Config) *environmentclient.Environments {
	transport, err := authn.GetAPIKeyAuthTransport(config.CdpApiEndpointUrl, "")
	if err != nil {
		panic(err)
	}
	environments := environmentclient.New(transport, nil)

	return environments
}