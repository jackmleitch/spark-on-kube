package main

import (
	"sparkOnKubernetes/internal/api"
	"sparkOnKubernetes/internal/model"
	"sparkOnKubernetes/internal/storage"
)

func main() {
	storageService := storage.NewInMemoryStorage[model.Application]()
	api.Router(storageService)
}
