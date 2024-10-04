package notifications

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"fmt"
	"net/http"
)

func CreateSocketNotication(w http.ResponseWriter, r *http.Request, notificationData models.NotificationData) error {
	fmt.Println(notificationData.IdGroup)
	n := models.Notifications{
		Type:        notificationData.Type,
		Sender:      notificationData.Sender,
		Receiver:    notificationData.Receiver,
		IdGroup:     notificationData.IdGroup,
		IdEvent:     notificationData.IdEvent,
		EventOption: "NULL",
	}
	return utils.CreateData(w, "create notification succesfull", n)
}
