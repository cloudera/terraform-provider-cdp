package authn

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
	"strings"
)

const (
	// Name of environment variable holding the users CDP access key ID.
	cdpAccessKeyIdEnvVar = "CDP_ACCESS_KEY_ID"

	// Name of environment variable holding the users CDP private key.
	cdpPrivateKeyEnvVar = "CDP_PRIVATE_KEY"

	// Name of property key holding the users CDP access key ID.
	cdpAccessKeyIdPropertyKey = "cdp_access_key_id"

	// Name of property key holding the users CDP private key.
	cdpPrivateKeyPropertyKey = "cdp_private_key"

	// Name of system environment variable holding the name of profile to use
	// when reading the credentials file. Overrides cdpDefaultProfile.
	cdpDefaultProfileEnvVar = "CDP_DEFAULT_PROFILE"

	// Python client uses both CDP_PROFILE and CDP_DEFAULT_PROFILE
	// versus Java SDK uses CDP_DEFAULT_PROFILE
	cdpProfileEnvVar = "CDP_PROFILE"

	// Name of the profile in the users credentials file to read.
	cdpDefaultProfile = "default"

	cdpDir          = ".cdp"
	credentialsFile = "credentials"
)

type Credentials struct {
	AccessKeyId string
	PrivateKey  string
}

// packages in go cannot have cyclic dependencies. Use this for now.
type InternalConfig struct {
	CdpApiEndpointUrl   string
	AltusApiEndpointUrl string
	Profile             string
	Credentials         *Credentials
	BaseAPIPath         string
}

func (credentials *Credentials) String() string {
	return fmt.Sprintf("{AccessKeyId: %s, PrivateKey: %s}", credentials.AccessKeyId, "*****")
}

type cdpProfile struct {
	properties map[string]string
}

func newCdpProfile() *cdpProfile {
	cdpProfile := new(cdpProfile)
	cdpProfile.properties = make(map[string]string)
	return cdpProfile
}

type cdpCredentialsConfig struct {
	profiles map[string]*cdpProfile
}

func newCdpCredentialsConfig() *cdpCredentialsConfig {
	credsConfig := new(cdpCredentialsConfig)
	credsConfig.profiles = make(map[string]*cdpProfile)
	return credsConfig
}

func getCdpProfile(cdpProfile string) string {
	if strings.TrimSpace(cdpProfile) != "" {
		return strings.TrimSpace(cdpProfile)
	}
	val, ok := os.LookupEnv(cdpDefaultProfileEnvVar)
	if ok && strings.TrimSpace(val) != "" {
		return strings.TrimSpace(val)
	}
	val, ok = os.LookupEnv(cdpProfileEnvVar)
	if ok && strings.TrimSpace(val) != "" {
		return strings.TrimSpace(val)
	}
	return cdpDefaultProfile
}

type cdpCredentialsProvider interface {
	getCredentials() (*Credentials, error)
}

type configCdpCredentialsProvider struct {
	config *InternalConfig
}

func (p *configCdpCredentialsProvider) getCredentials() (*Credentials, error) {
	if p.config == nil {
		return nil, fmt.Errorf("empty config provided for the config credentials provider")
	}
	if p.config.Credentials == nil {
		return nil, fmt.Errorf("empty credentials provided for the config credentials provider")
	}
	accessKey := p.config.Credentials.AccessKeyId
	if strings.TrimSpace(accessKey) == "" {
		return nil, fmt.Errorf("config does not contain AccessKeyId")
	}
	privateKey := p.config.Credentials.PrivateKey
	if strings.TrimSpace(privateKey) == "" {
		return nil, fmt.Errorf("config does not contain PrivateKey")
	}
	return p.config.Credentials, nil
}

type envCdpCredentialsProvider struct {
}

