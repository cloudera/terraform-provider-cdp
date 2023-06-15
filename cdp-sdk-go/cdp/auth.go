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
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	ed "golang.org/x/crypto/ed25519"
)

const (
	altusAuthHeader      = "x-altus-auth"
	altusDateHeader      = "x-altus-date"
	contentTypeHeader    = "content-type"
	applicationJson      = "application/json"
	userAgentHeaderKey   = "User-agent"
	userAgentHeaderValue = "cdp-terraform"
	signPattern          = "%s\n%s\n%s\n%s\n%s"
	layout               = "Mon, 02 Jan 2006 15:04:05 GMT"
	authAlgo             = "ed25519v1"
)

type metastr struct {
	AccessKey  string `json:"access_key_id"`
	AuthMethod string `json:"auth_method"`
}

func newMetastr(accessKeyID string) *metastr {
	return &metastr{accessKeyID, authAlgo}
}

func requestSigWriter(ctx context.Context, logger Logger, baseAPIPath string, credentials *Credentials) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		date := formatDate()
		auth, err := authHeader(ctx, logger, credentials.AccessKeyId, credentials.PrivateKey, r.GetMethod(), resourcePath(baseAPIPath, r.GetPath(), r.GetQueryParams().Encode()), date)
		if err != nil {
			return err
		}
		err = r.SetHeaderParam(altusAuthHeader, auth)
		logger.Debugf(ctx, "Request signature: %s", auth)
		if err != nil {
			return err
		}
		err = r.SetHeaderParam(contentTypeHeader, applicationJson)
		if err != nil {
			return err
		}
		err = r.SetHeaderParam(userAgentHeaderKey, userAgentHeaderValue)
		if err != nil {
			return err
		}
		return r.SetHeaderParam(altusDateHeader, date)
	})
}

func resourcePath(baseAPIPath, path, query string) string {
	base := escapePath(strings.ReplaceAll(baseAPIPath+path, "//", "/"))
	if len(query) > 0 {
		return fmt.Sprintf("%s?%s", base, query)
	}
	return base
}

func escapePath(path string) string {
	spl := strings.Split(path, "/")
	var encoded []string
	for _, e := range spl {
		encoded = append(encoded, url.PathEscape(e))
	}
	return strings.Join(encoded, "/")
}

func authHeader(ctx context.Context, logger Logger, accessKeyID, privateKey, method, path, date string) (string, error) {
	meta, err := urlSafeMeta(accessKeyID)
	if err != nil {
		return "", err
	}
	sig, err := urlSafeSignature(ctx, logger, privateKey, method, path, date)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s.%s", meta, sig), nil
}

func urlSafeSignature(ctx context.Context, logger Logger, seedBase64, method, path, date string) (string, error) {
	seed, err := base64.StdEncoding.DecodeString(seedBase64)
	if err != nil {
		return "", err
	}
	k := ed.NewKeyFromSeed(seed)
	message := fmt.Sprintf(signPattern, method, applicationJson, date, path, authAlgo)
	logger.Debugf(ctx, "HTTP Message to sign: \n%s\n", message)
	signature := ed.Sign(k, []byte(message))
	return urlSafeBase64Encode(signature), nil
}

func urlSafeMeta(accessKeyID string) (string, error) {
	b, err := json.Marshal(newMetastr(accessKeyID))
	if err != nil {
		return "", err
	}
	return urlSafeBase64Encode(b), nil
}

func urlSafeBase64Encode(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}

func formatDate() string {
	return time.Now().UTC().Format(layout)
}
