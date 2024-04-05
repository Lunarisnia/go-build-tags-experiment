package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lunarisnia/go-build-tags-experiment/internal/defaults"
	"lunarisnia/go-build-tags-experiment/internal/interfaces"
	"net/http"
)

var handler interfaces.BasicHandler

func init() {
	handler = defaults.DefaultHandler{}
}

func main() {
	r := gin.Default()
	r.POST("/execute", func(c *gin.Context) {
		requestBody, err := handler.GetRequestBody(c)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		executor := handler.GetExecutor()

		responseBody, err := executor.JSONExecute(requestBody)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, responseBody)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
