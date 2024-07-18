package models

type Status string

const (
	Complete   Status = "complete"
	Incomplete Status = "incomplete"
)

type CreateTaskRequest struct {
	Name string `validate:"required,min=1,max=200" json:"name"`
}

type UpdateTaskRequest struct {
	Name   string `validate:"required,max=200,min=1" json:"name"`
	Status string `validate:"enum" json:"status"`
}
