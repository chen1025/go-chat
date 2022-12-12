package v1

import (
	"fmt"
	"ginchat/utils"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

type attachController struct {
}

var Attach = new(attachController)

func (obj *attachController) Upload(c *gin.Context) {
	writer := c.Writer
	request := c.Request
	file, header, err := request.FormFile("file")
	if err != nil {
		utils.RespFail(writer, err.Error())
		return
	}
	suffix := ".png"
	fileName := header.Filename
	split := strings.Split(fileName, ".")
	if len(split) > 1 {
		suffix = "." + split[len(split)-1]
	}
	fileName = fmt.Sprintf("%d%04d%s", time.Now().Unix(), rand.Int31(), suffix)
	dist, err := os.Create("./asset/upload/" + fileName)
	if err != nil {
		utils.RespFail(writer, err.Error())
		return
	}
	_, err = io.Copy(dist, file)
	if err != nil {
		utils.RespFail(writer, err.Error())
		return
	}
	utils.RespOk(writer, "成功", "./asset/upload/"+fileName)
}
