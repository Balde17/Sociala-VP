package notifications

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"backend/lib/orm/ORM/queryBuilder"
	"net/http"
)

// id: 1, title: "Travel Moon", image: "/avatar.png", status: "Public Group", message: "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Nemo atque, odio soluta fuga accusamus itaque accusantium dolorem aut a nisi" },
func UpdateNotification(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}
	idNotif := r.FormValue("idNotif")
	OptType := r.FormValue("type")

	_, err := queryBuilder.NewUpdateBuilder().
		Update("Notifications").
		Set("eventOption").Values(OptType).
		Where("Id = " + idNotif).Build(models.OrmInstance.Db)

	if err != nil {
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return
	}

	utils.SendJSON(w, http.StatusOK, data)
}
