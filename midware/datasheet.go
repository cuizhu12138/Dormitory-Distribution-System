package midware

import "time"

// 登录表
type Login struct {
	UID          int64  `gorm:"column:UID;             type:INT PRIMARY KEY; autoIncrement"`
	Authority    int64  `gorm:"column:Authority;       type:INT"`
	SchoolNumber string `gorm:"column:SchoolNumber;    type:NVARCHAR"`
	PassWord     string `gorm:"column:PassWord;        type:NVARCHAR NOT NULL"`
}

func (Login) TableName() string {
	return "Login"
}

// 用户基本信息表
type UserInfo struct {
	UID      int64  `gorm:"column:UID;            	  type:INT PRIMARY KEY; autoIncrement"`
	UserName string `gorm:"column:UserName;       	  type:NVARCHAR"`
	Gender   string `gorm:"column:Gender;          	  type:NVARCHAR"`
	Age      int64  `gorm:"column:Age;       	  	  type:INT"`
	Ethnic   string `gorm:"column:Ethnic;          	  type:VARCHAR"`
	HomeTown string `gorm:"column:HomeTown;       	  type:VARCHAR"`
	Major    string `gorm:"column:Major;          	  type:VARCHAR"`
}

func (UserInfo) TableName() string {
	return "UserInfo"
}

// 用户习惯表
type UserHabbit struct {
	UID                     int64  `gorm:"column:UID;            	  		type:INT PRIMARY KEY;autoIncrement"`
	BedTime                 string `gorm:"column:BedTime;           		type:VARCHAR"`
	WakeUpTime              string `gorm:"column:WakeUpTime;           		type:VARCHAR"`
	SleepQuality            string `gorm:"column:SleepQuality;           	type:VARCHAR"`
	SynchronizedSchedule    string `gorm:"column:SynchronizedSchedule;      type:BIT"`
	DormStudy               string `gorm:"column:DormStudy;           		type:BIT"`
	Smoke                   string `gorm:"column:Smoke;          		 	type:BIT"`
	Drink                   string `gorm:"column:Drink;           			type:BIT"`
	Snore                   string `gorm:"column:Snore;           			type:BIT"`
	ChattingSharingThoughts string `gorm:"column:ChattingSharingThoughts;   type:BIT"`
	Leanliness              string `gorm:"column:Leanliness;          		type:NVARCHAR"`
	CleaningFrequency       string `gorm:"column:CleaningFrequency;         type:NVARCHAR"`
	ShowerFrequency         string `gorm:"column:ShowerFrequency;           type:NVARCHAR"`
	MonthlyBudget           string `gorm:"column:MonthlyBudget;          	type:NVARCHAR"`
	SpendingResponsibility  string `gorm:"column:SpendingResponsibility;    type:NVARCHAR"`
	JointOutings            string `gorm:"column:JointOutings;           	type:BIT"`
	SharedExpenses          string `gorm:"column:SharedExpenses;            type:BIT"`
	Interests               string `gorm:"column:Interests;          		type:NVARCHAR"`
	SharedInterests         string `gorm:"column:SharedInterests;           type:BIT"`
}

func (UserHabbit) TableName() string {
	return "UserHabbit"
}

// 分配结果表
type DistributionResult struct {
	AllocationID        int64  `gorm:"column:AllocationID;      			type:INT"`
	OptionInfo          string `gorm:"column:OptionInfo;          			type:NVARCHAR"`
	RoomNumber          string `gorm:"column:RoomNumber;          			type:NVARCHAR"`
	UID                 int64  `gorm:"column:UID;       					type:INT"`
	DecisionForReassign string `gorm:"column:DecisionForReassign;           type:BIT"`
	ReassignResult      string `gorm:"column:ReassignResult;          		type:NVARCHAR"`
}

func (DistributionResult) TableName() string {
	return "DistributionResult"
}

// 信息反馈表
type InformationFeedback struct {
	FeedbackID      int64     `gorm:"column:FeedbackID;       				type:INT"`
	UID             int64     `gorm:"column:UID;       						type:INT"`
	FeedbackContent string    `gorm:"column:FeedbackContent;          		type:NVARCHAR"`
	TimeStamp       time.Time `gorm:"column:TimeStamp;          			type:TIMESTAMP"`
}
type UserFeedback struct {
	ID              uint   `gorm:"primaryKey;autoIncrement"`
	UserID          int64  `gorm:"not null"`
	FeedbackContent string `gorm:"type:TEXT"`
	Timestamp       time.Time
}

func (UserFeedback) TableName() string {
	return "UserFeedback"
}

func (InformationFeedback) TableName() string {
	return "InformationFeedback"
}
