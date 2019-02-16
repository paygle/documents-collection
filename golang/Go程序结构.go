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
	1. 变量声明格式： var <变量名称> <变量类型>
	2. 变量赋值格式： <变量名称> = <表达式>
	3. 声明同时赋值： var <变量名称> <变量类型> = <表达式>
	4. 声明赋值简写： <变量名称> := <表达式> ， 冒号代替var关键字
	5. 多变量声明： var a, b, c = 1, 2, 3
	
	6. GO 中 不存在隐匿类型转换，所有转换必须显式声明，转换只能发生在两种相互兼容的类型之间
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
	1. 数组定义格式： var <varName> [n] <type>，其中 n >= 0；
	2. 数组长度也是类型的一部分，因此具有不同长度的数组为不同类型，数组为值类型；
	3. 注意区分指向数组的指针和指针数组；
	4. 数组之间可以使用 == 或 != 进行比较，但不可以使用 > 或 <，可以使用 new 来创建数组，此方法返回一个指向数组的指针； 
*/

var aa [100]int  // 声明一个数组
/*
 1. 数组定义各种形式， 使用符号 “:=” 定义，只能在函数内使用；

 ar := [2]int{1, 1}
 ar := [...]int{1, 2, 3, 4, 5, ...}
 ar := [...]int{0: 1, 1: 2, 2: 3, 3: 4}
 ar := [...]int{19: 1}
 pp := [...]*int{&name1, &name2}  // 指针数组, 保存指针
 pn := new([10]int) // 产生指向数组的指针

 2. 多维数组
 pm := [2][3]int {
	 {1, 1, 1},
	 {2, 2, 3}}

 3. 获取数组长度
 length := len(ar)

*/
var pa *[100]int = &aa // 数组的指针

/*
	【切片 Slice】
	1. 其本身并不是数组，它指向底层的数组；
	2. 作为变长数组的替代方案，可以关联底层数组的局部或全部，为引用类型，可以直接创建或从底层数组获取生成；
	3. 使用 len() 获取元素个数， cap() 预分配单位容量，一般用 make() 创建；
	4. 如果多个slice指向相同底层数组，其中一个值改变会影响全部；

		 make([]T, len, cap)
		 
	创建指向指针的指针，其中 cap 可省略，则和 len 相同

	ss := ar[2:5]    // 非正式获取 slice, 不包含终值 5

	sm := make([]int, 10, 10)  // slice 正式声明

	sm = append(sm, 1, 2, 3)  //  在 slice 中追加元素，扩充容量

	copy(targetArray, sourceArray)  // 以目标函数长度为准，源覆盖目标数组，多余丢弃

	copy(target[2:4], source[1:3])  // 也可以使用slice复制部分元素
*/

var slc []int  // 声明一个空的 slice 类型

/*
	【Map 数据类型】
	1. 其Key必须是支持 == 或 != 比较运算的类型，不可以是函数、map 或 slice；
	2. Map 查找比线性搜索快很多，但比使用索引访问数据的类型慢100倍；
	3. Map 使用 make() 创建， 支持 “ := ” 简写方式

	4. make([KeyType]valueType, cap), cap 表示容量可省略，超出容量会自动扩容，但尽量提供一个合理的初始值，使用 len() 获取元素个数；

	5. 键值对不存在时自动添加，使用 delete() 删除某键值对；
	6. 使用 for range 对 map 和 slice 进行迭代操作；

	mm := make(map[int]string)              // 简写
	mm := map[int]string{1: "a", 2: "b"}
	mm[1] = "OK"                            // 使用 map

	7. 多层 map 声明与使用

	var ml map[int]map[int]string
	ml = make(map[int]map[int]string)

	am, isOk := ml[1][1]  // isOk 为返回元素状态
	if !isOk { // 元素状态为 false 则初始化
		ml[1] = make(map[int]string)
	}

	ml[1][1] = "OK"
	am := ml[1][1]
	std.Println(am, isOk)

	8. Map 迭代操作, range 的赋值只是复制，不会影响原数据

	for key, _ := range ml {
		ml[k] = make(map[int]string)
		ml[k][1] = "OK"
		std.Println(ml[k])
	}

	9. 数组迭代操作

	for index, value := range arr {

	}
*/

