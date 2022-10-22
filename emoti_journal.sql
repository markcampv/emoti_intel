DROP DATABASE IF EXISTS JournalDB;
CREATE SCHEMA JournalDB;
CREATE TABLE JournalDB.Emoti (
    Emotion varchar(25) NOT NULL,
    Response varchar(100) NOT NULL,
    JournalDate datetime
 );