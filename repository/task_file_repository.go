package repository

import (
	"errors"
	"novanxyz/models"
	"novanxyz/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TaskFileRepositoryInterface interface {
	Save(taskFile models.TaskFile) uint
	FindById(fileId uint) (task models.TaskFile, err error)
	FindTaskFile(taskId uint) []models.TaskFile
	Delete(taskFileId uint) int
}

type TaskFileRepository struct {
	Db *gorm.DB
}

func NewTaskFileRepository(Db *gorm.DB) TaskFileRepositoryInterface {
	return &TaskFileRepository{Db: Db}
}

func (t *TaskFileRepository) Save(taskFile models.TaskFile) uint {
	result := t.Db.Clauses(clause.Returning{}).Create(&taskFile)
	utils.ErrorPanic(result.Error)
	return taskFile.Id
}

func (t *TaskFileRepository) FindById(fileId uint) (models.TaskFile, error) {
	var file models.TaskFile
	result := t.Db.Find(&file, fileId)
	if result != nil {
		return file, nil
	} else {
		return file, errors.New("file is not found")
	}
}

func (t *TaskFileRepository) FindTaskFile(taskId uint) []models.TaskFile {

	var files []models.TaskFile
	result := t.Db.Where("ParentTask = ? ", taskId).Find(&files)
	if result != nil {
		return files
	} else {
		return nil
	}
}

func (t *TaskFileRepository) Delete(fileId uint) int {
	var file models.TaskFile
	result := t.Db.Where("id = ?", fileId).Delete(&file)
	utils.ErrorPanic(result.Error)
	return int(result.RowsAffected)
}
