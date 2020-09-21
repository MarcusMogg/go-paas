package api

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// uploadFile 上传文件 filePath最后需要有一个/
func uploadFile(filePath string, c *gin.Context) (string, string, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return "", "", err
	}
	os.MkdirAll(filePath, os.ModePerm)
	_, suf := getFileInfo(file.Filename)
	fileName := fmt.Sprintf("%d", time.Now().Unix())
	if err := c.SaveUploadedFile(file, filePath+fileName+suf); err != nil {
		return "", "", err
	}

	return fileName, suf, nil
}

// getFileInfo 拆分文件名和格式
func getFileInfo(fileName string) (string, string) {
	tmp := strings.Split(fileName, ".")
	tmpLen := len(tmp)
	return strings.Join(tmp[:tmpLen-1], "."), "." + tmp[tmpLen-1]
}
