package main

import (
	"cloud_storage/handler"
	"net/http"
)

func main() {

	//router := gin.Default()
	////router.StaticFile()
	//router.LoadHTMLGlob("static/view/*")
	//router.Static("/static", "./static")
	//
	//router.GET("/", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "home.html", gin.H{})
	//})
	//
	//file := router.Group("file")
	//{
	//	this := new(controller.Handler)
	//	file.POST("upload", this.UploadFile)
	//	file.POST("suc", this.UploadSucHandler)
	//}

	//router.Run(":9090")

	http.HandleFunc("/file/upload", handler.UploadHandle)          //上传文件
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandle)   //上传文件成功
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)      //通过filehash获取文件信息
	http.HandleFunc("/file/query", handler.FileQueryHandler)       //批量查询文件
	http.HandleFunc("/file/download", handler.DownloadHandler)     //下载文件
	http.HandleFunc("/file/update", handler.FileMetaUploadHandler) //更新文件
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)     //删除文件

	http.HandleFunc("/user/signup", handler.SingupHandler)    //用户注册
	http.HandleFunc("/user/signin", handler.SiginInHandler)   //用户登录
	http.HandleFunc("/user/toSignin", handler.RedirectSignin) //重定向到登录页面

	http.ListenAndServe(":9090", nil)
}
