package mysql

import (
	"cloud_storage/conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
)

var db *gorm.DB

func init() {

	db, _ = gorm.Open("mysql", conf.DataSource)

	db.DB().SetMaxOpenConns(1000) //设置连接池大小

	db.SingularTable(true) //关闭表名复数形式

	//连接测试
	err := db.DB().Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql,err:" + err.Error())
		os.Exit(1)
	}
}

// 全局数据库连接对象
func DBConn() *gorm.DB {
	return db
}
