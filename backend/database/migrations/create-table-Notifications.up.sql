CREATE TABLE IF NOT EXISTS Notifications (
	Id INTEGER PRIMARY KEY AUTOINCREMENT,
	CreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
	Type TEXT CHECK( Type IN ('FOLLOW','INVITE','REQUEST','EVENT','PRIVATE','Groups') ),
	Sender INTEGER ,
	Receiver INTEGER ,
	FOREIGN KEY (Sender) REFERENCES User (Id)
)