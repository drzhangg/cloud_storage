package db

import (
	mydb "cloud_storage/db/mysql"
	"database/sql"
	"fmt"
)

// OpenFileUploadFinished: 文件上传成功，保存meta
func OpenFileUploadFinished(fileHash string, fileName string, fileSize int64, fileAddr string) bool {
	sql := `insert ignore into tbl_file (file_sha1,file_name,file_size,file_addr,status) values (?,?,?,?,1)`
	db := mydb.DBConn().Exec(sql, fileHash, fileName, fileSize, fileAddr)
	if db.Error != nil {
		fmt.Println("Failed to exec sql ,err :" + db.Error.Error())
		return false
	}

	if db.RowsAffected <= 0 {
		fmt.Printf("File with hash:%s has been upload before", fileHash)
	}

	fmt.Println(db.RowsAffected)

	return true
}

type TableFile struct {
	FileSha1 string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

//GetFileMeta ：从mysql获取文件元信息
func GetFileMeta(fileHash string) (*TableFile, error) {
	var (
		sql       string
		tableFile TableFile
	)

	sql = `select file_sha1,file_name,file_size,file_addr from tbl_file where file_sha1 = '%s' and status=1 limit 1;`
	sql = fmt.Sprintf(sql, fileHash)
	err := mydb.DBConn().Raw(sql).Scan(&tableFile).Error
	if err != nil {
		fmt.Println("GetFileMeta failed, err: ", err.Error())
		return nil, err
	}
	return &tableFile, nil
}
