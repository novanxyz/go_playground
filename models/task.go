package models

type TaskStatus string

const (
	COMPLETE   TaskStatus = "complete"
	INCOMPLETE TaskStatus = "incomplete"
)

type TaskFile struct {
	Id       uint   `gorm:"type:int;primary_key"`
	Filename string `gorm:"type:varchar(255)"`
	Mime     string `gorm:"type:varchar(32)"`
	Content  []byte `gorm:"type:longblob"`
	TaskId   uint   `gorm:"ForeignKey:Id"`
}

type Task struct {
	Id     uint        `gorm:"type:int;primary_key"`
	Name   string      `gorm:"type:varchar(255) unique" `
	Status string      `gorm:"type:ENUM('complete', 'incomplete') default 'incomplete' not null"`
	Files  []*TaskFile `gorm:"embedded foreignKey:TaskId"`
}

func TaskFileResponse(elements []*TaskFile) map[uint]string {
	elementMap := make(map[uint]string)
	for _, f := range elements {
		elementMap[f.Id] = f.Filename
	}
	return elementMap
}

func (task *Task) ToResponse() TaskResponse {
	return TaskResponse{
		Id:        task.Id,
		Name:      task.Name,
		Status:    task.Status,
		TaskFiles: TaskFileResponse(task.Files),
	}
}
