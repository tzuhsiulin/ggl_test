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
			{1, "test", 1},
			{2, "test2", 0},
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

func TestTaskService_Add(t *testing.T) {
	t.Run("shouldAddCorrectly", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		appCtx := &dto.AppContext{GinContext: &gin.Context{}}
		addTaskPayload := &entity.Task{Name: "test"}
		taskId := int64(1)

		mockTaskRepo := mock_repo.NewMockITaskRepo(ctrl)
		mockTaskRepo.EXPECT().Add(gomock.Eq(appCtx), gomock.Eq(addTaskPayload)).Return(taskId, nil)
		mockTaskRepo.EXPECT().GetById(gomock.Eq(appCtx), gomock.Eq(taskId)).Return(&entity.Task{
			Id:     1,
			Name:   "test",
			Status: 1,
		}, nil)

		taskSvc := NewTaskService(mockTaskRepo)
		taskInfo, err := taskSvc.Add(appCtx, addTaskPayload)
		assert.Nil(t, err)
		assert.NotNil(t, taskInfo)
		assert.Equal(t, taskInfo.Id, int64(1))
		assert.Equal(t, taskInfo.Name, "test")
		assert.Equal(t, taskInfo.Status, 1)
	})

	t.Run("shouldHandleAddErrorCorrectly", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		appCtx := &dto.AppContext{GinContext: &gin.Context{}}
		addTaskPayload := &entity.Task{Name: "test"}

		mockTaskRepo := mock_repo.NewMockITaskRepo(ctrl)
		mockTaskRepo.EXPECT().Add(gomock.Eq(appCtx), gomock.Eq(addTaskPayload)).Return(int64(0), errors.New("test"))

		taskSvc := NewTaskService(mockTaskRepo)
		taskInfo, err := taskSvc.Add(appCtx, addTaskPayload)
		assert.NotNil(t, err)
		assert.Nil(t, taskInfo)
	})

	t.Run("shouldHandleGetInfoErrorCorrectly", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		appCtx := &dto.AppContext{GinContext: &gin.Context{}}
		addTaskPayload := &entity.Task{Name: "test"}
		taskId := int64(1)

		mockTaskRepo := mock_repo.NewMockITaskRepo(ctrl)
		mockTaskRepo.EXPECT().Add(gomock.Eq(appCtx), gomock.Eq(addTaskPayload)).Return(taskId, nil)
		mockTaskRepo.EXPECT().GetById(gomock.Eq(appCtx), gomock.Eq(taskId)).Return(nil, errors.New("test"))

		taskSvc := NewTaskService(mockTaskRepo)
		taskInfo, err := taskSvc.Add(appCtx, addTaskPayload)
		assert.NotNil(t, err)
		assert.Nil(t, taskInfo)
	})
}

func TestTaskService_Update(t *testing.T) {
	t.Run("shouldUpdateCorrectly", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		appCtx := &dto.AppContext{GinContext: &gin.Context{}}
		updateTaskPayload := &entity.Task{Name: "test", Status: 1}
		taskId := int64(1)

		mockTaskRepo := mock_repo.NewMockITaskRepo(ctrl)
		mockTaskRepo.EXPECT().GetById(gomock.Eq(appCtx), gomock.Eq(taskId)).Return(
			&entity.Task{Id: 1, Name: "test", Status: 1}, nil)
		mockTaskRepo.EXPECT().UpdateById(gomock.Eq(appCtx), gomock.Eq(taskId), gomock.Eq(updateTaskPayload)).Return(nil)
		mockTaskRepo.EXPECT().GetById(gomock.Eq(appCtx), gomock.Eq(taskId)).Return(
			&entity.Task{Id: 1, Name: "test", Status: 1}, nil)

		taskSvc := NewTaskService(mockTaskRepo)
		taskInfo, err := taskSvc.Update(appCtx, taskId, updateTaskPayload)
		assert.Nil(t, err)
		assert.NotNil(t, taskInfo)
		assert.Equal(t, taskInfo.Id, taskId)
		assert.Equal(t, taskInfo.Name, "test")
		assert.Equal(t, taskInfo.Status, 1)
	})

	t.Run("shouldHandleDataNotFoundError", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		appCtx := &dto.AppContext{GinContext: &gin.Context{}}
		updateTaskPayload := &entity.Task{Name: "test", Status: 1}
		taskId := int64(1)

		mockTaskRepo := mock_repo.NewMockITaskRepo(ctrl)
		mockTaskRepo.EXPECT().GetById(gomock.Eq(appCtx), gomock.Eq(taskId)).Return(nil, nil)

		taskSvc := NewTaskService(mockTaskRepo)
		taskInfo, err := taskSvc.Update(appCtx, taskId, updateTaskPayload)
		assert.NotNil(t, err)
		assert.Nil(t, taskInfo)
		assert.Equal(t, err.ErrorCode, customerror.ErrorCodeInvalidParam)
	})

	t.Run("shouldHandleUpdateErrorCorrectly", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		appCtx := &dto.AppContext{GinContext: &gin.Context{}}
		updateTaskPayload := &entity.Task{Name: "test", Status: 1}
		taskId := int64(1)

		mockTaskRepo := mock_repo.NewMockITaskRepo(ctrl)
		mockTaskRepo.EXPECT().GetById(gomock.Eq(appCtx), gomock.Eq(taskId)).Return(
			&entity.Task{Id: 1, Name: "test", Status: 1}, nil)
		mockTaskRepo.EXPECT().UpdateById(gomock.Eq(appCtx), gomock.Eq(taskId), gomock.Eq(updateTaskPayload)).Return(errors.New("test"))

		taskSvc := NewTaskService(mockTaskRepo)
		taskInfo, err := taskSvc.Update(appCtx, taskId, updateTaskPayload)
		assert.NotNil(t, err)
		assert.Nil(t, taskInfo)
	})

	t.Run("shouldHandleGetInfoErrorCorrectly", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		appCtx := &dto.AppContext{GinContext: &gin.Context{}}
		updateTaskPayload := &entity.Task{Name: "test", Status: 1}
		taskId := int64(1)

		mockTaskRepo := mock_repo.NewMockITaskRepo(ctrl)
		mockTaskRepo.EXPECT().GetById(gomock.Eq(appCtx), gomock.Eq(taskId)).Return(
			&entity.Task{Id: 1, Name: "test", Status: 1}, nil)
		mockTaskRepo.EXPECT().UpdateById(gomock.Eq(appCtx), gomock.Eq(taskId), gomock.Eq(updateTaskPayload)).Return(nil)
		mockTaskRepo.EXPECT().GetById(gomock.Eq(appCtx), gomock.Eq(taskId)).Return(nil, errors.New("test"))

		taskSvc := NewTaskService(mockTaskRepo)
		taskInfo, err := taskSvc.Update(appCtx, taskId, updateTaskPayload)
		assert.NotNil(t, err)
		assert.Nil(t, taskInfo)
	})
}
