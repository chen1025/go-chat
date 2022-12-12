package models

import (
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string     `json:"name" gorm:"index:idx_name,unique;type:varchar(20)"`
	Password      string     `json:"password" gorm:"type:varchar(128)"`
	Phone         string     `json:"phone"  gorm:"index:idx_phone;type:varchar(20)" valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string     `json:"email" gorm:"type:varchar(128)" valid:"email" `
	ClientIp      string     `json:"clientIp" gorm:"type:varchar(32)"`
	Identity      string     `json:"identity" gorm:"type:varchar(32)"`
	ClientPort    string     `json:"clientPort" gorm:"type:varchar(32)"`
	LoginTime     *time.Time `json:"loginTime"`
	HeartbeatTime *time.Time `json:"heartbeatTime"`
	LoginOutTime  *time.Time `json:"loginOutTime"`
	IsLoginOut    bool       `json:"isLoginOut"`
	DeviceInfo    string     `json:"deviceInfo" gorm:"type:varchar(128)"`
	Icon          string     `json:"icon" gorm:"type:varchar(128)"`
	Nickname      string     `json:"nickname" gorm:"type:varchar(128)"`
}

func (user *UserBasic) TableName() string {
	return "user_basic"
}
