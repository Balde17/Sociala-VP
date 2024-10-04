package profile

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"backend/lib/orm/ORM/queryBuilder"
	"fmt"
	"net/http"
)

// id: 1, title: "Travel Moon", image: "/avatar.png", status: "Public Group", message: "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Nemo atque, odio soluta fuga accusamus itaque accusantium dolorem aut a nisi" },
func GetIdUser(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}
	_idUser := r.URL.Query().Get("idUser")

	var listPost []map[string]string
	rslt, err := queryBuilder.NewSelectBuilder().
		Select(
			"Id",
			"AboutMe",
		).
		From("User").
		Where("Id = " + _idUser).
		SelectQuery(models.OrmInstance.Db)

	if err == nil {
		for _, row := range rslt {
			listPost = append(listPost, map[string]string{
				"Id":      row[0],
				"AboutMe": row[1],
			})
		}
	}

	data.Content = listPost
	fmt.Println(listPost)
	utils.SendJSON(w, http.StatusOK, data)
}
