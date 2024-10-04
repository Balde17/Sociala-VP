package notifications

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"backend/lib/orm/ORM/queryBuilder"
	"log"
	"net/http"
)

func DeleteNotif(w http.ResponseWriter, r *http.Request) {

	data := models.Data{}
	r.ParseMultipartForm(10 << 20)
	sender := r.FormValue("sender")
	receiver := r.FormValue("receiver")
	_, _, err := queryBuilder.NewDeleteBuilder().
		DeleteValues("Notifications").
		Where("Sender = " + sender + " And Receiver = " + receiver + " And Type = 'REQUEST'").
		DeleteQuery(models.OrmInstance.Db)
	if err != nil {
		log.Println(err.Error())
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusNotFound, data)
		return
	}

	// Envoyez une réponse JSON indiquant le succès de l'opération
	utils.SendJSON(w, http.StatusOK, models.Data{Message: "Delete notification successful json"})

}
