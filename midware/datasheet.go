package midware

import "time"

// 登录表
type Login struct {
	UID          int64   `gorm:"column:uid;auto_increment;primary_key"`
	Authority    int64  `gorm:"column:authority;       type:INT"`
	SchoolNumber string `gorm:"column:schoolNumber;    type:NVARCHAR NOT NULL"`
	PassWord     string `gorm:"column:passWord;        type:NVARCHAR NOT NULL"`
}

type UserBaseInfo struct {
	UID                    int64  `gorm:"column:uid							type:INT"`						
	Name                   string `gorm:"column:name						type:NVARCHAR"`
	Sex                    string `gorm:"column:sex							type:NVARCHAR"`
	Major                  string `gorm:"column:major						type:NVARCHAR"`
	Age                    string `gorm:"column:age							type:NVARCHAR"`
	Homestr                string `gorm:"column:home						type:NVARCHAR"`
	SychronizedSchedule    string `gorm:"column:sychronizedSchedule			type:NVARCHAR"`
	SpendingResponsibility string `gorm:"column:spendingResponsibility		type:NVARCHAR"`
	Interests              string `gorm:"column:interests					type:NVARCHAR"`
}

type UserQuestionnaireData struct {
	UID                     int64   `gorm:"column:uid						type:INT"`
	BedTime                 string `gorm:"column:bedTime					type:NVARCHAR"`
	WakeUpTime              string `gorm:"column:wakeUpTime					type:NVARCHAR"`
	SleepQuality            string `gorm:"column:sleepQuality				type:NVARCHAR"`
	DomStudy                string `gorm:"column:domStudy					type:NVARCHAR"`
	Smoke                   string `gorm:"column:smoke						type:NVARCHAR"`
	Drink                   string `gorm:"column:drink						type:NVARCHAR"`
	Snore                   string `gorm:"column:snore						type:NVARCHAR"`
	ChattingSharinsThoushts string `gorm:"column:chattingSharinsThoushts	type:NVARCHAR"`
	Leanliness              string `gorm:"column:leanliness					type:NVARCHAR"`
	Cleaningfrsgueney       string `gorm:"column:cleaningfrsgueney			type:NVARCHAR"`
	ShowerFrequency         string `gorm:"column:showerkrequency			type:NVARCHAR"`
	MonthlyBudget           string `gorm:"column:monthlyBudset				type:NVARCHAR"`
	JointOutings            string `gorm:"column:jointOutings				type:NVARCHAR"`
	SharedExpenses          string `gorm:"column:sharedExpenses				type:NVARCHAR"`
	SharedInterests         string `gorm:"column:sharedInterests			type:NVARCHAR"`
}

// 分配结果表
type DistributionResult struct {
	RoomNumber          string `gorm:"column:RoomNumber;          			type:NVARCHAR"`
	UID                 int64  `gorm:"column:UID;       					type:INT"`
	DecisionForReassign string `gorm:"column:DecisionForReassign;           type:NVARCHAR"`
	ReassignResult      string `gorm:"column:ReassignResult;          		type:NVARCHAR"`
}

// 信息反馈表
type InformationFeedback struct {
	FeedbackID      int64     `gorm:"column:feedbackid;auto_increment;primary_key"`
	UID             int64     `gorm:"column:uid;       						type:INT"`
	FeedbackContent string    `gorm:"column:feedbackcontent;          		type:NVARCHAR"`
	TimeStamp       time.Time `gorm:"column:timestamp;          			type:TIMESTAMP"`
}

//反馈信息表
type UserFeedback struct {
	ID              uint   `gorm:"column:id;auto_increment;primary_key"`
	UserID          int64  `gorm:"not null"`
	FeedbackContent string `gorm:"type:TEXT"`
	Timestamp       time.Time
}

func (Login) TableName() string {
	return "logins"
}

func (UserBaseInfo) TableName() string {
	return "user_base_infos"
}

func (UserQuestionnaireData) TableName() string {
	return "user_questionnaire_datas"
}

func (DistributionResult) TableName() string {
	return "distribution_results"
}

func (InformationFeedback) TableName() string {
	return "information_feedbacks"
}

func (UserFeedback) TableName() string {
	return "user_feedbacks"
}
