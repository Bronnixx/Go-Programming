package main
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)
func main(){
	sigs := make(chan os.Signal, 1)
	done := make(chan struct{})
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTSTP)
	go func(){
		for {
			s := <-sigs
			switch s {
				case syscall.SIGINT:
				fmt.Println()
				fmt.Println("My process has been interrupted.CTRL+C")
				fmt.Println("Some clean up is occuring")
				cleanUp()
				done <- struct{}{}
				return
				case syscall.SIGTSTP:
				fmt.Println()
				fmt.Println("someone pressed CTRL+Z")
				fmt.Println("Some clean up is occuring")
				cleanUp()
				done <- struct{}{}
				return
			}
		}
	}()
	fmt.Println("program is blocked until a signal is caught\nCTRL+Z,CTRL+C")
	<-done
	fmt.Println("out of here")
}
func cleanUp(){
	fmt.Println("simulating clean up")
	for i := 0; i <= 10; i++{
		fmt.Println("Deleting Files.. Not really.", i)
		time.Sleep(1 * time.Second)
	}
}