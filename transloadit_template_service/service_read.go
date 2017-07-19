package transloadit_template_service

import (
	"net/http"

	"github.com/bocodigitalmedia/go-transloadit/transloadit_api"
)

func (s *Service) Read(id string) (*Template, *http.Response, error) {
	params := new(ReadParams)
	path := s.Path(id)
	result := new(Template)
	resp, err := s.Api.Get(path, params, result)

	if errResp, ok := err.(*transloadit_api.ErrorResponse); ok && errResp.Code == "SERVER_404" {
		return nil, nil, &NotFound{id}
	} else {
		return result, resp, err
	}
}

type ReadParams struct {
	Auth *transloadit_api.Auth `json:"auth"`
}

func (p *ReadParams) SetAuth(auth *transloadit_api.Auth) {
	p.Auth = auth
}
