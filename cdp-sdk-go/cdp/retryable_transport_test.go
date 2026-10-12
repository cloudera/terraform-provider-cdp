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
	"context"
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
	closed  bool
	readErr error
}

func (m *mockReadCloser) Read(p []byte) (int, error) {
	if m.readErr != nil {
		return 0, m.readErr
	}
	return 0, io.EOF
}

func (m *mockReadCloser) Close() error {
	m.closed = true
	return nil
}

func setupRetryTest(t *testing.T) *RetryableTransport {
	t.Helper()
	_ = os.Setenv(testRetryCountEnvVariableKey, fmt.Sprintf("%v", testRetryCount))
	t.Cleanup(func() { _ = os.Unsetenv(testRetryCountEnvVariableKey) })
	httpmock.Activate()
	t.Cleanup(httpmock.DeactivateAndReset)
	return &RetryableTransport{transport: http.DefaultTransport}
}

func registerErrorResponder() {
	httpmock.RegisterResponder(testGetHttpMethod, testUrl,
		httpmock.NewErrorResponder(fmt.Errorf("network error")))
}

func registerSuccessResponder() {
	httpmock.RegisterResponder(testGetHttpMethod, testUrl,
		httpmock.NewStringResponder(http.StatusOK, testEmptyBody))
}

func newTestGetRequest(t *testing.T) *http.Request {
	t.Helper()
	req, _ := http.NewRequest(testGetHttpMethod, testUrl, nil)
	return req
}

func assertCallCount(t *testing.T, expected int, msg string) {
	t.Helper()
	if httpmock.GetTotalCallCount() != expected {
		t.Fatalf("%s: expected %v calls, got %v", msg, expected, httpmock.GetTotalCallCount())
	}
}

func TestShouldRetryWithNilErrorAndNilResponse(t *testing.T) {
	if shouldRetry(nil, nil) {
		t.Fatalf("Expected false for nil error and nil response")
	}
}

func TestRoundTripRetriesOnError(t *testing.T) {
	transport := setupRetryTest(t)
	registerErrorResponder()

	resp, err := transport.RoundTrip(newTestGetRequest(t))
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
	if resp != nil {
		t.Fatalf("Expected nil response, got %v", resp)
	}
	assertCallCount(t, testRetryCount+1, "retries on error")
}

func TestRoundTripRetriesOnRetryableStatusCode(t *testing.T) {
	transport := setupRetryTest(t)
	httpmock.RegisterResponder(testGetHttpMethod, testUrl,
		httpmock.NewStringResponder(http.StatusTooManyRequests, testEmptyBody))

	resp, err := transport.RoundTrip(newTestGetRequest(t))
	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}
	if resp.StatusCode != http.StatusTooManyRequests {
		t.Fatalf("Expected status code %v, got %v", http.StatusTooManyRequests, resp.StatusCode)
	}
	assertCallCount(t, testRetryCount+1, "retries on 429")
}

func TestRoundTripNoRetryOnSuccess(t *testing.T) {
	transport := setupRetryTest(t)
	registerSuccessResponder()

	resp, err := transport.RoundTrip(newTestGetRequest(t))
	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %v, got %v", http.StatusOK, resp.StatusCode)
	}
	assertCallCount(t, 1, "no retry on success")
}

func TestRoundTripRetriesWithBody(t *testing.T) {
	transport := setupRetryTest(t)
	httpmock.RegisterResponder(testPostHttpMethod, testUrl,
		httpmock.NewErrorResponder(fmt.Errorf("network error")))

	body := "test body"
	req, _ := http.NewRequest(testPostHttpMethod, testUrl, bytes.NewBufferString(body))

	resp, err := transport.RoundTrip(req)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
	if resp != nil {
		t.Fatalf("Expected nil response, got %v", resp)
	}
	assertCallCount(t, testRetryCount+1, "retries with body")
	if req.Body == nil {
		t.Fatalf("Expected request body to be retried, but it was not")
	}
}

func TestDrainBodyWithNilResponse(t *testing.T) {
	drainBody(nil)
}

func TestDrainBodyWithNilBody(t *testing.T) {
	resp := &http.Response{}
	drainBody(resp)
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

func TestDrainBodyWithReadError(t *testing.T) {
	mockBody := &mockReadCloser{readErr: fmt.Errorf("read error")}
	resp := &http.Response{Body: mockBody}
	drainBody(resp)

	if !mockBody.closed {
		t.Fatalf("Expected body to be closed even after read error")
	}
}

func TestRoundTripRecoversOnIntermittentFailure(t *testing.T) {
	transport := setupRetryTest(t)

	callCount := 0
	httpmock.RegisterResponder(testGetHttpMethod, testUrl,
		func(req *http.Request) (*http.Response, error) {
			callCount++
			if callCount <= 2 {
				return nil, fmt.Errorf("network error")
			}
			return httpmock.NewStringResponse(http.StatusOK, ""), nil
		})

	resp, err := transport.RoundTrip(newTestGetRequest(t))
	if err != nil {
		t.Fatalf("Expected nil error after recovery, got %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status %v, got %v", http.StatusOK, resp.StatusCode)
	}
	assertCallCount(t, 3, "intermittent failure recovery")
}

func TestRoundTripRetryCounterPersistsAcrossCalls(t *testing.T) {
	transport := setupRetryTest(t)
	registerErrorResponder()

	_, _ = transport.RoundTrip(newTestGetRequest(t))
	assertCallCount(t, testRetryCount+1, "call 1")

	httpmock.Reset()
	registerErrorResponder()

	_, _ = transport.RoundTrip(newTestGetRequest(t))
	assertCallCount(t, 1, "call 2 (budget exhausted)")
}

func TestRoundTripBackoffResetsAfterSuccessfulCall(t *testing.T) {
	transport := setupRetryTest(t)

	registerErrorResponder()
	_, _ = transport.RoundTrip(newTestGetRequest(t))
	assertCallCount(t, testRetryCount+1, "call 1 (all fail)")

	httpmock.Reset()
	registerSuccessResponder()
	resp, err := transport.RoundTrip(newTestGetRequest(t))
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("Call 2: expected success, got err=%v status=%v", err, resp.StatusCode)
	}
	assertCallCount(t, 1, "call 2 (success resets counter)")

	httpmock.Reset()
	registerErrorResponder()
	_, _ = transport.RoundTrip(newTestGetRequest(t))
	assertCallCount(t, testRetryCount+1, "call 3 (full budget after reset)")
}

func TestRoundTripStopsOnCancelledContext(t *testing.T) {
	transport := setupRetryTest(t)
	registerErrorResponder()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	req, _ := http.NewRequestWithContext(ctx, testGetHttpMethod, testUrl, nil)

	_, err := transport.RoundTrip(req)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
	assertCallCount(t, 1, "no retries with cancelled context")
}
