package post

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	//user := jwt.GetPayload(r)
	idUser := utils.GetInt(r.FormValue("idUser"), w)
	idGroup := utils.GetInt(r.FormValue("idGroup"), w)
	fmt.Println(idUser, idGroup, r.FormValue("title"), r.FormValue("imageUrlFromChild"), r.FormValue("status"))
	createFunc := func() error {
		p := models.Post{
			Title:    r.FormValue("title"),
			ImgUrl:   r.FormValue("imageUrlFromChild"),
			Content:  r.FormValue("content"),
			IdUser:   idUser,
			IdGroups: idGroup,
			Status:   r.FormValue("status"),
		}
		ok, err := utils.ValidateStruct(p)
		if !ok {
			return err
		}

		postId, err := models.OrmInstance.Insert1(p)
		if err != nil {
			return err
		}

		if p.Status == "ALMOST" {
			_selected := strings.Split(r.FormValue("selectedFollowers"), ",")
			for _, followerID := range _selected {
				followerIdInt, err := strconv.Atoi(followerID)
				if err != nil {
					continue
				}
				v := models.Visibility{
					IdPost:   int(postId),
					Type:     "POST",
					IdObject: followerIdInt,
				}

				ok, err := utils.ValidateStruct(v)
				if !ok {
					return err
				}

				_, err = models.OrmInstance.Insert1(v)
				if err != nil {
					return err
				}
			}
		}

		return nil
	}

	utils.CreateEntity(w, r, createFunc, "Create Post successful json")
}
