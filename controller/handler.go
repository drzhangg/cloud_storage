package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

type Handler struct {
}

func (*Handler) UploadFile(ctx *gin.Context) {
	file, handler, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("get file err: %s", err))
		return
	}

	//获取文件名
	fileName := handler.Filename
	//写入文件
	out, err := os.Create("data_file/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	filepath := "http://127.0.0.1:8000/file/" + fileName

	ctx.JSON(http.StatusOK, gin.H{
		"filepath": filepath,
	})
}
