package main
import(
	"flag"
	"fmt"
)
func main(){
	v := flag.Int("value", -1, "Needs a value for the flag.")
	flag.Parse()
	fmt.Println(*v)
}