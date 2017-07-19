package transloadit_template_service

import (
	"log"
	"net/http"

	"github.com/bocodigitalmedia/go-transloadit/transloadit_api"
)

func (s *Service) Update(id string, params *UpdateParams) (*Template, *http.Response, error) {
	if exists, err := s.Exists(id); err != nil {
		return nil, nil, err
	} else if !exists {
		return nil, nil, &TemplateNotFound{id}
	}
	log.Printf("EXISTS")
	path := s.Path(id)
	result := new(Template)

	if resp, err := s.Api.Put(path, params, result); err != nil {
		return nil, resp, err
	} else {
		return s.Read(id)
	}
}

type UpdateParams struct {
	Auth    *transloadit_api.Auth `json:"auth"`
	Name    string                `json:"name"`
	Content string                `json:"template"`
}

func (p *UpdateParams) SetAuth(auth *transloadit_api.Auth) {
	p.Auth = auth
}
