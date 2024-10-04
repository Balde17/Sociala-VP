package models

import (
	orm "backend/lib/orm/ORM"
)

var (
	OrmInstance = orm.NewORM()
)

type User struct {
	orm.Model
	Username    string
	Email       string `orm-go:"NOT NULL UNIQUE" validate:"email"`
	Password    string `validate:"password"`
	FirstName   string `orm-go:"NOT NULL" validate:"required,min=4,max=20"`
	LastName    string `orm-go:"NOT NULL" validate:"required,min=4,max=20"`
	DateOfBirth string `orm-go:"NOT NULL"`
	ImageURL    string
	AboutMe     string
	Profil      string `orm-go:"CHECK:IN:('PUBLIC','PRIVATE')"`
	JWT         string `orm-go:"DEFAULT ''"`
}

type Groups struct {
	orm.Model
	Title       string `orm-go:"NOT NULL"`
	Description string `orm-go:"NOT NULL"`
	ImgUrl      string
	IdUser      int `orm-go:"FOREIGN_KEY:User:Id"`
}

type Post struct {
	orm.Model
	Title    string `orm-go:"NOT NULL" validate:"required"`
	ImgUrl   string
	Content  string `orm-go:"NOT NULL" validate:"required"`
	Status   string `orm-go:"CHECK:IN:('PUBLIC','PRIVATE','ALMOST','GROUP')"`
	IdUser   int    `orm-go:"FOREIGN_KEY:User:Id"`
	IdGroups int    `orm-go:"FOREIGN_KEY:Groups:Id"`
}

type Visibility struct {
	orm.Model
	IdPost   int    `orm-go:"FOREIGN_KEY:Post:Id"`
	Type     string `orm-go:"CHECK:IN:('GROUP','POST')"`
	IdObject int
}

type Comment struct {
	orm.Model
	ImgUrl  string
	Content string `orm-go:"NOT NULL" validate:"required,min=5"`
	IdUser  int    `orm-go:"FOREIGN_KEY:User:Id"`
	IdPost  int
}

type Like struct {
	orm.Model
	Type     string `orm-go:"CHECK:IN:('COMMENT','POST')"`
	Liked    int
	IdObject int
}

type Chat struct {
	orm.Model
	Type     string `orm-go:"CHECK:IN:('Groups','CHAT')" validate:"required"`
	Sender   int    `orm-go:"FOREIGN_KEY:User:Id" validate:"required"`
	Receiver int    `validate:"required"`
	Content  string `orm-go:"NOT NULL" validate:"required"`
}

type Event struct {
	orm.Model
	IdUser      int    `orm-go:"FOREIGN_KEY:User:Id"`
	Title       string `orm-go:"NOT NULL" validate:"required"`
	Description string `orm-go:"NOT NULL" validate:"required"`
	Option      string `orm-go:"CHECK:IN:('GOING','NOT_GOING')" validate:"required"`
	Date        string
	IdGroups    int `orm-go:"FOREIGN_KEY:Groups:Id"`
}

type EventOptions struct {
	orm.Model
	IdUser  int    `orm-go:"FOREIGN_KEY:User:Id"`
	Type    string `orm-go:"CHECK:IN:('GOING','NOTGOING')" validate:"required"`
	IdEvent int    `orm-go:"FOREIGN_KEY:Event:Id"`
}
type GroupMember struct {
	orm.Model
	IdUser   int `orm-go:"FOREIGN_KEY:User:Id"`
	IdGroups int `orm-go:"FOREIGN_KEY:Groups:Id"`
	IsAdmin  int `orm-go:"CHECK:IN:(0,1)"`
}

type Notifications struct {
	orm.Model
	Type        string `orm-go:"CHECK:IN:('FOLLOW','INVITE','REQUEST','EVENT','PRIVATE','Groups')"`
	Sender      int    `orm-go:"FOREIGN_KEY:User:Id"`
	Receiver    int    `orm-go:"FOREIGN_KEY:User:Id"`
	IdGroup     int    `orm-go:"FOREIGN_KEY:Groups:Id"`
	IdEvent     int
	EventOption string
}

type Followers struct {
	orm.Model
	IdFollower int    `orm-go:"FOREIGN_KEY:User:Id" validate:"required"`
	IdFollowee int    `orm-go:"FOREIGN_KEY:User:Id" validate:"required"`
	Status     string `orm-go:"CHECK:IN:('PENDING','VALIDATE')"`
}

type Invite struct {
	orm.Model
	IdUser        int    `orm-go:"FOREIGN_KEY:User:Id"`
	IdUserInvited int    `orm-go:"FOREIGN_KEY:User:Id"`
	IdGroup       int    `orm-go:"FOREIGN_KEY:Groups:Id"`
	Accepted      string `orm-go:"CHECK:IN:('PENDING','VALIDATE')"`
}
