package transloadit_template_service

import "fmt"

type TemplateNotFound struct {
	Id string
}

func (e TemplateNotFound) Error() string {
	return fmt.Sprintf("Template not found: %s", e.Id)
}

func NewTemplateNotFound(id string) *TemplateNotFound {
	return &TemplateNotFound{id}
}
