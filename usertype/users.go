package usertype

type Group struct {
	GroupId  string
	UserList []User
}

type User struct {
	UserId       int
	Name         string
	Email        string
	Login        string
	PasswordHash string
	Role         Role
}
