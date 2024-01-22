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