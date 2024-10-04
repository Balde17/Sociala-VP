package ws

import (
	"backend/app/internals/models"
	"backend/lib/orm/ORM/queryBuilder"
	"fmt"
)

// id: 1, title: "Travel Moon", image: "/avatar.png", status: "Public Group", message: "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Nemo atque, odio soluta fuga accusamus itaque accusantium dolorem aut a nisi" },
func IsFriend(sender, receiver string) bool {
	rslt, err := queryBuilder.NewSelectBuilder().
		Select("count (*) as friend").
		From("Followers F").
		Where("(F.IdFollowee =" + sender + " AND F.IdFollower =" + receiver + ") OR (F.IdFollowee =" + receiver + " AND F.IdFollower =" + sender + ")").
		SelectQuery(models.OrmInstance.Db)
	fmt.Println("--------------is sender and receiver friends: ", rslt[0][0] != "0")
	if err != nil {
		return false
	}
	fmt.Println(rslt[0], len(rslt))
	return rslt[0][0] != "0"
}
