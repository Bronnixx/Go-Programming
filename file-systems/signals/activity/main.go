package main
import(
	"fmt"
	"os"
	"os/signal"
	"syscall"
)
func main(){
	sigs := make(chan os.Signal, 1)
	done := make(chan struct{})
	signal.Notify(sigs, syscall.SIGINT)
	go func(){
		for {
			s := <- sigs
			switch s {
				case syscall.SIGINT:
				fmt.Println()
				fmt.Println("My process has been interrupted. someone might pressed CTRL+C")
				fmt.Println("Some clean up is occuring")
				done <- struct{}{}
			}
		}
	}()
	fmt.Println("Program is blocked until a signal is caught")
	done <- struct{}{}
	fmt.Println("Out of here...I'm done")
}