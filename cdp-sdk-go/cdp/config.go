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
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/mitchellh/go-homedir"
)

// exported constants
const (
	RegionUsWest1 = "us-west-1"
	RegionEu1     = "eu-1"
	RegionAp1     = "ap-1"
	RegionUsg1    = "usg-1"
)

const (
	// These environment variables (and the ones below in the code are shared between the python CDP CLI, the Java SDK,
	// GoLang SDK and the Terraform provider for CDP. These should be treated as a compatibility surface.

	defaultBaseApiPath = ""
	// Name of system environment variable holding the name of profile to use
	// when reading the credentials file. Overrides cdpDefaultProfile.
	cdpDefaultProfileEnvVar = "CDP_DEFAULT_PROFILE"

	// Python client uses both CDP_PROFILE and CDP_DEFAULT_PROFILE
	// versus Java SDK uses CDP_DEFAULT_PROFILE
	cdpProfileEnvVar = "CDP_PROFILE"

	// Name of the profile in the users credentials file to read.
	cdpDefaultProfile = "default"

	defaultLocalEnvironment = "false"

	cdpApiEndpoint    = "https://api.%s.cdp.cloudera.com/"
	cdpApiEndpointGov = "https://api.%s.cdp.clouderagovt.com/"
	altusApiEndpoint  = "https://%%sapi.%s.altus.cloudera.com/"

	cdpDir          = ".cdp"
	credentialsFile = "credentials"
	configFile      = "config"

	// CDP defaults to using 'us-west-1'
	defaultCdpRegion = RegionUsWest1
)

type Config struct {
	CdpApiEndpointUrl     string
	AltusApiEndpointUrl   string
	Profile               string
	CdpRegion             string
	Credentials           *Credentials
	BaseApiPath           string
	ConfigFile            string
	CredentialsFile       string
	LocalEnvironment      bool
	Logger                Logger
	Context               context.Context
	UserAgent             string
	ClientApplicationName string
	Version               string

	properties map[string]map[string]string

	credentialsProvider CredentialsProvider
}

func NewConfig() *Config {
	return &Config{
		Logger:  NewDefaultLogger(),
		Context: context.Background(),
	}
}

func (config *Config) loadConfig() error {
	if config.BaseApiPath == "" {
		config.BaseApiPath = defaultBaseApiPath
	}
	properties, err := config.loadConfigFile()
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		// silently ignore to match the behavior from CDP CLI
	}
	config.properties = config.convertProfileMap(properties)

	credentialsFile, err := config.GetCdpCredentialsFile()
	if err != nil {
		return err
	}
	cdpProfile, err := config.GetCdpProfile()
	if err != nil {
		return err
	}
	// Default provider chain. By default, it first checks whether the given Config contains any, then it
	// checks the environment variables, and lastly it checks the credentials from the shared credentials file
	// under ~/.cdp/credentials.
	config.credentialsProvider = &ChainCredentialsProvider{
		ProviderChain: []CredentialsProvider{
			&StaticCredentialsProvider{Credentials: config.Credentials},
			&EnvCredentialsProvider{},
			NewFileCredentialsProvider(credentialsFile, cdpProfile),
		},
	}

	return nil
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

func (config *Config) WithCdpRegion(cdpRegion string) *Config {
	config.CdpRegion = cdpRegion
	return config
}

func (config *Config) WithCredentials(credentials *Credentials) *Config {
	config.Credentials = credentials
	return config
}

func (config *Config) WithConfigFile(configFile string) *Config {
	config.ConfigFile = configFile
	return config
}

func (config *Config) WithCredentialsFile(credentialsFile string) *Config {
	config.CredentialsFile = credentialsFile
	return config
}

func (config *Config) WithLocalEnvironment(localEnvironment bool) *Config {
	config.LocalEnvironment = localEnvironment
	return config
}

func (config *Config) WithLogger(logger Logger) *Config {
	config.Logger = logger
	return config
}

func (config *Config) WithContext(ctx context.Context) *Config {
	config.Context = ctx
	return config
}

func (config *Config) WithUserAgent(userAgent string) *Config {
	config.UserAgent = userAgent
	return config
}

func (config *Config) WithClientApplicationName(clientApplicationName string) *Config {
	config.ClientApplicationName = clientApplicationName
	return config
}

