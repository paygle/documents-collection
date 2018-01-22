# Go 语言基础

## Go 语言结构

```go
package main

import "fmt"

func main() {
   /* 这是我的第一个简单的程序 */
   fmt.Println("Hello, World!")
}
```

## Go 语言变量

* 声明变量的一般形式是使用 var 关键字

```go
var identifier type
```

* 第一种声明，指定变量类型，声明后若不赋值，使用默认值。

```go
var v_name v_type
v_name = value
```

* 第二种声明，根据值自行判定变量类型。

```go
var v_name = value
```

* 第三种声明，省略var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误。

```go
v_name := value
```

* 多变量声明

```go
//类型相同多个变量, 非全局变量
var vname1, vname2, vname3 type
vname1, vname2, vname3 = v1, v2, v3

var vname1, vname2, vname3 = v1, v2, v3 //和python很像,不需要显示声明类型，自动推断

vname1, vname2, vname3 := v1, v2, v3 //出现在:=左侧的变量不应该是已经被声明过的，否则会导致编译错误


// 这种因式分解关键字的写法一般用于声明全局变量
var (
    vname1 v_type1
    vname2 v_type2
)
```

## Go 语言常量

* 常量的定义格式：

```go
const identifier [type] = value

const b string = "abc"  // 显式类型定义
const b = "abc"    // 隐式类型定义

// 多个相同类型的声明可以简写为：
const c_name1, c_name2 = value1, value2
```

```go
// 声明示例
package main

import "fmt"

func main() {
   const LENGTH int = 10
   const WIDTH int = 5
   var area int
   const a, b, c = 1, false, "str" //多重赋值

   area = LENGTH * WIDTH
   fmt.Printf("面积为 : %d", area)
   println()
   println(a, b, c)   
}
```

### iota 特殊常量

> 可以认为是一个可以被编译器修改的常量。

> 在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1。

```go
// iota 可以被用作枚举值
const (
    a = iota
    b = iota
    c = iota
)

// 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；
// 所以 a=0, b=1, c=2 可以简写为如下形式：

const (
    a = iota
    b
    c
)
```

## Go 语言运算符

### 算术运算符

|运算符|	描述|	实例(假定 A 值为 10，B 值为 20。)|
|-----|------|----------------------------|
|+	|相加|	A + B 输出结果 30|
|-	|相减|	A - B 输出结果 -10|
|*	|相乘|	A * B 输出结果 200|
|/	|相除|	B / A 输出结果 2|
|%	|求余|	B % A 输出结果 0|
|++	|自增|	A++ 输出结果 11|
|--	|自减|	A-- 输出结果 9|

### 关系运算符

|运算符	|描述(假定 A 值为 10，B 值为 20。)	|实例|
|-----|------|----------------------------|
|==	|检查两个值是否相等，如果相等返回 True 否则返回 False。	|(A == B) 为 False|
|!=	|检查两个值是否不相等，如果不相等返回 True 否则返回 False。	|(A != B) 为 True|
|>	|检查左边值是否大于右边值，如果是返回 True 否则返回 False。	|(A > B) 为 False|
|<	|检查左边值是否小于右边值，如果是返回 True 否则返回 False。	|(A < B) 为 True|
|>=	|检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。	|(A >= B) 为 False|
|<=	|检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。	|(A <= B) 为 True|

### 逻辑运算符

|运算符|	描述|	实例|
|-----|------|----------------------------|
|&&	|逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。	|(A && B) 为 False|
|&brvbar;&brvbar;|	逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。	|(A &brvbar;&brvbar; B) 为 True|
|!	|逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。	|!(A && B) 为 True|

### 位运算符

对整数在内存中的二进制位进行操作。

