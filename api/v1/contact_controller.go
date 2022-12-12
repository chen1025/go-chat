package v1

import (
	"ginchat/models"
	service2 "ginchat/mvc/service"
	"ginchat/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type contactController struct {
}

var Contact = new(contactController)

// AddFriend
// @Summary 添加好友
// @Tags 添加好友
// @accept json
// @param userId formData int true "用户id"
// @param targetName formData string true "好友名称"
// @Success 200 {string} json{message,data} 成功
// @Router /contact/addFriend [post]
func (obj *contactController) AddFriend(c *gin.Context) {
	id := c.PostForm("userId")
	name := c.PostForm("targetName")
	uid, _ := strconv.Atoi(id)
	flag := service2.ContactService.AddFriend(uint(uid), name)
	if flag == 0 {
		utils.RespOk(c.Writer, "添加成功", nil)
	} else if flag == 1 {
		utils.RespFail(c.Writer, "添加失败，已经是你的好友")
	} else {
		utils.RespFail(c.Writer, "添加失败，用户不存在")
	}

}

// CreateGroup
// @Summary 创建群组
// @Tags 创建群组
// @accept json
// @param ownerId formData int true "用户id"
// @param name formData string true "群名称"
// @param icon formData string true "群头像"
// @Success 200 {string} json{message,data} 成功
// @Router /contact/createGroup [post]
func (obj *contactController) CreateGroup(c *gin.Context) {
	id := c.PostForm("ownerId")
	name := c.PostForm("name")
	icon := c.PostForm("icon")
	desc := c.PostForm("desc")
	uid, _ := strconv.Atoi(id)
	if uid == 0 || len(name) < 3 {
		utils.RespFail(c.Writer, "参数错误")
		return
	}
	group := models.GroupBasic{
		Name:        name,
		OwnerId:     uint(uid),
		Icon:        icon,
		Description: desc,
	}
	flag := service2.GroupService.CreateGroup(group)
	if flag == 0 {
		utils.RespOk(c.Writer, "创建成功", nil)
	} else {
		utils.RespFail(c.Writer, "创建失败")
	}

}

// FindGroup
// @Summary 查询群组
// @Tags 查询群组
// @accept json
// @param ownerId formData int true "用户id"
// @Success 200 {string} json{message,data} 成功
// @Router /contact/findGroup [post]
func (obj *contactController) FindGroup(c *gin.Context) {
	id := c.PostForm("ownerId")
	uid, _ := strconv.Atoi(id)
	data := service2.GroupService.FindGroup(uint(uid))
	utils.RespOkList(c.Writer, data, len(data))
}

// AddGroup
// @Summary 添加群组
// @Tags 添加群组
// @accept json
// @param comId formData int true "群id"
// @param userId formData int true "用户id"
// @Success 200 {string} json{message,data} 成功
// @Router /contact/addGroup [post]
func (obj *contactController) AddGroup(c *gin.Context) {
	id := c.PostForm("comId")
	u := c.PostForm("userId")
	gid, _ := strconv.Atoi(id)
	uid, _ := strconv.Atoi(u)
	flag := service2.GroupService.AddGroup(uint(uid), uint(gid))
	if flag == 0 {
		utils.RespOk(c.Writer, "添加成功", nil)
	} else if flag == 1 {
		utils.RespFail(c.Writer, "请勿重复添加")
	} else {
		utils.RespFail(c.Writer, "添加失败，群不存在")
	}
}
