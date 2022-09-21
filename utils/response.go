package utils

import (
	"net/http"

	"ggl_test/models/dto"
	"ggl_test/utils/customerror"
	"github.com/gin-gonic/gin"
)

var httpStatusCodeMapping = map[int]int{
	customerror.ErrorCodeUnknown:      http.StatusInternalServerError,
	customerror.ErrorCodeInvalidParam: http.StatusBadRequest,
}

var customErrMessage = map[int]string{
	customerror.ErrorCodeUnknown:      "unknown error",
	customerror.ErrorCodeInvalidParam: "invalid params",
}

func Response(c *dto.AppContext, obj interface{}, args ...int) {
	statusCode := http.StatusOK
	if len(args) > 0 {
		statusCode = http.StatusOK
	}
	if obj == nil {
		obj = gin.H{}
	}
	c.GinContext.JSON(statusCode, obj)
}

func ResponseError(c *dto.AppContext, err *customerror.CustomError) {
	errCode := err.ErrorCode
	errMsg := err.ErrorMsg
	httpStatusCode := http.StatusBadRequest

	if val, ok := httpStatusCodeMapping[errCode]; ok {
		httpStatusCode = val
	} else {
		if errCode >= 4000 && errCode < 5000 {
			httpStatusCode = http.StatusBadRequest
		} else if errCode >= 5000 && errCode < 6000 {
			httpStatusCode = http.StatusInternalServerError
		}
	}

	if len(errMsg) == 0 {
		if msg, ok := customErrMessage[errCode]; ok {
			errMsg = msg
		} else {
			errMsg = "unknown error"
		}
	}

	c.GinContext.JSON(httpStatusCode, &dto.CommonErrorResponse{
		Status:  "error",
		ErrCode: errCode,
		ErrMsg:  errMsg,
	})
}
