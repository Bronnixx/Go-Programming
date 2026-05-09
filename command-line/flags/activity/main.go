package main 
import(
	"fmt"
	"flag"
)
var(
	nameFlag = flag.String("name", "erick", "Name of the person to say hello")
	quietFlag = flag.Bool("quiet", false ,"Toggle to be quiet when saying hello")
)
func main(){
	flag.Parse()
	if !*quietFlag{
		greeting := fmt.Sprintf("Hello, %s! welcome to the command line.", *nameFlag)
		fmt.Println(greeting)
	}
}
//to see the available help message for the preceding code,you can run go run main.go --help