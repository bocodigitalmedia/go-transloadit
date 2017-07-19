package transloadit_template_service

import (
	"strings"

	"github.com/bocodigitalmedia/go-transloadit/transloadit_api"
)

const BasePath = "templates"

type Service struct {
	Api *transloadit_api.Api
}

func (s *Service) Path(paths ...string) string {
	strs := append([]string{BasePath}, paths...)
	return strings.Join(strs, "/")
}

func New(api *transloadit_api.Api) *Service {
	return &Service{Api: api}
}
