package transloadit_template_service

import (
	"net/http"

	"github.com/bocodigitalmedia/go-transloadit/transloadit_api"
)

func (s *Service) Create(params *CreateParams) (*Template, *http.Response, error) {
	path := s.Path()
	result := new(Template)
	resp, err := s.Api.Post(path, params, result)

	return result, resp, err
}

type CreateParams struct {
	Auth    *transloadit_api.Auth `json:"auth"`
	Name    string                `json:"name"`
	Content *TemplateContent      `json:"template"`
}

func (p *CreateParams) SetAuth(auth *transloadit_api.Auth) {
	p.Auth = auth
}
