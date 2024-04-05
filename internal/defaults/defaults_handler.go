package defaults

import (
	"github.com/gin-gonic/gin"
	"lunarisnia/go-build-tags-experiment/internal/defaults/structs"
	"lunarisnia/go-build-tags-experiment/internal/interfaces"
)

type DefaultHandler struct{}

func (DefaultHandler) GetRequestBody(ctx *gin.Context) (interface{}, error) {
	return defaultstructs.DefaultRequestBody{
		Message: "pong",
	}, nil
}

func (DefaultHandler) GetExecutor() interfaces.BasicExecutor {
	return DefaultExecutor{}
}
