package serializer

import (
	"MyBlog/model"
	"time"
)

type Comments struct {
	model.Page
	Id       string    `json:"id"`
	Content  string    `json:"content"`
	Superior string    `json:"superior"`
	SuperiorUser string `json:"superior_user"`
	Uid      string    `json:"uid"`
	Bid      string    `json:"bid"`
	Created  time.Time `json:"created"`
}

func BuildComments(comments model.Comments) Comments {
	return Comments{
		/*Page:     Page{
			PageSize:      comments.PageSize,
			CurrentPage:   comments.CurrentPage,
			TotalCount:    comments.TotalCount,
			TotalPagesNum: comments.TotalPagesNum,
		},*/
		Id:       comments.Id,
		Content:  comments.Content,
		Superior: comments.Superior,
		SuperiorUser:comments.SuperiorUser,
		Uid:      comments.Uid,
		Bid:      comments.Bid,
		Created:  comments.Created,
	}
}

func BuildCommentsResponse(comments model.Comments) Response {
	return Response{
		Code:  200,
		Data:  BuildComments(comments),
	}
}