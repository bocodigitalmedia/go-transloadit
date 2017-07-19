package transloadit_template_service

import (
	"net/http"

	"github.com/bocodigitalmedia/go-transloadit/transloadit_api"
)

func (s *Service) Read(id string, payload *ReadPayload) (*ReadResult, *http.Response, error) {
	if payload == nil {
		payload = new(ReadPayload)
	}

	path := s.Path(id)
	result := new(ReadResult)
	resp, err := s.Api.Get(path, payload, result)

	if errResp, ok := err.(*transloadit_api.ErrorResponse); ok && errResp.Code == "SERVER_404" {
		return nil, nil, &NotFound{id}
	} else {
		return result, resp, err
	}
}

type ReadPayload struct {
	Auth *transloadit_api.Auth `json:"auth"`
}

func (p *ReadPayload) SetAuth(auth *transloadit_api.Auth) {
	p.Auth = auth
}

type ReadResult struct {
	Ok      string      `json:"ok"`
	Id      string      `json:"id"`
	Name    string      `json:"name"`
	Content interface{} `json:"content"`
}
