package models

import (
	"main/dao"
)

type User struct {
	Id   int    `gorm:"primary_key;AUTO_INCREMENT;not null" json:"id"`
	Email  string    `gorm:"DEFAULT NULL;not null" json:"email"`
	Password string `gorm:"DEFAULT NULL;not null" json:"password"`
	//gorm后添加约束，json后为对应mysql里的字段
}


// 查询用户密码
func GetUser(email string) (id int,password string){
	var user User
	dao.MysqlDB.Where("email=?",email).First(&user)
	return user.Id,user.Password
}


// 创建新用户
func CreateUser(email,password string)  {
	user :=User{
		Email:email,
		Password:password,
	}
	dao.MysqlDB.Create(&user)
}
//更新用户信息
func UpdateUser(email,password string)  {
	var user User
	dao.MysqlDB.Model(&user).Where("email=?",email).Update("password",password)
	
}
//删除用户
func DeleteuUser(email,password string){
	i,p :=GetUser(email)
	user :=User{
		Id:i,
		Email:email,
		Password:p,
	}
	dao.MysqlDB.Delete(&user)
}

