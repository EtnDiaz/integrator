package services

import (
    "errors" // Import the errors package
    "gitlab.com/roneeSoft/integrator/pkg/shared/services"
)

type Template struct {
    ID      string `json:"id"`
    Name    string `json:"name"`
    Content string `json:"content"`
}

type TemplateService struct {
    baseService *services.BaseService[Template]
}

// NewTemplateService initializes a new TemplateService
func NewTemplateService() *TemplateService {
    return &TemplateService{
        baseService: services.NewBaseService[Template](),
    }
}

// GetByID retrieves a template by its ID
func (s *TemplateService) GetByID(id string) (*Template, error) {
    // Use baseService to get the template
    template, err := s.baseService.Get(id)
    if err != nil {
        return nil, errors.New("template not found")
    }
    return &template, nil
}

func (s *TemplateService) Create(template Template) error {
    return s.baseService.Create(template.ID, template)
}

func (s *TemplateService) Get(id string) (Template, error) {
    return s.baseService.Get(id)
}

func (s *TemplateService) Update(id string, template Template) error {
    return s.baseService.Update(id, template)
}

func (s *TemplateService) Delete(id string) error {
    return s.baseService.Delete(id)
}
