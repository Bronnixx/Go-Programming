package main
import(
	"os"
	"log"
	"time"
)
func main(){
	message := []byte("hey yooh come look at this")
	err := os.WriteFile(cmd(), message,0644)
	if err != nil{
		log.Println(err)
	}
}
func cmd()string{
	time.Sleep(1*time.Second)
	filename := os.Args[1]
	log.Printf("file %s created\n", os.Args[1])
	return filename
}