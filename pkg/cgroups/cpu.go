package cgroups

type cpuSubsystem struct{}

func (s *cpuSubsystem) Name() string {
	return "cpu"
}

func (s *cpuSubsystem) Set(path string, conf *ResourceConfig) error {
	return nil
}

func (s *cpuSubsystem) Apply(path string, pid int, conf *ResourceConfig) error {
	return nil
}

func (s *cpuSubsystem) Remove(path string) error {
	return nil
}
