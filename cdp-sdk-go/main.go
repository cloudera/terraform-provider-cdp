package main

import (
	"fmt"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/authn"
	environmentclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

func main() {
	cdpEndpoint := "https://api.us-west-1.cdp.cloudera.com"
	transport, err := authn.GetAPIKeyAuthTransport(cdpEndpoint, "")
	if err != nil {
		panic(err)
	}
	environments := environmentclient.New(transport, nil)

	params := operations.NewListCredentialsParams()
	params.WithInput(&environmentmodels.ListCredentialsRequest{})
	resp, err := environments.Operations.ListCredentials(params)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Error())
	fmt.Println(resp.Payload.Credentials)
}
