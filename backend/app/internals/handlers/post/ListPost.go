package post

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"backend/lib/orm/ORM/queryBuilder"
	"fmt"
	"net/http"
)

// id: 1, title: "Travel Moon", image: "/avatar.png", status: "Public Group", message: "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Nemo atque, odio soluta fuga accusamus itaque accusantium dolorem aut a nisi" },
func ListPost(w http.ResponseWriter, r *http.Request) {

	data := models.Data{}
	_idUser := r.URL.Query().Get("idUser")
	_idGroup := r.URL.Query().Get("idGroup")
	var listPost []map[string]string

	fmt.Println(_idGroup, _idUser)
	if _idGroup == "0" {
		rslt, err := queryBuilder.NewSelectBuilder().

			// Union().
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
			Where("IdGroups = "+_idGroup).
			And("Status = 'PUBLIC'").
			Union().
			Select(
				"p.Id as Id",
				"Title",
				"Content",
				"ImgUrl",
				"p.Status as Status",
				"IdUser",
				"IdGroups",
				"p.CreatedAt as CreatedAt",
				"FirstName",
				"LastName",
				"ImageURL",
			).
			From("Post as p").
			Join("inner", "User as u", "u.Id = p.IdUser").
			Where("IdGroups = "+_idGroup).
			And("(u.Id = "+_idUser+")").
			And("p.Status = 'PRIVATE'").
			Union().
			Select(
				"p.Id as Id",
				"Title",
				"Content",
				"ImgUrl",
				"p.Status as Status",
				"IdUser",
				"IdGroups",
				"p.CreatedAt as CreatedAt",
				"FirstName",
				"LastName",
				"ImageURL",
			).
			From("Post as p").
			Join("inner", "Followers as f", "f.IdFollowee = p.IdUser").
			Join("inner", "User as u", "u.Id = p.IdUser").
			Where("IdGroups = "+_idGroup).
			And("(IdFollower = "+_idUser+" or IdFollowee = "+_idUser+")").
			And("p.Status = 'PRIVATE'").
			And("f.Status = 'VALIDATE'").
			Union().
			Select(
				"p.Id as Id",
				"Title",
				"Content",
				"ImgUrl",
				"p.Status as Status",
				"IdUser",
				"IdGroups",
				"p.CreatedAt as CreatedAt",
				"FirstName",
				"LastName",
				"ImageURL",
			).
			From("Post as p").
			Join("inner", "Visibility as v", "v.IdPost= p.Id").
			Join("inner", "User as u", "u.Id = p.IdUser").
			Where("IdGroups = " + _idGroup).
			And("(v.IdObject = " + _idUser + " or p.IdUser = " + _idUser + ")").
			And("p.Status = 'ALMOST'").
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
	} else {
		rslt, err := queryBuilder.NewSelectBuilder().

			// Union().
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
			Where("IdGroups = " + _idGroup).
			And("Status = 'GROUP'").
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
	}
	data.Content = listPost
	utils.SendJSON(w, http.StatusOK, data)
}
