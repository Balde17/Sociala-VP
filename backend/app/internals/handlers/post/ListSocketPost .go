package post

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"net/http"

	"github.com/gorilla/websocket"
)

// id: 1, title: "Travel Moon", image: "/avatar.png", status: "Public Group", message: "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Nemo atque, odio soluta fuga accusamus itaque accusantium dolorem aut a nisi" },

func ListSocketPost(w http.ResponseWriter, r *http.Request, clients map[*websocket.Conn]string) error {

	data := models.Data{}

	models.OrmInstance.Custom.OrderBy("CreatedAt", "DESC").Limit(1)
	_posts, err := models.OrmInstance.Scan(models.Post{}, "Id", "Title", "Content", "ImgUrl", "Status", "IdUser", "IdGroups", "CreatedAt")
	models.OrmInstance.Custom.Clear()
	
	if err != nil {
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return err
	}

	return utils.EncodeJson(w, "createPost", _posts, data, clients)
}
