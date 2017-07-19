package transloadit_template_service

import (
	"net/http"

	"github.com/bocodigitalmedia/go-transloadit/transloadit_api"
)

func (s *Service) Delete(id string) (*DeleteResult, *http.Response, error) {
	path := s.Path(id)
	result := new(DeleteResult)
	params := new(DeleteParams)
	resp, err := s.Api.Delete(path, params, result)

	if errResp, ok := err.(*transloadit_api.ErrorResponse); ok && errResp.Code == "SERVER_404" {
		return nil, nil, &TemplateNotFound{id}
	} else {
		return result, resp, err
	}
}

type DeleteResult struct {
	Ok string `json:"ok"`
}

type DeleteParams struct {
	Auth *transloadit_api.Auth `json:"auth"`
}

func (p *DeleteParams) SetAuth(auth *transloadit_api.Auth) {
	p.Auth = auth
}
