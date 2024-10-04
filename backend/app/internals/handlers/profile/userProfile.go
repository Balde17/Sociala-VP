package profile

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"net/http"
	"time"
)

func UserProfile(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}

	models.OrmInstance.Custom.OrderBy("CreatedAt", "DESC")
	_users, err := models.OrmInstance.Scan(models.User{}, "Id", "Username", "LastName", "FirstName", "Email", "ImageURL", "AboutMe", "Profil", "DateOfBirth", "CreatedAt")
	models.OrmInstance.Custom.Clear()

	if err != nil {
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return
	}
	users, ok := _users.([]struct {
		Id          int64
		Username    string
		LastName    string
		FirstName   string
		Email       string
		ImageURL    string
		AboutMe     string
		Profil      string
		DateOfBirth string
		CreatedAt   time.Time
	})

	if !ok {
		data.Error = "Error conversion from struct"
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return
	}
	data.Content = users
	utils.SendJSON(w, http.StatusOK, data)
}
