package profile

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func UpdateProfil(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}
	r.ParseMultipartForm(10 << 20)
	Id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Println(err.Error())
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusNotFound, data)
		return
	}
	err = models.OrmInstance.UpdateStatut(r.FormValue("status"), Id)
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
