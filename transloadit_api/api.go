package transloadit_api

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

type Api struct {
	sling         *sling.Sling
	authKey       string
	authSecret    string
	authExpiresIn time.Duration
}

func (a *Api) Base() *sling.Sling {
	return a.sling.New()
}

func (a *Api) Auth() *Auth {
	return NewAuthExpiresIn(a.authKey, a.authExpiresIn)
}

func (a *Api) Signature(b []byte) string {
	hash := hmac.New(sha1.New, []byte(a.authSecret))
	hash.Write(b)
	sum := hash.Sum(nil)

	return hex.EncodeToString(sum)
}

func (a *Api) Receive(decorate SlingDecorator, result *interface{}) (*http.Response, error) {
	errResp := new(ErrorResponse)
	base := decorate(a.Base())
	resp, err := base.Receive(result, errResp)

	if err == nil && !errResp.IsEmpty() {
		err = errResp
	}

	return resp, err
}

func (a *Api) Payload(params interface{}) (*Payload, error) {
	params.(AuthorizableParams).SetAuth(a.Auth())

	if bytes, err := json.Marshal(params); err != nil {
		return nil, err
	} else {
		signature := a.Signature(bytes)

		payload := &Payload{
			Params:    string(bytes),
			Signature: signature,
		}

		return payload, nil
	}
}

func (a *Api) Post(path string, params AuthorizableParams, result interface{}) (*http.Response, error) {
	if payload, err := a.Payload(params); err != nil {
		return nil, err
	} else {

		decorate := func(s *sling.Sling) *sling.Sling {
			return s.Post(path).BodyJSON(payload)
		}

		return a.Receive(decorate, &result)
	}
}

func (a *Api) Get(path string, params AuthorizableParams, result interface{}) (*http.Response, error) {
	if payload, err := a.Payload(params); err != nil {
		return nil, err
	} else {

		decorate := func(s *sling.Sling) *sling.Sling {
			return s.Get(path).QueryStruct(payload)
		}

		return a.Receive(decorate, &result)
	}
}

func (a *Api) Put(path string, params AuthorizableParams, result interface{}) (*http.Response, error) {
	if payload, err := a.Payload(params); err != nil {
		return nil, err
	} else {

		decorate := func(s *sling.Sling) *sling.Sling {
			return s.Put(path).BodyJSON(payload)
		}

		return a.Receive(decorate, &result)
	}
}

func (a *Api) Delete(path string, params AuthorizableParams, result interface{}) (*http.Response, error) {
	if payload, err := a.Payload(params); err != nil {
		return nil, err
	} else {
		decorate := func(s *sling.Sling) *sling.Sling {
			return s.Delete(path).BodyJSON(payload)
		}

		return a.Receive(decorate, &result)
	}
}

type NewApiParams struct {
	AuthKey       string
	AuthSecret    string
	AuthExpiresIn time.Duration
	HttpClient    *http.Client
	BaseUrl       string
}

func New(params *NewApiParams) *Api {
	httpClient := params.HttpClient
	baseUrl := params.BaseUrl
	authExpiresIn := params.AuthExpiresIn

	if httpClient == nil {
		httpClient = &http.Client{Timeout: time.Minute * 1}
	}

	if baseUrl == "" {
		baseUrl = BaseUrl
	}

	if authExpiresIn == 0 {
		authExpiresIn = AuthExpiresIn
	}

	sling := sling.New().Base(baseUrl).Client(httpClient)

	return &Api{
		sling:         sling,
		authKey:       params.AuthKey,
		authSecret:    params.AuthSecret,
		authExpiresIn: authExpiresIn,
	}
}
