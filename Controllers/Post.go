package Controllers

import (
	"net/http"

	"github.com/michaelpege/arc/Models"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)


func GetAllPosts(c *gin.Context){
	var posts []Models.Post
	err := Models.GetAllPosts(&posts)
	if err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, posts)
	}
}

//get specific posts
func GetPost(c *gin.Context){
	id := c.Param("id")
	var post Models.Post
	err := Models.GetPostById(&post, id)
		if err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, post)
	}
}