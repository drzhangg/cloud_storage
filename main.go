package main

import (
	"cloud_storage/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()
	//router.StaticFile()
	router.LoadHTMLGlob("static/view/*")
	router.StaticFile("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	file := router.Group("file")
	{
		this := new(controller.Handler)
		file.POST("upload", this.UploadFile)
		file.POST("suc", this.UploadSucHandler)
	}

	router.Run(":9090")
}