var mm map[int]string = make(map[int]string) // 正式声明初始化一个map对象


// 一般类型声明, 严格讲只是声明类型的别名（自定义类型）
type newType int

// 组声明
type (
	type1 int
	type2 float32
	type3 byte
)

/*
	【结构 struct】
	1. Go 中的 struct 与 C 中的 struct 非常相似，并且Go 没有 class；
	2. 使用 type <Name> struct{} 定义结构，名称遵循可见性规则；
	3. 支持指向自身的指针类型成员；
	4. 支持匿名结构，可用作成员或定义成员变量；
	5. 匿名结构也可以用于Map的值；
	6. 可以使用字面值对结构进行初始化；
	7. 允许直接通过指针来读写结构成员；
	8. 相同类型的成员可进行直接拷贝赋值；
	9. 支持 == 与 != 比较运算符，但不支持 > 或 <；
	10. 支持匿名字段，本质上是定义了以某个类型名为名称的字段；
	11. 嵌入结构作为匿名字段看起来像继承，但不是继承；
	12. 可以使用匿名字段指针；

*/

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

/*
	【函数 Function】
	1. Go 函数 “不支持” 嵌套、重载和默认参数；可以作为一种类型作用
	2. Go 函数无需声明原型、不定长度变参、多返回值、命名返回值参数、匿名函数、闭包；
*/

func examFunc1(param1 int, param2 string) int {
	return 1
}

func examFunc2(param1, param2 string) (int, string) {
	return 2, "a"
}

// 给定返回类型，renturn 可以省略返回表达式
func examFunc3() (b, c int) {
	b, c = 1, 2
	return
}

// 不定长变参
func examFunc4(param ...int){
	av := 8
	funtype := examfunType  // 函数类型
	nonameFun := func() {   // 匿名函数
		std.Println("nonameFun")
	}

	funtype(&av)
	nonameFun()
	std.Println(param)
}

// 参数指针
func examfunType(p *int){
	*p = 5             // 改变原始值
	std.Println(*p)
}

// 闭包函数
func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

/*
	【defer 类似析构函数】
	1. 在函数体执行结束后按照调用顺序的“相反顺序”逐个执行；
	2. 即使函数发生严重错误也会执行；
	3. 支持匿名函数的调用
	4. 常用于资源清理、文件关闭、解锁以及记录时间等操作；
	5. 通过与匿名函数配合可在 return 之后修改函数计算结果；
	6. 如果函数体内某个变量作为 defer 的匿名函数的参数，则在定义 defer 时即已经获得了拷贝，否则则是引用某个变量的地址；

	7. Go 没有异常机制，但有 panic/recover 模式来处理错误；
	8. Painc 可以在任何地方引发，但 recover 只有在 defer 调用的函数中有效；
*/

func deferFunExam() {

	for i := 0; i < 3; i++ {
		defer func() {
			std.Println(i)      // 打印的结果全是 3
		}()
	}

	std.Println("A")
	defer std.Println("B")    // 第2个调用
	defer std.Println("C")    // 第1个调用

	// panic/recover 模式来处理错误
	AAA()
	BBB()
	CCC()
}

func AAA() {
	std.Println("Func AAA")
}

func BBB() {
	// panic/recover 模式来处理错误
	defer func() {  // defer 要在 panic 之前执行
		if err := recover(); err != nil {
			std.Println("Recover in BBB")
		}
	}()
	panic("Panic in BBB")
}

func CCC() {
	std.Println("Func AAA")
}


// 由 main 函数作为程序的入口
func main() {

	/*
		1. Go 虽然保留了指针，但与其它语言不同的是，直接采用 "." 选择符来操作指针目标对象的成员；
		2. 操作符“ & ” 取变量地址， 使用 “ * ” 通过指针间接访问目标对象；
		3. 默认值为 nil 而非 NULL；
	*/

	a := 1

	a++ // ++， -- 操作符只能单独作为一条语句使用，不能放在 = 号右边
	a--

	var p *int = &a
	f := closure(10)  // 闭包函数

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
	std.Println(f(1))
	std.Println("Hello world! 你好！")
	deferFunExam()
}
