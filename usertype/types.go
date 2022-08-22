package usertype

type Role int

const (
	Undefined Role = iota
	Admin
	ReadOnly
	Constrain
)

func (r Role) String() string {
	return [...]string{"UNDEFINED", "ADMIN", "READONLY", "CONSTRAIN"}[r]
}