func (*envCdpCredentialsProvider) getCredentials() (*Credentials, error) {
	accessKey, ok := os.LookupEnv(cdpAccessKeyIdEnvVar)
	if !ok || strings.TrimSpace(accessKey) == "" {
		return nil, fmt.Errorf("env variable %s not defined", cdpAccessKeyIdEnvVar)
	}
	privateKey, ok := os.LookupEnv(cdpPrivateKeyEnvVar)
	if !ok || strings.TrimSpace(privateKey) == "" {
		return nil, fmt.Errorf("env variable %s not defined", cdpPrivateKeyEnvVar)
	}
	return &Credentials{AccessKeyId: accessKey, PrivateKey: privateKey}, nil
}

type fileCdpCredentialsProvider struct {
	path    string
	profile string
}

func newFileCdpCredentialsProvider(path string, profile string) (*fileCdpCredentialsProvider, error) {
	var err error
	if path == "" {
		path, err = defaultCdpCredentialsFile()
		if err != nil {
			return nil, err
		}

	}
	return &fileCdpCredentialsProvider{path: path, profile: profile}, nil
}

func (p *fileCdpCredentialsProvider) getCredentials() (*Credentials, error) {
	credsConfig, err := loadCdpCredentialsFile(p.path)
	if err != nil {
		return nil, err
	}

	profile := getCdpProfile(p.profile)
	// fmt.Printf("CDP Profile to use: %s\n", profile) // TODO: switch to proper logging

	profileData, ok := credsConfig.profiles[profile]
	if !ok || strings.TrimSpace(profile) == "" {
		return nil, fmt.Errorf("cannot find profile %s in the credentials file", profile)
	}

	accessKeyId, ok := profileData.properties[cdpAccessKeyIdPropertyKey]
	if !ok || strings.TrimSpace(accessKeyId) == "" {
		return nil, fmt.Errorf("cannot find %s in profile %s in the credentials file", cdpAccessKeyIdPropertyKey, profile)
	}

	privateKey, ok := profileData.properties[cdpPrivateKeyPropertyKey]
	if !ok || strings.TrimSpace(privateKey) == "" {
		return nil, fmt.Errorf("cannot find %s in profile %s in the credentials file", cdpPrivateKeyPropertyKey, profile)
	}

	return &Credentials{AccessKeyId: accessKeyId, PrivateKey: privateKey}, nil
}

func defaultCdpCredentialsFile() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(home, cdpDir, credentialsFile)
	return path, nil
}

func loadCdpCredentialsFile(path string) (*cdpCredentialsConfig, error) {
	credsConfig := newCdpCredentialsConfig()
	cfg, err := ini.InsensitiveLoad(path)
	if err != nil {
		return nil, err
	}
	for _, section := range cfg.Sections() {
		profile := newCdpProfile()
		credsConfig.profiles[section.Name()] = profile
		for _, key := range section.Keys() {
			profile.properties[key.Name()] = key.Value()
		}
	}
	return credsConfig, nil
}

// getCdpCredentials returns CDP credentials by using a chain of credential providers in order. By default, it first
// checks whether the given InternalConfig contains any, then it checks the environment variables, and lastly it checks
// the credentials from the shared credentials file under ~/.cdp/credentials.
func getCdpCredentials(config *InternalConfig, credentialsFile string) (*Credentials, error) {
	fileCdpCredentialsProvider, err := newFileCdpCredentialsProvider(credentialsFile, config.Profile)
	if err != nil {
		return nil, err
	}

	// default provider chain
	providerChain := [...]cdpCredentialsProvider{
		&configCdpCredentialsProvider{config: config},
		&envCdpCredentialsProvider{},
		fileCdpCredentialsProvider,
	}

	for _, provider := range providerChain {
		credentials, err := provider.getCredentials()
		if err == nil {
			return credentials, nil
		} else {
			// TODO: switch to proper logging
			//fmt.Printf("error: %v\n", err)
		}
	}

	return nil, fmt.Errorf("no CDP Credentials from the providers in the default chain")
}
