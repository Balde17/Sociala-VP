package auth

import (
	"backend/app/internals/utils"
	"backend/lib/jwt"
	"fmt"
	"net/http"
	"os"
)

func CheckToken(w http.ResponseWriter, r *http.Request) bool {

	_jwt, err := jwt.Parse(utils.GetJWT(r))

	if err != nil {
		return false
	}

	_payload, _ := _jwt.Payload("payload")

	SECRET := os.Getenv("SECRET_KEY")

	payloadMap, ok := _payload.(map[string]interface{})
	if !ok {
		fmt.Println("Erreur lors de la conversion du payload en map.")
		return false
	}

	return SECRET == jwt.DecryptCode(payloadMap["Sign"].(string))
}