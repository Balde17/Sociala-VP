package comment

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"net/http"
)

func CreateSocketComment(w http.ResponseWriter, r *http.Request, comment models.CommentData) error {
	c := models.Comment{
		Content: comment.Content,
		//ImgUrl:      Comment.ImgUrl,
		IdPost: comment.IdPost,
	}
	return utils.CreateData(w, "Create Comment successful json", c)
}
