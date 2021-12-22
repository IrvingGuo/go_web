package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success      bool        `json:"success"`
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"errorMessage"`
}

const (
	FAIL    = false
	SUCCESS = true
)

const (
	OK     = "ok"
	NOT_OK = "fail"
)

func Result(success bool, data interface{}, errMsg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		success,
		data,
		errMsg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "", c)
}

func OkWithData(c *gin.Context, data interface{}) {
	Result(SUCCESS, data, "", c)
}

func Fail(c *gin.Context) {
	Result(FAIL, map[string]interface{}{}, NOT_OK, c)
}

func FailWithMsg(c *gin.Context, msg string) {
	Result(FAIL, map[string]interface{}{}, msg, c)
}

func FailWithErr(c *gin.Context, err error) {
	Result(FAIL, map[string]interface{}{}, err.Error(), c)
}

func FailWithData(c *gin.Context, data interface{}) {
	Result(FAIL, data, NOT_OK, c)
}

func FailWithDataMsg(c *gin.Context, data interface{}, msg string) {
	Result(FAIL, data, msg, c)
}

func Abort(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(code, err.Error())
}

func AbortInternalServerError(c *gin.Context, err error) {
	Abort(c, http.StatusInternalServerError, err)
}

func AbortSaveFailed(c *gin.Context, err error) {
	Abort(c, http.StatusInternalServerError, err)
}

func AbortFindFailed(c *gin.Context, err error) {
	Abort(c, http.StatusInternalServerError, err)
}

func AbortDeleteFailed(c *gin.Context, err error) {
	Abort(c, http.StatusInternalServerError, err)
}

func AbortBadRequest(c *gin.Context, err error) {
	Abort(c, http.StatusBadRequest, err)
}

// 403
func AbortForbidden(c *gin.Context, err error) {
	Abort(c, http.StatusForbidden, err)
}

func AbortNotImplementedMethod(c *gin.Context) {
	Abort(c, http.StatusInternalServerError, fmt.Errorf("not implemented method"))
}

func MiddlewareAbortWithJson(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusOK, Response{
		FAIL,
		map[string]interface{}{},
		msg,
	})
}
