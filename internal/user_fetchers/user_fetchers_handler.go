package user_fetchers

import (
	"fmt"
)

type UserFetchers struct {
}

type UserFetchersRequestStruct struct {
	Foo string `json:"foo"`
}

func (u UserFetchers) GetRequestStruct() (UserFetchersRequestStruct, error) {
	fmt.Println("Run")
	us := UserFetchersRequestStruct{
		Foo: "ADasdasd",
	}

	return us, nil
}
