package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	StatusOk                = 0
	StatusBindingError      = 1
	StatusDBError           = 2
	StatusNoImageFoundError = 3
	StatusNotAuthError      = 4
	StatusInternalError     = 5
	StatusS3Error           = 6
)

// var StatusMessageMap = map[int]string{
// 	StatusOk:                "ok",
// 	StatusBindingError:      "request param wrong",
// 	StatusDBError:           "something wrong when query in db",
// 	StatusNoImageFoundError: "no image find the condition",
// 	StatusNotAuthError:      "not auth",
// 	StatusInternalError:     "internal server error",
// 	StatusS3Error:           "something wrong when using s3 sdk",
// }

type UniformResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func respOK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, UniformResponse{
		StatusCode: StatusOk,
		Message:    "ok",
		Data:       data,
	})
}

func respParamBindingError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, UniformResponse{
		StatusCode: StatusBindingError,
		Message:    err.Error(),
	})
}

func respDBError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, UniformResponse{
		StatusCode: StatusDBError,
		Message:    err.Error(),
	})
}

func respS3Error(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, UniformResponse{
		StatusCode: StatusS3Error,
		Message:    err.Error(),
	})
}

func respNotFoundError(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, UniformResponse{
		StatusCode: StatusNoImageFoundError,
		Message:    "no image found",
	})
}

func respUnknownError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, UniformResponse{
		StatusCode: StatusInternalError,
		Message:    err.Error(),
	})
}

func respAbortWithForbiddenError(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusForbidden, UniformResponse{
		StatusCode: StatusNotAuthError,
		Message:    err.Error(),
	})
}
