package group

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"backend/lib/orm/ORM/queryBuilder"
	"net/http"
)

func UserIsInvited(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}
	idUser := r.FormValue("idUser")
	accepted := "PENDING"
	rslt, err := queryBuilder.NewSelectBuilder().

		// Union().
		Select(
			"g.Id as Id",
			"Title",
			"Description",
			"ImgUrl",
			"g.CreatedAt as CreatedAt",
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
				"Title":       row[1],
				"Description": row[2],
				"ImgUrl":      row[3],
				"CreatedAt":   row[4],
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
