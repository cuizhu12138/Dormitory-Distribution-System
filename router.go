package main

import (
	"Dormitory-Distribution-System/controller"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
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
            {0, "2021 Freshman First Questionnaire", "Disabled", 235, "3173"},
            {1, "2022 Freshman First Questionnaire", "Disabled", 235, "ac1e"},
            {2, "2022 Freshman First Questionnaire", "Submitted", 1, "6fdc"},
            {3, "2023 Freshman Second Questionnaire", "Enabled", 0, "0daa"},
			{4, "2023 Freshman Second Questionnaire", "Enabled", 0, "0dasaa"},
			{5, "2023 Freshman Second Questionnaire", "Enabled", 0, "0daaas"},
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
		fmt.Printf("%+v\n", requestData)
	})
}
