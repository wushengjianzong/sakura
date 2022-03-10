package cgroups

type memorySubsystem struct{}

func (s *memorySubsystem) Name() string {
	return "cpu"
}

func (s *memorySubsystem) Set(path string, conf *ResourceConfig) error {
	return nil
}

func (s *memorySubsystem) Apply(path string, pid int, conf *ResourceConfig) error {
	return nil
}

func (s *memorySubsystem) Remove(path string) error {
	return nil
}
