package models

import (
	"gorm.io/gorm"
)

// Contact 人员关系
type Contact struct {
	gorm.Model
	OwnerId     uint   `json:"ownerId" gorm:"index:fid;type:bigint(20);comment:谁的关系id"`
	TargetId    uint   `json:"targetId" gorm:"index:tid;type:bigint(20);comment:对应id"`
	Type        int    `json:"type" gorm:"type:tinyint(1);comment:聊天类型(1私聊,2群聊)"`
	Description string `json:"description" gorm:"type:text;comment:描述"`
}

func (msg *Contact) TableName() string {
	return "contact"
}
