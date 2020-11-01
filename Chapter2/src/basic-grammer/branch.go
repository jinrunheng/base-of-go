package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	ifTest1()
	ifTest2()
	fmt.Println(eval(4, 5, "*"))
}

func ifTest1() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func ifTest2() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func eval(a, b int, op string) int {
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

func grade(score int) string {
	grade := ""
	switch {
	case score < 0 || score > 100 :
		panic(fmt.Sprintf("Wrong score : %d",score))
	case score < 60 :
		grade = "F"
	case score < 80 :
		grade = "C"
	case score < 90 :
		grade = "A"
	}
	return grade
}