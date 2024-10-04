package profile

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"fmt"
	"net/http"
)

// id: 1, title: "Travel Moon", image: "/avatar.png", status: "Public Group", message: "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Nemo atque, odio soluta fuga accusamus itaque accusantium dolorem aut a nisi" },
func ListFollower(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}
	_idUser := r.URL.Query().Get("idUser")
	idUser := utils.GetInt(_idUser, w)

	models.OrmInstance.Custom.Join("Followers", "Followers.IdFollower = User.Id").Where("Followers.IdFollowee", idUser).And("Followers.Status ", "VALIDATE")
	_followers, err := models.OrmInstance.Scan1(models.User{}, "User.Id", "Username", "LastName", "FirstName", "Email", "Password", "ImageURL", "AboutMe", "Profil", "DateOfBirth")

	fmt.Println("------------------------------------>", err)

	if err != nil {
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return
	}
	followers, ok := _followers.([]struct {
		Id          int64
		Username    string
		LastName    string
		FirstName   string
		Email       string
		Password    string
		ImageURL    string
		AboutMe     string
		Profil      string
		DateOfBirth string
	})
	if !ok {
		data.Error = "Error conversion from struct"
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return
	}
	data.Content = followers
	utils.SendJSON(w, http.StatusOK, data)
}
