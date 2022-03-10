package cgroups

type Subsystem interface {
	Name() string
	Set(path string, conf *ResourceConfig) error
	Apply(path string, pid int, conf *ResourceConfig) error
	Remove(path string) error
}

var Subsystems = []Subsystem{
	&cpuSubsystem{},
	&memorySubsystem{},
}
