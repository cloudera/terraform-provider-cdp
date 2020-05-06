package cdp

import (
	"testing"
)

func TestDefaultEndpoints(t *testing.T) {
	config := NewConfig()
	if config.CdpApiEndpointUrl != defaultCdpApiEndpointUrl {
		t.Errorf("Expected default CDP endpoint to be %s", defaultCdpApiEndpointUrl)
	}
	if config.AltusApiEndpointUrl != defaultAltusApiEndpointUrl {
		t.Errorf("Expected default Altus endpoint to be %s", defaultAltusApiEndpointUrl)
	}
}
