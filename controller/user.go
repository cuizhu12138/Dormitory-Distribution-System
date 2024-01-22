package controller

import (
	"Dormitory-Distribution-System/midware"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginResponse struct {
	midware.Response
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	schoolnumber := c.Query("schoolnumber")
	password := c.Query("password")

	result := midware.Login{}
	// 查询学号，根据返回的err判断是否存在，token等于学号加盐保证唯一且安全
	if err := midware.DB.Where("SchoolNumber = ?", schoolnumber).Take(&result).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, LoginResponse{
			midware.Response{
				StatusCode: 0, StatusMsg: "User not exist or PassWord wrong",
			}, ""})
	} else {
		if password == result.PassWord {
			if midware.Online[schoolnumber] == 1 {
				c.JSON(http.StatusOK, LoginResponse{ // 固定token（暂时没有退出登录功能）
					midware.Response{
						StatusCode: 1, StatusMsg: "ok",
					}, midware.SToT[schoolnumber]})
				return
			}
			token := schoolnumber + midware.RandSalt()
			midware.Tokens[token] = 1
			midware.Online[schoolnumber] = 1
			midware.SToT[schoolnumber] = token
			c.JSON(http.StatusOK, LoginResponse{ //
				midware.Response{
					StatusCode: 1, StatusMsg: "ok",
				}, token})
		} else {
			c.JSON(http.StatusOK, LoginResponse{
				midware.Response{
					StatusCode: 0, StatusMsg: "User not exist or PassWord wrong",
				}, ""})
		}
	}

}
