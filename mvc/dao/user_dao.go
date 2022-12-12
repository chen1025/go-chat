package dao

import (
	"ginchat/models"
	"ginchat/utils"
	"gorm.io/gorm"
)

type user struct {
}

var User = new(user)

// CreateUser 创建用户
func (us *user) CreateUser(basic models.UserBasic) *gorm.DB {
	//basic.ID = uint(utils.SN.Generate())
	return utils.DB.Create(&basic)
}

// DeleteUser 删除用户
func (us *user) DeleteUser(basic models.UserBasic) *gorm.DB {
	return utils.DB.Delete(&basic, basic.ID)
}

// UpdateUser 修改用户
func (us *user) UpdateUser(basic models.UserBasic) *gorm.DB {
	return utils.DB.Model(&basic).Updates(models.UserBasic{Icon: basic.Icon, Nickname: basic.Nickname, Email: basic.Email, Phone: basic.Phone})
}

// GetByName 根据名称查询用户
func (us *user) GetByName(name string) models.UserBasic {
	basic := models.UserBasic{}
	utils.DB.Where("name", name).First(&basic)
	return basic
}

// GetByPhone 根据手机号查询用户
func (us *user) GetByPhone(phone string) models.UserBasic {
	basic := models.UserBasic{}
	utils.DB.Where("phone", phone).First(&basic)
	return basic
}

// GetByEmail 根据邮箱查询用户
func (us *user) GetByEmail(email string) models.UserBasic {
	basic := models.UserBasic{}
	utils.DB.Where("email", email).First(&basic)
	return basic
}

// FindByUserId 根据id查询用户
func (us *user) FindByUserId(userId uint) models.UserBasic {
	basic := models.UserBasic{}
	utils.DB.Raw("select id,name,icon from user_basic where id=?", userId).Scan(&basic)
	return basic
}
