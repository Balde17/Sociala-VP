package chat

import (
	"backend/app/internals/models"
	"backend/lib/orm/ORM/queryBuilder"
)

func RegisterMessage(privateMessage *models.ChatData) error {
	_, _, err := queryBuilder.NewInsertBuilder().
		InsertInto("chat", "Receiver", "Sender", "Type", "Content").
		Values(privateMessage.Receiver, privateMessage.Sender, privateMessage.Type, privateMessage.Content).
		InsertQuery(models.OrmInstance.Db)

	return err
}
