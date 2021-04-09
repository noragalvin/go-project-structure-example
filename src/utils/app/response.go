package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gempages-admin-server/src/utils/logger"
)

// Gin struct
type Gin struct {
	C *gin.Context
}

// Response struct
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode int, data interface{}) {
	// g.C.JSON(httpCode, Response{
	// 	Code: httpCode,
	// 	Msg:  e.GetMsg(httpCode),
	// 	Data: data,
	// })
	g.C.JSON(httpCode, data)
}

// Ok setting gin.JSON
func (g *Gin) Ok(data interface{}) {
	g.Response(http.StatusOK, data)
}

// Error setting gin.JSON
func (g *Gin) Error(data interface{}) {
	logger.Logger.Error(data)
	g.Response(http.StatusInternalServerError, data)
}

// BadRequest setting gin.JSON
func (g *Gin) BadRequest(data interface{}) {
	logger.Logger.Error(data)
	g.Response(http.StatusBadRequest, data)
}

// Unauthorized setting gin.JSON
func (g *Gin) Unauthorized(data interface{}) {
	logger.Logger.Error(data)
	g.Response(http.StatusUnauthorized, data)
}

// ResponseMessage get response message
func ResponseMessage(message string) map[string]interface{} {
	return map[string]interface{}{"message": message}
}
