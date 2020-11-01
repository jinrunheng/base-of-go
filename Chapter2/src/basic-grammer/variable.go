package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 3
var bb bool = false
var cc = "ss"

var (
	aaa = 2
	bbb = true
	ccc = "sss"
)

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc)
	fmt.Println(aaa, bbb, ccc)

	testComplex()
	euler()

	triangle()
	// 常量
	consts()
	// 枚举
	enums1()
	enums2()
	enumTestBit()
}

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 1, 2
	var s string = "abc"

	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, d = 1, 2, true, "def"
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	a, b, c, d := 1, 2, true, "def"
	fmt.Println(a, b, c, d)
}

func testComplex() {
	a := 3 + 4i
	fmt.Println(cmplx.Abs(a))
}

// 验证欧拉公式
func euler() {
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

func triangle() {
	var a int = 3
	var b int = 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums1() {
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)
	fmt.Println(cpp, java, python, golang)
}

func enums2() {
	const (
		cpp = iota
		java
		python
		golang
	)
	fmt.Println(cpp, java, python, golang)
}

func enumTestBit() {
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
