package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	UserId      int64  `json:"userId" gorm:"index:fid;type:bigint(20);comment:发送者id"`
	TargetId    int64  `json:"targetId" gorm:"index:tid;type:bigint(20);comment:接收者id"`
	Type        int    `json:"type" gorm:"type:tinyint(1);comment:聊天类型(1私聊，2群聊，3心跳)"`
	Media       int    `json:"media" gorm:"type:tinyint(1);comment:消息类型(1文字，2表情包，3音频，4图片)"`
	Content     string `json:"content" gorm:"type:text;comment:内容"`
	Pic         string `json:"pic" gorm:"type:text;comment:图片"`
	Url         string `json:"url" gorm:"type:text;comment:url"`
	Description string `json:"description" gorm:"type:text;comment:描述"`
	Amount      int    `json:"amount" gorm:"type:int(10);comment:其他统计字段"`
	Icon        string `json:"icon" gorm:"type:varchar(128)"`
}

func (msg *Message) TableName() string {
	return "message"
}
