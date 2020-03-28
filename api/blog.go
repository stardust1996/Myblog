package api

import (
	"MyBlog/model"
	"MyBlog/service"
	"github.com/gin-gonic/gin"
	"strconv"
)
//获取所有博客并返回到主页
func GetList(c *gin.Context) {
	var blogService service.BlogService
	var page model.Page
	edit := c.Query("edit")
	page.CurrentPage, _ = strconv.Atoi(c.DefaultQuery("pNum", "0"))//主页列表查询
	page.PageSize = 5
	re := blogService.ShowList(page,edit)
	if edit != "" {
		c.JSON(200,re)//返回json格式数据
	}else {//返回到主页
		c.HTML(re.Code,"index.html",gin.H{
			"result" : re.Data,
		})
	}

}
//获取博客并返回
func GetBlogById(c *gin.Context) {
	var blogService service.BlogService
	id := c.Param("id")
	edit := c.Query("edit")
	re := blogService.GetBlog(c,id)
	if edit != "" {//返回json格式数据
		c.JSON(200, re)
	}else {//返回到查看页面
		c.HTML(re.Code,"blog.html",gin.H{
			"resultData" : re.Data,
		})
	}
}
//获取博客标题列表
func GetBlogTitles(c *gin.Context) {
	var blogService service.BlogService
	re := blogService.GetTitleList()
	c.JSON(re.Code,gin.H{
		"titleList" :re.Data,
	})
}
//创建博客
func CreateBlog(c *gin.Context) {
	var blogService service.BlogService
	if err := c.ShouldBind(&blogService); err == nil{
		res := blogService.CreateBlog()
		c.JSON(200,res)
	}else {
		c.JSON(200, ErrorResponse(err))
	}
}
//删除博客
func DeleteBlog(c *gin.Context) {
	var blogService service.BlogService
	id := c.Param("id")
	re := blogService.DeleteBlog(id)
	c.JSON(200,re)
}

//更新修改博客
func UpdateBlog(c *gin.Context) {
	var blogService service.BlogService
	if err := c.ShouldBind(&blogService); err == nil{
		res := blogService.UpdateBlog()
		c.JSON(200,res)
	}else {
		c.JSON(200, ErrorResponse(err))
	}
}