func (config *Config) WithVersion(version string) *Config {
	// TODO: this function should not be exposed to SDK end-users. When the golang SDK is taken out of
	// the terraform provider to be its own project, set this from the goreleaser config and do not let users override it.
	config.Version = version
	return config
}

func (config *Config) String() string {
	return fmt.Sprintf("{CdpApiEndpointUrl: %s, AltusApiEndpointUrl: %s, Profile: %s, Credentials: %s}",
		config.CdpApiEndpointUrl, config.AltusApiEndpointUrl, config.Profile, config.Credentials.String())
}

// represents a configurable property with keys to different lookup methods like environment variables, configuration
// keys and a default value.
type propertySchema struct {
	// The name(s) of environment variables that this property can be configured
	envVars     []string
	configKey   string
	defaultFunc func() (string, error)
}

// The environment variables and config file keys below are shared between the python CDP CLI, the Java SDK,
// GoLang SDK and the Terraform provider for CDP. These should be treated as a compatibility surface.
var propertySchemas = map[string]propertySchema{
	"cdp_config_file": {
		envVars:     []string{"CDP_CONFIG_FILE"},
		configKey:   "",
		defaultFunc: defaultCdpConfigFile,
	},
	"cdp_credentials_file": {
		envVars:     []string{"CDP_SHARED_CREDENTIALS_FILE"},
		configKey:   "",
		defaultFunc: defaultCdpCredentialsFile,
	},
	"cdp_profile": {
		envVars:     []string{cdpDefaultProfileEnvVar, cdpProfileEnvVar},
		configKey:   "",
		defaultFunc: stringSupplier(cdpDefaultProfile),
	},
	"cdp_region": {
		envVars:     []string{"CDP_REGION"},
		configKey:   "cdp_region",
		defaultFunc: stringSupplier(defaultCdpRegion),
	},
	"cdp_endpoint_url": {
		envVars:     []string{"CDP_ENDPOINT_URL"},
		configKey:   "cdp_endpoint_url",
		defaultFunc: stringSupplier(""),
	},
	"altus_endpoint_url": {
		envVars:     []string{"ENDPOINT_URL"},
		configKey:   "endpoint_url",
		defaultFunc: stringSupplier(""),
	},
	"local_environment": {
		envVars:     []string{"LOCAL_ENVIRONMENT"},
		configKey:   "local_environment",
		defaultFunc: stringSupplier(defaultLocalEnvironment),
	},
}

func (config *Config) GetCdpProfile() (string, error) {
	return config.getVal(config.Profile, propertySchemas["cdp_profile"])
}

func (config *Config) GetCdpRegion() (string, error) {
	return config.getVal(config.CdpRegion, propertySchemas["cdp_region"])
}

func (config *Config) GetCdpApiEndpoint() (string, error) {
	val, err := config.getVal(config.CdpApiEndpointUrl, propertySchemas["cdp_endpoint_url"])
	if err != nil {
		return val, err
	}
	if val == "" {
		return defaultCdpEndpoint(config)
	}
	return val, err
}

func (config *Config) GetAltusApiEndpoint() (string, error) {
	val, err := config.getVal(config.AltusApiEndpointUrl, propertySchemas["altus_endpoint_url"])
	if err != nil {
		return val, err
	}
	if val == "" {
		return defaultAltusEndpoint(config)
	}
	return val, err
}

func (config *Config) GetCdpConfigFile() (string, error) {
	return config.getVal(config.ConfigFile, propertySchemas["cdp_config_file"])
}

func (config *Config) GetCdpCredentialsFile() (string, error) {
	return config.getVal(config.CredentialsFile, propertySchemas["cdp_credentials_file"])
}

func (config *Config) GetLocalEnvironment() bool {
	val, _ := config.getVal(strconv.FormatBool(config.LocalEnvironment), propertySchemas["local_environment"])
	boolVal, _ := strconv.ParseBool(val)
	return boolVal
}

func (config *Config) GetEndpoint(serviceName string, isAltusService bool) (string, error) {
	if isAltusService {
		altusEndpoint, err := config.GetAltusApiEndpoint()
		if err != nil {
			return "", err
		}
		if strings.Contains(altusEndpoint, "%s") {
			return fmt.Sprintf(altusEndpoint, serviceName), nil
		}
		return altusEndpoint, nil
	} else {
		return config.GetCdpApiEndpoint()
	}
}

