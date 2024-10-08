CREATE TABLE IF NOT EXISTS GroupMember (
	Id INTEGER PRIMARY KEY AUTOINCREMENT,
	CreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
	IdUser INTEGER ,
	IdGroups INTEGER ,
	IsAdmin INTEGER CHECK( IsAdmin IN (0,1) ),
	FOREIGN KEY (IdUser) REFERENCES User (Id),
FOREIGN KEY (IdGroups) REFERENCES Groups (Id)
)