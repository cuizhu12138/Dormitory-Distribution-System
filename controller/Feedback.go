package controller

import (
	"Dormitory-Distribution-System/midware"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//用户还未关联(调试限制，后续关联token调试)
func Feedback(c *gin.Context) {
	db, err := gorm.Open(sqlite.Open("feedback.db"), &gorm.Config{})
	if err != nil {
		panic("连接失败")
	}
	// 连接数据库
	db.AutoMigrate(&midware.UserFeedback{})
	//确保UserFeedback表存在
	var feedback midware.UserFeedback
	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 插入用户反馈信息
	result := db.Create(&feedback)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "插入错误"})
		return
	}
	c.JSON(200, gin.H{"message": "反馈成功"})
}

/*
{
  "UserID": 123,
  "FeedbackContent": "This is a sample feedback",
  "Timestamp": "2024-01-21T12:34:56Z"
}
*/
