package main

import (
	"Dormitory-Distribution-System/midware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// init some service like database or some other connections
	go midware.InitService()

	r := gin.Default()

	InitRouter(r)

	r.Run()

	fmt.Println("Init Success")
}
