package main

import (
	"Dormitory-Distribution-System/controller"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type QuestionnaireInfo struct {
    QID    int    `json:"qid"`
    Title  string `json:"title"`
    State  string `json:"state"`
    Number int    `json:"number"`
    ID     string `json:"id"`
}
type QuestionnaireData struct {
	ID               string  `json:"id"`
	Qid              struct {
		IsTrusted bool `json:"isTrusted"`
	} `json:"qid"`
	StudentId        interface{} `json:"studentId"`
	Name             string      `json:"name"`
	Sex              string      `json:"sex"`
	Major            string      `json:"major"`
	Age              string      `json:"age"`
	Home             []string    `json:"home"`
	Ethnic           string      `json:"ethnic"`
	SleepTime        interface{} `json:"sleepTime"`
	GetupTime        interface{} `json:"getupTime"`
	SameRoutine      interface{} `json:"sameRoutine"`
	LearnInDorm      interface{} `json:"learnInDorm"`
	NeatExpection    interface{} `json:"neatExpection"`
	CleanPeriod      interface{} `json:"cleanPeriod"`
	BathePeriod      interface{} `json:"bathePeriod"`
	Expense          interface{} `json:"expense"`
	CostType         []string    `json:"costType"`
	OutCost          interface{} `json:"outCost"`
	ShareCost        interface{} `json:"shareCost"`
	Hobby            []string    `json:"hobby"`
	HobbySameExpection interface{} `json:"hobbySameExpection"`
	WantCommunicate  interface{} `json:"wantCommunicate"`
	Smoke            interface{} `json:"smoke"`
	Drink            interface{} `json:"drink"`
	Snore            interface{} `json:"snore"`
	SleepQuality     interface{} `json:"sleepQuality"`
}
type QuestionnaireDataAF struct {
    UID 						uint 		   `gorm:"primaryKey;autoIncrement"`
	Name             			string         `gorm:"name"`
	Sex             			string         `gorm:"sex"`
	Major            			string         `gorm:"major"`
	Age              			string         `gorm:"age"`
	Homestr            			string         `gorm:"home"`
	Ethnic           			string         `gorm:"ethnic"`
    BedTime   		  			string         `gorm:"column:bedTime"`
    WakeUpTime        			string         `gorm:"column:wakeUpTime"`
    SleepQuality     			string         `gorm:"column:sleepQuality"`
    DomStudy             		string         `gorm:"column:domStudy"`
    Smoke               		string         `gorm:"column:smoke"`
    Drink              			string         `gorm:"column:drink"`
    Snore            			string         `gorm:"column:snore"`
    ChattingSharinsThoushts     string     	   `gorm:"column:chattingSharinsThoushts"`
    Leanliness         			string     	   `gorm:"column:leanliness"`
    Cleaningfrsgueney       	string         `gorm:"column:cleaningfrsgueney"`
    Showerkrequency       		string         `gorm:"column:showerkrequency"`
    MonthlyBudset   			string         `gorm:"column:monthlyBudset"`
    JointOutings       			string         `gorm:"column:jointOutings"`
    SharedExpenses      	    string         `gorm:"column:sharedExpenses"`
    SharedInterests             string         `gorm:"column:sharedInterests"`
}

func InitRouter(r *gin.Engine) {
	
	g1 := r.Group("/user")
	{
		g1.POST("/login/", controller.Login)
	}
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})
	r.GET("/questionnaireInfo", func(c *gin.Context) {
        questionnaireInfo := []QuestionnaireInfo{
			{0, "2023 Freshman Second Questionnaire", "Enabled", 0, "0daaas"},
        }
        c.JSON(http.StatusOK, questionnaireInfo)
    })
	r.OPTIONS("/questionnaire", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Status(http.StatusOK)
	})
	r.POST("/questionnaire", func(c *gin.Context) {
		var requestData QuestionnaireData
		
		if err := c.BindJSON(&requestData); err != nil {
			c.JSON(400, gin.H{"error": "Failed to parse JSON"})
			return
		}
		c.JSON(200, gin.H{"message": "Data received"})
		dsn := "root:11111111@(localhost)/Questionnaire_data?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err!= nil{
			panic(err)
		}
		db.AutoMigrate(&QuestionnaireDataAF{})
		var data QuestionnaireDataAF
		data.Age = requestData.Age
		data.Name = requestData.Name
		data.Major = requestData.Major
		data.Sex = requestData.Sex
		data.Homestr = strings.Join(requestData.Home , ",")
		data.BedTime = 		requestData.SleepTime.(string)
		data.WakeUpTime = 	requestData.GetupTime.(string)
		data.SleepQuality =	requestData.SleepQuality.(string)
		data.DomStudy = 	requestData.LearnInDorm.(string)
		data.Smoke = 		requestData.Smoke.(string)
		data.Drink = 		requestData.Drink.(string)
		data.Snore =		requestData.Snore.(string)
		data.ChattingSharinsThoushts = requestData.WantCommunicate.(string)
		data.Leanliness = requestData.NeatExpection.(string)
		data.Cleaningfrsgueney = requestData.CleanPeriod.(string)
		data.Showerkrequency = requestData.BathePeriod.(string)
		data.MonthlyBudset = requestData.Expense.(string)
		data.JointOutings = requestData.OutCost.(string)
		data.SharedExpenses = requestData.ShareCost.(string)
		data.SharedInterests =  requestData.HobbySameExpection.(string)
		db.Create(&data)
		db.Debug().Create(&data) 
		var u = new(QuestionnaireDataAF)
		db.First(u)
		fmt.Printf("%#v\n", u)
	})
}
