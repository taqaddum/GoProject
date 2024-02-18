package component

import (
	"github.com/gin-gonic/gin"
	"sync"
)

var onceRouter sync.Once
var router *gin.Engine

func NewRouter() *gin.Engine {
	onceRouter.Do(func() {
		router = gin.Default()
	})

	return router
}
