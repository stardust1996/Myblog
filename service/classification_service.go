package service

import (
	"MyBlog/model"
	"MyBlog/serializer"
	"time"
)

type ClassificationService struct {
	KindName string    `json:"kind_name" form:"kind_name"`
}

func (service *ClassificationService) CreateClassification() serializer.Response {
	classification := model.Classification{
		Id:       model.GetNewId("02"),
		KindName: service.KindName,
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	if _, err := model.CreateClassification(classification); err != nil {
		return serializer.DBErr("连接数据库失败",err)
	}

	return serializer.BuildClassificationResponse(classification)
}

func (service *ClassificationService) GetAll() serializer.Response  {
	if list, err := model.GetAllClass(); err != nil {
		return serializer.DBErr("连接数据库失败",err)
	} else {
		return serializer.Response{
			Code:  200,
			Data:  list,
		}
	}
}

func (service *ClassificationService) DeleteClass(id string) serializer.Response {
	var count int
	if result := model.DB.Where("kind=?",id).Model(model.Blog{}).Count(&count);count != 0 {
		return serializer.ParamErr("该类型下现有文章，不能删除！",result.Error)
	}else {
		if eRow , err :=model.DeleteClassification(model.Classification{Id:       id,});err!=nil{
				return serializer.DBErr("连接数据失败",err)
		}else {
			return serializer.Response{
				Code:  200,
				Data:  eRow,
			}
		}
	}
}