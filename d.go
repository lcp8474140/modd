package modd

import "fmt"
import "github.com/lcp8474140/modc"

var Version="v1.0.2"

func Print() {
	fmt.Println("modd", Version)
	modc.Print()
}
