package database

import (
	orm "backend/lib/orm/ORM"
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
	ImgUrl      string
	Title       string `orm-go:"NOT NULL"`
	Description string `orm-go:"NOT NULL"`
}

type Post struct {
	orm.Model
	Title    string `orm-go:"NOT NULL" validate:"required,min=4"`
	ImgUrl   string
	Content  string `orm-go:"NOT NULL" validate:"required,min=2"`
	IdUser   int    `orm-go:"FOREIGN_KEY:User:Id"`
	Status   string `orm-go:"CHECK:IN:('PUBLIC','PRIVATE','ALMOST')"`
	IdGroups int    `orm-go:"FOREIGN_KEY:Groups:Id"`
}

type Comment struct {
	orm.Model
	Type     string `orm-go:"CHECK:IN:('COMMENT','POST')"`
	ImgUrl   string
	Content  string `orm-go:"NOT NULL" validate:"required,min=5"`
	IdUser   int    `orm-go:"FOREIGN_KEY:User:Id"`
	IdObject int
}

type Like struct {
	orm.Model
	Type     string `orm-go:"CHECK:IN:('COMMENT','POST')"`
	Liked    int    `orm-go:"CHECK:IN:(0,1)"`
	IdObject int
}

type Category struct {
	orm.Model
	Name string `orm-go:"NOT NULL"`
}

type PostCategory struct {
	orm.Model
	IdPost     int `orm-go:"FOREIGN_KEY:Post:Id" validate:"required"`
	IdCategory int `orm-go:"FOREIGN_KEY:Category:Id" validate:"required"`
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
	Title       string `orm-go:"NOT NULL" validate:"required"`
	Description string `orm-go:"NOT NULL" validate:"required"`
	Option      string `orm-go:"CHECK:IN:('GOING','NOT_GOING')" validate:"required"`
}

type GroupMember struct {
	orm.Model
	IdUser   int `orm-go:"FOREIGN_KEY:User:Id"`
	IdGroups int `orm-go:"FOREIGN_KEY:Groups:Id"`
	IsAdmin  int `orm-go:"CHECK:IN:(0,1)"`
}

type Notifications struct {
	orm.Model
	Type     string `orm-go:"CHECK:IN:('FOLLOW','INVITE','REQUEST','EVENT','PRIVATE','Groups')"`
	Sender   int    `orm-go:"FOREIGN_KEY:User:Id"`
	Receiver int
}

type Followers struct {
	orm.Model
	IdFollower int    `orm-go:"FOREIGN_KEY:User:Id"`
	IdFollowee int    `orm-go:"FOREIGN_KEY:User:Id"`
	Status     string `orm-go:"CHECK:IN:('PENDING','VALIDATE')"`
}
