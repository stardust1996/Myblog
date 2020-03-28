package api

import (
	"MyBlog/service"
	"github.com/gin-gonic/gin"
)

func GetClasses(c *gin.Context) {
	var class service.ClassificationService
	re := class.GetAll()
	c.JSON(re.Code,gin.H{
		"resultData" : re.Data,
	})
}

func CreateClass(c *gin.Context) {
	var class service.ClassificationService
	if err := c.ShouldBind(&class); err == nil {
		re := class.CreateClassification()
		c.JSON(200,re)
	}else {
		c.JSON(200, ErrorResponse(err))
	}
}

func DeleteClass(c *gin.Context) {
	var class  service.ClassificationService
	id := c.Param("id")
	rep  := class.DeleteClass(id)
	c.JSON(200,rep)
}