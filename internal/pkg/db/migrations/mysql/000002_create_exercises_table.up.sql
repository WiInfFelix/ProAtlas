CREATE TABLE IF NOT EXISTS Exercises(
    ID INT NOT NULL AUTO_INCREMENT,
    ExerciseName VARCHAR (255),
    ExerciseDescription VARCHAR (255),
    TargetMuscle VARCHAR(255),
    CreatorID INT,
    FOREIGN KEY (CreatorID) REFERENCES Users(ID),
    PRIMARY KEY (ID)
)