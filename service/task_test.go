package service

import (
	"errors"
	"testing"

	mock_repo "ggl_test/mocks/models/repo"
	"ggl_test/models/dto"
	"ggl_test/models/entity"
	"ggl_test/utils/customerror"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestTaskService_GetAll(t *testing.T) {
	t.Run("shouldGetListCorrectly", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		appCtx := &dto.AppContext{
			GinContext: &gin.Context{},
		}

		mockTaskRepo := mock_repo.NewMockITaskRepo(ctrl)
		mockTaskRepo.EXPECT().GetList(gomock.Eq(appCtx)).Return(&[]entity.Task{
			{1, "test", true},
			{2, "test2", false},
		}, nil)

		taskSvc := NewTaskService(mockTaskRepo)
		taskList, err := taskSvc.GetAll(appCtx)
		assert.Nil(t, err)
		assert.NotNil(t, taskList)
		assert.Equal(t, len(*taskList), 2)
	})

	t.Run("shouldHandleErrorCorrectly", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		appCtx := &dto.AppContext{
			GinContext: &gin.Context{},
		}

		mockTaskRepo := mock_repo.NewMockITaskRepo(ctrl)
		mockTaskRepo.EXPECT().GetList(gomock.Eq(appCtx)).Return(nil, errors.New("test"))

		taskSvc := NewTaskService(mockTaskRepo)
		taskList, err := taskSvc.GetAll(appCtx)
		assert.NotNil(t, err)
		assert.Nil(t, taskList)
		assert.Equal(t, err.ErrorCode, customerror.ErrorCodeUnknown)
	})
}
