package followers

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"net/http"
)

func FollowUser(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}
	f := models.Followers{
		IdFollower: utils.GetInt(r.FormValue("idFollower"), w),
		IdFollowee: utils.GetInt(r.FormValue("idFollowee"), w),
		Status:     r.FormValue("status"),
	}
	err := utils.CreateData(w, "followers succesful", f)
	if err != nil {
		data.Error = err.Error()
		utils.SendJSON(w, http.StatusInternalServerError, data)
		return
	}
}
