package services

type CRUDService[T any] interface {
    Create(item T) error
    Get(id string) (T, error)
    Update(id string, item T) error
    Delete(id string) error
}
