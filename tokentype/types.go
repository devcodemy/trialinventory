package tokentype

type Purpose int

const (
	Undefined Purpose = iota
	PaaS
	Cluster
	Environment
)

func (p Purpose) String() string {
	return [...]string{"UNDEFINED", "PAAS", "CLUSTER", "ENVIRONMENT"}[p]
}
