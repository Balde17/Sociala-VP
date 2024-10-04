package group

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"net/http"
)

func CreateSocketGroup(w http.ResponseWriter, r *http.Request, group models.GroupData) error {
	//idUser, err := strconv.Atoi("1")
	//idGroup, err0 := strconv.Atoi("1")
	// if err != nil || err0 != nil {
	// 	return err
	// }
	g := models.Groups{
		Title:       group.Title,
		ImgUrl:      group.ImgUrl,
		Description: group.Description,
		IdUser:      group.IdUser,
	}
	return utils.CreateData(w, "Create Group successful json", g)
}
