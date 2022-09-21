package repo

import (
	"errors"
	"testing"

	"ggl_test/models/dto"
	"ggl_test/models/entity"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTaskRepo_GetList(t *testing.T) {
	t.Run("shouldReturnTaskListCorrectly", func(t *testing.T) {
		sqlDb, mock, _ := sqlmock.New()
		mock.ExpectQuery(`SELECT id, name, status FROM tasks`).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "status"}).
					AddRow(1, "test", 1).
					AddRow(2, "test", 0),
			)
		defer sqlDb.Close()

		taskRepo := NewTaskRepo(sqlDb)
		taskList, err := taskRepo.GetList(&dto.AppContext{GinContext: &gin.Context{}})
		assert.Nil(t, err)
		assert.NotNil(t, taskList)
		assert.Equal(t, len(*taskList), 2)
	})

	t.Run("shouldHandleErrorCorrectly", func(t *testing.T) {
		sqlDb, mock, _ := sqlmock.New()
		mock.ExpectQuery(`SELECT id, name, status FROM tasks`).
			WillReturnError(errors.New("test"))
		defer sqlDb.Close()

		taskRepo := NewTaskRepo(sqlDb)
		taskList, err := taskRepo.GetList(&dto.AppContext{GinContext: &gin.Context{}})
		assert.NotNil(t, err)
		assert.Nil(t, taskList)
	})
}

func TestTaskRepo_Add(t *testing.T) {
	t.Run("shouldAddCorrectly", func(t *testing.T) {
		sqlDb, mock, _ := sqlmock.New()
		mock.ExpectExec("INSERT INTO tasks (.+) VALUES (.+)").WithArgs("test").
			WillReturnResult(sqlmock.NewResult(1, 1))
		defer sqlDb.Close()

		c := &dto.AppContext{
			GinContext: &gin.Context{},
		}
		taskRepo := NewTaskRepo(sqlDb)
		id, err := taskRepo.Add(c, &entity.Task{Name: "test"})
		assert.Nil(t, err)
		assert.Equal(t, id, int64(1))
	})

	t.Run("shouldHandleErrorCorrectly", func(t *testing.T) {
		sqlDb, mock, _ := sqlmock.New()
		mock.ExpectExec("INSERT INTO tasks (.+) VALUES (.+)").WithArgs("test").
			WillReturnError(errors.New("test"))
		defer sqlDb.Close()

		c := &dto.AppContext{
			GinContext: &gin.Context{},
		}
		taskRepo := NewTaskRepo(sqlDb)
		id, err := taskRepo.Add(c, &entity.Task{Name: "test"})
		assert.NotNil(t, err)
		assert.Equal(t, id, int64(0))
	})
}

func TestTaskRepo_GetById(t *testing.T) {
	t.Run("shouldGetTaskCorrectly", func(t *testing.T) {
		sqlDb, mock, _ := sqlmock.New()
		mock.ExpectQuery("SELECT id, name, status FROM tasks WHERE id = ?").WithArgs(1).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "status"}).AddRow(1, "test", 0))
		defer sqlDb.Close()

		c := &dto.AppContext{
			GinContext: &gin.Context{},
		}
		taskRepo := NewTaskRepo(sqlDb)
		taskInfo, err := taskRepo.GetById(c, 1)
		assert.Nil(t, err)
		assert.NotNil(t, taskInfo)
		assert.Equal(t, taskInfo.Name, "test")
		assert.Equal(t, taskInfo.Status, 0)
	})

	t.Run("shouldHandleDataNotFoundErrorCorrectly", func(t *testing.T) {
		sqlDb, mock, _ := sqlmock.New()
		mock.ExpectQuery("SELECT id, name, status FROM tasks WHERE id = ?").WithArgs(1).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "status"}))
		defer sqlDb.Close()

		c := &dto.AppContext{
			GinContext: &gin.Context{},
		}
		taskRepo := NewTaskRepo(sqlDb)
		taskInfo, err := taskRepo.GetById(c, 1)
		assert.Nil(t, err)
		assert.Nil(t, taskInfo)
	})

	t.Run("shouldHandleErrorCorrectly", func(t *testing.T) {
		sqlDb, mock, _ := sqlmock.New()
		mock.ExpectQuery("SELECT id, name, status FROM tasks WHERE id = ?").WithArgs(1).
			WillReturnError(errors.New("test"))
		defer sqlDb.Close()

		c := &dto.AppContext{
			GinContext: &gin.Context{},
		}
		taskRepo := NewTaskRepo(sqlDb)
		taskInfo, err := taskRepo.GetById(c, 1)
		assert.NotNil(t, err)
		assert.Nil(t, taskInfo)
	})
}

func TestTaskRepo_UpdateById(t *testing.T) {
	t.Run("shouldUpdateCorrectly", func(t *testing.T) {
		sqlDb, mock, _ := sqlmock.New()
		mock.ExpectExec("UPDATE tasks").WithArgs("test", 0, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		defer sqlDb.Close()

		c := &dto.AppContext{
			GinContext: &gin.Context{},
		}
		taskRepo := NewTaskRepo(sqlDb)
		err := taskRepo.UpdateById(c, 1, &entity.Task{
			Name:   "test",
			Status: 0,
		})
		assert.Nil(t, err)
	})

	t.Run("shouldHandleErrorCorrectly", func(t *testing.T) {
		sqlDb, mock, _ := sqlmock.New()
		mock.ExpectExec("UPDATE tasks").WithArgs("test", 0, 1).
			WillReturnError(errors.New("test"))
		defer sqlDb.Close()

		c := &dto.AppContext{
			GinContext: &gin.Context{},
		}
		taskRepo := NewTaskRepo(sqlDb)
		err := taskRepo.UpdateById(c, 1, &entity.Task{
			Name:   "test",
			Status: 0,
		})
		assert.NotNil(t, err)
	})
}

func TestTaskRepo_DeleteById(t *testing.T) {
	t.Run("shouldDeleteCorrectly", func(t *testing.T) {
		sqlDb, mock, _ := sqlmock.New()
		mock.ExpectExec("DELETE FROM tasks WHERE id = ?").WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		defer sqlDb.Close()

		c := &dto.AppContext{
			GinContext: &gin.Context{},
		}
		taskRepo := NewTaskRepo(sqlDb)
		err := taskRepo.DeleteById(c, 1)
		assert.Nil(t, err)
	})

	t.Run("shouldHandleErrorCorrectly", func(t *testing.T) {
		sqlDb, mock, _ := sqlmock.New()
		mock.ExpectExec("DELETE FROM tasks WHERE id = ?").WithArgs(1).
			WillReturnError(errors.New("test"))
		defer sqlDb.Close()

		c := &dto.AppContext{
			GinContext: &gin.Context{},
		}
		taskRepo := NewTaskRepo(sqlDb)
		err := taskRepo.DeleteById(c, 1)
		assert.NotNil(t, err)
	})
}
