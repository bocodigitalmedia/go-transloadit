package transloadit_template_service

import (
	"net/http"

	"github.com/bocodigitalmedia/go-transloadit/transloadit_api"
)

func (s *Service) Update(id string, params *UpdateParams) (*UpdateResult, *http.Response, error) {
	path := s.Path(id)
	result := new(UpdateResult)
	resp, err := s.Api.Put(path, params, result)

	if errResp, ok := err.(*transloadit_api.ErrorResponse); ok && errResp.Code == "SERVER_404" {
		return nil, nil, &TemplateNotFound{id}
	} else {
		return result, resp, err
	}
}

type UpdateResult struct {
	Ok string `json:"ok"`
}

type UpdateParams struct {
	Auth    *transloadit_api.Auth `json:"auth"`
	Name    string                `json:"name"`
	Content string                `json:"template"`
}

func (p *UpdateParams) SetAuth(auth *transloadit_api.Auth) {
	p.Auth = auth
}
