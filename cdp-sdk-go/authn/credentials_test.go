package authn

import (
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/common"
	"os"
	"testing"
)

func TestEnvCdpCredentialsProviderWithUnsetEnv(t *testing.T) {
	os.Unsetenv(cdpAccessKeyIdEnvVar)
	os.Unsetenv(cdpPrivateKeyEnvVar)
	credentials, err := envCdpCredentialsProvider()
	if err == nil {
		t.Errorf("Env credentials provider should have returned error with empty environment %s", credentials)
	}
}

func TestEnvCdpCredentialsProviderWithEmptyEnv(t *testing.T) {
	os.Setenv(cdpAccessKeyIdEnvVar, "")
	os.Setenv(cdpPrivateKeyEnvVar, "")
	credentials, err := envCdpCredentialsProvider()
	if err == nil {
		t.Errorf("Env credentials provider should have returned error with empty environment %s", credentials)
	}
}

func TestEnvCdpCredentialsProviderWithWhitespaceEnv(t *testing.T) {
	os.Setenv(cdpAccessKeyIdEnvVar, " ")
	os.Setenv(cdpPrivateKeyEnvVar, "\t")
	credentials, err := envCdpCredentialsProvider()
	if err == nil {
		t.Errorf("Env credentials provider should have returned error with empty environment %s", credentials)
	}
}

func TestEnvCdpCredentialsProviderWithAccessKeyId(t *testing.T) {
	os.Setenv(cdpAccessKeyIdEnvVar, "foo")
	os.Unsetenv(cdpPrivateKeyEnvVar)
	credentials, err := envCdpCredentialsProvider()
	if err == nil {
		t.Errorf("Env credentials provider should have returned error with just access key %s", credentials)
	}
}

func TestEnvCdpCredentialsProviderWithPrivateKey(t *testing.T) {
	os.Unsetenv(cdpAccessKeyIdEnvVar)
	os.Setenv(cdpPrivateKeyEnvVar, "bar")
	credentials, err := envCdpCredentialsProvider()
	if err == nil {
		t.Errorf("Env credentials provider should have returned error with just private key %s", credentials)
	}
}

func TestEnvCdpCredentialsProviderWithAccessKeyAndPrivateKey(t *testing.T) {
	os.Setenv(cdpAccessKeyIdEnvVar, "foo")
	os.Setenv(cdpPrivateKeyEnvVar, "bar")
	res, err := envCdpCredentialsProvider()
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

	path := "testdata/test-credentials"
	credsConfig, err := loadCdpCredentialsFile(path)
	if err != nil {
		t.Fatal(err)
	}

	common.AssertEquals(t, expected, credsConfig)
}

func TestFileCdpCredentialsProviderDefaultProfile(t *testing.T) {
	path := "testdata/test-credentials"
	profile := ""
	cdpCredentials, err := fileCdpCredentialsProvider(path, profile)
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
	cdpCredentials, err := fileCdpCredentialsProvider(path, profile)
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
	cdpCredentials, err := fileCdpCredentialsProvider(path, profile)
	if err == nil || cdpCredentials != nil {
		t.Fatal(err)
	}
}

func TestFileCdpCredentialsProviderNonExistingProfile(t *testing.T) {
	path := "testdata/test-credentials"
	profile := "profile_non_existing"
	cdpCredentials, err := fileCdpCredentialsProvider(path, profile)
	if err == nil || cdpCredentials != nil {
		t.Fatal(err)
	}
}

// TODO: test getCdpCredentials()
