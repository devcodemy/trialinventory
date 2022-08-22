package tmp

type Check int

const (
	Undefined Check = iota
	SomeHow
	SomeOne
	SomeWhere
	SomeBody
)

type Behavior struct {
	Name   string
	Action Check
}

func (c Check) String() string {
	return [...]string{"UNDEFINED", "SOMEHOW", "SOME1", "SOMEWHERE", "SOMEBODY"}[c]
}

func (c Check) EnumIdx() int {
	return int(c)
}
