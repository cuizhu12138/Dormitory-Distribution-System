package main

import (
	"Dormitory-Distribution-System/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// reflect URL "/static" to path "/public"
	r.Static("/static", "./public")

	// Users control
	g1 := r.Group("/user")
	{
		g1.POST("/login/", controller.Login)
	}
}
