package transloadit_template_service

type Template struct {
	Id      string           `json:"id,omitempty"`
	Name    string           `json:"name,omitempty"`
	Content *TemplateContent `json:"content,omitempty"`
}

type TemplateContent struct {
	Steps map[string]interface{} `json:"steps"`
}

func NewTemplate(name string, content *TemplateContent) *Template {
	return &Template{
		Name:    name,
		Content: content,
	}
}