|运算符	|描述(假定 A 为60，B 为13)	|实例|
|-----|--------------|--------------|
|&	|按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。|	(A & B) 结果为 12, 二进制为 0000 1100|
|	&brvbar;|按位或运算符"&brvbar;"是双目运算符。 其功能是参与运算的两数各对应的二进位相或|	(A &brvbar; B) 结果为 61, 二进制为 0011 1101|
|^	|按位异或运算符"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。|	(A ^ B) 结果为 49, 二进制为 0011 0001|
|<<	|左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。|	A << 2 结果为 240 ，二进制为 1111 0000|
|>>|	右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。|	A >> 2 结果为 15 ，二进制为 0000 1111|

|p|q|p & q|	p &brvbar; q	|p ^ q|
|---|---|----|----|----|
|0	|0	|0	|0	|0|
|0	|1	|0	|1	|1|
|1	|1	|1	|1	|0|
|1	|0	|0	|1	|1|

### 赋值运算符

|运算符	|描述	|实例|
|-----|--------------|--------------|
|=	|简单的赋值运算符，将一个表达式的值赋给一个左值	|C = A + B 将 A + B 表达式结果赋值给 C|
|+=	|相加后再赋值	|C += A 等于 C = C + A|
|-=	|相减后再赋值	|C -= A 等于 C = C - A|
|*=	|相乘后再赋值	|C *= A 等于 C = C * A|
|/=	|相除后再赋值	|C /= A 等于 C = C / A|
|%=	|求余后再赋值	|C %= A 等于 C = C % A|
|<<=|	左移后赋值	|C <<= 2 等于 C = C << 2|
|>>=|	右移后赋值	|C >>= 2 等于 C = C >> 2|
|&=	|按位与后赋值	|C &= 2 等于 C = C & 2|
|^=	|按位异或后赋值	|C ^= 2 等于 C = C ^ 2|
|&brvbar;=	|按位或后赋值	|C &brvbar;= 2 等于 C = C &brvbar; 2|


### 其他运算符

|运算符	|描述	|实例|
|-----|--------------|--------------|
|&	|返回变量存储地址	|&a， 将给出变量的实际地址。|
|*	|指针变量。	|*a， 是一个指针变量|

## Go 语言条件语句

|语句	|描述|
|-----|--------------|
|if 语句|	if 语句 由一个布尔表达式后紧跟一个或多个语句组成。|
|if...else| 语句	if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。|
|if 嵌套语句|	你可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。|
|switch 语句|	switch 语句用于基于不同条件执行不同动作。|
|select 语句|	select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。|

## Go 语言循环语句

|循环类型	|描述|
|-----|--------------|
|for 循环	|重复执行语句块|
|循环嵌套	|在 for 循环中嵌套一个或多个 for 循环|

* 循环控制语句

|控制语句	|描述|
|-----|--------------|
|break 语句|	经常用于中断当前 for 循环或跳出 switch 语句|
|continue 语句|	跳过当前循环的剩余语句，然后继续进行下一轮循环。|
|goto 语句|	将控制转移到被标记的语句。|

```go
// 无限循环
for true  {
    fmt.Printf("这是无限循环。\n");
}

// 输出 1-100 素数
package main
import "fmt"
func main() {
    var C, c int//声明变量
    C=1 /*这里不写入FOR循环是因为For语句执行之初会将C的值变为1，当我们goto A时for语句会重新执行（不是重新一轮循环）*/
    A: for C < 100 {
           C++ //C=1不能写入for这里就不能写入
           for c=2; c < C ; c++ {
               if C%c==0 {
                   goto A //若发现因子则不是素数
               }
           }
           fmt.Println(C,"是素数")
    }
}
```

## Go 语言函数

Go 语言最少有个 main() 函数。

```go
// 函数定义格式
func function_name( [parameter list] ) [return_types] {
   函数体
}
```
* func：函数由 func 开始声明
* function_name：函数名称，函数名和参数列表一起构成了函数签名。
* parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
* return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
* 函数体：函数定义的代码集合。

