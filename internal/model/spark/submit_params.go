package spark

type SubmitParams struct {
	Name           string `json:"name"`
	File           string `json:"file"`
	NumExecutors   int    `json:"num_executors"`
	ExecutorCores  int    `json:"executor_cores"`
	ExecutorMemory string `json:"executor_memory"`
	DriverCores    int    `json:"driver_cores"`
	DriverMemory   string `json:"driver_memory"`
}
