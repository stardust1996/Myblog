package api

import (
	"MyBlog/service"
	"github.com/gin-gonic/gin"
)

func GetComments(c *gin.Context) {
	var commentsService service.CommentsService
	//s := sessions.Default(c)
	id := c.Query("bid")
	re := commentsService.GetCommentList(id)
	c.JSON(re.Code,gin.H{
		"resultData" : re.Data,
	})
}

func CreateComments(c *gin.Context) {
	var commentsService service.CommentsService
	if err := c.ShouldBind(&commentsService); err == nil {
		res := commentsService.CreatedComment()
		c.JSON(200,res)
	}else {
		c.JSON(200, ErrorResponse(err))
	}
}