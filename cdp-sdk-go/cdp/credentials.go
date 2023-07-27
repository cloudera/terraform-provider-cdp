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
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"strings"
)

const (
	// These environment variables (and the ones below in the code are shared between the python CDP CLI, the Java SDK,
	// GoLang SDK and the Terraform provider for CDP. These should be treated as a compatibility surface.

	// CdpAccessKeyIdEnvVar Name of environment variable holding the users CDP access key ID.
	CdpAccessKeyIdEnvVar = "CDP_ACCESS_KEY_ID"

	// CdpPrivateKeyEnvVar Name of environment variable holding the users CDP private key.
	CdpPrivateKeyEnvVar = "CDP_PRIVATE_KEY"

	// Name of property key holding the users CDP access key ID.
	cdpAccessKeyIdPropertyKey = "cdp_access_key_id"

	// Name of property key holding the users CDP private key.
	cdpPrivateKeyPropertyKey = "cdp_private_key"
)

type Credentials struct {
	AccessKeyId string
	PrivateKey  string
}

func (credentials *Credentials) String() string {
	return fmt.Sprintf("{AccessKeyId: %s, PrivateKey: %s}", credentials.AccessKeyId, "*****")
}

type CredentialsProvider interface {
	GetCredentials() (*Credentials, error)
}

// StaticCredentialsProvider provides credentials from pre-configured static Credentials
type StaticCredentialsProvider struct {
	Credentials *Credentials
}

func (p *StaticCredentialsProvider) GetCredentials() (*Credentials, error) {
	if p.Credentials == nil {
		return nil, fmt.Errorf("empty credentials provided for the config credentials provider")
	}
	accessKey := p.Credentials.AccessKeyId
	if strings.TrimSpace(accessKey) == "" {
		return nil, fmt.Errorf("config does not contain AccessKeyId")
	}
	privateKey := p.Credentials.PrivateKey
	if strings.TrimSpace(privateKey) == "" {
		return nil, fmt.Errorf("config does not contain PrivateKey")
	}
	return p.Credentials, nil
}

// EnvCredentialsProvider provides credentials from environment variables CDP_ACCESS_KEY_ID and CDP_PRIVATE_KEY
type EnvCredentialsProvider struct {
}

func (*EnvCredentialsProvider) GetCredentials() (*Credentials, error) {
	accessKey, ok := os.LookupEnv(CdpAccessKeyIdEnvVar)
	if !ok || strings.TrimSpace(accessKey) == "" {
		return nil, fmt.Errorf("env variable %s not defined", CdpAccessKeyIdEnvVar)
	}
	privateKey, ok := os.LookupEnv(CdpPrivateKeyEnvVar)
	if !ok || strings.TrimSpace(privateKey) == "" {
		return nil, fmt.Errorf("env variable %s not defined", CdpPrivateKeyEnvVar)
	}
	return &Credentials{AccessKeyId: accessKey, PrivateKey: privateKey}, nil
}

// FileCredentialsProvider provides credentials by reading the profile and credentials from shared configuration and
// credentials files. By default, it uses ~/.cdp/config and ~/.cdp/credentials. These files are shared between CDP CLI
// and all CDP SDKs.
type FileCredentialsProvider struct {
	path    string
	profile string
}

func NewFileCredentialsProvider(path string, profile string) *FileCredentialsProvider {
	return &FileCredentialsProvider{path: path, profile: profile}
}

func (p *FileCredentialsProvider) GetCredentials() (*Credentials, error) {
	profileMap, err := rawParseConfigFile(p.path)
	if err != nil {
		return nil, err
	}

	profile := p.profile

	profileData, ok := profileMap[profile]
	if !ok || strings.TrimSpace(profile) == "" {
		return nil, fmt.Errorf("cannot find profile %s in the credentials file", profile)
	}

	accessKeyId, ok := profileData[cdpAccessKeyIdPropertyKey]
	if !ok || strings.TrimSpace(accessKeyId) == "" {
		return nil, fmt.Errorf("cannot find %s in profile %s in the credentials file", cdpAccessKeyIdPropertyKey, profile)
	}

	privateKey, ok := profileData[cdpPrivateKeyPropertyKey]
	if !ok || strings.TrimSpace(privateKey) == "" {
		return nil, fmt.Errorf("cannot find %s in profile %s in the credentials file", cdpPrivateKeyPropertyKey, profile)
	}

	return &Credentials{AccessKeyId: accessKeyId, PrivateKey: privateKey}, nil
}

// ChainCredentialsProvider returns CDP credentials by using a chain of credential providers in order.
type ChainCredentialsProvider struct {
	ProviderChain []CredentialsProvider
}

func (p *ChainCredentialsProvider) GetCredentials() (*Credentials, error) {
	for _, provider := range p.ProviderChain {
		credentials, err := provider.GetCredentials()
		if err == nil {
			return credentials, nil
		}
	}
	return nil, fmt.Errorf("no CDP Credentials from the providers in the default chain")
}

// rawParseConfigFile parses and returns the given INI file as a map of maps where the first maps keys are profiles, and each profile
// maps to a set of key value pairs.
func rawParseConfigFile(path string) (map[string]map[string]string, error) {
	properties := make(map[string]map[string]string)
	cfg, err := ini.LoadSources(
			ini.LoadOptions{
				InsensitiveKeys: true,
			},
		path)
	if err != nil {
		return nil, err
	}
	for _, section := range cfg.Sections() {
		profile := make(map[string]string)
		properties[section.Name()] = profile
		for _, key := range section.Keys() {
			profile[key.Name()] = key.Value()
		}
	}
	return properties, nil
}
