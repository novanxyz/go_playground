package models

type TaskResponse struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	TaskFiles []uint `json:"task_files,omitempty"`
}

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}
