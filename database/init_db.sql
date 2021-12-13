CREATE TABLE IF NOT EXISTS users (
                                     ID INTEGER PRIMARY KEY AUTOINCREMENT,
                                     Username TEXT NOT NULL UNIQUE,
                                     Email TEXT NOT NULL UNIQUE,
                                     Password TEXT NOT NULL
)