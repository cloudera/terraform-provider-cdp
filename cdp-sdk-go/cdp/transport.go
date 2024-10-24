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
	"net/http"
	"strings"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

var (
	prefixTrim = []string{"http://", "https://"}
)

const (
	defaultLinearBackoffStep = 2
)

type ClientTransport struct {
	Runtime *client.Runtime
}

type DelegatingRoundTripper struct {
	delegate http.RoundTripper
}

type LoggingRoundTripper struct {
	DelegatingRoundTripper
	logger Logger
}

// RequestHeadersRoundTripper sets the User-Agent and other custom headers
// see https://github.com/go-swagger/go-swagger/blob/701e7f3ee85df9d47fcf639dd7a279f7ab6d94d7/docs/faq/faq_client.md?plain=1#L28
type RequestHeadersRoundTripper struct {
	DelegatingRoundTripper
	headers map[string]string
}

func (t *ClientTransport) Submit(operation *runtime.ClientOperation) (interface{}, error) {
	response, err := t.Runtime.Submit(operation)
	return response, err
}

func getDefaultTransport(config *Config) (http.RoundTripper, error) {
	tlsClientOptions := client.TLSClientOptions{
		InsecureSkipVerify: config.GetLocalEnvironment(),
	}
	cfg, err := client.TLSClientAuth(tlsClientOptions)
	if err != nil {
		return nil, err
	}

	retryableTransport := &RetryableTransport{
		transport: &http.Transport{Proxy: http.ProxyFromEnvironment, TLSClientConfig: cfg},
	}

	return retryableTransport, nil
}

func buildClientTransportWithDefaultHttpTransport(config *Config, endpoint string) (*ClientTransport, error) {
	roundTripper, err := getDefaultTransport(config)
	if err != nil {
		return nil, err
	}
	return buildClientTransport(config, endpoint, roundTripper)
}

func buildClientTransport(config *Config, endpoint string, roundTripper http.RoundTripper) (*ClientTransport, error) {
	credentials, err := config.GetCredentials()
	if err != nil {
		return nil, err
	}
	baseApiPath := config.BaseApiPath
	address, basePath := cutAndTrimAddress(endpoint)

	rtChain := buildInterceptorChain(config, roundTripper)

	transport := &ClientTransport{client.NewWithClient(address, basePath+baseApiPath, []string{"https"}, &http.Client{Transport: rtChain})}
	// TODO: Look into whether this should be done as a RoundTripper in the chain to unify the request interceptors.
	transport.Runtime.DefaultAuthentication = requestSigWriter(config.Context, config.Logger, baseApiPath, credentials)

	return transport, nil
}

// TODO: this should use a proper URL parser
func cutAndTrimAddress(address string) (string, string) {
	for _, v := range prefixTrim {
		address = strings.TrimPrefix(address, v)
	}
	address = strings.TrimRight(address, "/ ")
	basePath := ""
	slashIndex := strings.Index(address, "/")
	if slashIndex != -1 {
		basePath = address[slashIndex:]
		address = address[0:slashIndex]
	}
	return address, basePath
}

// buildInterceptorChain builds a chain of RoundTripper objects that modify the request and delegates
// to the next one in the chain.
func buildInterceptorChain(config *Config, rt0 http.RoundTripper) http.RoundTripper {
	rt1 := buildRequestHeadersRoundTripper(config, rt0)
	rt2 := buildLoggingRoundTripper(config, rt1)
	return rt2
}

func buildLoggingRoundTripper(config *Config, delegate http.RoundTripper) *LoggingRoundTripper {
	return &LoggingRoundTripper{
		DelegatingRoundTripper: DelegatingRoundTripper{delegate: delegate},
		logger:                 config.Logger,
	}
}

func buildRequestHeadersRoundTripper(config *Config, delegate http.RoundTripper) *RequestHeadersRoundTripper {
	reqHeadersRT := &RequestHeadersRoundTripper{
		DelegatingRoundTripper: DelegatingRoundTripper{delegate: delegate},
	}
	reqHeadersRT.AddHeader("User-Agent", config.GetUserAgentOrDefault())
	reqHeadersRT.AddHeader("x-altus-client-app", config.ClientApplicationName)
	return reqHeadersRT
}

func (t *LoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	startTime := time.Now()
	resp, err := t.delegate.RoundTrip(req)
	duration := time.Since(startTime)

	if err != nil {
		t.logger.Debugf(req.Context(), "HTTP Request URL=%s method=%s error=%s resp=%v durationMs=%d", req.URL, req.Method, resp, err.Error(), duration)
	} else {
		t.logger.Debugf(req.Context(), "HTTP Request URL=%s method=%s status=%d resp=%v durationMs=%d", req.URL, req.Method, resp.StatusCode, resp, duration)
	}

	return resp, err
}

func (r *RequestHeadersRoundTripper) AddHeader(key, value string) {
	if key == "" || value == "" {
		return
	}
	if r.headers == nil {
		r.headers = make(map[string]string)
	}
	r.headers[key] = value
}

func (r *RequestHeadersRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if r.headers != nil {
		for k, v := range r.headers {
			req.Header.Set(k, v)
		}
	}

	return r.delegate.RoundTrip(req)
}
