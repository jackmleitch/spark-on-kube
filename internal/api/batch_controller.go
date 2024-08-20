package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"sparkOnKubernetes/internal/model"
	"sparkOnKubernetes/internal/model/spark"
	"sparkOnKubernetes/internal/storage"
)

type BatchController struct {
	storageService storage.StorageService[model.Application]
}

func registerBatchController(router *gin.Engine, storageService storage.StorageService[model.Application]) {
	controller := &BatchController{storageService: storageService}
	batch := router.Group("/batch")
	{
		batch.POST("", controller.createJob)
		batch.GET("", controller.getAllJobs)
		batch.GET("/:id", controller.getJobNyName)
		batch.DELETE("/:id", controller.deleteJobByName)
	}
}

func (bc *BatchController) createJob(c *gin.Context) {
	var newParams spark.SubmitParams
	if err := c.BindJSON(&newParams); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	application := model.Application{Id: uuid.New(), SubmitParams: newParams}
	bc.storageService.Put(application.Id, application)
	c.IndentedJSON(http.StatusCreated, application)
}

func (bc *BatchController) getAllJobs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, bc.storageService.GetAll())
}

func (bc *BatchController) getJobNyName(c *gin.Context) {
	id := c.Param("id")
	parsedUUID, ok := uuid.Parse(id)
	if ok != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unable to parse string id to UUID", "id": id})
		return
	}
	maybeSubmitParams, found := bc.storageService.Get(parsedUUID)
	if !found {
		c.Status(http.StatusNotFound)
		return
	}
	c.IndentedJSON(http.StatusOK, maybeSubmitParams)
}

func (bc *BatchController) deleteJobByName(c *gin.Context) {
	id := c.Param("id")
	parsedUUID, ok := uuid.Parse(id)
	if ok != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unable to parse string id to UUID", "id": id})
		return
	}
	bc.storageService.Delete(parsedUUID)
	c.Status(http.StatusNoContent)
}
