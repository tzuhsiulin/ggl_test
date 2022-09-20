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
