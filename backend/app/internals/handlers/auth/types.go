package auth

import (
	"backend/app/internals/models"
)

type Payload struct {
	User       models.User
	RemoteAddr string
	Sign       string
}
