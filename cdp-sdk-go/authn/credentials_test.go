package authn

import (
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/common"
	"os"
	"testing"
)

func TestEnvCdpCredentialsProviderWithUnsetEnv(t *testing.T) {
	os.Unsetenv(cdpAccessKeyIdEnvVar)
	os.Unsetenv(cdpPrivateKeyEnvVar)
	credentials, err := (&envCdpCredentialsProvider{}).getCredentials()
	if err == nil {
		t.Errorf("Env credentials provider should have returned error with empty environment %s", credentials)
	}
}

func TestEnvCdpCredentialsProviderWithEmptyEnv(t *testing.T) {
	os.Setenv(cdpAccessKeyIdEnvVar, "")
	os.Setenv(cdpPrivateKeyEnvVar, "")
	credentials, err := (&envCdpCredentialsProvider{}).getCredentials()
	if err == nil {
		t.Errorf("Env credentials provider should have returned error with empty environment %s", credentials)
	}
}

func TestEnvCdpCredentialsProviderWithWhitespaceEnv(t *testing.T) {
	os.Setenv(cdpAccessKeyIdEnvVar, " ")
	os.Setenv(cdpPrivateKeyEnvVar, "\t")
	credentials, err := (&envCdpCredentialsProvider{}).getCredentials()
	if err == nil {
		t.Errorf("Env credentials provider should have returned error with empty environment %s", credentials)
	}
}

func TestEnvCdpCredentialsProviderWithAccessKeyId(t *testing.T) {
	os.Setenv(cdpAccessKeyIdEnvVar, "foo")
	os.Unsetenv(cdpPrivateKeyEnvVar)
	credentials, err := (&envCdpCredentialsProvider{}).getCredentials()
	if err == nil {
		t.Errorf("Env credentials provider should have returned error with just access key %s", credentials)
	}
}

func TestEnvCdpCredentialsProviderWithPrivateKey(t *testing.T) {
	os.Unsetenv(cdpAccessKeyIdEnvVar)
	os.Setenv(cdpPrivateKeyEnvVar, "bar")
	credentials, err := (&envCdpCredentialsProvider{}).getCredentials()
	if err == nil {
		t.Errorf("Env credentials provider should have returned error with just private key %s", credentials)
	}
}

func TestEnvCdpCredentialsProviderWithAccessKeyAndPrivateKey(t *testing.T) {
	os.Setenv(cdpAccessKeyIdEnvVar, "foo")
	os.Setenv(cdpPrivateKeyEnvVar, "bar")
	res, err := (&envCdpCredentialsProvider{}).getCredentials()
	if err != nil {
		t.Errorf("Env credentials provider should returned successfully")
	}
	if res.AccessKeyId != "foo" && res.PrivateKey != "bar" {
		t.Errorf("Wrong values returned as CDP credentials: accesKey: %s privateKey:%s", res.AccessKeyId, res.PrivateKey)
	}
}

func TestGetCdpProfileFromArg(t *testing.T) {
	os.Setenv(cdpProfileEnvVar, "foo")
	os.Setenv(cdpDefaultProfileEnvVar, "bar")
	common.AssertEquals(t, "baz", getCdpProfile("baz"))
}

func TestGetCdpProfileFromEnv(t *testing.T) {
	os.Setenv(cdpProfileEnvVar, "foo")
	os.Setenv(cdpDefaultProfileEnvVar, "bar")
	common.AssertEquals(t, "bar", getCdpProfile(""))
}

func TestGetCdpProfileFromEnv2(t *testing.T) {
	os.Setenv(cdpProfileEnvVar, "")
	os.Setenv(cdpDefaultProfileEnvVar, "bar")
	common.AssertEquals(t, "bar", getCdpProfile(""))
}

func TestGetCdpProfileFromDefault(t *testing.T) {
	os.Setenv(cdpProfileEnvVar, "")
	os.Setenv(cdpDefaultProfileEnvVar, "")
	common.AssertEquals(t, cdpDefaultProfile, getCdpProfile(""))
}

func TestLoadCdpCredentialsFile(t *testing.T) {
	expected := newCdpCredentialsConfig()
	expected.profiles["default"] = newCdpProfile()
	expected.profiles["default"].properties = map[string]string{
		"cdp_access_key_id": "value1",
		"cdp_private_key":   "value2",
	}
	expected.profiles["profile_foo"] = newCdpProfile()
	expected.profiles["profile_foo"].properties = map[string]string{
		"key1": "value3",
		"key2": "value4",
	}
	expected.profiles["profile_bar"] = newCdpProfile()
	expected.profiles["profile_bar"].properties = map[string]string{
		"key3":              "value5",
		"cdp_access_key_id": "value6",
		"cdp_private_key":   "value7",
	}
	expected.profiles["file_cdp_credentials_provider_profile"] = newCdpProfile()
	expected.profiles["file_cdp_credentials_provider_profile"].properties = map[string]string{
		"cdp_access_key_id": "value-from-file",
		"cdp_private_key":   "value-from-file",
	}

	path := "testdata/test-credentials"
	credsConfig, err := loadCdpCredentialsFile(path)
	if err != nil {
		t.Fatal(err)
	}

	common.AssertEquals(t, expected, credsConfig)
}

