package service

import (
	"ginchat/models"
	"ginchat/mvc/dao"
	"ginchat/utils"
	"gorm.io/gorm"
)

type contactService struct {
}

var ContactService = new(contactService)

func (cs *contactService) AddFriend(userId uint, targetName string) int {
	user := UserService.GetByName(targetName)
	// 用户存在
	if user.Name != "" {
		// 是否已经是你的好友
		friend := ContactService.IsFriend(userId, user.ID, 1)
		if friend > 0 || userId == user.ID {
			return 1
		}

		err := utils.DB.Transaction(func(tx *gorm.DB) error {
			con := models.Contact{}
			con.OwnerId = userId
			con.TargetId = user.ID
			con.Type = 1

			target := models.Contact{}
			target.OwnerId = user.ID
			target.TargetId = userId
			target.Type = 1
			// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
			if err := tx.Create(&con).Error; err != nil {
				// 返回任何错误都会回滚事务
				return err
			}
			if err := tx.Create(&target).Error; err != nil {
				return err
			}

			// 返回 nil 提交事务
			return nil
		})
		if err == nil {
			return 0
		}
	}
	return -1
}

func (cs *contactService) SearchFriend(userId uint) []models.UserBasic {
	return dao.Contact.SearchFriend(userId)
}

func (cs *contactService) IsFriend(userId, targetId uint, _type int) int {
	return dao.Contact.IsFriend(userId, targetId, _type)
}
