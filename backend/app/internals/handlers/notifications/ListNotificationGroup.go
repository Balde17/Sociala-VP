package notifications

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"backend/lib/orm/ORM/queryBuilder"
	"net/http"
)

// id: 1, title: "Travel Moon", image: "/avatar.png", status: "Public Group", message: "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Nemo atque, odio soluta fuga accusamus itaque accusantium dolorem aut a nisi" },
func ListNotificationGroup(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}
	_idUser := r.URL.Query().Get("idUser")

	var notifications []map[string]string

	rslt, err := queryBuilder.NewSelectBuilder().
		Select(
			"u.Id as Id",
			"Username",
			"LastName",
			"FirstName",
			"Email",
			"AboutMe",
			"Profil",
			"Title",
			"n.Type as Type",
			"n.CreatedAt as CreatedAt",
			"n.IdGroup as IdGroup",
			"n.Id as IdNotif",
		).
		From("User u").
		Join("inner", "Notifications as n", "n.Sender = u.Id").
		Join("inner", "Groups as g", "n.IdGroup=g.Id").
		Where("n.Receiver = " + _idUser).
		And("(n.Type = 'INVITE' OR n.Type = 'Groups')").
		OrderBy("n.CreatedAt DESC").
		SelectQuery(models.OrmInstance.Db)

	if err == nil {
		for _, row := range rslt {
			notifications = append(notifications, map[string]string{
				"Id":        row[0],
				"Username":  row[1],
				"LastName":  row[2],
				"FirstName": row[3],
				"Email":     row[4],
				"AboutMe":   row[5],
				"Profil":    row[6],
				"Title":     row[7],
				"Type":      row[8],
				"CreatedAt": row[9],
				"IdGroup":   row[10],
				"IdNotif":   row[11],
			})
		}
	}

	data.Content = notifications
	utils.SendJSON(w, http.StatusOK, data)
}
