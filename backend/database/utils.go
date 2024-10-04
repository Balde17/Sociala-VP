package database

import (
	"backend/app/internals/models"
	"backend/lib/validators"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

func LoadEnv(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err.Error())
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Println("Your env file must be set")
		}
		key := parts[0]
		value := parts[1]
		err := os.Setenv(key, value)
		if err != nil {
			return err
		}
	}
	return scanner.Err()
}

func SendJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func EncodeJson(w http.ResponseWriter, typ string, _posts interface{}, data models.Data, clients map[*websocket.Conn]string) error {
	jsonData, err := json.Marshal(map[string]interface{}{"type": typ, "data": _posts})
	if err != nil {
		data.Error = err.Error()
		SendJSON(w, http.StatusInternalServerError, data)
		return err
	}
	for client := range clients {
		if err := client.WriteMessage(websocket.TextMessage, jsonData); err != nil {
			log.Printf("Erreur lors de l'envoi des posts au client: %v", err)
		}
	}
	return nil
}

func CreateData(w http.ResponseWriter, message string, nameModel interface{}) error {
	var err error
	data := models.Data{}
	_errors := validators.Validate(nameModel)
	if len(_errors) > 0 {
		fmt.Println("Validation errors:")
		__errors := ""
		for _, err := range _errors {
			fmt.Println(err)
			__errors += err
		}
		data.Error = __errors
		SendJSON(w, http.StatusInternalServerError, data)
		return errors.New(__errors)
	} else {
		err = models.OrmInstance.Insert(nameModel)
	}
	if err != nil {
		data.Error = err.Error()
		SendJSON(w, http.StatusInternalServerError, data)
		return err
	}
	data.Message = message
	SendJSON(w, http.StatusOK, data)
	return nil
}

func GetJWT(r *http.Request) string {

	cookie, err := r.Cookie("jwt")
	if err != nil {
		return ""
	}
	return cookie.Value
}

func GetInt(valueStr string, w http.ResponseWriter) int {
	data := models.Data{}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		data.Error = err.Error()
		SendJSON(w, http.StatusInternalServerError, data)
		return 0
	}
	return value
}

func CreateEntity(w http.ResponseWriter, r *http.Request, createFunc func() error, successMessage string) {
	r.ParseMultipartForm(10 << 20)

	err := createFunc()
	if err != nil {
		SendJSON(w, http.StatusInternalServerError, models.Data{Error: err.Error()})
		return
	}

	SendJSON(w, http.StatusOK, models.Data{Message: successMessage})
}
