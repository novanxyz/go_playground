package models

type TaskResponse struct {
	Id        uint            `json:"id"`
	Name      string          `json:"name"`
	Status    string          `json:"status"`
	TaskFiles map[uint]string `json:"files,omitempty"`
}

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}
