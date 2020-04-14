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

	http.HandleFunc("/file/upload", handler.UploadHandle)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandle)

	http.ListenAndServe(":9090", nil)
}
