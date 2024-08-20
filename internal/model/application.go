package model

import (
	"github.com/google/uuid"
	"sparkOnKubernetes/internal/model/spark"
)

type Application struct {
	Id           uuid.UUID          `json:"id"`
	SubmitParams spark.SubmitParams `json:"submit_params"`
}
