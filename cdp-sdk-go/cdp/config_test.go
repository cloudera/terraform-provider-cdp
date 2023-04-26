package cdp

import (
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/common"
	"os"
	"testing"
)

func TestGetCdpProfileFromConfig(t *testing.T) {
	os.Setenv(cdpProfileEnvVar, "foo")
	os.Setenv(cdpDefaultProfileEnvVar, "bar")
	config := Config{
		Profile: "baz",
	}
	common.AssertEquals(t, "baz", config.GetCdpProfile())
}

func TestGetCdpProfileFromEnv(t *testing.T) {
	os.Setenv(cdpProfileEnvVar, "foo")
	os.Setenv(cdpDefaultProfileEnvVar, "bar")
	config := Config{
		Profile: "",
	}
	common.AssertEquals(t, "bar", config.GetCdpProfile())
}

func TestGetCdpProfileFromEnv2(t *testing.T) {
	os.Setenv(cdpProfileEnvVar, "")
	os.Setenv(cdpDefaultProfileEnvVar, "bar")
	config := Config{}
	common.AssertEquals(t, "bar", config.GetCdpProfile())
}

func TestGetCdpProfileFromDefault(t *testing.T) {
	os.Setenv(cdpProfileEnvVar, "")
	os.Setenv(cdpDefaultProfileEnvVar, "")
	config := Config{}
	common.AssertEquals(t, cdpDefaultProfile, config.GetCdpProfile())
}

func TestGetCredentialsNotFound(t *testing.T) {
	// empty env vars.
	os.Setenv(CdpAccessKeyIdEnvVar, "")
	os.Setenv(CdpPrivateKeyEnvVar, "")
	os.Setenv(cdpProfileEnvVar, "")
	os.Setenv(cdpDefaultProfileEnvVar, "")

	path := "testdata/test-credentials"
	profile := "profile_non_existing"

	emptyConfig := Config{
		Profile:         profile,
		CredentialsFile: path,
	}
	emptyConfig.initConfig()
	_, err := emptyConfig.GetCredentials()
	if err == nil {
		t.Fatal(err)
	}
}

func TestGetCdpCredentials(t *testing.T) {
	// empty env vars.
	os.Setenv(CdpAccessKeyIdEnvVar, "value-from-env")
	os.Setenv(CdpPrivateKeyEnvVar, "value-from-env")
	os.Setenv(cdpProfileEnvVar, "")
	os.Setenv(cdpDefaultProfileEnvVar, "")

	path := "testdata/test-credentials"
	profile := "file_cdp_credentials_provider_profile"

	config := Config{
		Profile:         profile,
		CredentialsFile: path,
		Credentials: &Credentials{
			AccessKeyId: "value-from-config",
			PrivateKey:  "value-from-config"},
	}
	config.initConfig()
	credentials, err := config.GetCredentials()
	if err != nil {
		t.Fatal(err)
	}

	if credentials.AccessKeyId != "value-from-config" && credentials.PrivateKey != "value-from-config" {
		if err != nil {
			t.Errorf("Expected the provider chain to prioritize credentials from config")
		}
	}

	config.Credentials = nil
	credentials, err = config.GetCredentials()
	if err != nil {
		t.Fatal(err)
	}
	if credentials.AccessKeyId != "value-from-env" && credentials.PrivateKey != "value-from-env" {
		if err != nil {
			t.Errorf("Expected the provider chain to prioritize credentials from env after config")
		}
	}

	os.Setenv(CdpAccessKeyIdEnvVar, "")
	os.Setenv(CdpPrivateKeyEnvVar, "")
	credentials, err = config.GetCredentials()
	if err != nil {
		t.Fatal(err)
	}
	if credentials.AccessKeyId != "value-from-file" && credentials.PrivateKey != "value-from-file" {
		if err != nil {
			t.Errorf("Expected the provider chain to use credentials from file last")
		}
	}
}

func TestLoadConfigFileNotFound(t *testing.T) {
	os.Setenv("CDP_CONFIG_FILE", "")
	config := Config{
		ConfigFile: "testdata/not-existing-file",
	}
	err := config.initConfig()
	if err == nil {
		t.Errorf("Expected error from initConfig()")
	}
}

func TestLoadConfigFile(t *testing.T) {
	os.Setenv("CDP_CONFIG_FILE", "")
	config := Config{
		ConfigFile: "testdata/test-config",
	}
	err := config.initConfig()
	if err != nil {
		t.Fatal(err)
	}
	common.AssertEquals(t, "value1", config.GetCdpApiEndpoint())
}

func TestLoadConfigFileFromEnv(t *testing.T) {
	os.Setenv("CDP_CONFIG_FILE", "testdata/test-config")
	config := Config{
		ConfigFile: "testdata/test-config",
	}
	err := config.initConfig()
	if err != nil {
		t.Fatal(err)
	}
	common.AssertEquals(t, "value1", config.GetCdpApiEndpoint())
}

func TestGetCdpApiEndpointWithProfile(t *testing.T) {
	config := Config{
		Profile:    "foo",
		ConfigFile: "testdata/test-config",
	}
	err := config.initConfig()
	if err != nil {
		t.Fatal(err)
	}
	common.AssertEquals(t, "value3", config.GetCdpApiEndpoint())
}

func TestDefaultEndpoints(t *testing.T) {
	config := Config{
		ConfigFile: "testdata/test-config",
		Profile:    "non-existing",
	}
	err := config.initConfig()
	if err != nil {
		t.Fatal(err)
	}
	if config.GetCdpApiEndpoint() != defaultCdpApiEndpointUrl {
		t.Errorf("Expected default CDP endpoint to be %s", defaultCdpApiEndpointUrl)
	}
	if config.GetAltusApiEndpoint() != defaultAltusApiEndpointUrl {
		t.Errorf("Expected default Altus endpoint to be %s", defaultAltusApiEndpointUrl)
	}
}
