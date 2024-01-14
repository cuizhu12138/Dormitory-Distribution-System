package main

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// reflect URL "/static" to path "/public"
	r.Static("/static", "./public")

	// part1
	g1 := r.Group("/g1")
	{
		g1.GET("")
		g1.POST("")
	}
	// part2
	g2 := r.Group("/g2")
	{
		g2.GET("")
		g2.POST("")
	}
	// part3
	g3 := r.Group("/g3")
	{
		g3.GET("")
		g3.POST("")
	}
}
