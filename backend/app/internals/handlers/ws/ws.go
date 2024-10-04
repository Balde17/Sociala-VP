package ws

import (
	"backend/app/internals/handlers/chat"
	"backend/app/internals/handlers/comment"
	"backend/app/internals/handlers/group"
	"backend/app/internals/handlers/like"
	"backend/app/internals/handlers/notifications"
	"backend/app/internals/utils"
	"fmt"
	"strconv"

	//"backend/app/internals/handlers/like"
	"backend/app/internals/models"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

var (
	Clients  = make(map[*websocket.Conn]string)
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func extractUserIdFromWebSocketURL(url *url.URL) string {
	// Récupérer la valeur du paramètre userId de l'URL
	queryParams := url.Query()
	userId := queryParams.Get("userId")
	return userId
}

func findClientByID(ID string, clients map[*websocket.Conn]string) *websocket.Conn {

	for conn, _ := range Clients {

		if ID == Clients[conn] {
			return conn
		}
	}
	return nil
}

func WS(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if ws == nil {
		if err != nil {
			fmt.Println("conn nil")
			return
		}
	}
	if err != nil {
		log.Println(err.Error())
	}
	defer ws.Close()
	fmt.Println("connected")
	// TODO: Token
	if err != nil {
		log.Printf("Échec de récupération de l'utilisateur: %v", err)
		return
	}
	userId := extractUserIdFromWebSocketURL(r.URL)

	conn := findClientByID(userId, Clients)
	if conn == nil {
		Clients[ws] = userId
		fmt.Println("new connexion")
	} else {
		delete(Clients, conn)
		Clients[ws] = userId
		fmt.Println("connexion reintialisée")
	}
	fmt.Println("Clients[ws] - ", Clients[ws])
	fmt.Println("Clients - ", Clients)
	for {
		var generic models.Generic
		err := ws.ReadJSON(&generic)
		if err != nil {
			log.Printf("error: %v", err)
			delete(Clients, ws)
			break
		}
		switch generic.Type {
		case "addComment":
		case "post":
		case "getAllPosts":
		case "getComments":
		case "notification":
			fmt.Println("--------------------------notification data received:")

			var notificationData models.NotificationData
			if err := json.Unmarshal(generic.Data, &notificationData); err != nil {
				log.Println("Unmarshal notificationData error:", err)
				return
			}

			fmt.Println("--------------------------notification data received:", notificationData)
			notifications.CreateSocketNotication(w, r, notificationData)
			Client := findClientByID(strconv.Itoa(notificationData.Receiver), Clients)
			if Client != nil {
				notifications.ListSocketNotification(w, r, Client, notificationData.IdGroup, notificationData.Receiver, notificationData.Type)
			}
			// post.ListSocketPost(w, r, Clients)
		case "notificationEvent":
			var eventDataMessage models.EventData
			if err := json.Unmarshal(generic.Data, &eventDataMessage); err != nil {
				log.Println("Unmarshal likeData error:", err)
				continue
			}
			idSender := strconv.Itoa(eventDataMessage.Sender)
			idGroup := strconv.Itoa(eventDataMessage.IdGroup)

			members := group.GetMemberGroup(idGroup)
			isGroupMember := IsGroupMember(idSender, members)

			//client := findClientByID(idReceiver, Clients)
			fmt.Println("isGroupMember : ", isGroupMember)
			fmt.Println("err : ", err)

			if err == nil && idSender != "" && idGroup != "" && isGroupMember {
				jsonData, err2 := json.Marshal(eventDataMessage)

				for _, member := range members {
					notif := models.NotificationData{
						Type:     eventDataMessage.Type,
						Sender:   eventDataMessage.Sender,
						Receiver: utils.GetInt(member["Id"], w),
						IdGroup:  eventDataMessage.IdGroup,
						IdEvent:  eventDataMessage.IdEvent,
					}
					err1 := notifications.CreateSocketNotication(w, r, notif)
					if err1 != nil && err2 != nil {
						client := findClientByID(member["Id"], Clients)
						if client != nil && member["Id"] != idSender {
							client.WriteMessage(websocket.TextMessage, jsonData)
						}
					}
				}
			} else {
				errorMsg := map[string]string{
					"error":    "Erreur d'envoi au groupe",
					"sender":   idSender,
					"receiver": idSender,
				}
				jsonData, _ := json.Marshal(errorMsg)

				clientSender := findClientByID(idSender, Clients)
				if clientSender != nil {
					if err := clientSender.WriteMessage(websocket.TextMessage, jsonData); err != nil {
						log.Printf("Erreur lors de l'envoi du message d'erreur au client: %v", err)
					}
				}

			}
		case "message":
			fmt.Println(string(generic.Data))
		case "createPost":

		case "createGroup":
			var groupData models.GroupData
			if err := json.Unmarshal(generic.Data, &groupData); err != nil {
				log.Println("Unmarshal postData error:", err)
				return
			}
			group.CreateSocketGroup(w, r, groupData)
			group.ListSocketGroup(w, r, Clients)

		case "createComment":
			var commentData models.CommentData
			if err := json.Unmarshal(generic.Data, &commentData); err != nil {
				log.Println("Unmarshal postData error:", err)
				return
			}
			fmt.Println("Post data received:", commentData)
			comment.CreateSocketComment(w, r, commentData)
			comment.ListSocketComment(w, r, Clients, commentData.IdPost)
		case "createLike":
			var likeData models.LikeData

			if err := json.Unmarshal(generic.Data, &likeData); err != nil {
				log.Println("Unmarshal likeData error:", err)
				return
			}
			fmt.Println("like post data received:", likeData)
			like.CreateSocketLike(w, r, likeData)
			like.ListSocketLike(w, r, Clients)
		case "privateMessage":
			var privateMessage models.ChatData
			if err := json.Unmarshal(generic.Data, &privateMessage); err != nil {
				log.Println("Unmarshal likeData error:", err)
				return
			}
			idSender := strconv.Itoa(privateMessage.Sender)
			idReceiver := strconv.Itoa(privateMessage.Receiver)

			client := findClientByID(idReceiver, Clients)
			fmt.Println(client, " found for ", idReceiver)
			if err == nil && idSender != "" && idReceiver != "" {
				socketOk := SocketOK(idSender, idReceiver)
				isFriend := IsFriend(idSender, idReceiver)

				if !isFriend {
					// Préparer un message d'erreur
					errorMsg := map[string]string{
						"error":    "Vous ne pouvez envoyer de messages privés qu'à vos amis.",
						"sender":   idSender,
						"receiver": idReceiver,
					}
					jsonData, err := json.Marshal(errorMsg)
					if err != nil {
						log.Printf("Erreur lors de la sérialisation du message d'erreur: %v", err)

					}

					// Envoyer le message d'erreur à l'expéditeur
					clientSender := findClientByID(idSender, Clients)
					if clientSender != nil {
						if err := clientSender.WriteMessage(websocket.TextMessage, jsonData); err != nil {
							log.Printf("Erreur lors de l'envoi du message d'erreur au client: %v", err)
						}
					} else {
						log.Printf("Client expéditeur non trouvé.")
					}

				} else {
					err1 := chat.RegisterMessage(&privateMessage)
					fmt.Println("insertion-------", err1)
					jsonData, err2 := json.Marshal(privateMessage)
					if err1 == nil && err2 == nil && client != nil && socketOk {
						if err := client.WriteMessage(websocket.TextMessage, jsonData); err != nil {
							log.Printf("Erreur lors de l'envoi des posts au client: %v", err)
						}
					}
				}

				fmt.Println("like post data received:", privateMessage)
			}
		case "groupMessage":
			var groupMessage models.ChatData
			if err := json.Unmarshal(generic.Data, &groupMessage); err != nil {
				log.Println("Unmarshal likeData error:", err)
				continue
			}
			idSender := strconv.Itoa(groupMessage.Sender)
			idGroup := strconv.Itoa(groupMessage.Receiver)

			members := group.GetMemberGroup(idGroup)
			isGroupMember := IsGroupMember(idSender, members)

			//client := findClientByID(idReceiver, Clients)
			fmt.Println("isGroupMember : ", isGroupMember)
			fmt.Println("err : ", err)
			if err == nil && idSender != "" && idGroup != "" && isGroupMember {

				err1 := chat.RegisterMessage(&groupMessage)
				jsonData, err2 := json.Marshal(groupMessage)

				if err1 == nil && err2 == nil {
					for _, member := range members {
						fmt.Println("member : ", member["Id"], " user :", member["Email"], string(jsonData))
						client := findClientByID(member["Id"], Clients)
						if client != nil && member["Id"] != idSender {
							client.WriteMessage(websocket.TextMessage, jsonData)
						}
					}
				}
			} else {
				errorMsg := map[string]string{
					"error":    "Erreur d'envoi au groupe",
					"sender":   idSender,
					"receiver": idSender,
				}
				jsonData, _ := json.Marshal(errorMsg)

				clientSender := findClientByID(idSender, Clients)
				if clientSender != nil {
					if err := clientSender.WriteMessage(websocket.TextMessage, jsonData); err != nil {
						log.Printf("Erreur lors de l'envoi du message d'erreur au client: %v", err)
					}
				}

			}

		}

	}
	delete(Clients, ws)
}
