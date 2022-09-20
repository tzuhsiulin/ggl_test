package repo

import (
	"errors"
	"testing"

	"ggl_test/models/dto"
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
