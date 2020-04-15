package mysql

import (
	"cloud_storage/conf"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var db *sql.DB

func init() {

	db, err := sql.Open("mysql", conf.DataSource)
	if err != nil {
		log.Fatal("gorm.Open failed, err:", err)
		return
	}

	db.SetMaxOpenConns(1000) //设置连接池大小

	//连接测试
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql,err:" + err.Error())
		os.Exit(1)
	}
}

// 全局数据库连接对象
func DBConn() *sql.DB {
	return db
}
