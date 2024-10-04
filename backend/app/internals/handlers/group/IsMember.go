package group

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"backend/lib/orm/ORM/queryBuilder"
	"net/http"
)

func IsMemberGroup(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}
	_idGroup := r.URL.Query().Get("idGroup")
	_idUser := r.URL.Query().Get("idUser")

	accepted := "VALIDATE"

	rslt, err := queryBuilder.NewSelectBuilder().
		Select(
			"u.Id as Id",
			"Username",
			"LastName",
			"FirstName",
			"Email",
			"AboutMe",
			"Profil",
			"u.CreatedAt as CreatedAt").
		From("User u").
		Join("inner", "Groups as g", "g.IdUser = u.Id").
		Where("g.Id = "+_idGroup).
		And("u.Id= "+_idUser).
		Union().
		Select(
			"u.Id as Id",
			"Username",
			"LastName",
			"FirstName",
			"Email",
			"AboutMe",
			"Profil",
			"u.CreatedAt as CreatedAt").
		From("User u").
		Join("inner", "Invite as i", "i.IdUserInvited = u.Id or i.IdUser = u.Id").
		Where("i.IdGroup = " + _idGroup).
		And("Accepted = '" + accepted + "'").
		And("u.Id= " + _idUser).
		SelectQuery(models.OrmInstance.Db)
	var notifications []map[string]string

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
				"CreatedAt": row[7],
			})
		}
	}

	data.Content = notifications
	utils.SendJSON(w, http.StatusOK, data)
}
