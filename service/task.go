package service

import (
	"ggl_test/models/dto"
	"ggl_test/models/entity"
	"ggl_test/models/repo"
	"ggl_test/utils/customerror"
	"ggl_test/utils/log"
)

//go:generate mockgen -source=task.go -destination=../mocks/service/task.go
type ITaskService interface {
	GetAll(c *dto.AppContext) (*[]entity.Task, *customerror.CustomError)
}

type TaskService struct {
	taskRepo repo.ITaskRepo
}

func NewTaskService(taskRepo repo.ITaskRepo) ITaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) GetAll(c *dto.AppContext) (*[]entity.Task, *customerror.CustomError) {
	taskList, err := s.taskRepo.GetList(c)
	if err != nil {
		log.GetLoggerWithCtx(c).Error("failed to get task list")
		return nil, customerror.NewErr(customerror.ErrorCodeUnknown)
	}
	return taskList, nil
}
