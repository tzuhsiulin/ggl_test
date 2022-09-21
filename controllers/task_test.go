package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"ggl_test/config"
	mock_service "ggl_test/mocks/service"
	"ggl_test/models/dto"
	"ggl_test/models/entity"
	"ggl_test/utils/customerror"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestTaskController_GetTaskList(t *testing.T) {
	t.Run("shouldGetListCorrectly", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		appCfg := config.GetAppCfg()
		taskList := &[]entity.Task{
			{1, "test", 1},
			{2, "test2", 0},
		}

		mockTaskSvc := mock_service.NewMockITaskService(ctrl)
		mockTaskSvc.EXPECT().GetAll(gomock.Any()).Return(taskList, nil)

		taskCtrl := NewTaskController(appCfg, mockTaskSvc)
		router := gin.Default()
		router.GET("/tasks", taskCtrl.GetTaskList)
		r := httptest.NewRequest(http.MethodGet, "/tasks", bytes.NewReader([]byte{}))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)

		var resp dto.GetTaskListResp
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Nil(t, err)
		assert.Equal(t, len(resp.Result), 2)
	})

	t.Run("shouldHandleErrorCorrectly", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		appCfg := config.GetAppCfg()
		customErr := customerror.NewErr(customerror.ErrorCodeUnknown)

		mockTaskSvc := mock_service.NewMockITaskService(ctrl)
		mockTaskSvc.EXPECT().GetAll(gomock.Any()).Return(nil, customErr)

		taskCtrl := NewTaskController(appCfg, mockTaskSvc)
		router := gin.Default()
		router.GET("/tasks", taskCtrl.GetTaskList)
		r := httptest.NewRequest(http.MethodGet, "/tasks", bytes.NewReader([]byte{}))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)

		var resp dto.CommonErrorResponse
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Nil(t, err)
		assert.Equal(t, resp.Status, "error")
		assert.Equal(t, resp.ErrCode, customerror.ErrorCodeUnknown)
		assert.Equal(t, resp.ErrMsg, "unknown error")
	})
}

func TestTaskController_Add(t *testing.T) {
	t.Run("shouldAddCorrectly", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		appCfg := config.GetAppCfg()
		taskInfo := &entity.Task{Id: 1, Name: "test", Status: 1}
		reqPayload := &dto.CreateTaskReq{
			Name: "test",
		}

		mockTaskSvc := mock_service.NewMockITaskService(ctrl)
		mockTaskSvc.EXPECT().Add(gomock.Any(), gomock.Any()).Return(taskInfo, nil)

		taskCtrl := NewTaskController(appCfg, mockTaskSvc)
		router := gin.Default()
		router.POST("/task", taskCtrl.Add)
		data, _ := json.Marshal(reqPayload)
		r := httptest.NewRequest(http.MethodPost, "/task", bytes.NewReader(data))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)

		var resp dto.CreateTaskResp
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Nil(t, err)
		assert.Equal(t, resp.Result.Id, int64(1))
		assert.Equal(t, resp.Result.Name, "test")
		assert.Equal(t, resp.Result.Status, 1)
	})

	t.Run("shouldHandleInvalidParamErrorCorrectly", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		appCfg := config.GetAppCfg()
		reqPayload := &dto.CreateTaskReq{
			Name: "",
		}

		mockTaskSvc := mock_service.NewMockITaskService(ctrl)

		taskCtrl := NewTaskController(appCfg, mockTaskSvc)
		router := gin.Default()
		router.POST("/task", taskCtrl.Add)
		data, _ := json.Marshal(reqPayload)
		r := httptest.NewRequest(http.MethodPost, "/task", bytes.NewReader(data))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)

		var resp dto.CommonErrorResponse
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Nil(t, err)
		fmt.Println(resp)
		assert.Equal(t, resp.Status, "error")
		assert.Equal(t, resp.ErrCode, customerror.ErrorCodeInvalidParam)
		assert.Equal(t, resp.ErrMsg, "`name` cannot be empty")
	})

	t.Run("shouldHandleAddErrorCorrectly", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		appCfg := config.GetAppCfg()
		reqPayload := &dto.CreateTaskReq{
			Name: "test",
		}
		customErr := customerror.NewErr(customerror.ErrorCodeUnknown)

		mockTaskSvc := mock_service.NewMockITaskService(ctrl)
		mockTaskSvc.EXPECT().Add(gomock.Any(), gomock.Any()).Return(nil, customErr)

		taskCtrl := NewTaskController(appCfg, mockTaskSvc)
		router := gin.Default()
		router.POST("/task", taskCtrl.Add)
		data, _ := json.Marshal(reqPayload)
		r := httptest.NewRequest(http.MethodPost, "/task", bytes.NewReader(data))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)

		var resp dto.CommonErrorResponse
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Nil(t, err)
		assert.Equal(t, resp.Status, "error")
		assert.Equal(t, resp.ErrCode, customerror.ErrorCodeUnknown)
		assert.Equal(t, resp.ErrMsg, "unknown error")
	})
}

