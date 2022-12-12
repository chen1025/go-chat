package service

import (
	"ginchat/models"
	"ginchat/mvc/dao"
	"gorm.io/gorm"
)

type userService struct {
}

var UserService = new(userService)

// CreateUser 创建用户
func (us *userService) CreateUser(basic models.UserBasic) *gorm.DB {
	//basic.ID = uint(utils.SN.Generate())
	return dao.User.CreateUser(basic)
}

// DeleteUser 删除用户
func (us *userService) DeleteUser(basic models.UserBasic) *gorm.DB {
	return dao.User.DeleteUser(basic)
}

// UpdateUser 修改用户
func (us *userService) UpdateUser(basic models.UserBasic) *gorm.DB {
	return dao.User.UpdateUser(basic)
}

// GetByName 根据名称查询用户
func (us *userService) GetByName(name string) models.UserBasic {
	return dao.User.GetByName(name)
}

// GetByPhone 根据手机号查询用户
func (us *userService) GetByPhone(phone string) models.UserBasic {
	return dao.User.GetByPhone(phone)
}

// GetByEmail 根据邮箱查询用户
func (us *userService) GetByEmail(email string) models.UserBasic {
	return dao.User.GetByEmail(email)
}

// FindByUserId 根据id查询用户
func (us *userService) FindByUserId(userId uint) models.UserBasic {
	return dao.User.FindByUserId(userId)
}
