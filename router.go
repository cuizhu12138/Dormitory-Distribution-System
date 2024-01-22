package main

import (
	"Dormitory-Distribution-System/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// reflect URL "/static" to path "/public"
	r.Static("/static", "./")
	// web访问根
	// Users control
	g1 := r.Group("/user")
	{
		g1.POST("/login/", controller.Login)
		g1.POST("/Feedback", controller.Feedback)

	}
	// g2 :=r.Group("/Administrator")
	// {
	// 	g2.GET("/Feedback_list" , ##)
	// }
	// 管理员
}
