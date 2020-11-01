package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	loopTest1()
	fmt.Println(convertToBin(13))
	printFile("abc.txt")
}

func loopTest1() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// 将一个十进制数字转换为二进制数字的字符串形式
// 13 -> "1101"
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func forever()  {
	// 死循环相当于 while(true)
	for {
		fmt.Println("forever")
	}
}