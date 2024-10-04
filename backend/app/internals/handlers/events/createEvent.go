package events

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"backend/lib/orm/ORM/queryBuilder"
	"fmt"
	"net/http"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}
	e := models.Event{
		IdUser:      utils.GetInt(r.FormValue("idUser"), w),
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Date:        r.FormValue("datetime"),
		Option:      "GOING",
		IdGroups:    utils.GetInt(r.FormValue("idGroup"), w),
	}
	_, _, lastInsertID, err := queryBuilder.NewInsertBuilder().
		InsertInto("Event",
			"IdUser",
			"Title",
			"Description",
			"Date",
			"Option",
			"IdGroups").
		Values(
			e.IdUser,
			e.Title,
			e.Description,
			e.Date,
			e.Option, e.IdGroups,
		).
		InsertQueryLastID(models.OrmInstance.Db)

	fmt.Println(err)
	if err != nil {
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusUnauthorized, data)
		return
	}
	data.Message = "Registeration successful"
	data.Content = lastInsertID
	utils.SendJSON(w, http.StatusOK, data)
}
