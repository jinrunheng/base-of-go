package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	myEval(1, 3, "+")
	fmt.Println(div(13, 2))
	// q, r := div2(13, 2)
	// fmt.Printf("q : %d,r : %d", q, r)
	q, _ := div2(13, 2)
	fmt.Println(q)

	if result, err := eval2(3, 4, "x"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result := apply(op, 12, 3)
	fmt.Println(result)
	fmt.Println(apply(func(a int, b int) int {
		return a * b
	}, 3, 4))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func myEval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator :" + op)
	}
	return result
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func eval2(a, b int, operation string) (result int, err error) {
	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return -1, fmt.Errorf("unsupported operation : %s", operation)
	}
}

func op(a, b int) int {
	return a + b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d,%d)", opName, a, b)
	fmt.Println()
	return op(a, b)
}
