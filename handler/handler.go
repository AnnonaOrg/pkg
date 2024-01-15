package handler

import (
	"net/http"

	"github.com/AnnonaOrg/pkg/errno"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 返回 Json
//
//	{
//	  "code": 0,
//	  "message": "",
//	  "data": {
//	    ...其他字段
//	  }
//	}
//
//	code: 返回 0，表示当前接口正确返回，否则按错误请求处理；
//	message: 返回接口处理信息，主要用于表单提交或请求失败时的 toast 显示；
//	data: 必须返回一个结构的对象。
func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	//always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

type ResponseEx struct {
	Status int `json:"status"`
	// Code   int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 返回 Json
//
//	{
//	  "status": 0,
//	  "msg": "",
//	  "data": {
//	    ...其他字段
//	  }
//	}
//
//	status: 返回 0，表示当前接口正确返回，否则按错误请求处理；
//	msg: 返回接口处理信息，主要用于表单提交或请求失败时的 toast 显示；
//	data: 必须返回一个结构的对象。
func SendResponseEx(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	//always return http.StatusOK
	c.JSON(http.StatusOK, ResponseEx{
		Status: code,
		// Code:   code,
		Msg:  message,
		Data: data,
	})
}

// Redirect 301
func SendRedirect(c *gin.Context, data string) {
	c.Redirect(http.StatusMovedPermanently, data)
}

// Redirect 302
func SendRedirect302(c *gin.Context, data string) {
	c.Redirect(http.StatusFound, data)
}

// String
func SendString(c *gin.Context, code int, data string) {
	//http.StatusOK
	c.String(code, data)
}
