package defaults

import (
	"errors"
	"github.com/gin-gonic/gin"
	defaultstructs "lunarisnia/go-build-tags-experiment/internal/defaults/structs"
)

type DefaultExecutor struct{}

func (DefaultExecutor) JSONExecute(requestBody interface{}) (gin.H, error) {
	defaultBody, ok := requestBody.(defaultstructs.DefaultRequestBody)
	if ok != true {
		return gin.H{}, errors.New("error")
	}

	return gin.H{
		"message": defaultBody.Message,
	}, nil
}
