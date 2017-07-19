package transloadit_template_service

import "fmt"

type NotFound struct {
	Id string
}

func (e *NotFound) Error() string {
	return fmt.Sprintf("Template not found: %s", e.Id)
}

func NewNotFound(id string) *NotFound {
	return &NotFound{id}
}
