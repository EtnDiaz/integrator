// shared/services/version_control_service.go
package services

import "errors"

// VersionControlService управляет версиями объектов.
type VersionControlService[T any] struct {
    versions map[string][]T
}

// NewVersionControlService создает новый VersionControlService.
func NewVersionControlService[T any]() *VersionControlService[T] {
    return &VersionControlService[T]{
        versions: make(map[string][]T),
    }
}

func (s *VersionControlService[T]) SaveVersion(id string, item T) {
    s.versions[id] = append(s.versions[id], item)
}

func (s *VersionControlService[T]) GetVersions(id string) ([]T, error) {
    versions, exists := s.versions[id]
    if !exists {
        return nil, errors.New("no versions found")
    }
    return versions, nil
}

func (s *VersionControlService[T]) RevertToVersion(id string, version int) (T, error) {
    versions, exists := s.versions[id]
    if !exists || version >= len(versions) {
        return *new(T), errors.New("invalid version")
    }
    return versions[version], nil
}
