package repo

import (
	"ggl_test/models/dto"
	"ggl_test/models/entity"
	"ggl_test/utils"
	"ggl_test/utils/log"
	"github.com/Masterminds/squirrel"
)

//go:generate mockgen -source=task.go -destination=../../mocks/models/repo/task.go
type ITaskRepo interface {
	GetList(c *dto.AppContext) (*[]entity.Task, error)
	Add(c *dto.AppContext, data *entity.Task) (int64, error)
	GetById(c *dto.AppContext, id int64) (*entity.Task, error)
	UpdateById(c *dto.AppContext, id int64, data *entity.Task) error
	DeleteById(c *dto.AppContext, id int64) error
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

func (s *TaskRepo) Add(c *dto.AppContext, data *entity.Task) (int64, error) {
	res, err := squirrel.Insert("tasks").Columns("name").Values(data.Name).RunWith(s.db).Exec()
	if err != nil {
		log.GetLoggerWithCtx(c).Error(err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.GetLoggerWithCtx(c).Error(err)
		return 0, err
	}

	return id, nil
}

func (s *TaskRepo) GetById(c *dto.AppContext, id int64) (*entity.Task, error) {
	var task entity.Task
	err := squirrel.Select("id", "name", "status").From("tasks").
		Where(squirrel.Eq{"id": id}).
		RunWith(s.db).
		Scan(&task.Id, &task.Name, &task.Status)
	if err != nil {
		if utils.IsDataNotFoundErr(err) {
			log.GetLoggerWithCtx(c).Infof("task not found: %d", id)
			return nil, nil
		}
		log.GetLoggerWithCtx(c).Error(err)
		return nil, err
	}
	return &task, nil
}

func (s *TaskRepo) UpdateById(c *dto.AppContext, id int64, data *entity.Task) error {
	_, err := squirrel.Update("tasks").
		Set("name", data.Name).
		Set("status", data.Status).
		Where(squirrel.Eq{"id": id}).
		RunWith(s.db).Exec()
	if err != nil {
		log.GetLoggerWithCtx(c).Error(err)
		return err
	}
	return nil
}

func (s *TaskRepo) DeleteById(c *dto.AppContext, id int64) error {
	_, err := squirrel.Delete("tasks").
		Where(squirrel.Eq{"id": id}).RunWith(s.db).Exec()
	if err != nil {
		log.GetLoggerWithCtx(c).Error(err)
		return err
	}
	return nil
}
