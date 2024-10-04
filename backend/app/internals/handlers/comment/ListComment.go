package comment

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"backend/lib/orm/ORM/queryBuilder"
	"net/http"
)

func ListComment(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}
	_idPost := r.URL.Query().Get("idPost")
	var ListComment []map[string]string
	rslt, err := queryBuilder.NewSelectBuilder().

		// Union().
		Select(
			"c.Id as Id",
			"Content",
			"ImgUrl",
			"IdUser",
			"IdPost",
			"c.CreatedAt as CreatedAt ",
			"FirstName",
			"LastName",
			"ImageURL",
		).
		From("Comment as c").
		Join("inner", "User as u", "u.Id = c.IdUser").
		Where("IdPost = " + _idPost).
		OrderBy("c.CreatedAt DESC").
		SelectQuery(models.OrmInstance.Db)

	if err == nil {
		for _, row := range rslt {
			ListComment = append(ListComment, map[string]string{
				"Id":        row[0],
				"Content":   row[1],
				"ImgUrl":    row[2],
				"IdUser":    row[3],
				"IdPost":    row[4],
				"CreatedAt": row[5],
				"FirstName": row[6],
				"LastName":  row[7],
				"ImageURL":  row[8],
			})
		}
	}
	data.Content = ListComment
	utils.SendJSON(w, http.StatusOK, data)
}
