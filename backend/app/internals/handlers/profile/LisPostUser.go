package profile

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"backend/lib/orm/ORM/queryBuilder"
	"net/http"
)

// id: 1, title: "Travel Moon", image: "/avatar.png", status: "Public Group", message: "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Nemo atque, odio soluta fuga accusamus itaque accusantium dolorem aut a nisi" },
func ListPostUser(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}
	_idUser := r.URL.Query().Get("idUser")

	var listPost []map[string]string
	rslt, err := queryBuilder.NewSelectBuilder().
		Select(
			"p.Id as Id",
			"Title",
			"Content",
			"ImgUrl",
			"Status",
			"IdUser",
			"IdGroups",
			"p.CreatedAt as CreatedAt",
			"FirstName",
			"LastName",
			"ImageURL",
		).
		From("Post as p").
		Join("inner", "User as u", "u.Id = p.IdUser").
		Where("IdUser = " + _idUser).
		And("Status != 'GROUP'").
		OrderBy("p.CreatedAt DESC").
		SelectQuery(models.OrmInstance.Db)

	if err == nil {
		for _, row := range rslt {
			listPost = append(listPost, map[string]string{
				"Id":        row[0],
				"Title":     row[1],
				"Content":   row[2],
				"ImgUrl":    row[3],
				"Status":    row[4],
				"IdUser":    row[5],
				"IdGroups":  row[6],
				"CreatedAt": row[7],
				"FirstName": row[8],
				"LastName":  row[9],
				"ImageURL":  row[10],
			})
		}
	}

	data.Content = listPost
	utils.SendJSON(w, http.StatusOK, data)
}
