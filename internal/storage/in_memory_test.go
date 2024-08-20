package storage

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInMemoryStorage(t *testing.T) {
	storage := NewInMemoryStorage[string]()

	assert.NotNil(t, storage)
	assert.NotNil(t, storage.cache)
}

func TestInMemoryStorage_Get(t *testing.T) {
	storage := NewInMemoryStorage[string]()
	key := uuid.New()
	value := "testValue"
	storage.Put(key, value)

	result, found := storage.Get(key)

	assert.NotNil(t, found)
	assert.Equal(t, value, result)
}

func TestInMemoryStorage_GetAll(t *testing.T) {
	storage := NewInMemoryStorage[string]()
	key1 := uuid.New()
	key2 := uuid.New()
	value1 := "testValue1"
	value2 := "testValue2"
	storage.Put(key1, value1)
	storage.Put(key2, value2)

	results := storage.GetAll()

	assert.Contains(t, results, value1)
	assert.Contains(t, results, value2)
}

func TestInMemoryStorage_Put(t *testing.T) {
	storage := NewInMemoryStorage[string]()
	key := uuid.New()
	value := "testValue"

	result := storage.Put(key, value)
	assert.Equal(t, value, result)

	storedValue, found := storage.Get(key)
	assert.True(t, found)
	assert.Equal(t, value, storedValue)
}

func TestInMemoryStorage_Update(t *testing.T) {
	storage := NewInMemoryStorage[string]()
	key := uuid.New()
	value := "testValue"
	updatedValue := "updatedValue"

	storage.Put(key, value)
	result := storage.Update(key, updatedValue)
	assert.Equal(t, updatedValue, result)

	storedValue, found := storage.Get(key)
	assert.True(t, found)
	assert.Equal(t, updatedValue, storedValue)
}

func TestInMemoryStorage_Delete(t *testing.T) {
	storage := NewInMemoryStorage[string]()
	key := uuid.New()
	value := "testValue"
	storage.Put(key, value)
	_, found := storage.Get(key)
	assert.True(t, found)

	storage.Delete(key)

	_, found = storage.Get(key)
	assert.False(t, found)
}
