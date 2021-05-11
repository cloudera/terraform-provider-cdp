package main

import (
	"context"
	"flag"
	"github.com/cloudera/terraform-provider-cdp/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"log"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "Set to true to run provider with debugger")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return provider.Provider()
		},
	}
	if debugMode {
		err := plugin.Debug(context.Background(), "terraform.cloudera.com/cloudera/cdp", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}
	plugin.Serve(opts)
}
