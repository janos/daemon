package daemon

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

type Daemon struct {
	PidFileName string
	PidFileMode os.FileMode
}

func (d *Daemon) Daemonize(workDir string, inFile *os.File, outFile *os.File, errFile *os.File) error {
	if syscall.Getppid() != 1 {
		path, err := filepath.Abs(os.Args[0])
		if err != nil {
			return err
		}
		cmd := exec.Command(path, os.Args[1:]...)
		cmd.Stdin = inFile
		cmd.Stdout = outFile
		cmd.Stderr = errFile
		if err := cmd.Start(); err != nil {
			return err
		}
		os.Exit(0)
	}
	if workDir != "" {
		if err := os.Chdir(workDir); err != nil {
			return err
		}
		os.Chdir(workDir)
	}

	syscall.Umask(0)

	s_ret, s_err := syscall.Setsid()
	if s_err != nil {
		return s_err
	}

	if err := ioutil.WriteFile(d.PidFileName, []byte(strconv.Itoa(s_ret)), d.PidFileMode); err != nil {
		return err
	}

	return nil
}

func (d *Daemon) Cleanup() error {
	if d.PidFileName == "" {
		return nil
	}
	return os.Remove(d.PidFileName)
}

func (d *Daemon) Pid() int {
	pid, _ := ioutil.ReadFile(d.PidFileName)
	p, _ := strconv.Atoi(string(pid))
	return p
}

func (d *Daemon) Process() (*os.Process, error) {
	return os.FindProcess(d.Pid())
}

func (d *Daemon) Signal(sig os.Signal) error {
	process, err := d.Process()
	if err != nil {
		return err
	}
	return process.Signal(sig)
}

func (d *Daemon) Status() (pid int, err error) {
	p, err := d.Process()
	if err != nil {
		return 0, err
	}
	return p.Pid, p.Signal(syscall.Signal(0x0))
}

func (d *Daemon) Stop() error {
	process, err := d.Process()
	if err != nil {
		return err
	}
	if err := process.Signal(syscall.SIGTERM); err != nil {
		return process.Kill()
	}
	return nil
}
