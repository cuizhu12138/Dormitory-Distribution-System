package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// init some service like database or some other connections
	go InitService()

	r := gin.Default()

	InitRouter(r)

	r.Run()

	fmt.Println("Init Success")
}

func InitService() {

}
