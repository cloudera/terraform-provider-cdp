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

// TODO: terraform should be able to pass provider, and credentials

func envCdpCredentialsProvider() (*Credentials, error) {
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

func sharedFileCdpCredentialsProvider() (*Credentials, error) {
	path, err := defaultCdpCredentialsFile()
	if err != nil {
		return nil, err
	}

	// TODO: pass in profile
	return fileCdpCredentialsProvider(path, "")
}

func fileCdpCredentialsProvider(path string, profile string) (*Credentials, error) {
	credsConfig, err := loadCdpCredentialsFile(path)
	if err != nil {
		return nil, err
	}

	profile = getCdpProfile(profile)

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

func getCdpCredentials() (*Credentials, error) {
	// default provider chain
	providerChain := [...]func() (*Credentials, error){envCdpCredentialsProvider, sharedFileCdpCredentialsProvider}

	var err error
	for _, provider := range providerChain {
		credentials, err := provider()
		if err == nil {
			return credentials, nil
		}
	}

	return nil, err
}
