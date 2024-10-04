package ws

// id: 1, title: "Travel Moon", image: "/avatar.png", status: "Public Group", message: "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Nemo atque, odio soluta fuga accusamus itaque accusantium dolorem aut a nisi" },
func IsGroupMember(idUser string, members []map[string]string) bool {
	for _, member := range members {
		if member["Id"] == idUser {
			return true
		}
	}
	return false
}
