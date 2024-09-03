package services

import "errors"

type VersionControlService struct {
    versions map[string][]Template
}

func NewVersionControlService() *VersionControlService {
    return &VersionControlService{
        versions: make(map[string][]Template),
    }
}

func (s *VersionControlService) SaveVersion(id string, template Template) {
    s.versions[id] = append(s.versions[id], template)
}

func (s *VersionControlService) GetVersions(id string) ([]Template, error) {
    versions, exists := s.versions[id]
    if !exists {
        return nil, errors.New("no versions found for template")
    }
    return versions, nil
}

func (s *VersionControlService) RevertToVersion(id string, version int) (Template, error) {
    versions, exists := s.versions[id]
    if !exists || version >= len(versions) {
        return Template{}, errors.New("invalid version")
    }
    return versions[version], nil
}
