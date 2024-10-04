package models

type Data struct {
	Id      int64
	Error   string
	Message string
	Content interface{}
	Token   string
}
