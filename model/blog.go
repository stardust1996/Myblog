package model

import (
	"time"
)

type Blog struct {
	//gorm.Model
	//Page
	Id      string `gorm:"column:id"`
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content;type:text"`
	Kind    string `gorm:"column:kind"`
	Hits    int    `gorm:"column:hits"`
	Created time.Time `gorm:"column:created"`
	Updated time.Time `gorm:"column:updated"`
}

type BlogShow struct {
	//gorm.Model
	//Page
	Id      string `gorm:"column:id"`
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content;type:text"`
	Kind    string `gorm:"column:kind"`
	KindName string `gorm:"column:kindName"`
	Hits    int    `gorm:"column:hits"`
	Created time.Time `gorm:"column:created"`
	Updated time.Time `gorm:"column:updated"`
}


//根据blog的id获取详细信息
func GetBlog(ID interface{}) (BlogShow,error) {
	blogList := make([]BlogShow,0)
	result := DB.Table("blogs as b").Select("b.id, b.title, b.content, b.kind, c.kindName, b.hits, b.created, b.updated").
		Joins("left join classifications c on b.kind = c.id").Where("b.id=?",ID).Scan(&blogList)
	return blogList[0],result.Error
}
//管理界面获取blog列表
func GetBlogShowList() ([]BlogShow,error) {
	blogList := make([]BlogShow,0)
	result := DB.Table("blogs as b").Select("b.id, b.title, b.content, b.kind, c.kindName, b.hits, b.created, b.updated").Joins("left join classifications c on b.kind = c.id").Scan(&blogList)
	return blogList,result.Error
}
//获取blog列表
func GetBlogList(page Page) ([]Blog,error) {
	blogList := make([]Blog,0)
	result := DB.Limit(page.PageSize).Offset((page.CurrentPage-1)*page.PageSize).Find(&blogList)
	return blogList,result.Error
}
//添加博客
func CreateBlog(blog Blog) (int64, error) {
	result := DB.Create(&blog)
	return result.RowsAffected, result.Error
}
//删除博客
func DeleteBlog(blog Blog) (int64, error) {
	result := DB.Delete(&blog)
	return result.RowsAffected, result.Error
}
//更新博客
func UpdateBlog(blog Blog) (int64, error) {
	result := DB.Model(&blog).Update(blog)
	return result.RowsAffected, result.Error
}
//获取数量
func GetTotalCount() (int, error)  {
	var total int
	result := DB.Model(Blog{}).Count(&total)
	return total,result.Error
}
//获取标题
func GetBlogTitles() ([]Blog, error) {
	//var titles []string
	blogList := make([]Blog,0)
	result := DB.Limit(10).Find(&blogList)
	return blogList,result.Error
	//rows,err := DB.Raw("select title from blogs limit 10").Rows()

	/*if err != nil {
		return titles,err
	}else {
		for rows.Next()  {
			var title string
			_ = rows.Scan(&title)
			titles = append(titles, title)
		}
		return titles,err
	}*/
}