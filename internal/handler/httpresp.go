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
)

var StatusMessageMap = map[int]string{
	StatusOk:                "ok",
	StatusBindingError:      "request param wrong",
	StatusDBError:           "something wrong when query in db",
	StatusNoImageFoundError: "no image find the condition",
}

type UniformResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func respOK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, UniformResponse{
		StatusCode: StatusOk,
		Message:    StatusMessageMap[StatusOk],
		Data:       data,
	})
}

func respParamBindingError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, UniformResponse{
		StatusCode: StatusBindingError,
		Message:    StatusMessageMap[StatusBindingError],
		Data:       err,
	})
}

func respDBError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, UniformResponse{
		StatusCode: StatusDBError,
		Message:    StatusMessageMap[StatusDBError],
		Data:       err,
	})
}

func respNotFoundError(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, UniformResponse{
		StatusCode: StatusNoImageFoundError,
		Message:    StatusMessageMap[StatusNoImageFoundError],
		Data:       nil,
	})
}
