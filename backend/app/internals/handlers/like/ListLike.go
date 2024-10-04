package like

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"log"
	"net/http"
	"strconv"
	"time"
)

// id: 1, title: "Travel Moon", image: "/avatar.png", status: "Public Group", message: "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Nemo atque, odio soluta fuga accusamus itaque accusantium dolorem aut a nisi" },
func ListLike(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}
	_idPost := r.URL.Query().Get("idPost")
	typeLike := r.URL.Query().Get("type")
	idPost, err := strconv.Atoi(_idPost)
	if err != nil {
		log.Println(err.Error())
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusNotFound, data)
		return
	}
	models.OrmInstance.Custom.Where("IdObject", idPost).And("Type", typeLike).OrderBy("CreatedAt", "DESC").Limit(1)
	_like, err := models.OrmInstance.Scan(models.Like{}, "Id", "IdObject", "Type", "Liked", "CreatedAt")
	if err != nil {
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return
	}
	likes, ok := _like.([]struct {
		Id        int64
		IdObject  int
		Type      string
		Liked     int
		CreatedAt time.Time
	})
	if !ok {
		data.Error = "Error conversion from struct"
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return
	}
	data.Content = likes
	utils.SendJSON(w, http.StatusOK, data)
}
