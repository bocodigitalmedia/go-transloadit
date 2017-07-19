package transloadit_template_service

import "encoding/json"

type Template struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type TemplateUnmarshal struct {
	Id      string                 `json:"id"`
	Name    string                 `json:"name"`
	Content map[string]interface{} `json:"content"`
}

func (u *Template) UnmarshalJSON(data []byte) error {
	aux := new(TemplateUnmarshal)

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	if bytes, err := json.Marshal(aux.Content); err != nil {
		return err
	} else {
		u.Id = aux.Id
		u.Name = aux.Name
		u.Content = string(bytes)
		return nil
	}
}
