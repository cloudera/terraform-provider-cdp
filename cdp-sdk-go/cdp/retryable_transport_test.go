// Copyright 2024 Cloudera. All Rights Reserved.
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
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
)

const (
	testUrl                      = "https://example.com"
	testGetHttpMethod            = "GET"
	testPostHttpMethod           = "POST"
	testEmptyBody                = ""
	testRetryCountEnvVariableKey = "CDP_TF_CALL_RETRY_COUNT"
	testRetryCount               = 3
)

type mockReadCloser struct {
	closed bool
}

func (m *mockReadCloser) Read(p []byte) (int, error) {
	return 0, io.EOF
}

func (m *mockReadCloser) Close() error {
	m.closed = true
	return nil
}

func TestRoundTripRetriesOnError(t *testing.T) {
	_ = os.Setenv(testRetryCountEnvVariableKey, fmt.Sprintf("%v", testRetryCount))
	defer func() {
		_ = os.Unsetenv(testRetryCountEnvVariableKey)
	}()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(testGetHttpMethod, testUrl,
		httpmock.NewErrorResponder(fmt.Errorf("network error")))

	req, _ := http.NewRequest(testGetHttpMethod, testUrl, nil)
	retryableTransport := &RetryableTransport{transport: http.DefaultTransport}

	resp, err := retryableTransport.RoundTrip(req)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
	if resp != nil {
		t.Fatalf("Expected nil response, got %v", resp)
	}
	if httpmock.GetTotalCallCount() != testRetryCount+1 {
		t.Fatalf("Expected %v retries, got %v", testRetryCount, httpmock.GetTotalCallCount()-1)
	}
}

func TestRoundTripRetriesOnRetryableStatusCode(t *testing.T) {
	_ = os.Setenv(testRetryCountEnvVariableKey, fmt.Sprintf("%v", testRetryCount))
	defer func() {
		_ = os.Unsetenv(testRetryCountEnvVariableKey)
	}()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(testGetHttpMethod, testUrl,
		httpmock.NewStringResponder(http.StatusTooManyRequests, testEmptyBody))

	req, _ := http.NewRequest(testGetHttpMethod, testUrl, nil)
	retryableTransport := &RetryableTransport{transport: http.DefaultTransport}

	resp, err := retryableTransport.RoundTrip(req)
	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}
	if resp.StatusCode != http.StatusTooManyRequests {
		t.Fatalf("Expected status code %v, got %v", http.StatusTooManyRequests, resp.StatusCode)
	}
	if httpmock.GetTotalCallCount() != testRetryCount+1 {
		t.Fatalf("Expected %v retries, got %v", testRetryCount, httpmock.GetTotalCallCount()-1)
	}
}

func TestRoundTripNoRetryOnSuccess(t *testing.T) {
	_ = os.Setenv(testRetryCountEnvVariableKey, fmt.Sprintf("%v", testRetryCount))
	defer func() {
		_ = os.Unsetenv(testRetryCountEnvVariableKey)
	}()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(testGetHttpMethod, testUrl,
		httpmock.NewStringResponder(http.StatusOK, testEmptyBody))

	req, _ := http.NewRequest(testGetHttpMethod, testUrl, nil)
	retryableTransport := &RetryableTransport{transport: http.DefaultTransport}

	resp, err := retryableTransport.RoundTrip(req)
	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %v, got %v", http.StatusOK, resp.StatusCode)
	}
	if httpmock.GetTotalCallCount() != 1 {
		t.Fatalf("Expected 1 call, got %v", httpmock.GetTotalCallCount())
	}
}

func TestRoundTripRetriesWithBody(t *testing.T) {
	_ = os.Setenv(testRetryCountEnvVariableKey, fmt.Sprintf("%v", testRetryCount))
	defer func() {
		_ = os.Unsetenv(testRetryCountEnvVariableKey)
	}()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(testPostHttpMethod, testUrl,
		httpmock.NewErrorResponder(fmt.Errorf("network error")))

	body := "test body"
	req, _ := http.NewRequest(testPostHttpMethod, testUrl, bytes.NewBufferString(body))
	retryableTransport := &RetryableTransport{transport: http.DefaultTransport}

	resp, err := retryableTransport.RoundTrip(req)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
	if resp != nil {
		t.Fatalf("Expected nil response, got %v", resp)
	}
	if httpmock.GetTotalCallCount() != testRetryCount+1 {
		t.Fatalf("Expected %v retries, got %v", testRetryCount, httpmock.GetTotalCallCount()-1)
	}
	if req.Body == nil {
		t.Fatalf("Expected request body to be retried, but it was not")
	}
}

func TestDrainBodyWithNilResponse(t *testing.T) {
	drainBody(nil)
	// No assertions needed, just ensure no panic occurs
}

func TestDrainBodyWithNilBody(t *testing.T) {
	resp := &http.Response{}
	drainBody(resp)
	// No assertions needed, just ensure no panic occurs
}

func TestDrainBodyWithNonNilBody(t *testing.T) {
	mockBody := &mockReadCloser{}
	resp := &http.Response{
		Body: mockBody,
	}
	drainBody(resp)

	if !mockBody.closed {
		t.Fatalf("Expected body to be closed, but it was not")
	}
}
