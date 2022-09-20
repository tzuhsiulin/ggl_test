package repo

import (
	"ggl_test/models/dto"
	"ggl_test/models/entity"
	"ggl_test/utils/log"
	"github.com/Masterminds/squirrel"
)

//go:generate mockgen -source=task.go -destination=../../mocks/models/repo/task.go
type ITaskRepo interface {
	GetList(c *dto.AppContext) (*[]entity.Task, error)
}

type TaskRepo struct {
	db squirrel.BaseRunner
}

func NewTaskRepo(db squirrel.BaseRunner) *TaskRepo {
	return &TaskRepo{
		db: db,
	}
}

func (s *TaskRepo) GetList(c *dto.AppContext) (*[]entity.Task, error) {
	rows, err := squirrel.Select("id", "name", "status").
		From("tasks").
		RunWith(s.db).
		Query()
	if err != nil {
		log.GetLoggerWithCtx(c).Error(err)
		return nil, err
	}
	defer rows.Close()

	result := make([]entity.Task, 0)
	for rows.Next() {
		var task entity.Task
		err := rows.Scan(&task.Id, &task.Name, &task.Status)
		if err != nil {
			log.GetLoggerWithCtx(c).Error(err)
			continue
		}
		result = append(result, task)
	}
	return &result, nil
}
