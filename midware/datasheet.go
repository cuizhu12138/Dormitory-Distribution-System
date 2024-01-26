package midware

import (
	"time"
)

type Login struct {
	Authority    int    `gorm:"column:Authority;type:INT NOT NULL"`
	Uid          int    `gorm:"column:Uid;type:INT;PRIMARY KEY;AUTO_INCREMENT;primaryKey"`
	SchoolNumber string `gorm:"column:SchoolNumber"`
	PassWord     string `gorm:"column:PassWord"`
}

// 限定表名为Login
func (v Login) TableName() string {
	return "Login"
}

type UserFeedback struct {
	ID              uint   `gorm:"primaryKey;autoIncrement"`
	UserID          int64  `gorm:"not null"`
	FeedbackContent string `gorm:"type:TEXT"`
	Timestamp       time.Time
}

func (UserFeedback) TableName() string {
	return "UserFeedback"
}
