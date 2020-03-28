package serializer

import (
	"MyBlog/model"
	"time"
)

type Blog struct {
	model.Page
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Kind    string `json:"kind"`
	KindName string `gorm:"column:kingName"`
	Hits    int    `json:"hits"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`

}

// BuildBlog 序列化博客
func BuildBlogShow(blog model.BlogShow) Blog {
	return Blog{
		/*Page: Page{
			PageSize:      blog.PageSize,
			CurrentPage:   blog.CurrentPage,
			TotalCount:    blog.TotalCount,
			TotalPagesNum: blog.TotalPagesNum,
		},*/
		ID:blog.Id,
		Title : blog.Title,
		Content:blog.Content,
		Kind:blog.Kind,
		KindName:blog.KindName,
		Hits:blog.Hits,
		Created:blog.Created,
		Updated:blog.Updated,
	}
}

// BuildUserResponse 序列化博客响应
func BuildBlogShowResponse(blog model.BlogShow) Response {
	return Response{
		Code: 200,
		Data: BuildBlogShow(blog),
	}
}

// BuildBlog 序列化博客
func BuildBlog(blog model.Blog) Blog {
	return Blog{
		/*Page: Page{
			PageSize:      blog.PageSize,
			CurrentPage:   blog.CurrentPage,
			TotalCount:    blog.TotalCount,
			TotalPagesNum: blog.TotalPagesNum,
		},*/
		ID:blog.Id,
		Title : blog.Title,
		Content:blog.Content,
		Kind:blog.Kind,
		Hits:blog.Hits,
		Created:blog.Created,
		Updated:blog.Updated,
	}
}

// BuildUserResponse 序列化博客响应
func BuildBlogResponse(blog model.Blog) Response {
	return Response{
		Code: 200,
		Data: BuildBlog(blog),
	}
}