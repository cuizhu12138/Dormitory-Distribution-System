package controller

import (
	"Dormitory-Distribution-System/midware"
	// "errors"
	"net/http"
	"crypto/sha256"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
	"encoding/hex"
)
func Register(c *gin.Context){
	var userJSON midware.RegisterRequest
	if err := c.ShouldBindJSON(&userJSON); err != nil{
		c.JSON(http.StatusBadRequest , gin.H{"resgiste erro" : err.Error()})
		return
	}
	var user midware.Login
	user.SchoolNumber = userJSON.Username
	user.PassWord = userJSON.Password
	user.Authority = 0
	user.PassWord = hashPassword(user.PassWord)
	if err := midware.DB.Create(&user).Error; err!=nil{
		c.JSON(http.StatusInternalServerError , gin.H{"error":err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
func Login(c *gin.Context){
	var userJSON midware.LoginRequest
	if err := c.ShouldBindJSON(&userJSON); err != nil{
		c.JSON(http.StatusBadRequest , gin.H{"login erro" : err.Error()})
		return
	}
	var user midware.Login
	if err := midware.DB.Where("schoolNumber = ? AND passWord = ?" , userJSON.Username , hashPassword(userJSON.Password)).First(&user).Error; err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
	}
	userID := user.UID

	token , err := midware.GenToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK , gin.H{"token" : token})
}


func hashPassword(password string) string{
	hasher := sha256.New()
	hasher.Write([]byte(password))
	hashedPassword := hasher.Sum(nil)
	return hex.EncodeToString(hashedPassword)
}
