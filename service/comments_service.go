package service

import (
	"MyBlog/model"
	"MyBlog/serializer"
	"time"
)

type CommentsService struct {
	//Id       string    `json:"id" form:"id"`
	Content  string    `json:"content" form:"content"`
	Superior string    `json:"superior" form:"superior"`
	SuperiorUser string `json:"superior_user" form:"superior_user"`
	Uid      string    `json:"uid" form:"uid"`
	Bid      string    `json:"bid" form:"bid"`
}

func (service *CommentsService) GetCommentList(ID interface{}) serializer.Response {
	//是否应该加入对评论的归类,将内容进行排序,及评论嫁接在评论下
	if list,err := model.GetByBid(ID);err != nil {
		return serializer.DBErr("连接数据库失败",err)
	} else {
		return serializer.Response{
			Code:  200,
			Data:  list,
		}
	}
}

func (service *CommentsService)CreatedComment() serializer.Response {
	comment := model.Comments{
		Id:       model.GetNewId("03"),
		Content:  service.Content,
		Superior: service.Superior,
		SuperiorUser: service.SuperiorUser,
		Uid:      service.Uid,
		Bid:      service.Bid,
		Created:  time.Now(),
	}


	if  _,err := model.CreateComments(comment); err != nil{
		return serializer.DBErr("连接数据库失败",err)
	}
	return serializer.BuildCommentsResponse(comment)
}