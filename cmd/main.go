package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lunarisnia/go-build-tags-experiment/internal/interfaces"
	"lunarisnia/go-build-tags-experiment/internal/user_fetchers"
	"net/http"
)

// TODO: Find out how to call this dynamically
func initiateRequest[ST interfaces.BasicHandler[ST]](s ST) {
	s.GetRequestStruct()
}

func main() {
	r := gin.Default()

	r.POST("/execute", func(c *gin.Context) {
		userFetchers := user_fetchers.UserFetchers{}

		var destStruct any
		destStruct, err := userFetchers.GetRequestStruct()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Hello: ", destStruct)

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
