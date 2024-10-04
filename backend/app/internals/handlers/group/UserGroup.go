package group

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"backend/lib/orm/ORM/queryBuilder"
	"net/http"
)

func UserGroup(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}
	idUser := r.FormValue("idUser")
	accepted := "VALIDATE"
	rslt, err := queryBuilder.NewSelectBuilder().

		// Union().
		Select(
			"Id",
			"CreatedAt",
			"Title",
			"Description",
			"ImgUrl",
			"IdUser",
		).
		From("Groups").
		Where("idUser = "+idUser).
		Union().
		Select(
			"g.Id as Id",
			"g.CreatedAt as CreatedAt",
			"Title",
			"Description",
			"ImgUrl",
			"g.IdUser",
		).
		From("Groups g").
		Join("inner", "Invite as i", "i.Idgroup = g.Id").
		Where("i.IdUserInvited = " + idUser).And("i.Accepted = '" + accepted + "'").
		SelectQuery(models.OrmInstance.Db)
	var listOtherGroup []map[string]string

	if err == nil {
		for _, row := range rslt {
			listOtherGroup = append(listOtherGroup, map[string]string{
				"Id":          row[0],
				"CreatedAt":   row[1],
				"Title":       row[2],
				"Description": row[3],
				"ImgUrl":      row[4],
				"IdUser":      row[5],
			})
		}
	}

	// if listOtherGroup == nil {
	// 	rslt, err := queryBuilder.NewSelectBuilder().
	// 		Select(
	// 			"Id",
	// 			"Title",
	// 			"Description",
	// 			"ImgUrl",
	// 			"CreatedAt",
	// 		).
	// 		From("Groups").
	// 		Where("IdUser != " + idUser).
	// 		SelectQuery(models.OrmInstance.Db)

	// 	if err == nil {
	// 		for _, row := range rslt {
	// 			listOtherGroup = append(listOtherGroup, map[string]string{
	// 				"Id":          row[0],
	// 				"Title":       row[1],
	// 				"Description": row[2],
	// 				"CreatedAt":   row[3],
	// 			})
	// 		}
	// 	}
	// }

	data.Content = listOtherGroup
	utils.SendJSON(w, http.StatusOK, data)
}
