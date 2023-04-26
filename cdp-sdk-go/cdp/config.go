package cdp

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mitchellh/go-homedir"
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

	defaultCdpApiEndpointUrl   = "https://api.us-west-1.cdp.cloudera.com/"
	defaultAltusApiEndpointUrl = "https://%sapi.us-west-1.altus.cloudera.com/"

	cdpDir          = ".cdp"
	credentialsFile = "credentials"
	configFile      = "config"
)

type Config struct {
	CdpApiEndpointUrl   string
	AltusApiEndpointUrl string
	Profile             string
	Credentials         *Credentials
	BaseApiPath         string
	ConfigFile          string
	CredentialsFile     string
	LocalEnvironment    bool
	Logger              Logger
	Context             context.Context

	properties map[string]map[string]string

	credentialsProvider CdpCredentialsProvider
}

func NewConfig() *Config {
	return &Config{
		Logger:  NewDefaultLogger(),
		Context: context.Background(),
	}
}

func (config *Config) initConfig() error {
	if config.BaseApiPath == "" {
		config.BaseApiPath = defaultBaseApiPath
	}
	properties, err := config.loadConfigFile()
	if err != nil {
		return err
	}
	config.properties = config.convertProfileMap(properties)

	// Default provider chain. By default, it first checks whether the given Config contains any, then it
	// checks the environment variables, and lastly it checks the credentials from the shared credentials file
	// under ~/.cdp/credentials.
	config.credentialsProvider = &ChainCdpCredentialsProvider{
		ProviderChain: []CdpCredentialsProvider{
			&ConfigCdpCredentialsProvider{Credentials: config.Credentials},
			&EnvCdpCredentialsProvider{},
			NewFileCdpCredentialsProvider(config.GetCdpCredentialsFile(), config.GetCdpProfile()),
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
		envVars:   []string{cdpDefaultProfileEnvVar, cdpProfileEnvVar},
		configKey: "",
		defaultFunc: func() (string, error) {
			return cdpDefaultProfile, nil
		},
	},
	"cdp_endpoint_url": {
		envVars:   []string{"CDP_ENDPOINT_URL"},
		configKey: "cdp_endpoint_url",
		defaultFunc: func() (string, error) {
			return defaultCdpApiEndpointUrl, nil
		},
	},
	"altus_endpoint_url": {
		envVars:   []string{"ENDPOINT_URL"},
		configKey: "endpoint_url",
		defaultFunc: func() (string, error) {
			return defaultAltusApiEndpointUrl, nil
		},
	},
	"local_environment": {
		envVars:   []string{"LOCAL_ENVIRONMENT"},
		configKey: "local_environment",
		defaultFunc: func() (string, error) {
			return defaultLocalEnvironment, nil
		},
	},
}

func (config *Config) GetCdpProfile() string {
	val, _ := config.getVal(config.Profile, propertySchemas["cdp_profile"])
	return val
}

func (config *Config) GetCdpApiEndpoint() string {
	val, _ := config.getVal(config.CdpApiEndpointUrl, propertySchemas["cdp_endpoint_url"])
	return val
}

func (config *Config) GetAltusApiEndpoint() string {
	val, _ := config.getVal(config.AltusApiEndpointUrl, propertySchemas["altus_endpoint_url"])
	return val
}

func (config *Config) GetCdpConfigFile() string {
	val, _ := config.getVal(config.ConfigFile, propertySchemas["cdp_config_file"])
	return val
}

func (config *Config) GetCdpCredentialsFile() string {
	val, _ := config.getVal(config.CredentialsFile, propertySchemas["cdp_credentials_file"])
	return val
}

func (config *Config) GetLocalEnvironment() bool {
	val, _ := config.getVal(strconv.FormatBool(config.LocalEnvironment), propertySchemas["local_environment"])
	boolVal, _ := strconv.ParseBool(val)
	return boolVal
}

func (config *Config) GetEndpoint(serviceName string, isAltusService bool) string {
	if isAltusService {
		return fmt.Sprintf(config.GetAltusApiEndpoint(), serviceName)
	} else {
		return config.GetCdpApiEndpoint()
	}
}

func (config *Config) GetCredentials() (*Credentials, error) {
	return config.credentialsProvider.GetCredentials()
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

func (config *Config) loadConfigFile() (map[string]map[string]string, error) {
	properties, err := RawParseConfigFile(config.GetCdpConfigFile())
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
		} else {
			// silently ignore. We do not yet support config keys that are not profiles. Can be added later.
		}
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
		scopedConfig := config.properties[config.GetCdpProfile()]
		val, ok := scopedConfig[meta.configKey]
		if ok && strings.TrimSpace(val) != "" {
			return strings.TrimSpace(val), nil
		}
	}

	return meta.defaultFunc()
}
