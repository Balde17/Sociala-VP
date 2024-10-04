package group

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"log"
	"net/http"
	"time"
)

func ListIdGroup(w http.ResponseWriter, r *http.Request) {
	idGroup := utils.GetInt(r.FormValue("idGroup"), w)

	data := models.Data{}
	models.OrmInstance.Custom.Where("Id", idGroup)
	_groups, err := models.OrmInstance.Scan(models.Groups{}, "Id", "Title", "Description", "ImgUrl", "CreatedAt")
	models.OrmInstance.Custom.Clear()

	if err != nil {
		log.Println(err.Error())
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return
	}
	groups, ok := _groups.([]struct {
		Id          int64
		Title       string
		Description string
		ImgUrl      string
		CreatedAt   time.Time
	})
	if !ok {
		data.Error = "Error conversion from struct"
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return
	}
	data.Content = groups
	utils.SendJSON(w, http.StatusOK, data)
}
