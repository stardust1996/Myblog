package model

import (
	"time"
)

type Classification struct {
	//gorm.Model
	Id       string    `gorm:"column:id"`
	KindName string    `gorm:"column:kindName"`
	Created  time.Time `gorm:"column:created"`
	Updated  time.Time `gorm:"column:updated"`
}

//查找分类
func GetAllClass() (classes []Classification,err error)  {
	err = DB.Order("id").Find(&classes).Error
	return
}
//添加分类
func CreateClassification(classification Classification) (int64, error) {
	result := DB.Create(&classification)
	return result.RowsAffected,result.Error
}
//删除分类
func DeleteClassification(classification Classification) (int64, error) {
	result := DB.Delete(&classification)
	return result.RowsAffected,result.Error
}