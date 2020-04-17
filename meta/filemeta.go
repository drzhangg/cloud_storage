package meta

import (
	"cloud_storage/db"
	"sort"
)

/**
文件元信息：
*/

type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

// UpdateFileMeta: 新增/更新文件元信息
func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

//UpdateFileMetaDB：新增/更新文件元信息到mysql中
func UpdateFileMetaDB(fmeta FileMeta) bool {
	return db.OpenFileUploadFinished(fmeta.FileSha1, fmeta.FileName, fmeta.FileSize, fmeta.Location)
}

// GetFileMeta:通过sha1值获取文件的元信息对象
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

//GetFileMetaDB：从mysql中获取文件元信息
func GetFileMetaDB(fileSha1 string) (FileMeta, error) {
	tableFile, err := db.GetFileMeta(fileSha1)
	if err != nil {
		return FileMeta{}, nil
	}

	return FileMeta{
		FileSha1: tableFile.FileSha1,
		FileName: tableFile.FileName.String,
		FileSize: tableFile.FileSize.Int64,
		Location: tableFile.FileAddr.String,
	}, nil
}

// GetLastFileMeta：获取批量的文件元信息列表
func GetLastFileMeta(count int) []FileMeta {
	fileMetaArr := make([]FileMeta, len(fileMetas))
	for _, v := range fileMetas {
		fileMetaArr = append(fileMetaArr, v)
	}

	sort.Sort(ByUploadTime(fileMetaArr))
	return fileMetaArr[0:count]
}

// RemoveFileMeta：删除元信息
func RemoveFileMeta(fileSha1 string) {
	delete(fileMetas, fileSha1)
}
