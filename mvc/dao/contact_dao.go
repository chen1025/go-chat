package dao

import (
	"ginchat/models"
	"ginchat/utils"
)

type contact struct {
}

var Contact = new(contact)

func (cs *contact) SearchFriend(userId uint) []models.UserBasic {
	user := make([]models.UserBasic, 0)
	utils.DB.Raw("select id,name,icon from user_basic "+
		"where id in (select target_id from contact where owner_id=? and type = 1)", userId).Scan(&user)
	return user
}

func (cs *contact) IsFriend(userId, targetId uint, _type int) int {
	result := 0
	utils.DB.Raw("select count(1) from contact where owner_id=? and target_id =? and type = ?", userId, targetId, _type).Scan(&result)
	return result
}
