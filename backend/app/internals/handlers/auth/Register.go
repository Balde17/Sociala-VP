package auth

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"backend/lib/orm/ORM/queryBuilder"
	"backend/lib/validators"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	u := models.User{
		Email:       r.FormValue("email"),
		FirstName:   r.FormValue("firstname"),
		LastName:    r.FormValue("lastname"),
		DateOfBirth: r.FormValue("birthday"),
		Profil:      r.FormValue("profil"),
		Username:    r.FormValue("username"),
		ImageURL:    r.FormValue("imageUrlFromChild"),
		AboutMe:     r.FormValue("aboutme"),
		Password:    r.FormValue("password"),
	}

	
	data := models.Data{}

	passwordCrypted, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return
	}

	u.Password = string(passwordCrypted)

	errors := validators.Validate(u)

	if len(errors) > 0 {
		fmt.Println("Validation errors:")
		for _, err := range errors {
			fmt.Println(err)
		}
	} else {
		fmt.Println("User is valid!")
		//err = models.OrmInstance.Insert(u)

		_, _, err := queryBuilder.NewInsertBuilder().
			InsertInto("User", "Email",
				"FirstName",
				"LastName",
				"DateOfBirth",
				"Profil",
				"Username",
				"ImageURL",
				"AboutMe",
				"Password").
			Values(u.Email,
				u.FirstName,
				u.LastName,
				u.DateOfBirth,
				u.Profil,
				u.Username,
				u.ImageURL,
				u.AboutMe,
				u.Password).
			InsertQuery(models.OrmInstance.Db)

		fmt.Println(err)
		if err != nil {
			data.Error = err.Error()
			utils.SendJSON(w, http.StatusUnauthorized, data)
			return
		}
		data.Message = "Registeration successful"

		utils.SendJSON(w, http.StatusOK, data)
	}

}
