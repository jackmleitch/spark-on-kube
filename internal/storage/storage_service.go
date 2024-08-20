package storage

import "github.com/google/uuid"

type StorageService[T any] interface {
	Get(key uuid.UUID) (T, bool)
	Put(key uuid.UUID, value T) T
	Delete(key uuid.UUID)
	GetAll() []T
	Update(key uuid.UUID, value T) T
}
