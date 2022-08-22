package usertype

type Role int

const (
	Undefined Role = iota
	Admin
	ReadOnly
	Constrain
)

func (r Role) String() string {
	switch r {
	case Admin:
		return "admin"
	case ReadOnly:
		return "readonly"
	case Constrain:
		return "constrain"
	}
	return "undefined"
}
