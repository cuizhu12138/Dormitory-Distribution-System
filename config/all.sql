CREATE TABLE Login(
    Uid integer PRIMARY KEY AUTO_INCREMENT,
    Authority INT NOT NULL,
    SchoolNumber text,
    PassWord text
);
CREATE TABLE UserFeedback (
    FeedbackID INT AUTO_INCREMENT,
    UserID BIGINT,
    FeedbackContent TEXT,
    Timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (UserID) REFERENCES Users(UserID) 
);
CREATE TABLE Questionnaire_data (
    id SERIAL PRIMARY KEY,
    id VARCHAR(255),
    qid JSONB,
    student_id VARCHAR(255),
    name VARCHAR(255),
    sex VARCHAR(255),
    major VARCHAR(255),
    age VARCHAR(255),
    home VARCHAR(255)[],
    ethnic VARCHAR(255),
    sleep_time VARCHAR(255),
    getup_time VARCHAR(255),
    same_routine VARCHAR(255),
    learn_in_dorm VARCHAR(255),
    neat_expection VARCHAR(255),
    clean_period VARCHAR(255),
    bathe_period VARCHAR(255),
    expense VARCHAR(255),
    cost_type VARCHAR(255)[],
    out_cost VARCHAR(255),
    share_cost VARCHAR(255),
    hobby VARCHAR(255)[],
    hobby_same_expection VARCHAR(255),
    want_communicate VARCHAR(255),
    smoke VARCHAR(255),
    drink VARCHAR(255),
    snore VARCHAR(255),
    sleep_quality VARCHAR(255)
);
