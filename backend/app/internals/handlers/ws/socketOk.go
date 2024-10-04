package ws

import (
	"backend/app/internals/models"
	"backend/lib/orm/ORM/queryBuilder"
	"fmt"
)

// id: 1, title: "Travel Moon", image: "/avatar.png", status: "Public Group", message: "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Nemo atque, odio soluta fuga accusamus itaque accusantium dolorem aut a nisi" },
func SocketOK(sender, receiver string) bool {

	rslt1, err := queryBuilder.NewSelectBuilder().
		Select("count (*) as friend").
		From("Followers F").
		Join("INNER", "User U", "U.Id = F.IdFollower").
		Where("(F.IdFollowee =" + sender + " AND F.IdFollower =" + receiver + ")").
		SelectQuery(models.OrmInstance.Db)
	fmt.Println("------------is receiver follower: ", rslt1[0][0])

	if err != nil {
		return false
	}

	rslt2, err := queryBuilder.NewSelectBuilder().
		Select("count (U.Id) as friend").
		From("User U").
		Where("U.Id =" + sender).
		And("U.Profil='PUBLIC'").
		SelectQuery(models.OrmInstance.Db)
	fmt.Println("-----------is sender has a public profil: ", rslt1[0][0])

	if err != nil {
		return false
	}
	fmt.Println("is socket : ", rslt1[0][0] != "0" || rslt2[0][0] != "0")
	return rslt1[0][0] != "0" || rslt2[0][0] != "0"
}
