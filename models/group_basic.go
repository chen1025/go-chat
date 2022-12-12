package models

import (
	"gorm.io/gorm"
)

// GroupBasic 群信息
type GroupBasic struct {
	gorm.Model
	OwnerId     uint   `json:"ownerId" gorm:"index:fid;type:bigint(20);comment:拥有者id"`
	Name        string `json:"name" gorm:"type:varchar(20);comment:名称"`
	Type        int    `json:"type" gorm:"type:tinyint(1);comment:类型"`
	Icon        string `json:"icon" gorm:"type:varchar(128);comment:图片"`
	Description string `json:"description" gorm:"type:text;comment:描述"`
}

func (msg *GroupBasic) TableName() string {
	return "group_basic"
}
