package dao

import(
	// "database/sql"
	_"github.com/go-sql-driver/mysql"
	// "log"
	// "fmt"
	"github.com/jinzhu/gorm"
	
)
var MysqlDB *gorm.DB



func Initdb()  {
	var err error
	MysqlDB, err = gorm.Open("mysql", "user:Uplooking_123@tcp(localhost:3306)/login")
	if err != nil {
		MysqlDB.Close()
		panic(err)
	}
	MysqlDB.SingularTable(true)
	// 设置连接池，空闲连接
	MysqlDB.DB().SetMaxIdleConns(50)
	// 打开链接
	MysqlDB.DB().SetMaxOpenConns(100)

}



























