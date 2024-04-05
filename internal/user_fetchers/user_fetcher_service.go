package userfetchers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	userfetcherstructs "lunarisnia/go-build-tags-experiment/internal/user_fetchers/structs"
	"net/http"
)

type UserFetcherService struct{}

func (UserFetcherService) JSONExecute(requestBody interface{}) (gin.H, error) {
	body, ok := requestBody.(userfetcherstructs.UserFetcherRequest)
	if !ok {
		return gin.H{}, errors.New("error")
	}

	response, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/users?username=%s", body.Username))
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response")
		}
	}(response.Body)
	if err != nil {
		return gin.H{}, err
	}

	decoder := json.NewDecoder(response.Body)

	type UserInfo struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
	}
	var userInfos []UserInfo
	err = decoder.Decode(&userInfos)
	if err != nil {
		return gin.H{}, err
	}

	var responseBody []interface{}
	inRequest, err := json.Marshal(userInfos)
	if err != nil {
		return gin.H{}, err
	}
	if err := json.Unmarshal(inRequest, &responseBody); err != nil {
		return gin.H{}, err
	}
	return gin.H{
		"users": responseBody,
	}, nil
}
