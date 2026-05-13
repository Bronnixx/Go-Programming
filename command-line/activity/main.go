package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"

	u "golang.org/x/sys/unix"
)

func main() {
	startTime := time.Now() 
	killSignal, pid := psKiller()
	taskTimer(startTime, killSignal,pid)
	
	
}
func psKiller() (syscall.Signal ,int){// returns a termination signal and pid
	fmt.Println("Initiating processKiller() on process: \t", u.Getpid())
	timeLimit := 5 *time.Second
	time.Sleep(timeLimit)
    signal := syscall.SIGTERM
    return  signal, u.Getpid()
	
}
func taskTimer(t time.Time, s syscall.Signal, pid int) {
	fmt.Println("\nStarting process",pid,"at Start Time", t)
	f := func(){
		fmt.Println("executing a remote command")
	cmd := exec.Command("watch","-n", "1", "lsof")//will never show execution on the terminal...there's termination signal
    timelimit := 5* time.Second
   time.Sleep(timelimit)	
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	}
	f()
	
}