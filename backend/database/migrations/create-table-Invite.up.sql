CREATE TABLE IF NOT EXISTS Invite (
	Id INTEGER PRIMARY KEY AUTOINCREMENT,
	CreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
	IdGroup INTEGER,
	IdUser INTEGER,
	IdUserInvited INTEGER,
	Accepted INTEGER
)

