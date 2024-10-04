package post

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"net/http"
)

func CreateSocketPost(w http.ResponseWriter, r *http.Request, post models.PostData) error {
	p := models.Post{
		Title:    post.Title,
		ImgUrl:   post.ImgUrl,
		Content:  post.Content,
		IdUser:   post.IdUser,
		Status:   post.Status,
		IdGroups: utils.GetInt("1", w),
	}
	return utils.CreateData(w, "create post succesfull", p)
}
