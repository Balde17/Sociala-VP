package profile

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// id: 1, title: "Travel Moon", image: "/avatar.png", status: "Public Group", message: "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Nemo atque, odio soluta fuga accusamus itaque accusantium dolorem aut a nisi" },
func GroupUserList(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}
	_idUser := r.URL.Query().Get("idUser")
	idUser, err := strconv.Atoi(_idUser)
	if err != nil {
		log.Println(err.Error())
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusNotFound, data)
		return
	}
	models.OrmInstance.Custom.Join("GroupMember", "GroupMember.IdGroups = Groups.Id").Where("GroupMember.IdUser", idUser)
	_groups, err := models.OrmInstance.Scan(models.Groups{}, "Title", "Description", "ImgUrl")
	models.OrmInstance.Custom.Clear()

	fmt.Println("------------------------------------>", err)

	if err != nil {
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return
	}
	groups, ok := _groups.([]struct {
		Title       string
		Description string
		ImgUrl      string
	})
	if !ok {
		data.Error = "Error conversion from struct"
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return
	}
	data.Content = groups
	utils.SendJSON(w, http.StatusOK, data)
}
