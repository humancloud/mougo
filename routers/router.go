package routers

import (
	"fmt"
	"mougo/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// Set Mode debug or release.
	gin.SetMode(utils.AppMode)

	mougo := gin.Default()

	// add router to router group.
	router := mougo.Group("api/version1")
	{
		router.GET("hello", hello) // api/version_x/hello
	}

	return mougo
}

func hello(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Hello mougo!")
}
