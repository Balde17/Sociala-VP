package invite

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"backend/lib/orm/ORM/queryBuilder"
	"net/http"
)

func DeleteInvite(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}
	r.ParseMultipartForm(10 << 20)
	idUser := r.FormValue("idUser")
	idGroup := r.FormValue("idGroup")
	idUserInvited := r.FormValue("idUserInvited")
	_, _, err := queryBuilder.NewDeleteBuilder().
		DeleteValues("Invite").
		Where("IdUser = " + idUser + " And IdUserInvited = " + idUserInvited + " And idGroup = " + idGroup).
		DeleteQuery(models.OrmInstance.Db)
	if err != nil {
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return
	}
	data.Message = "Updatesuccessful json"
	utils.SendJSON(w, http.StatusOK, data)
}
