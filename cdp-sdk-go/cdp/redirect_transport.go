// Copyright 2025 Cloudera. All Rights Reserved.
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
	"io"
	"net/http"
	"net/url"
)

// RedirectSigningTransport handles HTTP 308 redirects by re-signing the request
// for the new URL. This is needed for DF APIs like importFlowDefinition that
// redirect to a different host (e.g., console.*.cdp.cloudera.com) where the
// original signature is invalid.
type RedirectSigningTransport struct {
	transport   http.RoundTripper
	credentials *Credentials
	logger      Logger
	baseAPIPath string
}

func (t *RedirectSigningTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Buffer the body so we can replay it after redirect
	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = io.ReadAll(req.Body)
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	resp, err := t.transport.RoundTrip(req)
	if err != nil {
		return resp, err
	}

	if resp.StatusCode != http.StatusPermanentRedirect {
		return resp, err
	}

	location := resp.Header.Get("Location")
	if location == "" {
		return resp, err
	}

	drainBody(resp)

	redirectURL, err := url.Parse(location)
	if err != nil {
		return resp, err
	}

	newReq, err := http.NewRequestWithContext(req.Context(), req.Method, location, nil)
	if err != nil {
		return nil, err
	}

	if bodyBytes != nil {
		newReq.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		newReq.ContentLength = int64(len(bodyBytes))
	}

	// Copy original headers
	for k, v := range req.Header {
		newReq.Header[k] = v
	}

	// Re-sign for the new URL path
	date := FormatDate()
	path := redirectURL.Path
	if redirectURL.RawQuery != "" {
		path = path + "?" + redirectURL.RawQuery
	}
	auth, err := authHeader(req.Context(), t.logger, t.credentials.AccessKeyId, t.credentials.PrivateKey, newReq.Method, path, date)
	if err != nil {
		return nil, err
	}
	newReq.Header.Set(altusAuthHeader, auth)
	newReq.Header.Set(altusDateHeader, date)

	return t.transport.RoundTrip(newReq)
}
