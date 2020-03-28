package model

import (
	"math/rand"
	"strconv"
	"time"
)

//判断ID是否存在
func IsExist(rType string,id string) bool{
	count := 0
	switch rType {
	case "01":
		DB.Model(Blog{}).Where("id=?",id).Count(&count)
		break
	case "02":
		DB.Model(Classification{}).Where("id=?",id).Count(&count)
		break
	case "03":
		DB.Model(Comments{}).Where("id=?",id).Count(&count)
		break
	}

	if count > 0 {
		return true
	}
	return false
}

// RandStringRunes 返回随机字符串
func GetNewId(rType string) (newID string) {
LOOP:newID = time.Now().Format("20060102")
	newID += rType

	rand.Seed(time.Now().UnixNano())
	var randString string
	for ; len(randString) != 4;  {
		randString = strconv.Itoa(rand.Intn(10000))
	}
	newID += randString
	if IsExist(rType, newID) {
		goto LOOP
	}
	return
}

func GetNewUserId(rType string) (newID string) {
LOOP:newID = time.Now().Format("20060102")
	newID = rType + newID

	rand.Seed(time.Now().UnixNano())
	var randString string
	for ; len(randString) != 2;  {
		randString = strconv.Itoa(rand.Intn(100))
	}
	newID += randString
	count := 0
	DB.Model(User{}).Where("id=?",newID).Count(&count)
	if count > 0 {
		goto LOOP
	}
	return
}