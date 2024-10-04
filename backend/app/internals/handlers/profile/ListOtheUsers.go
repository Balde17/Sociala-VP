package profile

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"backend/lib/orm/ORM/queryBuilder"
	"net/http"
)

// id: 1, title: "Travel Moon", image: "/avatar.png", status: "Public Group", message: "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Nemo atque, odio soluta fuga accusamus itaque accusantium dolorem aut a nisi" },
func ListOtherUsers(w http.ResponseWriter, r *http.Request) {
	data := models.Data{}
	idUser := r.URL.Query().Get("idUser")
	var posts []map[string]string

	rslt, err := queryBuilder.NewSelectBuilder().
		Select("Id", "Username", "FirstName", "Email", "Password", "AboutMe", "Profil", "DateOfBirth", "ImageURL").
		From("User").
		Where("Id !="+idUser).
		GroupBy("Email").
		EXCEPT().
		Select("U.Id as UserId", "Username", "FirstName", "Email", "Password", "AboutMe", "Profil", "DateOfBirth", "ImageURL").
		From("User U").
		Join("INNER", "Followers F", "U.Id= F.IdFollower or U.Id=F.IdFollowee").
		Where("(F.IdFollower = " + idUser + " OR F.IdFollowee = " + idUser + ")").
		And("F.Status = 'VALIDATE'").
		GroupBy("U.Email").
		SelectQuery(models.OrmInstance.Db)

	if err == nil {
		for _, row := range rslt {
			posts = append(posts, map[string]string{
				"Id":          row[0],
				"Username":    row[1],
				"FirstName":   row[2],
				"Email":       row[3],
				"Password":    row[4],
				"AboutMe":     row[5],
				"Profil":      row[6],
				"DateOfBirth": row[7],
				"ImageURL":    row[8],
			})
		}
	}

	data.Content = posts
	utils.SendJSON(w, http.StatusOK, data)
}
