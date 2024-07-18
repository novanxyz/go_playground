package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"novanxyz/models"
	"novanxyz/repository"
	"novanxyz/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/szmcdull/glinq/garray"
)

type TaskServiceInterface interface {
	Create(task models.CreateTaskRequest) (models.Task, error)
	Update(taskid uint, task models.UpdateTaskRequest) (models.Task, error)
	Delete(taskId uint) (models.Task, error)
	FindById(taskId uint) (models.Task, error)
	FindAll(query interface{}) ([]models.Task, error)
	Mark(taskId uint, status string) int
	AssignTaskFile(taskId uint, uploadedFile *multipart.FileHeader) (uint, error)
	GetTaskFile(taskId uint, fileId uint) (models.TaskFile, error)
	DeleteTaskFile(taskId uint, fileId uint) (int, error)
	GetAllTaskFiles(taskId uint) ([]uint, error)
}

type TaskService struct {
	TaskRepository     repository.TaskRepositoryInterface
	TaskFileRepository repository.TaskFileRepositoryInterface
	Validator          *validator.Validate
}

func NewTaskService(taskRepository repository.TaskRepositoryInterface, taskFileRepository repository.TaskFileRepositoryInterface, validator *validator.Validate) TaskServiceInterface {
	return &TaskService{
		TaskRepository:     taskRepository,
		TaskFileRepository: taskFileRepository,
		Validator:          validator,
	}
}

// Create implements TaskService
func (t *TaskService) Create(taskRequest models.CreateTaskRequest) (models.Task, error) {
	err := t.Validator.Struct(taskRequest)
	utils.ErrorPanic(err)
	taskModel := models.Task{
		Name:   taskRequest.Name,
		Status: "incomplete", //# default value
	}
	task, err := t.TaskRepository.Save(taskModel)
	return task, err
}

// Delete implements TaskService
func (t *TaskService) Delete(taskId uint) (models.Task, error) {
	task, err := t.TaskRepository.FindById(taskId)
	utils.ErrorPanic(err)

	t.TaskRepository.Delete(taskId)
	return task, nil
}

// FindAll implements TaskService
func (t *TaskService) FindAll(filter interface{}) ([]models.Task, error) {

	filters := filter.(map[string]interface{})
	var page int
	var size int
	var err error
	if p, ok := filters["p"]; ok {
		tmp := p.(string)
		page, err = strconv.Atoi(tmp)
		utils.ErrorPanic(err)
		delete(filters, "p")
	} else {
		page = 1
	}

	if s, ok := filters["s"]; ok {
		tmp := s.(string)
		page, err = strconv.Atoi(tmp)
		utils.ErrorPanic(err)
		delete(filters, "s")
	} else {
		size = 10
	}

	fmt.Println(filter, filters, page, size)
	result := t.TaskRepository.FindAll(filters, page, size)

	return result, nil
}

// FindById implements TaskService
func (t *TaskService) FindById(taskId uint) (models.Task, error) {
	taskData, err := t.TaskRepository.FindById(taskId)
	return taskData, err
}

// Update implements TaskService
func (t *TaskService) Update(taskId uint, task models.UpdateTaskRequest) (models.Task, error) {
	taskData, err := t.TaskRepository.FindById(taskId)
	utils.ErrorPanic(err)
	taskData.Name = task.Name
	taskData.Status = task.Status
	t.TaskRepository.Update(taskData)
	return t.FindById(taskId)
}

func (t *TaskService) Mark(taskId uint, status string) int {
	taskData, err := t.TaskRepository.FindById(taskId)
	utils.ErrorPanic(err)
	taskData.Status = status
	affected := t.TaskRepository.Update(taskData)
	return affected
}

func (t *TaskService) AssignTaskFile(taskId uint, uploadFile *multipart.FileHeader) (uint, error) {

	fmt.Println(uploadFile.Header)
	task, err := t.TaskRepository.FindById(taskId)
	utils.ErrorPanic(err)

	fileHandler, err := uploadFile.Open()
	utils.ErrorPanic(err)

	content, err := io.ReadAll(fileHandler)
	utils.ErrorPanic(err)

	taskFile := models.TaskFile{Filename: uploadFile.Filename, Mime: uploadFile.Header.Get("Content-Type"), ParentTask: &task, Content: content}
	affected := t.TaskFileRepository.Save(taskFile)
	return affected, nil
}

func (t *TaskService) GetTaskFile(taskId uint, fileId uint) (models.TaskFile, error) {
	taskFile, err := t.TaskRepository.FindFileById(fileId)
	utils.ErrorPanic(err)
	return taskFile, nil
}

func (t *TaskService) DeleteTaskFile(taskId uint, fileId uint) (int, error) {
	taskData, err := t.TaskRepository.FindById(taskId)
	utils.ErrorPanic(err)
	return t.TaskRepository.Update(taskData), nil
}

func (t *TaskService) GetAllTaskFiles(taskId uint) ([]uint, error) {
	task, err := t.TaskRepository.FindById(taskId)
	utils.ErrorPanic(err)
	fileIds := garray.MapI(task.Files, func(i int) uint {
		return task.Files[i].Id
	})

	return fileIds, nil
}
