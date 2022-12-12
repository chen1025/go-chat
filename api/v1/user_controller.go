package v1

import (
	"ginchat/models"
	service2 "ginchat/mvc/service"
	"ginchat/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type userController struct {
}

var User = new(userController)

// CreateUser
// @Tags 注册
// @Summary 创建用户
// @param name formData string true "用户名"
// @param password formData string true "密码"
// @Success 200 {string} json{message,data}
// @Router /user/createUser [post]
func (uc *userController) CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.PostForm("name")
	pwd := c.PostForm("password")
	if user.Name == "" || pwd == "" {
		c.JSON(200, gin.H{
			"message": "参数不能为空",
			"code":    -1,
		})
		return
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(pwd), 0)
	user.Password = string(password)
	data := service2.UserService.GetByName(user.Name)
	if data.Name != "" {
		c.JSON(200, gin.H{
			"message": "用户已存在",
			"code":    -1,
		})
		return
	}
	service2.UserService.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "注册成功",
		"code":    0,
	})
}

// DeleteUser
// @Tags 注销
// @Summary 删除用户
// @param id query int true "id"
// @Success 200 {string} json{message,data}
// @Router /user/deleteUser [get]
func (uc *userController) DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	service2.UserService.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "注销成功",
		"code":    0,
	})
}

// swaggerignore:"true" 忽略字段
type UpdateModel struct {
	// id 必填
	ID       uint   `json:"id" example:"1" swaggertype:"integer"`
	Icon     string `json:"icon" swaggertype:"string"`
	Nickname string `json:"nickname" swaggertype:"string"`
	Phone    string `json:"phone" example:"15611111111" swaggertype:"string" `
	Email    string `json:"email" example:"abc@qq.com" swaggertype:"string" `
}

// UpdateUser
// @Summary 修改用户
// @Tags 修改用户
// @accept json
// @param id formData int true "用户id"
// @param icon formData string true "头像"
// @param name formData string true "名称"
// @Success 200 {string} json{message,data} 成功
// @Router /user/updateUser [post]
func (uc *userController) UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Nickname = c.PostForm("name")
	user.Icon = c.PostForm("icon")
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	validateStruct, err := govalidator.ValidateStruct(user)
	if err != nil || !validateStruct {
		c.JSON(200, gin.H{
			"message": "格式错误",
			"code":    -1,
		})
		return
	}
	service2.UserService.UpdateUser(user)
	c.JSON(200, gin.H{
		"message": "修改成功",
		"code":    0,
	})
}

type LoginModel struct {
	// name 必填
	Name string `json:"name" example:"account1" swaggertype:"string" `
	// 密码必填
	Password string `json:"password" example:"ppp" swaggertype:"string"`
}

// Login
// @Summary 登录
// @Tags 登录
// @accept json
// @param name formData string true "用户名"
// @param password formData string true "密码"
// @Success 200 {string} json{message,data} 成功
// @Router /user/login [post]
func (uc *userController) Login(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")

	u := service2.UserService.GetByName(user.Name)
	// 用户是否存在
	if u.Name == "" {
		c.JSON(200, gin.H{
			"message": "用户不存在或密码错误",
			"code":    -1,
		})
		return
	}
	// 密码是否正确
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "用户不存在或密码错误",
			"code":    -1,
		})
		return
	}
	u.Password = ""
	c.JSON(200, gin.H{
		"message": "登录成功",
		"data":    u,
		"code":    0,
	})
}

// 防止跨站点伪造请求
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// SearchFriend
// @Tags 查询好友
// @Summary 查询好友
// @param userId formData int true "用户ID"
// @Success 200 {string} json{message,data}
// @Router /user/searchFriend [post]
func (uc *userController) SearchFriend(c *gin.Context) {
	query := c.PostForm("userId")
	id, _ := strconv.Atoi(query)
	friend := service2.ContactService.SearchFriend(uint(id))
	utils.RespOkList(c.Writer, friend, len(friend))
}

// FindByUserId
// @Summary 根据id查询用户
// @Tags 根据id查询用户
// @accept json
// @param userId formData int true "用户id"
// @Success 200 {string} json{message,data} 成功
// @Router /user/FindByUserId [post]
func (uc *userController) FindByUserId(c *gin.Context) {
	id := c.PostForm("userId")
	uid, _ := strconv.Atoi(id)
	data := service2.UserService.FindByUserId(uint(uid))
	utils.RespOk(c.Writer, "查询成功", data)

}

// GetMsgByRedis
// @Summary 查询redis的消息
// @Tags 查询redis的消息
// @accept json
// @param userIdA formData int true "用户id"
// @param userIdB formData int true "发消息着id"
// @param start formData int true "开始位置"
// @param end formData int true "结束位置"
// @param isRev formData bool true "用户id"
// @Success 200 {string} json{message,data} 成功
// @Router /user/getMsgByRedis [post]
func (uc *userController) GetMsgByRedis(c *gin.Context) {
	uidA, _ := strconv.Atoi(c.PostForm("userIdA"))
	uidB, _ := strconv.Atoi(c.PostForm("userIdB"))
	start, _ := strconv.Atoi(c.PostForm("start"))
	end, _ := strconv.Atoi(c.PostForm("end"))
	isRev, _ := strconv.ParseBool(c.PostForm("isRev"))

	data := service2.MessageService.GetMsgByRedis(uidA, uidB, start, end, isRev)
	utils.RespOkList(c.Writer, "查询成功", data)
}
