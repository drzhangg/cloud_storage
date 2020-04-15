package db

import (
	"cloud_storage/conf"
	db "cloud_storage/db/mysql"
	"fmt"
)

// OpenFileUploadFinished: 文件上传成功，保存meta
func OpenFileUploadFinished(fileHash string, fileName string, fileSize int64, fileAddr string) bool {
	sql := `insert ignore into %s (file_sha1,file_name,file_size,file_addr,'status') values (%s,%s,%d,%s,1)`
	sql = fmt.Sprintf(sql, conf.Tbl_file, fileHash, fileName, fileSize, fileAddr)
	//db.DBConn().Prepare(sql)
	result, err := db.DBConn().Exec(sql)
	if err != nil {
		fmt.Println("Failed to exec sql ,err :" + err.Error())
		return false
	}

	if rf, err := result.RowsAffected(); err != nil {
		if rf <= 0 {
			fmt.Printf("File with hash:%s has been upload before", fileHash)
		}
		return true
	}
	return false
}
