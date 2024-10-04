package comment

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"net/http"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {

	//user := jwt.GetPayload(r)

	//fmt.Println(user)
	r.ParseMultipartForm(10 << 20)
	c := models.Comment{
		Content: r.FormValue("content"),
		IdPost:  utils.GetInt(r.FormValue("idPost"), w),
		ImgUrl:  r.FormValue("imageUrlFromChild"),
		IdUser:  utils.GetInt(r.FormValue("idUser"), w),
	}
	utils.CreateData(w, "Comment Created Succesfull", c)
}
