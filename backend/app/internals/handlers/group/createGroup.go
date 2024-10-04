package group

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"net/http"
)

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	//user := jwt.GetPayload(r)

	//fmt.Println(user)

		g := models.Groups{
			Title:       r.FormValue("title"),
			ImgUrl:      r.FormValue("imageUrlFromChild"),
			Description: r.FormValue("description"),
			//IdUser:      utils.GetInt(strconv.Itoa(int(user["Id"].(float64))), w),
			IdUser: utils.GetInt(r.FormValue("idUser"), w),
		}
		 utils.CreateData(w, "Post Created Succesfull", g)
	

}
