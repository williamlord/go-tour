package main

import "fmt"

func main() {
	i := 42
	f := 3.142
	g := 0.867 + 0.5i

	fmt.Printf("i is of type %T\n", i)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("g is of type %T\n", g)
}
