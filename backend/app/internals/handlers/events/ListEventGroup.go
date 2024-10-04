package events

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"backend/lib/orm/ORM/queryBuilder"
	"net/http"
)

func ListEventGroup(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}

	idGroup := r.URL.Query().Get("idGroup")

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
			"g.Title as Title",
			"e.Title as TitleEvent",
			"e.Id as IdEvent",
			"e.Description as Description",
		).
		From("User u").
		Join("inner", "Event as e", " e.IdUser=u.Id").
		Join("inner", "Groups as g", "e.IdGroups=g.Id").
		Where("e.IdGroups = " + idGroup).
		SelectQuery(models.OrmInstance.Db)

	if err == nil {
		for _, row := range rslt {
			notifications = append(notifications, map[string]string{
				"Id":          row[0],
				"Username":    row[1],
				"LastName":    row[2],
				"FirstName":   row[3],
				"Email":       row[4],
				"AboutMe":     row[5],
				"Profil":      row[6],
				"Title":       row[7],
				"TitleEvent":  row[8],
				"IdEvent":     row[9],
				"Description": row[10],
			})
		}
	}
	data.Content = notifications
	utils.SendJSON(w, 200, data)
}
