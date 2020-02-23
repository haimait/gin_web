package model

import (
	"fmt"
)

type User struct {
	ID           int64
	Name         string `gorm:"default:'小王子'"` // sql.NullString 实现了Scanner/Valuer接口
	Age          int64
}


func (u *User) CreateUser() bool{
	if err := Db.Create(u).Error; err != nil {
		fmt.Println(err)
		return false
	}
	return true
}


func (u *User) GetUserDetailById() (user User){
	err := Db.First(&user,u.ID).Error
	if(err!=nil){
		fmt.Println(err)
	}
	return
}

func (u *User) GetUserGetUserList() (user []User){
	//err := Db.Find(&user).Error
	str := "1 or 1=1"
	tx := Db
	tx = tx.Where("name = ?",str)
	tx.Find(&user)
	return
}



//
//func (u *user) GetAdDetailById(Id int64) (ads user) {
//	Db.Model(u).Where("id = ?", Id).First(&ads)
//	//db := Db.Model(a)
//	//db = db.Where("id = ?", Id)
//	//db = db.First(&ads)
//	return
//}


//func (a *Ads) AdUpdate2(data map[string]interface{}) bool {
//	if err := Db.Model(a).Update(data).Error; err != nil {
//		return false
//	}
//	return true
//}
