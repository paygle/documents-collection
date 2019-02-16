// 当前程序的包名，代码顺序比较重要
package main

// 导入其他的包
import std "fmt"
/*
import (
	"math"
	"net"
)
*/

/*
	iota 特殊常量，一个可以被编译器修改的常量
*/
// 常量的定义
const PI = 3.1415926
// 组声明
const (
	const1 = 1
	const2 = 2
	const3 = 3
)

/*
	变量声明格式： var <变量名称> <变量类型>
	变量赋值格式： <变量名称> = <表达式>
	声明同时赋值： var <变量名称> <变量类型> = <表达式>
	声明赋值简写： <变量名称> := <表达式> ， 冒号代替var关键字
	多变量声明： var a, b, c = 1, 2, 3
	
	GO 中 不存在隐匿类型转换，所有转换必须显式声明，转换只能发生在两种相互兼容的类型之间
	转换格式： <ValueA> [:]= <TypeOfValueA>(<ValueB>)
	var a float32 = 1.1
	b := int(a)
*/

// 全局变量的声明与赋值
var name = "gopher"
// 组声明, 由系统推断变量类型
var (
	name1 = 1
	name2 = 2
	name3 = 3
)

/*
	数组定义格式： var <varName> [n] <type>，其中 n >= 0
	数组长度也是类型的一部分，因此具有不同长度的数组为不同类型，数组为值类型
	注意区分指向数组的指针和指针数组
	数组之间可以使用 == 或 != 进行比较，但不可以使用 > 或 <，可以使用 new 来创建数组，此方法返回一个指向数组的指针
*/
var aa [100]int
/*
 ar := [2]int{1, 1}
 ar := [...]int{1, 2, 3, 4, 5, ...}
 ar := [...]int{0: 1, 1: 2, 2: 3, 3: 4}
 ar := [...]int{19: 1}
 pp := [...]*int{&name1, &name2}  // 指针数组, 保存指针
 pn := new([10]int) // 产生指向数组的指针

 // 多维数组
 pm := [2][3]int {
	 {1, 1, 1},
	 {2, 2, 3}}

 // 获取数组长度
 length := len(ar)
*/
var pa *[100]int = &aa // 数组的指针


// 一般类型声明, 严格讲只是声明类型的别名（自定义类型）
type newType int

// 组声明
type (
	type1 int
	type2 float32
	type3 byte
)

// 结构的声明
type gopher struct {}

// 接口的声明
type golang interface {}

// switch 语句例子
func switchGo() {
	/*
		选择语句 switch
		1. 可以使用任何类型或表达式作为条件语句
		2. 不需要写 break, 一旦条件符合自动终止
		3. 如希望继续执行下一个case，需要使用 fallthrough 语句
		4. 支持一个初始化表达式（可以是并行方式），右侧需跟分号
		5. 左大括号必须和条件语句在同一行
	*/
	a := 1
	switch {
		case a >= 0:
			std.Println("a=0")
			fallthrough
		case a >= 1:
			std.Println("a=1")
		default:
			std.Println("None")
	}
}

// 由 main 函数作为程序的入口
func main() {

	/*
		Go 虽然保留了指针，但与其它语言不同的是，直接采用 "." 选择符来操作指针目标对象的成员

		操作符“ & ” 取变量地址， 使用 “ * ” 通过指针间接访问目标对象
		默认值为 nil 而非 NULL
	*/

	a := 1

	a++ // ++， -- 操作符只能单独作为一条语句使用，不能放在 = 号右边
	a--

	var p *int = &a

	// 定义一个标记
	LABLE1:

	for i := 0; i < 3; i++ {

		// 条件不需要 () 号
		if b := 1; b > 1 {
			std.Println(b)
			break LABLE1  // 退出到标记， 或 continue LABLE1
			// goto LABLE1   // 调整程序执行位置到标记
		}

	}

	switchGo()
	std.Println(*p)
	std.Println("Hello world! 你好！")
}
