package like

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// id: 1, title: "Travel Moon", image: "/avatar.png", status: "Public Group", message: "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Nemo atque, odio soluta fuga accusamus itaque accusantium dolorem aut a nisi" },

func ListSocketLike(w http.ResponseWriter, r *http.Request, clients map[*websocket.Conn]string) error {

	data := models.Data{}

	models.OrmInstance.Custom.OrderBy("CreatedAt", "DESC").Limit(1)
	_like, err := models.OrmInstance.Scan(models.Like{}, "Id", "IdObject", "Type", "Liked", "CreatedAt")
	models.OrmInstance.Custom.Clear()

	if err != nil {
		fmt.Println(err)
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return err
	}

	return utils.EncodeJson(w, "createLike", _like, data, clients)
}