func getCredentialsFromFileProvider(t *testing.T, path string, profile string) (*Credentials, error) {
	provider, err := newFileCdpCredentialsProvider(path, profile)
	if err != nil {
		t.Fatal(err)
	}
	return provider.getCredentials()
}

func TestFileCdpCredentialsProviderDefaultProfile(t *testing.T) {
	path := "testdata/test-credentials"
	profile := ""

	cdpCredentials, err := getCredentialsFromFileProvider(t, path, profile)
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
	profile := "profile_bar"
	cdpCredentials, err := getCredentialsFromFileProvider(t, path, profile)
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
	profile := "profile_foo"
	cdpCredentials, err := getCredentialsFromFileProvider(t, path, profile)
	if err == nil || cdpCredentials != nil {
		t.Fatal(err)
	}
}

func TestFileCdpCredentialsProviderNonExistingProfile(t *testing.T) {
	path := "testdata/test-credentials"
	profile := "profile_non_existing"
	cdpCredentials, err := getCredentialsFromFileProvider(t, path, profile)
	if err == nil || cdpCredentials != nil {
		t.Fatal(err)
	}
}

func TestConfigCdpCredentialsProviderEmptyConfig(t *testing.T) {
	var provider = configCdpCredentialsProvider{config: &InternalConfig{}}
	res, err := provider.getCredentials()
	if err == nil || res != nil {
		t.Fatal(err)
	}
}

func TestConfigCdpCredentialsProviderNonEmptyConfig(t *testing.T) {
	var provider = configCdpCredentialsProvider{config: &InternalConfig{
		Credentials: &Credentials{AccessKeyId: "foo", PrivateKey: "bar"},
	}}
	res, err := provider.getCredentials()
	if err != nil {
		t.Errorf("Env credentials provider should returned successfully")
	}
	if res.AccessKeyId != "foo" && res.PrivateKey != "bar" {
		t.Errorf("Wrong values returned as CDP credentials: accesKey: %s privateKey:%s", res.AccessKeyId, res.PrivateKey)
	}
}

func TestGetCdpCredentialsNotFound(t *testing.T) {
	// empty env vars.
	os.Setenv(cdpAccessKeyIdEnvVar, "")
	os.Setenv(cdpPrivateKeyEnvVar, "")
	os.Setenv(cdpProfileEnvVar, "")
	os.Setenv(cdpDefaultProfileEnvVar, "")

	path := "testdata/test-credentials"
	profile := "profile_non_existing"

	emptyConfig := InternalConfig{Profile: profile}
	_, err := getCdpCredentials(&emptyConfig, path)
	if err == nil {
		t.Fatal(err)
	}
}

func TestGetCdpCredentials(t *testing.T) {
	// empty env vars.
	os.Setenv(cdpAccessKeyIdEnvVar, "value-from-env")
	os.Setenv(cdpPrivateKeyEnvVar, "value-from-env")
	os.Setenv(cdpProfileEnvVar, "")
	os.Setenv(cdpDefaultProfileEnvVar, "")

	path := "testdata/test-credentials"
	profile := "file_cdp_credentials_provider_profile"

	config := InternalConfig{
		Profile: profile,
		Credentials: &Credentials{
			AccessKeyId: "value-from-config",
			PrivateKey:  "value-from-config"},
	}
	credentials, err := getCdpCredentials(&config, path)
	if err != nil {
		t.Fatal(err)
	}

	if credentials.AccessKeyId != "value-from-config" && credentials.PrivateKey != "value-from-config" {
		if err != nil {
			t.Errorf("Expected the provider chain to prioritize credentials from config")
		}
	}

	credentials, err = getCdpCredentials(&InternalConfig{}, path)
	if err != nil {
		t.Fatal(err)
	}
	if credentials.AccessKeyId != "value-from-env" && credentials.PrivateKey != "value-from-env" {
		if err != nil {
			t.Errorf("Expected the provider chain to prioritize credentials from env after config")
		}
	}

	os.Setenv(cdpAccessKeyIdEnvVar, "")
	os.Setenv(cdpPrivateKeyEnvVar, "")
	credentials, err = getCdpCredentials(&InternalConfig{}, path)
	if err != nil {
		t.Fatal(err)
	}
	if credentials.AccessKeyId != "value-from-file" && credentials.PrivateKey != "value-from-file" {
		if err != nil {
			t.Errorf("Expected the provider chain to use credentials from file last")
		}
	}
}