```go
// 函数返回多个值
package main

import "fmt"

func swap(x, y string) (string, string) {
   return y, x
}

func main() {
   a, b := swap("Mahesh", "Kumar")
   fmt.Println(a, b)
}
```
* 函数参数

|传递类型	|描述|
|-------|--------------------------------|
|值传递	|值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。|
|引用传递	|引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。|

默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。

* 函数用法


|函数用法	|描述|
|-------|--------------------------------|
|函数作为值	|函数定义后可作为值来使用|
|闭包	|闭包是匿名函数，可在动态编程中使用|
|方法	|方法就是一个包含了接受者的函数|

## Go 语言变量作用域

* 函数内定义的变量称为局部变量
* 函数外定义的变量称为全局变量
* 函数定义中的变量称为形式参数

|数据类型	|初始化默认值|
|-------|-----------|
|int	|0|
|float32|	0|
|pointer|	nil|

```go
package main

import "fmt"

/* 声明全局变量 */
var a int = 20;

func main() {
   /* main 函数中声明局部变量 */
   var a int = 10
   var b int = 20
   var c int = 0

   fmt.Printf("main()函数中 a = %d\n",  a);
   c = sum( a, b);
   fmt.Printf("main()函数中 c = %d\n",  c);
}

/* 函数定义-两数相加 */
func sum(a, b int) int {
   fmt.Printf("sum() 函数中 a = %d\n",  a);
   fmt.Printf("sum() 函数中 b = %d\n",  b);

   return a + b;
}
```

## Go 语言数组

* 声明数组

```go
var variable_name [SIZE] variable_type

// 例如
var balance [10] float32
```

* 初始化数组

```go
// 初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

// 如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：
var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}


```

* Go 语言多维数组

```go
var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type

// 例如
var threedim [5][10][4]int

```

* Go 语言向函数传递数组

```go
// 方式一, 形参设定数组大小
void myFunction(param [10]int){}

// 方式二, 形参未设定数组大小
void myFunction(param []int){}
```

## Go 语言指针

一个指针变量可以指向任何一个值的内存地址它指向那个值的内存地址。类似于变量和常量，在使用指针前你需要声明指针。

如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。

Go 语言允许向函数传递指针，只需要在函数定义的参数上设置为指针类型即可。

```go
var var_name *var-type  // 声明形式

var ip *int        /* 指向整型*/
var fp *float32    /* 指向浮点型 */
```

```go
package main

import "fmt"

func main() {
   var a int= 20   /* 声明实际变量 */
   var ip *int        /* 声明指针变量 */

   ip = &a  /* 指针变量的存储地址 */

   fmt.Printf("a 变量的地址是: %x\n", &a  )

   /* 指针变量的存储地址 */
   fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

   /* 使用指针访问值 */
   fmt.Printf("*ip 变量的值: %d\n", *ip )
}
```

## Go 语言结构体

结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。

```go
type struct_variable_type struct {
   member definition;
   member definition;
   ...
   member definition;
}

// 一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：
variable_name := struct_variable_type {value1, value2...valuen}
```

```go
package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}

func main() {
   var Book1 Books        /* 声明 Book1 为 Books 类型 */
   var Book2 Books        /* 声明 Book2 为 Books 类型 */

   /* book 1 描述 */
   Book1.title = "Go 语言"
   Book1.author = "www.runoob.com"
   Book1.subject = "Go 语言教程"
   Book1.book_id = 6495407

   /* book 2 描述 */
   Book2.title = "Python 教程"
   Book2.author = "www.runoob.com"
   Book2.subject = "Python 语言教程"
   Book2.book_id = 6495700

   /* 打印 Book1 信息 */
   fmt.Printf( "Book 1 title : %s\n", Book1.title)
   fmt.Printf( "Book 1 author : %s\n", Book1.author)
   fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
   fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

   /* 打印 Book2 信息 */
   fmt.Printf( "Book 2 title : %s\n", Book2.title)
   fmt.Printf( "Book 2 author : %s\n", Book2.author)
   fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
   fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)
}
```

