package like

import (
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"backend/lib/validators"
	"fmt"
	"log"
	"net/http"
)

func CreateSocketLike(w http.ResponseWriter, r *http.Request, like models.LikeData) error {
	models.OrmInstance.Custom.Where("IdObject", like.IdObject).And("Type", like.Type)
	_like, err := models.OrmInstance.Scan(models.Like{}, "Type", "IdObject", "Liked")
	models.OrmInstance.Custom.Clear()

	fmt.Println("LIKE SOCKET", _like)
	if err != nil {
		fmt.Println(err)
		return err
	}
	likes, ok := _like.([]struct {
		Type     string
		IdObject int
		Liked    int
	})
	if len(likes) != 0 || !ok {
		fmt.Println("---------------------------->")
		err := models.OrmInstance.UpdateLike(like.Liked, like.IdObject, like.Type)
		fmt.Println("mmmmmmmm", err)
		if err != nil {
			fmt.Println(err)
		}
		return err
	} else {
		// Créez un nouveau like
		l := models.Like{
			Type:     like.Type,
			Liked:    like.Liked,
			IdObject: like.IdObject,
		}
		// Validez le nouveau like
		errors := validators.Validate(l)
		if len(errors) > 0 {
			fmt.Println("Validation errors:")
			for _, err := range errors {
				fmt.Println(err)
			}
			return err
		}
		// Insérez le nouveau like dans la base de données
		err = models.OrmInstance.Insert(l)
		if err != nil {
			log.Fatal(err)
		}
	}
	// Envoyez une réponse JSON indiquant le succès de l'opération
	utils.SendJSON(w, http.StatusOK, models.Data{Message: "Create Like successful json"})
	return nil
}
