package repository

import (
	"errors"
	"fmt"
	"novanxyz/models"
	"novanxyz/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TaskRepositoryInterface interface {
	Save(task models.Task) (models.Task, error)
	Update(task models.Task) int
	Delete(taskId uint) int
	FindById(taskId uint) (task models.Task, err error)
	FindAll(filter map[string]interface{}, page int, size int) []models.Task
	Mark(taskId uint, status string) int
	FindFileById(fileId uint) (taskFile models.TaskFile, err error)
}

type TaskRepository struct {
	Db *gorm.DB
}

func NewTaskRepository(Db *gorm.DB) TaskRepositoryInterface {
	return &TaskRepository{Db: Db}
}

func (t *TaskRepository) Delete(taskId uint) int {
	var task models.Task
	result := t.Db.Where("id = ?", taskId).Delete(&task)
	utils.ErrorPanic(result.Error)
	return int(result.RowsAffected)
}

// FindAll implements TaskRepository
func (t *TaskRepository) FindAll(filter map[string]interface{}, page int, size int) []models.Task {

	var tasks []models.Task
	result := t.Db.Where(filter).Find(&tasks).Limit(size).Offset((page - 1) * size)
	utils.ErrorPanic(result.Error)
	return tasks
}

// FindById implements TaskRepository
func (t *TaskRepository) FindById(taskId uint) (models.Task, error) {
	var task models.Task
	result := t.Db.Preload("Files").Find(&task, taskId)
	fmt.Sprintf("%v\n", result)

	if result.RowsAffected == 1 {
		return task, nil
	} else {
		return task, errors.New("task is not found")
	}
}

func (t *TaskRepository) Save(task models.Task) (models.Task, error) {
	result := t.Db.Clauses(clause.Returning{}).Create(&task)
	utils.ErrorPanic(result.Error)
	return t.FindById(task.Id)
}

func (t *TaskRepository) Update(task models.Task) int {
	var updateTask = models.Task{
		Name:   task.Name,
		Status: task.Status,
	}
	result := t.Db.Model(&task).Updates(updateTask)
	utils.ErrorPanic(result.Error)
	return int(result.RowsAffected)
}

func (t *TaskRepository) Mark(taskId uint, status string) int {
	task, err := t.FindById(taskId)
	utils.ErrorPanic(err)
	task.Status = status
	result := t.Db.Model(&task).Updates(task)
	utils.ErrorPanic(result.Error)
	return int(result.RowsAffected)
}
func (t *TaskRepository) FindFileById(fileId uint) (taskFile models.TaskFile, err error) {
	var file models.TaskFile
	result := t.Db.Find(&file, fileId)
	if result != nil {
		return file, nil
	} else {
		return file, errors.New("file is not found")
	}
}
