package invite

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"backend/lib/validators"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func InviteUser(w http.ResponseWriter, r *http.Request) {
	err := errors.New("")
	data := models.Data{}
	r.ParseMultipartForm(10 << 20)
	idGroup := utils.GetInt(r.FormValue("idGroup"), w)
	idUser := utils.GetInt(r.FormValue("idUser"), w)
	if idUser == 0 {
		models.OrmInstance.Custom.Where("Id", idGroup)
		_idUser, err := models.OrmInstance.Scan(models.Groups{}, "IdUser")
		if err != nil {
			log.Println(err.Error())
			data.Error = err.Error()
			utils.SendJSON(w, http.StatusInternalServerError, data)
			return
		}
		idUser_, ok := _idUser.([]struct {
			IdUser int
		})
		if !ok {
			data.Error = "Error conversion from struct"
			utils.SendJSON(w, http.StatusInternalServerError, data)
			return
		}
		idUser = idUser_[0].IdUser
	}
	idUserInvited := utils.GetInt(r.FormValue("idUserInvited"), w)
	fmt.Println(idGroup, idUserInvited, idUser, r.FormValue("accepted"))
	p := models.Invite{

		IdUser:        idUser,
		IdUserInvited: idUserInvited,
		IdGroup:       idGroup,
		Accepted:      r.FormValue("accepted"),
	}

	errors := validators.Validate(p)
	if len(errors) > 0 {
		fmt.Println("Validation errors:")
		for _, err := range errors {
			fmt.Println(err)
		}
	} else {
		err = models.OrmInstance.Insert(p)
		models.OrmInstance.Custom.Clear()
		fmt.Println(err)
		fmt.Println("Invitation successful")
	}
	if err != nil {
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return
	}
	data.Message = "Invitationsuccessful json"
	utils.SendJSON(w, http.StatusOK, data)
}