## Go 语言切片(Slice)

Go 语言切片是对数组的抽象。Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

```go
// 声明一个未指定大小的数组来定义切片：
var identifier []type

// 或使用make()函数来创建切片:
var slice1 []type = make([]type, len)

// 也可以简写为

slice1 := make([]type, len)

//也可以指定容量，其中capacity为可选参数。
make([]T, length, capacity)

// 直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3
s :=[] int {1,2,3 }

// 初始化切片s,是数组arr的引用
s := arr[:]

// 将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
s := arr[startIndex:endIndex]

// 缺省endIndex时将表示一直到arr的最后一个元素
s := arr[startIndex:]

// 缺省startIndex时将表示从arr的第一个元素开始
s := arr[:endIndex] 

// 通过切片s初始化切片s1
s1 := s[startIndex:endIndex] 

// 通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片
s :=make([]int,len,cap)
```

* len() 函数获取长度
* cap() 函数测量切片最长可以达到多少
* append() 向切片追加新元素
* copy() 拷贝切片函数

## Go 语言范围(Range)

range 关键字用于for循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引值，在集合中返回 key-value 对的 key 值。

```go
package main
import "fmt"
func main() {
    //这是我们使用range去求一个slice的和。使用数组跟这个很类似
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)
    //在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
    //range也可以用在map的键值对上。
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
    //range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
```

## Go 语言Map(集合)

Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。

```go
/* 声明变量，默认 map 是 nil */
var map_variable map[key_data_type]value_data_type

/* 使用 make 函数 */
map_variable := make(map[key_data_type]value_data_type)
```

```go
package main

import "fmt"

func main() {
   var countryCapitalMap map[string]string
   /* 创建集合 */
   countryCapitalMap = make(map[string]string)
   
   /* map 插入 key-value 对，各个国家对应的首都 */
   countryCapitalMap["France"] = "Paris"
   countryCapitalMap["Italy"] = "Rome"
   countryCapitalMap["Japan"] = "Tokyo"
   countryCapitalMap["India"] = "New Delhi"

    /* 删除元素 */
   delete(countryCapitalMap,"France");
   
   /* 使用 key 输出 map 值 */
   for country := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",countryCapitalMap[country])
   }
   
   /* 查看元素在集合中是否存在 */
   captial, ok := countryCapitalMap["United States"]
   /* 如果 ok 是 true, 则存在，否则不存在 */
   if(ok){
      fmt.Println("Capital of United States is", captial)  
   }else {
      fmt.Println("Capital of United States is not present") 
   }
}
```

* delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key。

## Go 语言递归函数

递归，就是在运行的过程中调用自己。
我们在使用递归时，开发者需要设置退出条件，否则递归将陷入无限循环中。

```go
func recursion() {
   recursion() /* 函数调用自身 */
}

func main() {
   recursion()
}
```

## Go 语言类型转换

```go
// 类型转换基本格式如下
type_name(expression)

// 例如
package main

import "fmt"

func main() {
   var sum int = 17
   var count int = 5
   var mean float32
   
   mean = float32(sum)/float32(count)
   fmt.Printf("mean 的值为: %f\n",mean)
}
```

## Go 语言接口

```go
/* 定义接口 */
type interface_name interface {
   method_name1 [return_type]
   method_name2 [return_type]
   method_name3 [return_type]
   ...
   method_namen [return_type]
}

/* 定义结构体 */
type struct_name struct {
   /* variables */
}

/* 实现接口方法 */
func (struct_name_variable struct_name) method_name1() [return_type] {
   /* 方法实现 */
}
...
func (struct_name_variable struct_name) method_namen() [return_type] {
   /* 方法实现*/
}
```

## Go 错误处理

error类型是一个接口类型，这是它的定义

```go
type error interface {
    Error() string
}
```

我们可以在编码中通过实现 error 接口类型来生成错误信息。

