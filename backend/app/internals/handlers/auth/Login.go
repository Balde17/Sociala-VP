package auth

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"backend/lib/jwt"
	orm "backend/lib/orm/ORM"
	"fmt"
	"os"

	//"backend/lib/jwt"

	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type MyData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	data := models.Data{}
	fmt.Println("-------------------------")
	r.ParseMultipartForm(10 << 20)

	// var loginData MyData

	// err := json.NewDecoder(r.Body).Decode(&loginData)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	emailLogin := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Println(r.Form, emailLogin, password)
	models.OrmInstance.Custom.Where("Email", emailLogin)
	_email, err := models.OrmInstance.Scan(models.User{}, "Id", "Username", "LastName", "FirstName", "Email", "Password", "AboutMe", "Profil", "ImageURL", "DateOfBirth")
	models.OrmInstance.Custom.Clear()

	if err != nil {
		log.Println(err.Error())
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return
	}

	email, ok := _email.([]struct {
		Id          int64
		Username    string
		LastName    string
		FirstName   string
		Email       string
		Password    string
		AboutMe     string
		Profil      string
		ImageURL    string
		DateOfBirth string
	})

	if !ok {
		data.Error = "Error conversion from struct"
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return
	}

	if len(email) == 0 {
		data.Error = "Invalid email address"
		fmt.Println(data.Error)
		utils.SendJSON(w, http.StatusBadRequest, data)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(email[0].Password), []byte(password))

	if err != nil {
		data.Error = "Invalid password"
		utils.SendJSON(w, http.StatusBadRequest, data)
		return
	}

	var payload Payload
	payload.User = models.User{
		Model:       orm.Model{Id: email[0].Id},
		Username:    email[0].Username,
		LastName:    email[0].LastName,
		FirstName:   email[0].FirstName,
		Email:       email[0].Email,
		AboutMe:     email[0].AboutMe,
		Profil:      email[0].Profil,
		ImageURL:    email[0].ImageURL,
		DateOfBirth: email[0].DateOfBirth,
	}
	payload.RemoteAddr = r.RemoteAddr

	SECRET := os.Getenv("SECRET_KEY")
	payload.Sign = jwt.CryptCode(SECRET)
	jwt := jwt.New()
	jwt.SetPayload("payload", payload)

	token, err := jwt.Sign(SECRET)

	if err != nil {
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return
	}

	data.Token = token
	data.Id = email[0].Id

	utils.SendJSON(w, http.StatusOK, data)

}
