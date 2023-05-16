package main

import (
	"context"
	"flag"
	"log"

	"github.com/cloudera/terraform-provider-cdp/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary.
	// https://goreleaser.com/cookbooks/using-main.version/
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		// TODO: To be changed to Terraform Registry address
		// Address: "registry.terraform.io/cloudera/cdp",
		Address: "terraform.cloudera.com/cloudera/cdp",
		Debug:   debug,
	}
	err := providerserver.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
