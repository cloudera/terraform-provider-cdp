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
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/common"
	"os"
	"testing"
)

func TestEnvCdpCredentialsProviderWithUnsetEnv(t *testing.T) {
	os.Unsetenv(CdpAccessKeyIdEnvVar)
	os.Unsetenv(CdpPrivateKeyEnvVar)
	credentials, err := (&EnvCdpCredentialsProvider{}).GetCredentials()
	if err == nil {
		t.Errorf("Env credentials provider should have returned error with empty environment %s", credentials)
	}
}

func TestEnvCdpCredentialsProviderWithEmptyEnv(t *testing.T) {
	os.Setenv(CdpAccessKeyIdEnvVar, "")
	os.Setenv(CdpPrivateKeyEnvVar, "")
	credentials, err := (&EnvCdpCredentialsProvider{}).GetCredentials()
	if err == nil {
		t.Errorf("Env credentials provider should have returned error with empty environment %s", credentials)
	}
}

func TestEnvCdpCredentialsProviderWithWhitespaceEnv(t *testing.T) {
	os.Setenv(CdpAccessKeyIdEnvVar, " ")
	os.Setenv(CdpPrivateKeyEnvVar, "\t")
	credentials, err := (&EnvCdpCredentialsProvider{}).GetCredentials()
	if err == nil {
		t.Errorf("Env credentials provider should have returned error with empty environment %s", credentials)
	}
}

func TestEnvCdpCredentialsProviderWithAccessKeyId(t *testing.T) {
	os.Setenv(CdpAccessKeyIdEnvVar, "foo")
	os.Unsetenv(CdpPrivateKeyEnvVar)
	credentials, err := (&EnvCdpCredentialsProvider{}).GetCredentials()
	if err == nil {
		t.Errorf("Env credentials provider should have returned error with just access key %s", credentials)
	}
}

func TestEnvCdpCredentialsProviderWithPrivateKey(t *testing.T) {
	os.Unsetenv(CdpAccessKeyIdEnvVar)
	os.Setenv(CdpPrivateKeyEnvVar, "bar")
	credentials, err := (&EnvCdpCredentialsProvider{}).GetCredentials()
	if err == nil {
		t.Errorf("Env credentials provider should have returned error with just private key %s", credentials)
	}
}

func TestEnvCdpCredentialsProviderWithAccessKeyAndPrivateKey(t *testing.T) {
	os.Setenv(CdpAccessKeyIdEnvVar, "foo")
	os.Setenv(CdpPrivateKeyEnvVar, "bar")
	res, err := (&EnvCdpCredentialsProvider{}).GetCredentials()
	if err != nil {
		t.Errorf("Env credentials provider should returned successfully")
	}
	if res.AccessKeyId != "foo" && res.PrivateKey != "bar" {
		t.Errorf("Wrong values returned as CDP credentials: accesKey: %s privateKey:%s", res.AccessKeyId, res.PrivateKey)
	}
}

func TestLoadCdpCredentialsFile(t *testing.T) {
	expected := map[string]map[string]string{
		"default": {
			"cdp_access_key_id": "value1",
			"cdp_private_key":   "value2",
		},
		"foo": {
			"key1": "value3",
			"key2": "value4",
		},
		"bar": {
			"key3":              "value5",
			"cdp_access_key_id": "value6",
			"cdp_private_key":   "value7",
		},
		"file_cdp_credentials_provider_profile": {
			"cdp_access_key_id": "value-from-file",
			"cdp_private_key":   "value-from-file",
		},
	}

	path := "testdata/test-credentials"
	credsConfig, err := RawParseConfigFile(path)
	if err != nil {
		t.Fatal(err)
	}

	common.AssertEquals(t, expected, credsConfig)
}

func GetCredentialsFromFileProvider(t *testing.T, path string, profile string) (*Credentials, error) {
	provider := NewFileCdpCredentialsProvider(path, profile)
	return provider.GetCredentials()
}

func TestFileCdpCredentialsProviderDefaultProfile(t *testing.T) {
	path := "testdata/test-credentials"
	profile := "default"

	cdpCredentials, err := GetCredentialsFromFileProvider(t, path, profile)
	if err != nil {
		t.Fatal(err)
	}
	if cdpCredentials.AccessKeyId != "value1" {
		t.Errorf("%s should have been %s for the profile: %s", cdpAccessKeyIdPropertyKey, "value1", profile)
	}
	if cdpCredentials.PrivateKey != "value2" {
		t.Errorf("%s should have been %s for the profile: %s", cdpPrivateKeyPropertyKey, "value2", profile)
	}
}

func TestFileCdpCredentialsProviderBarProfile(t *testing.T) {
	path := "testdata/test-credentials"
	profile := "bar"
	cdpCredentials, err := GetCredentialsFromFileProvider(t, path, profile)
	if err != nil {
		t.Fatal(err)
	}
	if cdpCredentials.AccessKeyId != "value6" {
		t.Errorf("%s should have been %s for the profile: %s", cdpAccessKeyIdPropertyKey, "value6", profile)
	}
	if cdpCredentials.PrivateKey != "value7" {
		t.Errorf("%s should have been %s for the profile: %s", cdpPrivateKeyPropertyKey, "value7", profile)
	}
}

func TestFileCdpCredentialsProviderFooProfile(t *testing.T) {
	path := "testdata/test-credentials"
	profile := "foo"
	cdpCredentials, err := GetCredentialsFromFileProvider(t, path, profile)
	if err == nil || cdpCredentials != nil {
		t.Fatal(err)
	}
}

func TestFileCdpCredentialsProviderNonExistingProfile(t *testing.T) {
	path := "testdata/test-credentials"
	profile := "profile_non_existing"
	cdpCredentials, err := GetCredentialsFromFileProvider(t, path, profile)
	if err == nil || cdpCredentials != nil {
		t.Fatal(err)
	}
}

func TestConfigCdpCredentialsProviderEmptyConfig(t *testing.T) {
	var provider = ConfigCdpCredentialsProvider{Credentials: &Credentials{}}
	res, err := provider.GetCredentials()
	if err == nil || res != nil {
		t.Fatal(err)
	}
}

func TestConfigCdpCredentialsProviderNonEmptyConfig(t *testing.T) {
	var provider = ConfigCdpCredentialsProvider{
		Credentials: &Credentials{AccessKeyId: "foo", PrivateKey: "bar"},
	}
	res, err := provider.GetCredentials()
	if err != nil {
		t.Errorf("Env credentials provider should returned successfully")
	}
	if res.AccessKeyId != "foo" && res.PrivateKey != "bar" {
		t.Errorf("Wrong values returned as CDP credentials: accesKey: %s privateKey:%s", res.AccessKeyId, res.PrivateKey)
	}
}
