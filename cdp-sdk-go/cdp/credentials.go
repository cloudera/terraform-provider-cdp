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

	// Name of environment variable holding the users CDP access key ID.
	CdpAccessKeyIdEnvVar = "CDP_ACCESS_KEY_ID"

	// Name of environment variable holding the users CDP private key.
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

type CdpCredentialsProvider interface {
	GetCredentials() (*Credentials, error)
}

type ConfigCdpCredentialsProvider struct {
	Credentials *Credentials
}

func (p *ConfigCdpCredentialsProvider) GetCredentials() (*Credentials, error) {
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

type EnvCdpCredentialsProvider struct {
}

func (*EnvCdpCredentialsProvider) GetCredentials() (*Credentials, error) {
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

type FileCdpCredentialsProvider struct {
	path    string
	profile string
}

func NewFileCdpCredentialsProvider(path string, profile string) *FileCdpCredentialsProvider {
	return &FileCdpCredentialsProvider{path: path, profile: profile}
}

func (p *FileCdpCredentialsProvider) GetCredentials() (*Credentials, error) {
	profileMap, err := RawParseConfigFile(p.path)
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

// returns CDP credentials by using a chain of credential providers in order.
type ChainCdpCredentialsProvider struct {
	ProviderChain []CdpCredentialsProvider
}

func (p *ChainCdpCredentialsProvider) GetCredentials() (*Credentials, error) {
	for _, provider := range p.ProviderChain {
		credentials, err := provider.GetCredentials()
		if err == nil {
			return credentials, nil
		}
	}
	return nil, fmt.Errorf("no CDP Credentials from the providers in the default chain")
}

// parses and returns the given INI file as a map of maps where the first maps keys are profiles, and each profile
// maps to a set of key value pairs.
func RawParseConfigFile(path string) (map[string]map[string]string, error) {
	properties := make(map[string]map[string]string)
	cfg, err := ini.InsensitiveLoad(path)
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
