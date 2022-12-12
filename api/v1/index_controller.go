package v1

import (
	"ginchat/models"
	"ginchat/mvc/service"
	"github.com/gin-gonic/gin"
	"strconv"
	"text/template"
)

type indexController struct {
}

var Index = new(indexController)

// GetIndex
// @Tags 跳转首页
// @Accept json
// @Produce json
// @Success 200 {string} welcome
// @Router /index [get]
func (in *indexController) GetIndex(c *gin.Context) {
	/*c.JSON(200, gin.H{
		"message": "welcome",
	})*/
	ind, err := template.ParseFiles("index.html", "views/chat/head.html")
	if err != nil {
		panic(interface{}(err))
	}
	err = ind.Execute(c.Writer, "index")

}

// ToRegister
// @Tags 跳转注册页面
// @Accept json
// @Produce json
// @Success 200 {string} welcome
// @Router /toRegister [get]
func (in *indexController) ToRegister(c *gin.Context) {

	ind, err := template.ParseFiles("views/user/register.html")
	if err != nil {
		panic(interface{}(err))
	}
	err = ind.Execute(c.Writer, "register")

}

// ToChat
// @Tags 跳转聊天页面
// @Accept json
// @Produce json
// @Success 200 {string} welcome
// @Router /toChat [get]
func (in *indexController) ToChat(c *gin.Context) {

	ind, err := template.ParseFiles("views/chat/index.html",
		"views/chat/head.html",
		"views/chat/foot.html",
		"views/chat/tabmenu.html",
		"views/chat/concat.html",
		"views/chat/group.html",
		"views/chat/profile.html",
		"views/chat/createcom.html",
		"views/chat/userinfo.html",
		"views/chat/main.html")
	if err != nil {
		panic(interface{}(err))
	}
	userId, _ := strconv.Atoi(c.Query("userId"))
	token := c.Query("token")
	user := models.UserBasic{}
	user.ID = uint(userId)
	user.Identity = token

	ind.Execute(c.Writer, user)

}

func (in *indexController) Chat(c *gin.Context) {
	service.MessageService.Chat(c.Writer, c.Request)
}
