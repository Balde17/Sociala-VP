package notifications

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"net/http"

	"github.com/gorilla/websocket"
)

// id: 1, title: "Travel Moon", image: "/avatar.png", status: "Public Group", message: "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Nemo atque, odio soluta fuga accusamus itaque accusantium dolorem aut a nisi" },

func ListSocketNotification(w http.ResponseWriter, r *http.Request, clients *websocket.Conn, idGroup, idUser int, typ string) error {

	data := models.Data{}

	models.OrmInstance.Custom.Join("Notifications", "Notifications.Sender = User.Id").Where("Notifications.Receiver", idUser).OrderBy("User.CreatedAt", "DESC").Limit(1)
	_notificarions, err := models.OrmInstance.Scan1(models.User{}, "User.Id", "Username", "LastName", "FirstName", "Email", "Password", "AboutMe", "Profil", "DateOfBirth", "Type", "User.CreatedAt")
	models.OrmInstance.Custom.Clear()

	if err != nil {
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return err
	}

	return utils.EncodeJson2(w, "notification", _notificarions, data, clients, idGroup, typ)
}
