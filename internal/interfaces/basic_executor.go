package interfaces

import "github.com/gin-gonic/gin"

type BasicExecutor interface {
	JSONExecute(requestBody interface{}) (gin.H, error)
}
