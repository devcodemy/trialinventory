package data

import "trial/tokentype"

type Token struct {
	Type  tokentype.Purpose
	Name  string
	Value string
}

type Usage struct {
	Type  tokentype.Device
	Name  string
	Ram   tokentype.Ram
	Specs []Spec
}

type Spec struct {
	Name  string
	Value string
}

type Gate struct {
	Name string
	IP   string
}

type Node struct {
	Name     string
	IP       string
	Settings string
}

type Environment struct {
	Nodes          []Node
	Gates          []Gate
	EnvironmentID  string
	MonitoredHosts int
}

type Cluster struct {
	Environments    []Environment
	InternalURL     string
	ManagedURL      string
	ClusterID       string
	License         string
	HostUnitsLimit  int // top bar
	HostUnitsSupply int // average
}
