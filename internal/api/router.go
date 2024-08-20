package api

import (
	"github.com/gin-gonic/gin"
	"sparkOnKubernetes/internal/model"
	"sparkOnKubernetes/internal/storage"
)

func Router(storageService storage.StorageService[model.Application]) {
	router := gin.Default()
	registerBatchController(router, storageService)
	router.Run(":8080")
}
