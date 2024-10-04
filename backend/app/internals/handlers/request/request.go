package request

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"log"
	"net/http"
	"strconv"
)

func Request(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}
	_idUser := r.URL.Query().Get("idUser")
	idUser, err := strconv.Atoi(_idUser)
	if err != nil {
		log.Println(err.Error())
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusNotFound, data)
		return
	}
	models.OrmInstance.Custom.Join("Followers", "Followers.IdFollower = User.Id").Where("Followers.IdFollowee", idUser).And("Status", "PENDING")
	_users, err := models.OrmInstance.Scan1(models.User{}, "User.Id", "Username", "LastName", "Email", "FirstName", "ImageURL")
	models.OrmInstance.Custom.Clear()

	if err != nil {
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return
	}
	users, ok := _users.([]struct {
		Id        int64
		Username  string
		LastName  string
		Email     string
		FirstName string
		ImageURL  string
	})
	if !ok {
		data.Error = "Error conversion from struct"
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return
	}
	data.Content = users
	utils.SendJSON(w, http.StatusOK, data)
}
