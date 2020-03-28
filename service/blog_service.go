package service

import (
	"MyBlog/model"
	"MyBlog/serializer"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"time"
)

type BlogService struct {
	Id      string   `form:"id" json:"id"`
	Title   string	 `form:"title" json:"title" binding:"required,min=5,max=30"`
	Content string	 `form:"content" json:"content" binding:"required"`
	Kind    string	 `form:"kind" json:"kind" binding:"required"`
	//Hits    int    `json:"hits" form:"hits" binding:"required"`
}
//获取全部博客内容
func (service *BlogService) ShowList(page model.Page,edit string) serializer.Response {
	//var blog model.Blog
	if  edit !="" {
		if re,err :=model.GetBlogShowList();err!=nil{
			return serializer.DBErr("读取数据库失败",err)
		}else {
			return serializer.Response{
				Code:  200,
				Data:  re,
			}
		}
	}else {
		result := make(map[interface{}]interface{})
		if page.CurrentPage == 0 && page.TotalCount == 0 {
			page.CurrentPage = 1
		}
		if total,err :=  model.GetTotalCount();err != nil{
			return serializer.DBErr("读取数量失败",err)
		}else {
			page.TotalCount = total
			if page.TotalCount%page.PageSize == 0 {
				page.TotalPagesNum =  page.TotalCount/page.PageSize
			} else {
				page.TotalPagesNum = page.TotalCount/page.PageSize + 1
			}
		}
		result["page"] = page
		if list,err := model.GetBlogList(page);err != nil{
			return serializer.DBErr("读取数据库失败",err)
		}else {
			result["bList"] = list
		}
		return serializer.Response{
			Code:  200,
			Data:  result,
		}
	}

}
//根据ID获取博客内容
func (service *BlogService)  GetBlog(c *gin.Context,ID interface{}) serializer.Response {
	if blog, err := model.GetBlog(ID);err!=nil {
		return serializer.DBErr("读取数据库操作失败",err)
	}else {

	//	service.setSession(c,blog)

		return serializer.BuildBlogShowResponse(blog)
	}
}
//获取标题列表
func (service *BlogService) GetTitleList() serializer.Response {
	if titleList,err := model.GetBlogTitles();err!=nil{
		return serializer.DBErr("读取数据库操作失败",err)
	}else {
		return serializer.Response{
			Code:  200,
			Data:  titleList,
		}
	}
}
//创建博客
func (service *BlogService) CreateBlog() serializer.Response {

	blog := model.Blog{
		Id:      model.GetNewId("01"),
		Title:   service.Title,
		Content: service.Content,
		Kind:    service.Kind,
		Hits:    0,
		Created: time.Now(),
		Updated: time.Now(),
	}

	if _, err := model.CreateBlog(blog); err != nil {
		return serializer.ParamErr("博客储存失败",err)
	}

	return serializer.BuildBlogResponse(blog)
}

//删除博客
func (service *BlogService) DeleteBlog(id string) serializer.Response {
	if eRow , err :=model.DeleteBlog(model.Blog{Id:       id,});err!=nil{
		return serializer.DBErr("连接数据失败",err)
	}else {
		//删除文章下的评论
		model.DB.Where("bid=?",id).Delete(model.Comments{})

		return serializer.Response{
			Code:  200,
			Data:  eRow,
		}
	}
}

//修改更新博客
func (service *BlogService)UpdateBlog() serializer.Response {
	blog := model.Blog{
		Id:      service.Id,
		Title:   service.Title,
		Content: service.Content,
		Kind:    service.Kind,
		Hits:    0,
		Updated: time.Now(),
	}

	if _, err := model.UpdateBlog(blog); err != nil {
		return serializer.ParamErr("博客更新失败",err)
	}

	return serializer.BuildBlogResponse(blog)
}


//设置session
func (service *BlogService) setSession(c *gin.Context, blog model.Blog) {
	s := sessions.Default(c)
//	s.Clear()
	s.Set("blog_id", blog.Id)
	_ = s.Save()
}