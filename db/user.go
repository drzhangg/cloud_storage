package db

import (
	db "cloud_storage/db/mysql"
	"cloud_storage/model"
	"fmt"
)

//UserSingnup：通过用户名和密码完成user表的注册操作
func UserSingnup(username string, password string) bool {
	var (
		sql string
	)

	sql = `insert ignore into tbl_user (user_name,user_pwd) value ("%s","%s")`
	sql = fmt.Sprintf(sql, username, password)

	stmt := db.DBConn().Exec(sql)
	if stmt.Error != nil {
		fmt.Println("Failed to insert,err:" + stmt.Error.Error())
		return false
	}

	if stmt.RowsAffected <= 0 {
		return false
	}
	return true
}

// UserSignin：用户登录
func UserSignin(username, encpwd string) bool {
	var (
		sql  string
		user model.User
		err  error
	)

	sql = `select user_name,user_pwd from tbl_user where user_name = "%s" and user_pwd = "%s" limit 1`
	sql = fmt.Sprintf(sql, username, encpwd)
	if err = db.DBConn().Raw(sql).Scan(&user).Error; err != nil {
		fmt.Println("Failed to select,err:", err)
		return false
	}

	fmt.Println("user:-----",user)

	return false
}
