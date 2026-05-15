package main
import(
	"flag"
	"fmt"
)
func main(){
	var v int
	flag.IntVar(&v, "value", -1, "Needs a value for the flag.")
	flag.Parse()
	fmt.Println(v)
}