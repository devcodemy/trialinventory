package tokentype

type (
	Purpose   int
	Device    int
	Ram       int
	Hdd       int
	Arch      int
	CpuVendor int
)

const (
	Undefined Purpose = iota
	PaaS
	Cluster
	Environment
)

const (
	_ Device = iota + 100
	Phone
	Tablet
	Laptop
)

const (
	_ Ram = iota
	GB_8
	GB_16
	GB_32
)

const (
	_ CpuVendor = iota
	Intel_x86
	Apple_M1
	Apple_M2
)

const (
	_ Arch = iota
	x86_64
	aarch64
)

func (p Purpose) String() string {
	return [...]string{"UNDEFINED", "PAAS", "CLUSTER", "ENVIRONMENT"}[p]
}

func (d Device) String() string {
	return [...]string{"-", "PHONE", "TABLET", "LAPTOP"}[d]
}

func (r Ram) String() string {
	return [...]string{"-", "8 Gb", "16 Gb", "32 Gb"}[r]
}

func (cv CpuVendor) String() string {
	return [...]string{"-", "INTEL_X86", "APPLE_M1", "APPLE_M2"}[cv]
}

func (a Arch) String() string {
	return [...]string{"-", "x86_64", "aarch64"}[a]
}

/*

const (
	Undefined Purpose = iota
	PaaS
	Cluster
	Environment
)

*/
