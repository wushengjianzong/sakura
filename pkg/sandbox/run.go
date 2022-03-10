package sandbox

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

const (
	EngineRoot = "/var/lib/sakura"
	SelfLink   = "/proc/self/exe"
	LogFile    = "container.log"
)

var (
	ContainersRoot = filepath.Join(EngineRoot, "containers")
)

func NewParentProcess(tty bool, name string, volume string, image string, envs []string) (*exec.Cmd, *os.File) {
	readPipe, writePipe, err := os.Pipe()
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	initLink, err := os.Readlink(SelfLink)
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	initCmd := exec.Command(initLink, "init")
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	initCmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWNET,
	}

	containerRoot := filepath.Join(ContainersRoot, name)
	if tty {
		initCmd.Stdin = os.Stdin
		initCmd.Stdout = os.Stdout
		initCmd.Stderr = os.Stderr
	} else {
		if err := os.MkdirAll(containerRoot, 0622); err != nil {
			log.Println(err)
			return nil, nil
		}
		containerLogFilePath := filepath.Join(containerRoot, LogFile)
		containerLogFile, err := os.Create(containerLogFilePath)
		if err != nil {
			log.Println(err)
			return nil, nil
		}
		initCmd.Stdout = containerLogFile
	}
	initCmd.ExtraFiles = []*os.File{readPipe}
	initCmd.Env = append(os.Environ(), envs...)
	initCmd.Dir = containerRoot
	return initCmd, writePipe
}
