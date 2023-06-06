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
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"net/http"
	"regexp"
	"testing"
)

var (
	testCredentials = Credentials{
		AccessKeyId: "auth_test",
		PrivateKey:  "37yMdtdkJANPn62X5KDKKI3iv5hbAAKvqxHdgIj22bo=",
	}
	testEndpoint = "https://api.us-west-1.cdp.cloudera.com"
)

type mockRoundTripper struct {
	req *http.Request
}

func (t *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	t.req = req
	return nil, errors.NotImplemented("mock")
}

func getTestOperation() *runtime.ClientOperation {
	return &runtime.ClientOperation{
		ID:                 "testApi",
		Method:             "POST",
		PathPattern:        "/api/v1/someservice/testApi",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Reader: runtime.ClientResponseReaderFunc(func(runtime.ClientResponse, runtime.Consumer) (interface{}, error) {
			return nil, nil
		}),

		Params: runtime.ClientRequestWriterFunc(func(req runtime.ClientRequest, reg strfmt.Registry) error {
			return req.SetBodyParam("test body")
		}),
		Context: context.Background(),
	}
}

func getTestConfig() *Config {
	config := NewConfig()
	config.WithContext(context.Background())
	config.WithCredentials(&testCredentials)
	config.WithLogger(NewDefaultLogger())
	config.WithVersion("0.1.0")
	config.initConfig()
	return config
}

func TestDefaultUserAgent(t *testing.T) {
	config := getTestConfig()

	mockRoundTripper := mockRoundTripper{}
	transport, err := buildClientTransport(config, testEndpoint, &mockRoundTripper)
	if err != nil {
		t.Fatalf("Failed to get transport: %v", err)
	}

	_, err = transport.Submit(getTestOperation())

	userAgent := mockRoundTripper.req.UserAgent()

	r, _ := regexp.Compile(`^CDPSDK_GO/.+ Go/.+ .+_.+$`)
	if !r.MatchString(userAgent) {
		t.Fatalf("Failed to match the User-Agent regex: %v", userAgent)
	}

	if err == nil {
		t.Fatalf("Should have failed with err from mock.")
	}
}

func TestCustomUserAgent(t *testing.T) {
	config := getTestConfig()
	config.WithUserAgent("test-user-agent")

	mockRoundTripper := mockRoundTripper{}
	transport, err := buildClientTransport(config, testEndpoint, &mockRoundTripper)
	if err != nil {
		t.Fatalf("Failed to get transport: %v", err)
	}

	_, err = transport.Submit(getTestOperation())

	userAgent := mockRoundTripper.req.UserAgent()

	if userAgent != "test-user-agent" {
		t.Fatalf("Failed to match the User-Agent: %v", userAgent)
	}

	if err == nil {
		t.Fatalf("Should have failed with err from mock.")
	}
}

func TestDefaultClientApplicationName(t *testing.T) {
	config := getTestConfig()

	mockRoundTripper := mockRoundTripper{}
	transport, err := buildClientTransport(config, testEndpoint, &mockRoundTripper)
	if err != nil {
		t.Fatalf("Failed to get transport: %v", err)
	}

	_, err = transport.Submit(getTestOperation())

	clientAppName := mockRoundTripper.req.Header.Get("x-altus-client-app")

	if clientAppName != "" {
		t.Fatalf("x-altus-client-app header should not have been set: %v", clientAppName)
	}

	if err == nil {
		t.Fatalf("Should have failed with err from mock.")
	}
}

func TestCustomClientApplicationName(t *testing.T) {
	config := getTestConfig()
	config.WithClientApplicationName("test-client-application-name")

	mockRoundTripper := mockRoundTripper{}
	transport, err := buildClientTransport(config, testEndpoint, &mockRoundTripper)
	if err != nil {
		t.Fatalf("Failed to get transport: %v", err)
	}

	_, err = transport.Submit(getTestOperation())

	clientAppName := mockRoundTripper.req.Header.Get("x-altus-client-app")

	if clientAppName != "test-client-application-name" {
		t.Fatalf("x-altus-client-app header does not match: %v", clientAppName)
	}

	if err == nil {
		t.Fatalf("Should have failed with err from mock.")
	}
}
