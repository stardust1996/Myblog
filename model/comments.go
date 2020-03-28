package model

import (
	"time"
)

type Comments struct {
	//gorm.Model
	//Page
	Id       string    `gorm:"column:id"`
	Content  string    `gorm:"column:content"`
	Superior string    `gorm:"column:superior"`
	SuperiorUser string    `gorm:"column:superiorUser"`
	Uid      string    `gorm:"column:uid"`
	Bid      string    `gorm:"column:bid"`
	Created  time.Time `gorm:"column:created"`
}

type result struct {
	CommentId string `json:"comment_id"`
	CommentContent string `json:"comment_content"`
	CommentSuperior string `json:"comment_superior"`
	CommentSuperiorUserId string `json:"comment_superior_user_id"`
	CommentUsername string `json:"comment_username"`
	CommentUserId string `json:"comment_user_id"`
	CommentSuperiorUser string `json:"comment_superior_user"`
	CommentBid string `json:"comment_bid"`
	CommentCreated string `json:"comment_created"`
}

//根据文章ID获取评论
func GetByBid(ID interface{}) (comList []result,err error) {
	//err = DB.Where("bid = ?",ID).Order("created").Find(&comList).Error
	//rows,err := DB.Table("comments").Joins("left join user on user.id = comments.uid")
	/* DB.Raw("select c.id as comment_id, c.content as comment_content, c.superior as comment_superior, " +
		"u2.username as comment_username, u1.username as comment_superior_user, c.bid as comment_bid, c.created as comment_created	" +
		"from comments c left join users u1 on c.uid = u1.id  left join  users u2 on c.superiorUser = u2.id " +
		"where c.bid = '" + ID +"' order by c.created").Scan(&comList)*/
	err = DB.Table("comments as c").Select("c.id as comment_id, c.content as comment_content, c.superior as comment_superior," +
		"u1.username as comment_username,u1.id as comment_user_id, u2.username as comment_superior_user, u2.id as comment_superior_user_id," +
		"c.bid as comment_bid, c.created as comment_created").
		Joins("left join users u1 on u1.id = c.uid").Joins("left join users u2 on u2.id = c.superiorUser where c.bid =?",ID).Order("c.created").Scan(&comList).Error

	return
}
//创建新的评论
func CreateComments(comments Comments) (aRows int64, err error) {
	result := DB.Create(&comments)
	aRows = result.RowsAffected
	err = result.Error
	return
}
//删除评论
func DeleteComments(comments Comments) (aRows int64, err error) {
	result := DB.Delete(&comments)
	aRows = result.RowsAffected
	err = result.Error
	return
}