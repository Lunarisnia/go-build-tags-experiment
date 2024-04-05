package interfaces

import "github.com/gin-gonic/gin"

type BasicHandler interface {
	GetRequestBody(ctx *gin.Context) interface{}
	GetExecutor() BasicExecutor
}
