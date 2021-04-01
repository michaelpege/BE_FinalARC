package Controllers

import (
	"net/http"

	"github.com/michaelpege/arc/Models"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)


func GetComment(c *gin.Context) {
	postid := c.Param("id")
	var comments []Models.Comment
	err := Models.GetCommentsByPostId(&comments, postid)
	if err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, comments)
	}
}

func PostComment(c *gin.Context) {
	postid := c.Param("id")
	var comment Models.Comment
	c.BindJSON(&comment)
	err := Models.PostCommentByPostId(&comment, postid)
	if err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, comment)
	}
}