func TestTaskController_Update(t *testing.T) {
	t.Run("shouldUpdateCorrectly", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		appCfg := config.GetAppCfg()
		taskId := int64(1)
		reqPayload := &dto.UpdateTaskReqData{
			Name:   "test",
			Status: 1,
		}

		mockTaskSvc := mock_service.NewMockITaskService(ctrl)
		mockTaskSvc.EXPECT().Update(gomock.Any(), gomock.Eq(taskId), gomock.Any()).Return(&entity.Task{
			Id:     1,
			Name:   "test",
			Status: 1,
		}, nil)

		taskCtrl := NewTaskController(appCfg, mockTaskSvc)
		router := gin.Default()
		router.PUT("/task/:id", taskCtrl.Update)
		data, _ := json.Marshal(reqPayload)
		r := httptest.NewRequest(http.MethodPut, "/task/1", bytes.NewReader(data))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)

		var resp dto.UpdateTaskResp
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Nil(t, err)
		assert.Equal(t, resp.Result.Id, int64(1))
		assert.Equal(t, resp.Result.Name, "test")
		assert.Equal(t, resp.Result.Status, 1)
	})

	t.Run("shouldHandleInvalidPathParamCorrectly", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		appCfg := config.GetAppCfg()
		reqPayload := &dto.UpdateTaskReqData{
			Name:   "test",
			Status: 1,
		}

		mockTaskSvc := mock_service.NewMockITaskService(ctrl)

		taskCtrl := NewTaskController(appCfg, mockTaskSvc)
		router := gin.Default()
		router.PUT("/task/:id", taskCtrl.Update)
		data, _ := json.Marshal(reqPayload)
		r := httptest.NewRequest(http.MethodPut, "/task/www", bytes.NewReader(data))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)

		var resp dto.CommonErrorResponse
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Nil(t, err)
		assert.Equal(t, resp.ErrCode, customerror.ErrorCodeInvalidParam)
	})

	t.Run("shouldHandleInvalidParamCorrectly", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		appCfg := config.GetAppCfg()
		reqPayload := &dto.UpdateTaskReqData{
			Name:   "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttest",
			Status: 2,
		}

		mockTaskSvc := mock_service.NewMockITaskService(ctrl)

		taskCtrl := NewTaskController(appCfg, mockTaskSvc)
		router := gin.Default()
		router.PUT("/task/:id", taskCtrl.Update)
		data, _ := json.Marshal(reqPayload)
		r := httptest.NewRequest(http.MethodPut, "/task/1", bytes.NewReader(data))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)

		var resp dto.CommonErrorResponse
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Nil(t, err)
		assert.Equal(t, resp.ErrCode, customerror.ErrorCodeInvalidParam)
	})

	t.Run("shouldHandleUpdateErrorCorrectly", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		appCfg := config.GetAppCfg()
		taskId := int64(1)
		reqPayload := &dto.UpdateTaskReqData{
			Name:   "test",
			Status: 1,
		}

		mockTaskSvc := mock_service.NewMockITaskService(ctrl)
		mockTaskSvc.EXPECT().Update(gomock.Any(), gomock.Eq(taskId), gomock.Any()).Return(
			nil, customerror.NewErr(customerror.ErrorCodeUnknown))

		taskCtrl := NewTaskController(appCfg, mockTaskSvc)
		router := gin.Default()
		router.PUT("/task/:id", taskCtrl.Update)
		data, _ := json.Marshal(reqPayload)
		r := httptest.NewRequest(http.MethodPut, "/task/1", bytes.NewReader(data))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)

		var resp dto.CommonErrorResponse
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Nil(t, err)
		assert.Equal(t, resp.Status, "error")
		assert.Equal(t, resp.ErrCode, customerror.ErrorCodeUnknown)
	})
}

func TestTaskController_Delete(t *testing.T) {
	t.Run("shouldDeleteCorrectly", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		appCfg := config.GetAppCfg()

		mockTaskSvc := mock_service.NewMockITaskService(ctrl)
		mockTaskSvc.EXPECT().Delete(gomock.Any(), gomock.Eq(int64(1))).Return(nil)

		taskCtrl := NewTaskController(appCfg, mockTaskSvc)
		router := gin.Default()
		router.DELETE("/task/:id", taskCtrl.Delete)
		r := httptest.NewRequest(http.MethodDelete, "/task/1", bytes.NewReader([]byte{}))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)

		assert.Equal(t, w.Code, http.StatusOK)
		assert.Equal(t, w.Body.String(), "{}")
	})

	t.Run("shouldHandleInvalidPathParamErrorCorrectly", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		appCfg := config.GetAppCfg()

		mockTaskSvc := mock_service.NewMockITaskService(ctrl)

		taskCtrl := NewTaskController(appCfg, mockTaskSvc)
		router := gin.Default()
		router.DELETE("/task/:id", taskCtrl.Delete)
		r := httptest.NewRequest(http.MethodDelete, "/task/dfdf", bytes.NewReader([]byte{}))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)

		var resp dto.CommonErrorResponse
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Nil(t, err)
		assert.Equal(t, w.Code, http.StatusBadRequest)
		assert.Equal(t, resp.ErrCode, customerror.ErrorCodeInvalidParam)
	})

	t.Run("shouldHandleInvalidPathParamErrorCorrectly", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		appCfg := config.GetAppCfg()

		mockTaskSvc := mock_service.NewMockITaskService(ctrl)
		mockTaskSvc.EXPECT().Delete(gomock.Any(), gomock.Eq(int64(1))).
			Return(customerror.NewErr(customerror.ErrorCodeUnknown))

		taskCtrl := NewTaskController(appCfg, mockTaskSvc)
		router := gin.Default()
		router.DELETE("/task/:id", taskCtrl.Delete)
		r := httptest.NewRequest(http.MethodDelete, "/task/1", bytes.NewReader([]byte{}))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)

		var resp dto.CommonErrorResponse
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Nil(t, err)
		assert.Equal(t, w.Code, http.StatusInternalServerError)
		assert.Equal(t, resp.ErrCode, customerror.ErrorCodeUnknown)
	})
}
