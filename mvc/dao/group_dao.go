package dao

import (
	"ginchat/models"
	"ginchat/utils"
)

type group struct {
}

var Group = new(group)

func (gs *group) FindGroup(userId uint) []models.GroupBasic {
	list := make([]models.GroupBasic, 0)
	utils.DB.Raw("select * from group_basic "+
		"where id in (select target_id from contact where owner_id=? and type = 2)", userId).Scan(&list)
	return list
}

func (gs *group) GetGroupById(groupId uint) *models.GroupBasic {
	group := models.GroupBasic{}
	utils.DB.First(&group, groupId)
	return &group
}

func (gs *group) GetGroupUserById(groupId uint) []uint {
	userIds := make([]uint, 0)
	utils.DB.Raw("select owner_id from contact where target_id=? and type =2", groupId).Scan(&userIds)
	return userIds
}
