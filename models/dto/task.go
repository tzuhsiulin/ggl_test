package dto

import "ggl_test/models/entity"

type GetTaskListResp struct {
	Result []entity.Task `json:"result"`
}

type CreateTaskReq struct {
	Name string `json:"name" binding:"required,min=1,max=64"`
}

type CreateTaskResp struct {
	Result entity.Task `json:"result"`
}

type UpdateTaskReq struct {
	Path struct {
		Id int64 `uri:"id" binding:"required"`
	}
	Data struct {
		Id     int64  `json:"id"`
		Name   string `json:"name" binding:"required"`
		Status int    `json:"status" binding:"min=0,max=1"`
	}
}

type UpdateTaskResp struct {
	Result entity.Task `json:"result"`
}

type DeleteTaskReq struct {
	Id int64 `uri:"id"`
}
