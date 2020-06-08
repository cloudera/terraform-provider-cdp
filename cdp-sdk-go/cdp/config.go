package cdp

import (
	"fmt"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/authn"
)

const (
	defaultAltusApiEndpointUrl = "https://%sapi.us-west-1.altus.cloudera.com/"
	defaultCdpApiEndpointUrl   = "https://api.us-west-1.cdp.cloudera.com/"
	defaultBaseApiPath         = ""
)

type Config struct {
	// TODO: alternatively read the endpoints from ~/.cdp/config file according to the profile?
	CdpApiEndpointUrl   string
	AltusApiEndpointUrl string
	Profile             string
	Credentials         *authn.Credentials
	BaseAPIPath         string
}

func NewConfig() *Config {
	config := Config{}
	config.AltusApiEndpointUrl = defaultAltusApiEndpointUrl
	config.CdpApiEndpointUrl = defaultCdpApiEndpointUrl
	config.BaseAPIPath = defaultBaseApiPath
	return &config
}

func (config *Config) WithCdpApiEndpointUrl(cdpApiEndpointUrl string) *Config {
	config.CdpApiEndpointUrl = cdpApiEndpointUrl
	return config
}

func (config *Config) WithAltusApiEndpointUrl(altusApiEndpointUrl string) *Config {
	config.AltusApiEndpointUrl = altusApiEndpointUrl
	return config
}

func (config *Config) WithProfile(profile string) *Config {
	config.Profile = profile
	return config
}

func (config *Config) WithCredentials(credentials *authn.Credentials) *Config {
	config.Credentials = credentials
	return config
}

func (config *Config) String() string {
	return fmt.Sprintf("{CdpApiEndpointUrl: %s, AltusApiEndpointUrl: %s, Profile: %s, Credentials: %s}",
		config.CdpApiEndpointUrl, config.AltusApiEndpointUrl, config.Profile, config.Credentials.String())
}

func (config *Config) toInternalConfig() *authn.InternalConfig {
	return &authn.InternalConfig{
		CdpApiEndpointUrl:   config.CdpApiEndpointUrl,
		AltusApiEndpointUrl: config.AltusApiEndpointUrl,
		Profile:             config.Profile,
		Credentials:         config.Credentials,
		BaseAPIPath:         config.BaseAPIPath,
	}
}