func (config *Config) GetCredentials() (*Credentials, error) {
	return config.credentialsProvider.GetCredentials()
}

func (config *Config) GetUserAgentOrDefault() string {
	if config.UserAgent == "" {
		return getDefaultUserAgent(config.Version)
	}
	return config.UserAgent
}

// getDefaultUserAgent returns a string to be set for the User-Agent header in HTTP requests. We follow the same format
// with the python based CDP CLI and Java based CDP SDK. However, there is no easy way to detect the OS version without
// running uname, so we don't do that. Can be added later if needed.
func getDefaultUserAgent(version string) string {
	return fmt.Sprintf("CDPSDK_GO/%s Go/%s %s_%s", version, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}

func defaultCdpCredentialsFile() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(home, cdpDir, credentialsFile)
	return path, nil
}

func defaultCdpConfigFile() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(home, cdpDir, configFile)
	return path, nil
}

func defaultCdpEndpoint(config *Config) (string, error) {
	cdpRegion, err := config.GetCdpRegion()
	if err != nil {
		return "", err
	}
	switch cdpRegion {
	case RegionUsg1:
		return fmt.Sprintf(cdpApiEndpointGov, cdpRegion), nil
	default:
		return fmt.Sprintf(cdpApiEndpoint, cdpRegion), nil
	}
}

func defaultAltusEndpoint(config *Config) (string, error) {
	cdpRegion, err := config.GetCdpRegion()
	if err != nil {
		return "", err
	}
	switch cdpRegion {
	case RegionUsWest1:
		return fmt.Sprintf(altusApiEndpoint, cdpRegion), nil
	default:
		return defaultCdpEndpoint(config)
	}
}

// stringSupplier returns a function that returns the input string and nil for error.
func stringSupplier(s string) func() (string, error) {
	return func() (string, error) {
		return s, nil
	}
}

func (config *Config) loadConfigFile() (map[string]map[string]string, error) {
	configFile, err := config.GetCdpConfigFile()
	if err != nil {
		return nil, err
	}
	properties, err := rawParseConfigFile(configFile)
	if err != nil {
		return nil, err
	}
	return properties, nil
}

// This function mimics how CDP CLI client behaves with respect to the profiles set in the cdp config file. The config
// is an INI formatted file, where every profile is put under a section named "[profile <name>]". The default profile
// can omit the profile prefix, and can just be put under section "[default]". An example file is:
//
//	[profile dev]
//	key1 = value1
//	key2 = value2
//
//	[profile test]
//	key1 = value1
//	key2 = value2
//
//	[default]
//	key1 = value1
//	key2 = value2
//
// We ignore the keys not under any profile, and just return a map of maps, where the first maps keys are the names of
// the profiles (without the "profile" prefix) and the values is a map of key value pairs with the mappings under the
// profile section from the file.
func (config *Config) convertProfileMap(properties map[string]map[string]string) map[string]map[string]string {
	ret := make(map[string]map[string]string)

	for profile, profileData := range properties {
		if strings.HasPrefix(profile, "profile ") {
			ret[profile[len("profile "):]] = profileData
		} else if profile == "default" {
			// Default is special, and considered a profile without having to write [profile default] as a section.
			ret[profile] = profileData
		}
		// else silently ignore. We do not yet support config keys that are not profiles. Can be added later.
	}

	return ret
}

// getVal returns the first non-empty value by doing these checks in order:
// 1. check the passed in parameter
// 2. check the environment variable(s) with the key
// 3. check the config variable with the key
// 4. value from the default function.
func (config *Config) getVal(val string, meta propertySchema) (string, error) {
	if strings.TrimSpace(val) != "" {
		return strings.TrimSpace(val), nil
	}

	for _, envVar := range meta.envVars {
		val = os.Getenv(envVar)
		if strings.TrimSpace(val) != "" {
			return strings.TrimSpace(val), nil
		}
	}

	if meta.configKey != "" {
		profile, err := config.GetCdpProfile()
		if err != nil {
			return "", err
		}
		scopedConfig := config.properties[profile]
		val, ok := scopedConfig[meta.configKey]
		if ok && strings.TrimSpace(val) != "" {
			return strings.TrimSpace(val), nil
		}
	}

	return meta.defaultFunc()
}
