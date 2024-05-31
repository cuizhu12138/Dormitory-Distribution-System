
-- Login 表
CREATE TABLE IF NOT EXISTS logins(
    uid INTEGER PRIMARY AUTOINCREMENT,
    Authority INTEGER,
    SchoolNumber NVARCHAR NOT NULL,
    PassWord NVARCHAR NOT NULL,
);

-- 用户信息表
CREATE TABLE IF NOT EXISTS user_base_infos(
    uid INTEGER PRIMARY AUTOINCREMENT,
    name INTEGER,
    sex NVARCHAR,
    major NVARCHAR,
    age NVARCHAR,
    home NVARCHAR,
    sychronizedSchedule NVARCHAR,
    spendingResponsibility NVARCHAR,
    interests NVARCHAR
);

-- 用户问卷表
CREATE TABLE IF NOT EXISTS user_questionnaire_datas(
    uid INTEGER,
    bedTime NVARCHAR,
    wakeUpTime NVARCHAR,
    sleepQuality NVARCHAR,
    domStudy NVARCHAR,
    smoke NVARCHAR,
    drink NVARCHAR,
    snore NVARCHAR,
    chattingSharinsThoushts NVARCHAR,
    leanliness NVARCHAR,
    cleaningfrsgueney NVARCHAR,
    showerkrequency NVARCHAR,
    monthlyBudset NVARCHAR,
    jointOutings NVARCHAR,
    sharedExpenses NVARCHAR,
    sharedInterests NVARCHAR
);

-- 寝室反馈表
CREATE TABLE distribution_results (
    AllocationID INTEGER PRIMARY KEY AUTOINCREMENT,
    OptionInfo NVARCHAR,
    RoomNumber NVARCHAR,
    UID INTEGER,
    DecisionForReassign NVARCHAR,
    ReassignResult NVARCHAR
);

-- 结果用户信息表
CREATE TABLE information_feedbacks (
    feedbackid INTEGER PRIMARY KEY AUTOINCREMENT,
    uid INTEGER,
    feedbackcontent NVARCHAR,
    timestamp TIMESTAMP
);

--反馈信息表
CREATE TABLE user_feedbacks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    userID INTEGER NOT NULL,
    feedbackContent TEXT,
    timestamp TIMESTAMP
);
