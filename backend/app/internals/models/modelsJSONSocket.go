package models

import "encoding/json"

type PostData struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	ImgUrl  string `json:"imgUrl"`
	Status  string `json:"status"`
	IdUser  int    `json:"idUser"`
}
type NotificationData struct {
	Type     string `json:"type"`
	Sender   int    `json:"sender"`
	Receiver int    `json:"receiver"`
	IdGroup  int    `json:"idGroup"`
	IdEvent  int    `json:"idEvent"`
}

type ChatData struct {
	Type        string `json:"type"`
	Sender      int    `json:"sender"`
	Receiver    int    `json:"receiver"`
	Content     string `json:"content"`
	MessageDate string `json:"messageDate"`
	Name        string `json:"name"`
}

type EventData struct {
	Type        string `json:"type"`
	Sender      int    `json:"sender"`
	IdGroup     int    `json:"idGroup"`
	IdEvent     int    `json:"idEvent"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Datetime    string `json:"datetime"`
}

type GroupData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImgUrl      string `json:"imgUrl"`
	IdUser      int    `json:"idUser"`
}

type Generic struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type CommentData struct {
	Content string `json:"content"`
	IdPost  int    `json:"idPost"`
	//ImgUrl      string `json:"imgUrl"`
}

type LikeData struct {
	Type     string `json:"type"`
	Liked    int    `json:"liked"`
	IdObject int    `json:"idObject"`
}

type InviteData struct {
	IdUser        int `json:"idUser"`
	IdUserInvited int `json:"idUserInvited"`
	IdGroup       int `json:"idGroup"`
	Accepted      int `json:"accepted"`
}
