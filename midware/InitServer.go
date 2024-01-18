package midware

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
	// 用于存储用户的token
	Tokens = map[string]int{
		"init": 0,
	}
	// 用于判断用户是否已经在线
	Online = map[string]int{
		"init": 0,
	}
	// 学号映射token
	SToT = map[string]string{
		"init": "0",
	}
)

func InitService() {
	// init DB
	if DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{}); err != nil {
		fmt.Printf("[DB Err]\t%v\n", err)
	}
	// init DataSheet
	if err = DB.AutoMigrate(&Login{}); err != nil {
		fmt.Printf("[DB Err]\t%v\n", err)
	}
}
