package userfetchers

import (
	"github.com/gin-gonic/gin"
	"lunarisnia/go-build-tags-experiment/internal/interfaces"
	userfetcherstructs "lunarisnia/go-build-tags-experiment/internal/user_fetchers/structs"
)

type UserFetcherHandler struct {
}

func (u UserFetcherHandler) GetExecutor() interfaces.BasicExecutor {
	return UserFetcherService{}
}

func (u UserFetcherHandler) GetRequestBody(ctx *gin.Context) (interface{}, error) {
	var requestBody userfetcherstructs.UserFetcherRequest
	err := ctx.BindJSON(&requestBody)
	if err != nil {
		return userfetcherstructs.UserFetcherRequest{}, err
	}

	return requestBody, nil
}
