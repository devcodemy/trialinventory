package tokentype

type Purpose int

const (
	Undefined Purpose = iota
	PaaS
	Cluster
	Environment
)

func (p Purpose) String() string {
	switch p {
	case PaaS:
		return "paas_demo"
	case Cluster:
		return "cluster"
	case Environment:
		return "env"
	}
	return "undefined"
}
