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
