package events

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"fmt"
	"net/http"
)

func CreateEventOption(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("type"), r.FormValue("idUser"), r.FormValue("idEvent"))
	createFunc := func() error {
		e := models.EventOptions{
			IdUser:  utils.GetInt(r.FormValue("idUser"), w),
			Type:    r.FormValue("type"),
			IdEvent: utils.GetInt(r.FormValue("idEvent"), w),
		}
		return utils.CreateData(w, "Event Created Succesfull", e)
	}

	utils.CreateEntity(w, r, createFunc, "Create Event successful json")
}
