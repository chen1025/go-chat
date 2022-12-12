package service

import (
	"ginchat/models"
	"ginchat/mvc/dao"
	"ginchat/utils"
	"gorm.io/gorm"
)

type groupService struct {
}

var GroupService = new(groupService)

func (gs *groupService) CreateGroup(group models.GroupBasic) int {

	err := utils.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&group).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		con := models.Contact{}
		con.OwnerId = group.OwnerId
		con.TargetId = group.ID
		con.Type = 2
		if err := tx.Create(&con).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return -1
	}
	return 0
}

func (gs *groupService) FindGroup(userId uint) []models.GroupBasic {
	return dao.Group.FindGroup(userId)
}

func (gs *groupService) AddGroup(userId, groupId uint) int {
	groupBasic := GroupService.GetGroupById(groupId)
	// 群是否存在
	if groupBasic.Name != "" {
		// 是否已经是你的好友
		friend := ContactService.IsFriend(userId, groupId, 2)
		if friend > 0 {
			return 1
		}
		con := models.Contact{}
		con.OwnerId = userId
		con.TargetId = groupId
		con.Type = 2
		err := utils.DB.Create(&con).Error
		if err == nil {
			return 0
		}
	}
	return -1
}

func (gs *groupService) GetGroupById(groupId uint) *models.GroupBasic {
	return dao.Group.GetGroupById(groupId)
}

func (gs *groupService) GetGroupUserById(groupId uint) []uint {
	return dao.Group.GetGroupUserById(groupId)
}
