package invite

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"fmt"
	"log"
	"net/http"
)

func UpdateInvite(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}
	r.ParseMultipartForm(10 << 20)
	idUser := utils.GetInt(r.FormValue("idUser"), w)
	idGroup := utils.GetInt(r.FormValue("idGroup"), w)
	idUserInvited := utils.GetInt(r.FormValue("idUserInvited"), w)
	fmt.Println(idUser, idGroup, idUserInvited)
	err := models.OrmInstance.UpdateInvite(idUser, idUserInvited, idGroup, "VALIDATE")
	fmt.Println("mmmmmmmm", err)
	if err != nil {
		log.Println(err.Error())
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusNotFound, data)
		return
	}

	// Envoyez une réponse JSON indiquant le succès de l'opération
	utils.SendJSON(w, http.StatusOK, models.Data{Message: "Create Update Profil successful json"})

}
