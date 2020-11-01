## Chapter2

#### 一：变量定义

##### 使用var关键字

- 使用var定义的变量可以赋初值或者不赋值；当不赋值时，需要指明具体类型,当赋初值时，可以省略具体类型让编译器自动选择类型或者使用语法糖:`:=`的简便式的写法，但是切记 `:=` 只能在函数内使用

```go
func main() {
  // 不赋初值 则需要指明具体的类型
	var a,b,c bool

  // 赋初值 可以带着类型声明，也可以省略具体类型；语法糖写法
  var s1,s2 string = "hello","world"
  var aa,bb = 1,"bb"
  // 使用 := 定义变量
  // := 只能在函数内使用
  cc,dd := true,"def"
}
```

- 使用var关键字的变量可以放在函数内，或直接放在包内

- 可以使用var()集中定义变量

完整的示例程序如下：

```go
package main

import "fmt"

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
```

#### 二：内建变量类型

- bool,string
- (u)int,(u)int8,(u)int16,(u)int32,(u)int64,uintptr
- byte,rune(相当于char,但是它是32位 4字节的)
- float32,float64,complex64,complex128 (complex是复数：实部 + 虚部)

##### 复数回顾

- i = √-1
- 复数：3 + 4i，3对应的就是实部，4对应的就是虚部

- 欧拉公式：e^(πi)+1=0

  ```go
  func euler() {
  	fmt.Println(cmplx.Exp(1i * math.Pi) + 1)
  }
  ```

##### 强制类型转换

- 类型转换是强制的

- 示例：已知三角形的直角边为3，4 求斜边

  代码如下：

  ```go
  func triangle() {
  	var a int = 3
  	var b int = 4
  	var c int
  	c = int(math.Sqrt(float64(a*a + b*b)))
  	fmt.Println(c)
  }
  ```


#### 三：常量与枚举

##### 常量

- `const filename = "abc.txt"`
- `const ` 数值可以作为各种类型使用
- `const a,b = 3,4`
- `var c int = int(math.Sqrt(a*a + b*b))` 这里面 a 和 b 都会自动被编译器识别成 float来用

```go
func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}
```

##### 枚举

- 普通枚举类型 : 通过 const 块来定义
- 自增值枚举类型 ：通过 iota定义自增值



示例程序

两个函数打印的结果是相同的

```go
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
```

#### 四：条件语句

##### if

- go语言中条件if后面是不需要写括号的
- if条件里面可以赋值
- if的条件里赋值的变量作用域就在这个if语句里

示例程序一：

```go
func ifTest1() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
```

示例程序二：

```go
func ifTest2() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
```

##### switch

- go语言中的switch语句和其他语言不同的地方是，switch会自动break，除非使用fallthrough
- switch语句后面可以没有表达式

示例程序

```go
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
```

#### 五：循环

- for的条件里面不需要（不能有）括号
- for的条件里可以省略初始条件，结束条件，递增表达式

示例一：

```go
func loopTest1() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

示例二：

```go
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
```

示例三：

```go
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
```

示例四：

```go
func forever()  {
	// 死循环相当于 while(true)
	for {
		fmt.Println("forever")
	}
}
```



#### 六：函数

- go语言可以有多个返回值
- 一个函数也可以作为另一个函数的参数
- 函数参数可以传入可变参数列表

示例程序一：

```go
func div(a, b int) (int, int) {
	return a / b, a % b
}
```

示例程序二：

```go
func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}
```

当一个函数有多个返回值，而我们只想用某个返回值当时候：

```go
func div(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	q, _ := div(13, 2)
	fmt.Println(q)
}
```

示例程序三：

```go
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
```



示例程序四：

```go
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

func main() {
	result := apply(op, 12, 3)
	fmt.Println(result)
  
  // 匿名函数
  // go语言还暂时不支持lambda表达式
	fmt.Println(apply(func(a int, b int) int {
		return a * b
	}, 3, 4))
}
```

示例程序五：

```go
// 函数参数可以传入可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
```



#### 七：指针		

##### 变量在内存中的地址

```go
import "fmt"

func main() {
	var a int = 10
	fmt.Println("a的地址为：", &a)
}
```

程序运行的结果为：

```
a的地址为：
 0xc000016080
```

##### 什么是指针

一个指针变量指向了一个值的内存地址

类似于变量和常量，在使用指针前，你需要声明指针。指针声明的格式如下：

```go
var var_name *var_type
```

##### 指针使用的示例程序

```go
func main() {
	var a int = 10
	
	var pa *int = &a
	fmt.Println("pa指针变量存储的指针地址：", pa)
	fmt.Println("*pa变量的值", *pa)
}
```

程序运行结果如下：

```
pa指针变量存储的指针地址： 0xc000016080
*pa变量的值 10
```

##### 参数传递

值传递？引用传递？

Go语言只有值传递一种方式，也就是说传递的方式就是将值拷贝一份

如示例程序：

```go
func main() {
	var a int = 10
	f(a)
	fmt.Println(a)
}
func f(a int) {
	a++
}
```

最后程序输出的结果为：

```
10
```

因为Go语言值传递是将原本的值拷贝了一份，而原值不变

如果想要实现“引用传递”，我们可以使用指针

如示例程序：

```go
func main() {
	var a int = 10
	f_pass_by_val(a)
	fmt.Println(a)

	var b int = 10
	var pb *int = &b
	f_pass_by_ref(pb)
	fmt.Println(b)
}
func f_pass_by_val(a int) {
	a++
}
func f_pass_by_ref(pb *int) {
	*pb++
}
```

该程序输出结果为：

```
10
11
```



那么程序是否真的是引用传递呢？很显然不是的，这个程序的本质上还是值传递，不过这次拷贝的是b的地址



对于这样的参数`var cache Cache`，并且有这样的方法`func f(cache Cache)` ；通过方法 f ，我们可以改变这个Cache这个对象，这是因为，本身这个对象 cache 它内部存储的数据就是指针。



我们用一个swap程序来演示下 在Java和Go两种语言下的做法：



示例程序：Java

```java
// java
// 在Java中也只有值传递，要实现swap功能其实并不简单
public class Swap {
    public static void main(String[] args) {
        int a = 1;
        int b = 2;
        int[] result = swap(a, b);
        a = result[0];
        b = result[1];
        System.out.println("a : " + a + "  " + "b : " + b);
    }

    public static int[] swap(int a, int b) {
        int tmp = a;
        a = b;
        b = tmp;
        return new int[]{a, b};
    }
}
```

示例程序：Go

```go
import "fmt"

func main() {
	a, b := 10, 20
	fmt.Println("a,b", a, b)
	swap(&a, &b)
	fmt.Println("a,b", a, b)
}

func swap(pa, pb *int) {
	*pa, *pb = *pb, *pa
}
```

有了指针，Go的程序显然要更加清晰简便

即便不使用指针，我们也可以这样做：

```go
func main() {
	a, b := 10, 20
	fmt.Println("a,b", a, b)
	a,b = swap2(a, b)
	fmt.Println("a,b", a, b)
}
func swap2(a, b int) (int, int) {
	return b, a
}
```

这样做也是正确的，并且这种做法也是值得提倡的























