package Routes

import (
	"github.com/michaelpege/arc/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	router := gin.Default()

	//get all posts
	router.GET("/posts", Controllers.GetAllPosts)

	//get specific post
	router.GET("/post/:id", Controllers.GetPost)

	//get commment from specific post
	router.GET("/comments/:id", Controllers.GetComment) 

	//create comment for specific post
	router.POST("/comments/:id", Controllers.PostComment)

	router.Run(":5000")

	return router
}