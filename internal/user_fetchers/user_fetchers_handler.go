package user_fetchers

import (
	"github.com/gin-gonic/gin"
	"lunarisnia/go-build-tags-experiment/internal/interfaces"
)

type UserFetcher struct {
	Foo string `json:"foo"`
}

func (u UserFetcher) GetExecutor() interfaces.BasicExecutor {
	return A{}
}

func (u UserFetcher) GetRequestBody(ctx *gin.Context) interface{} {
	us := UserFetcher{
		Foo: "ADasdasd",
	}

	return us
}

type A struct {
}

func (a A) JSONExecute(requestBody interface{}) gin.H {
	return gin.H{
		"foo": "bar",
	}
}
