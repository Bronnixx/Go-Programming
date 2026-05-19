package main
import(
	"fmt"
	"io"
	"log"
	"os"
)

func main(){
	message := "iambronnix....haha thought i would say i'm batman\n"
	f, err := os.Open(createFile(message))
	if err !=nil{
		log.Fatalf("Unable to read the file: %v", err)
	}
	buf := make([]byte, 1)
	for {
		n, err := f.Read(buf)
		if err == io.EOF{//breaks if the end of file is reached
			break
		}
		if err != nil {//if no EOF and no error i encountered code execution continues
			fmt.Println(err)
			continue
		}
		if n > 0{//prints one byte at a time
			fmt.Print(string(buf[:n]))
		}
	}
}
func createFile(message string)string{//it does that thing here
	err := os.WriteFile(os.Args[1],[]byte(message),0644)
	if err != nil{
		log.Panic(err)
	}
	return os.Args[1]
	
}