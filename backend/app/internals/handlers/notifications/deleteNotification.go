package notifications

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"fmt"
	"log"
	"net/http"
)

func DeleteNotification(w http.ResponseWriter, r *http.Request) {

	data := models.Data{}
	r.ParseMultipartForm(10 << 20)
	idNotif := utils.GetInt(r.FormValue("idNotif"), w)
	fmt.Println("idNotif", idNotif)
	err := models.OrmInstance.DeleteNotification(idNotif)
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
