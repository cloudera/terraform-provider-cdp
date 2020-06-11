package authn

// This file is mostly lifted from https://github.com/hortonworks/dp-cli-common

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
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

func GetAPIKeyAuthTransport(config *InternalConfig, serviceName string, isAltusService bool) (*Transport, error) {
	credentials, err := getCdpCredentials(config, "")
	if err != nil {
		return nil, err
	}
	var endpoint string
	if isAltusService {
		endpoint = fmt.Sprintf(config.AltusApiEndpointUrl, serviceName)
	} else {
		endpoint = config.CdpApiEndpointUrl
	}
	address, basePath := cutAndTrimAddress(endpoint)
	baseApiPath := config.BaseAPIPath
	transport := &Transport{client.New(address, basePath+baseApiPath, []string{"https"})}
	transport.Runtime.DefaultAuthentication = apiKeyAuth(baseApiPath, credentials)
	// transport.Runtime.Transport = utils.LoggedTransportConfig
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

func apiKeyAuth(baseAPIPath string, credentials *Credentials) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		date := formatDate()
		auth, err := authHeader(credentials.AccessKeyId, credentials.PrivateKey, r.GetMethod(), resourcePath(baseAPIPath, r.GetPath(), r.GetQueryParams().Encode()), date)
		if err != nil {
			return err
		}
		err = r.SetHeaderParam(altusAuthHeader, auth)
		if err != nil {
			return err
		}
		err = r.SetHeaderParam(contentTypeHeader, applicationJson)
		if err != nil {
			return err
		}
		fmt.Println(auth)
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

func authHeader(accessKeyID, privateKey, method, path, date string) (string, error) {
	meta, err := urlSafeMeta(accessKeyID)
	if err != nil {
		return "", err
	}
	sig, err := urlSafeSignature(privateKey, method, path, date)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s.%s", meta, sig), nil
}

func urlSafeSignature(seedBase64, method, path, date string) (string, error) {
	seed, err := base64.StdEncoding.DecodeString(seedBase64)
	if err != nil {
		return "", err
	}
	k := ed.NewKeyFromSeed(seed)
	message := fmt.Sprintf(signPattern, method, applicationJson, date, path, authAlgo)
	log.Debugf("Message to sign: \n%s\n", message)
	fmt.Printf("Message to sign: \n%s\n", message)
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
