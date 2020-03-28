package serializer

import (
	"MyBlog/model"
	"time"
)

type Classification struct {
	Id       string    `json:"id"`
	KindName string    `json:"kind_name"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}


//序列化分类
func BuildClassification(classification model.Classification) Classification {
	return Classification{
		Id:       classification.Id,
		KindName: classification.KindName,
		Created:  classification.Created,
		Updated:  classification.Updated,
	}
}

//序列化分类响应
func BuildClassificationResponse(classification model.Classification) Response {
	return Response{
		Data: BuildClassification(classification),
	}
}