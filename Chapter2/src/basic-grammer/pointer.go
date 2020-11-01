package main

import "fmt"

func main() {
	a, b := 10, 20
	fmt.Println("a,b", a, b)
	a, b = swap2(a, b)
	fmt.Println("a,b", a, b)
}

func swap(pa, pb *int) {
	*pa, *pb = *pb, *pa
}

func swap2(a, b int) (int, int) {
	return b, a
}
