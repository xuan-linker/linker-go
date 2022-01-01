package main

// #include<stdlib.h>
import "C"

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.unit(i))
}
/* default key
	break default func interface select
	case defer go map struck
	chan else goto package switch
	const fallthrouge if range type
	continue for import return var
*/
func main() {
	//var a int = 1
	//defer C.free(unsafe.Pointer())


}
