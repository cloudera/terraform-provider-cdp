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
	"log"
	"net/http"
	"time"
)

var (
	retryableStatusCodes = []int{
		http.StatusServiceUnavailable,
		http.StatusTooManyRequests,
		http.StatusGatewayTimeout,
		http.StatusBadGateway,
		http.StatusPreconditionFailed,
	}
)

type RetryableTransport struct {
	transport http.RoundTripper
}

func shouldRetry(err error, resp *http.Response) bool {
	if err != nil {
		return true
	} else if resp == nil {
		return false
	}
	return sliceContains(retryableStatusCodes, resp.StatusCode)
}

func drainBody(resp *http.Response) {
	if resp != nil && resp.Body != nil {
		_, err := io.Copy(io.Discard, resp.Body)
		if err != nil {
			log.Default().Println("Error while draining body: ", err)
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)
	}
}

func (t *RetryableTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = io.ReadAll(req.Body)
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}
	resp, err := t.transport.RoundTrip(req)
	retries := 0
	retryCount := intFromEnvOrDefault("CDP_TF_CALL_RETRY_COUNT", 10)
	for shouldRetry(err, resp) && retries < retryCount {
		log.Default().Printf("Retrying request (caused by: %+v;%+v)\n", err, resp)
		time.Sleep(backoff(retries))
		drainBody(resp)
		if req.Body != nil {
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
		resp, err = t.transport.RoundTrip(req)
		fmt.Printf("%v retry out of %v\n", retries+1, retryCount)
		retries++
	}
	return resp, err
}