函数通常在最后的返回值中返回错误信息。使用errors.New 可返回一个错误信息：

```go
func Sqrt(f float64) (float64, error) {
  if f < 0 {
      return 0, errors.New("math: square root of negative number")
  }
  // 实现
}

result, err:= Sqrt(-1)

if err != nil {
   fmt.Println(err)
}
```

```go
package main

import (
    "fmt"
)

// 定义一个 DivideError 结构
type DivideError struct {
    dividee int
    divider int
}

// 实现     `error` 接口
func (de *DivideError) Error() string {
    strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
    return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
    if varDivider == 0 {
        dData := DivideError{
            dividee: varDividee,
            divider: varDivider,
        }
        errorMsg = dData.Error()
        return
    } else {
        return varDividee / varDivider, ""
    }

}

func main() {

    // 正常情况
    if result, errorMsg := Divide(100, 10); errorMsg == "" {
        fmt.Println("100/10 = ", result)
    }
    // 当被除数为零的时候会返回错误信息
    if _, errorMsg := Divide(100, 0); errorMsg != "" {
        fmt.Println("errorMsg is: ", errorMsg)
    }

}
```

## 关键字

|25 个|关键字|或|保留字|:|
|-----|------|-------|------|------|
|break|	default|	func|	interface|	select|
|case	|defer	|go	|map	|struct|
|chan	|else	|goto	|package	|switch|
|const	|fallthrough	|if	|range	|type|
|continue	|for	|import	|return	|var|

|36 |个|预|定|义|标识|符|
|-----|-----|-----|-----|-----|-----|-----|
|append|	bool|	byte|	cap|	close	|complex|	complex64	|complex128|	uint16|
|copy	|false	|float32	|float64	|imag	|int	|int8	int16	|uint32|
|int32	|int64	|iota	|len	|make	|new	|nil	|panic|	uint64|
|print	|println	|real	|recover	|string|	true|	uint|	uint8|	uintptr|

## Go 语言数据类型

|序号	|类型和描述|
|---|---------------------------------------|
|1	|__布尔型__，只可以是常量 true 或者 false。一个简单的例子：var b bool = true。|
|2	|__数字类型__， 整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且原生支持复数，其中位的运算采用补码。|
|3	|__字符串类型__， 字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是由单个字节连接起来的。Go语言的字符串的字节使用UTF-8编码标识Unicode文本。|
|4	|__派生类型__， 包括：(a) 指针类型（Pointer）、(b) 数组类型、(c) 结构化类型(struct)、(d) Channel 类型、(e) 函数类型、(f) 切片类型、(g) 接口类型（interface）、(h) Map 类型|

### 数字类型

|序号	|类型和描述|
|---|-----------------------|
|1	|__uint8__，无符号 8 位整型 (0 到 255)|
|2	|__uint16__，无符号 16 位整型 (0 到 65535)|
|3	|__uint32__，无符号 32 位整型 (0 到 4294967295)|
|4	|__uint64__，无符号 64 位整型 (0 到 18446744073709551615)|
|5	|__int8__，有符号 8 位整型 (-128 到 127)|
|6	|__int16__，有符号 16 位整型 (-32768 到 32767)|
|7	|__int32__，有符号 32 位整型 (-2147483648 到 2147483647)|
|8	|__int64__，有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)|

#### 浮点型

|序号	|类型和描述|
|---|-----------------------|
|1	|__float32__，IEEE-754 32位浮点型数|
|2	|__float64__，IEEE-754 64位浮点型数|
|3	|__complex64__，32 位实数和虚数|
|4	|__complex128__，64 位实数和虚数|

#### 其他数字类型

|序号	|类型和描述|
|---|-----------------------|
|1	|__byte__，类似 uint8|
|2	|__rune__，类似 int32|
|3	|__uint__，32 或 64 位|
|4	|__int__，与 uint 一样大小|
|5	|__uintptr__，无符号整型，用于存放一个指针|
