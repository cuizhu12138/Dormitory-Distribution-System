package controller

import (
	"Dormitory-Distribution-System/midware"

	"github.com/gin-gonic/gin"
)

//用户还未关联(调试限制，后续关联token调试)
func Feedback(c *gin.Context) {
	var feedback midware.UserFeedback
	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 插入用户反馈信息
	result := midware.DB.Create(&feedback)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "插入错误"})
		return
	}
	c.JSON(200, gin.H{"message": "反馈成功"})
}
