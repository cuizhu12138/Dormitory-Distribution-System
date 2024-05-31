package controller

import (
	"Dormitory-Distribution-System/midware"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Feedback(c *gin.Context) {
	var feedback midware.FeedbackRequest
	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var fb midware.UserFeedback
	userID, exists := c.Get("UID")
	fmt.Println(userID)
	if !exists {
		fmt.Println("here is www")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "????"})
		return
	}
	fb.UserID = userID.(int64)
	fb.Timestamp = time.Now()
	fb.FeedbackContent = feedback.Context

	result := midware.DB.Create(&fb)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "插入错误"})
		return
	}
	c.JSON(200, gin.H{"message": "反馈成功"})
}