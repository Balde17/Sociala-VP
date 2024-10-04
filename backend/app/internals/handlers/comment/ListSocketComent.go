package comment

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"net/http"

	"github.com/gorilla/websocket"
)

// id: 1, title: "Travel Moon", image: "/avatar.png", status: "Public Group", message: "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Nemo atque, odio soluta fuga accusamus itaque accusantium dolorem aut a nisi" },
func ListSocketComment(w http.ResponseWriter, r *http.Request, clients map[*websocket.Conn]string, _idPost int) error {
	data := models.Data{}
	models.OrmInstance.Custom.Where("IdPost", _idPost).OrderBy("CreatedAt", "DESC")
	_comments, err := models.OrmInstance.Scan(models.Comment{}, "Id", "Content", "ImgUrl", "IdUser", "IdPost", "CreatedAt")
	models.OrmInstance.Custom.Clear()

	if err != nil {
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return err
	}
	return utils.EncodeJson(w, "createComment", _comments, data, clients)
}
