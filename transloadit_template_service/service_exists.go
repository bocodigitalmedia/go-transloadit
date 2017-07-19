package transloadit_template_service

func (s *Service) Exists(id string) (bool, error) {
	if _, _, err := s.Read(id); err != nil {
		if _, ok := err.(*TemplateNotFound); ok {
			return false, nil
		} else {
			return false, err
		}
	} else {
		return true, nil
	}
}
