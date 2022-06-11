package models

import (
	"main/dao"
)

type File struct {
	Id   int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name string `gorm:"not null" json:"name"`
	Size int64 `gorm:"not null" json:"size"`
	Path string `gorm:"not null" json:"path"`
	Username string `gorm:"not null" json:"username"`
	Uploaddate string `gorm:"not null" json:"uploaddate"`
}

func DeleteFile(id int){
	var file File
	dao.MysqlDB.Where("id=?",id).Delete(&file)
}

func FindFilebyname(username,name string)(id int,path string){
	var file File
	dao.MysqlDB.Where("username=? and name=?",username,name).Find(&file)
		return file.Id,file.Path
}
//查询当前登录用户文件
func FindFIle(username string)(rfile []File){
	var file []File
	dao.MysqlDB.Where("username=?",username).Find(&file)
	return file
}
//上传文件
func CreateFile(name string,path string,size int64,username string,uploaddate string){
	file :=File{
		Name:name,
		Size:size,
		Path:path,
		Username:username,
		Uploaddate:uploaddate,
	}
	dao.MysqlDB.Create(&file)

}
