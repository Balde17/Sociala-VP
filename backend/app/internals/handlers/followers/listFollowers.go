package followers

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"net/http"
)

func FollowerById(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}
	IdFollower := utils.GetInt(r.URL.Query().Get("idFollower"), w)
	IdFollowee := utils.GetInt(r.URL.Query().Get("idFollowee"), w)
	models.OrmInstance.Custom.Where("IdFollower", IdFollower).And("IdFollowee", IdFollowee)
	_followers, err := models.OrmInstance.Scan(models.Followers{}, "IdFollowee", "IdFollower", "Status")
	models.OrmInstance.Custom.Clear()
	
	if err != nil {
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return
	}
	followers := _followers.([]struct {
		IdFollowee int
		IdFollower int
		Status     string
	})
	data.Content = followers
	utils.SendJSON(w, http.StatusOK, data)
}
