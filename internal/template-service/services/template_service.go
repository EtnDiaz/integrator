// services/template_service.go
package services

import (
    "shared/services"
    "shared/utils"
)

type Template struct {
    ID      string `json:"id"`
    Name    string `json:"name"`
    Content string `json:"content"`
}

type TemplateService struct {
    baseService *services.BaseService[Template]
}

func NewTemplateService() *TemplateService {
    return &TemplateService{
        baseService: services.NewBaseService[Template](),
    }
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
