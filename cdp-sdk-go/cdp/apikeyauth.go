package cdp

// This file is mostly lifted from https://github.com/hortonworks/dp-cli-common

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	ed "golang.org/x/crypto/ed25519"
)

const (
	altusAuthHeader   = "x-altus-auth"
	altusDateHeader   = "x-altus-date"
	contentTypeHeader = "content-type"
	applicationJson   = "application/json"
	signPattern       = "%s\n%s\n%s\n%s\n%s"
	layout            = "Mon, 02 Jan 2006 15:04:05 GMT"
	authAlgo          = "ed25519v1"
)

type metastr struct {
	AccessKey  string `json:"access_key_id"`
	AuthMethod string `json:"auth_method"`
}

func newMetastr(accessKeyID string) *metastr {
	return &metastr{accessKeyID, authAlgo}
}

func GetAPIKeyAuthTransport(ctx context.Context, logger Logger, credentials *Credentials, endpoint string, baseApiPath string, insecureSkipVerify bool) (*Transport, error) {
	address, basePath := cutAndTrimAddress(endpoint)
	tlsClientOptions := client.TLSClientOptions{
		InsecureSkipVerify: insecureSkipVerify,
	}
	cfg, err := client.TLSClientAuth(tlsClientOptions)
	if err != nil {
		return nil, err
	}

	roundTripper := &LoggingRoundTripper{
		delegate: &http.Transport{TLSClientConfig: cfg},
		logger:   logger,
	}
	transport := &Transport{client.NewWithClient(address, basePath+baseApiPath, []string{"https"}, &http.Client{Transport: roundTripper})}
	transport.Runtime.DefaultAuthentication = apiKeyAuth(ctx, logger, baseApiPath, credentials)
	return transport, nil
}

var prefixTrim = []string{"http://", "https://"}

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

func apiKeyAuth(ctx context.Context, logger Logger, baseAPIPath string, credentials *Credentials) runtime.ClientAuthInfoWriter {
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
	encoded := []string{}
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

type Transport struct {
	Runtime *client.Runtime
}

func (t *Transport) Submit(operation *runtime.ClientOperation) (interface{}, error) {
	response, err := t.Runtime.Submit(operation)
	return response, err
}

type LoggingRoundTripper struct {
	delegate http.RoundTripper
	logger   Logger
}

func (t *LoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	startTime := time.Now()
	resp, err := t.delegate.RoundTrip(req)
	duration := time.Since(startTime)
	errMsg := ""
	if err != nil {
		errMsg = fmt.Sprintf("error=%s", err.Error())
	}

	t.logger.Debugf(req.Context(), "HTTP Request URL=%s method=%s status=%d %s resp=%v durationMs=%d", req.URL, req.Method, resp.StatusCode, errMsg, resp, duration)

	return resp, err
}
