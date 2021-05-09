package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// `go run main.go` run echo "hello from container"
//`go run main.go` run /bin/bash
func main() {
	switch os.Args[1] {
	case "run":
		run()

	case "kid": // naming as kid to show that /proc/self/exe just needs an arg to fork a new proc
		kid()

	default:
		panic("run args not provided")

	}
}

func run() {
	fmt.Printf("Running %v as PID %d in the host\n", os.Args[2:], os.Getpid())

	// cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd := exec.Command("/proc/self/exe", append([]string{"kid"}, os.Args[2:]...)...) // /proc/self/exe is a fork exec
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	//isolation starts  run as root  GOOS=Linux UTS=UnixTimeSharingSystem(hostname)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID}

	must(cmd.Run())
}

func kid() {
	fmt.Printf("Running %v as PID %d in the Continer\n", os.Args[2:], os.Getpid())

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	//No need to create ns since it's done already

	//chroot & mount proc
	must(syscall.Chroot("/root/ubuntu"))
	must(syscall.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
