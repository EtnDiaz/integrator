package services

import "errors"

type Integration struct {
    ID     string `json:"id"`
    Name   string `json:"name"`
    Config string `json:"config"`
}

type IntegrationService struct {
    integrations map[string]Integration
}

func NewIntegrationService() *IntegrationService {
    return &IntegrationService{
        integrations: make(map[string]Integration),
    }
}

func (s *IntegrationService) Create(integration Integration) error {
    if _, exists := s.integrations[integration.ID]; exists {
        return errors.New("integration already exists")
    }
    s.integrations[integration.ID] = integration
    return nil
}

func (s *IntegrationService) Get(id string) (Integration, error) {
    integration, exists := s.integrations[id]
    if !exists {
        return Integration{}, errors.New("integration not found")
    }
    return integration, nil
}

func (s *IntegrationService) Delete(id string) error {
    if _, exists := s.integrations[id]; !exists {
        return errors.New("integration not found")
    }
    delete(s.integrations, id)
    return nil
}
