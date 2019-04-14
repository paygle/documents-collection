# Kotlin 语言

* [关键字](#kotlin)

### 硬关键字 （始终解释为关键字，不能用作标识符）

| 关键字 | 说明 |
|--------------|--------------------------------------------------------|
|    as     | 用于类型转换；为导入指定一个别名|
|    as?    | 用于安全类型转换 |
|  break    | 终止循环的执行 |
| class |  声明一个类 |
| continue | 继续最近层循环的下一步 |
|  do  | 开始一个 do/while 循环（后置条件的循环） |
| else | 定义一个 if 表达式条件为 false 时执行的分支 |
| false | 指定布尔类型的“假”值 |
| for | 开始一个 for 循环 |
| fun | 声明一个函数 |
| if  | 开始一个 if 表达式 |
| in  | 指定在 for 循环中迭代的对象；用作中缀操作符以检查一个值属于一个区间、 一个集合或者其他定义“contains”方法的实体；在 when 表达式中用于上述目的；将一个类型参数标记为逆变 |
| !in | 用作中缀操作符以检查一个值不属于一个区间、 一个集合或者其他定义“contains”方法的实体； 在 when 表达式中用于上述目的 |
| interface | 声明一个接口 |
| is | 检查一个值具有指定类型； 在 when 表达式中用于上述目的 |
| !is | 检查一个值不具有指定类型； 在 when 表达式中用于上述目的 |
| null | 是表示不指向任何对象的对象引用的常量 |
| object | 同时声明一个类及其实例 |
| package | 指定当前文件的包 |
| return  | 从最近层的函数或匿名函数返回 |
| super | 引用一个方法或属性的超类实现；在次构造函数中调用超类构造函数 |
| this | 引用当前接收者； 在次构造函数中调用同一个类的另一个构造函数 |
| throw |  抛出一个异常 |
| true  | 指定布尔类型的“真”值 |
| try  | 开始一个异常处理块 |
| typealias  | 声明一个类型别名 |
| val  | 声明一个只读属性或局部变量 |
| var  |  声明一个可变属性或局部变量 |
| when  | 开始一个 when 表达式（执行其中一个给定分支）|
| while | 开始一个 while 循环（前置条件的循环） |

### 软关键字（在适用的上下文中充当关键字，而在其他上下文中可用作标识符）

| 关键字 | 说明 |
|--------------|---------------------------------------------------------------|
| by  |  将接口的实现委托给另一个对象； 将属性访问器的实现委托给另一个对象 |
| catch  |  开始一个处理指定异常类型的块 |
| constructor | 声明一个主构造函数或次构造函数 |
| delegate  |  用作注解使用处目标 |
| dynamic  | 引用一个 Kotlin/JS 代码中的动态类型 |
| field  | 用作注解使用处目标 |
| file | 用作注解使用处目标 |
| finally | 开始一个当 try 块退出时总会执行的块 |
| get  | 声明属性的 getter； 用作注解使用处目标 |
| import  | 将另一个包中的声明导入当前文件 |
| init | 开始一个初始化块 |
| param  | 用作注解使用处目标 |
| property | 用作注解使用处目标 |
| receiver | 用作注解使用处目标 |
|  set  | 声明属性的 setter； 用作注解使用处目标 |
| setparam | 用作注解使用处目标 |
| where  | 指定泛型类型参数的约束 |

### 修饰符关键字 （作为声明中修饰符列表中的关键字，并可用作其他上下文中的标识符）

| 修饰符 | 说明 |
|--------------|---------------------------------------------------------------|
| actual | 表示多平台项目中的一个平台相关实现 |
| abstract | 将一个类或成员标记为抽象 |
| annotation | 声明一个注解类 |
| companion  | 声明一个伴生对象 |
| const  | 将属性标记为编译期常量 |
| rossinline  | 禁止传递给内联函数的 lambda 中的非局部返回 |
| data | 指示编译器为类生成典型成员 |
| enum | 声明一个枚举 |
| expect | 将一个声明标记为平台相关，并期待在平台模块中实现 |
| external  | 将一个声明标记为不是在 Kotlin 中实现（通过JNI访问或者在JavaScript中实现） |
| final | 禁止成员覆盖 |
| infix  | 允许以中缀表示法调用函数 |
| inline  | 告诉编译器在调用处内联传给它的函数和 lambda 表达式 |
| inner | 允许在嵌套类中引用外部类实例 |
| internal | 将一个声明标记为在当前模块中可见 |
| lateinit | 允许在构造函数之外初始化非空属性 |
| noinline  | 关闭传给内联函数的 lambda 表达式的内联 |
|  open  | 允许一个类子类化或覆盖成员 |
| operator  | 将一个函数标记为重载一个操作符或者实现一个约定 |
|  out  | 将类型参数标记为协变 |
| override | 将一个成员标记为超类成员的覆盖 |
| private | 将一个声明标记为在当前类或文件中可见 |
| protected | 将一个声明标记为在当前类及其子类中可见 |
| public | 将一个声明标记为在任何地方可见 |
| reified | 将内联函数的类型参数标记为在运行时可访问 |
| sealed  | 声明一个密封类（限制子类化的类）|
| suspend |  将一个函数或 lambda 表达式标记为挂起式（可用做协程） |
| tailrec |  将一个函数标记为尾递归（允许编译器将递归替换为迭代） |
| vararg  | 允许一个参数传入可变数量的参数|

### 特殊标识符 （由编译器在指定上下文中定义，并且可以用作其他上下文中的常规标识符）

| 特殊标识符 | 说明 |
|--------------|---------------------------------------------------------------|
| field  | 用在属性访问器内部来引用该属性的幕后字段 |
| it  | 用在 lambda 表达式内部来隐式引用其参数 |

### 操作符和特殊符号

| 操作符 | 说明 |
|--------------|-----------------------------------------------------|
| +、-、*、/、%  | 数学操作符；* 也用于将数组传递给 vararg 参数 |
|  =  | 值操作符，也用于指定参数的默认值 |
| =、-=、*=、/=、%= | 广义赋值操作符 |
| ++、-- | 增与递减操作符 |
| &&、\|\|、 ! | 逻辑“与”、“或”、“非”操作符（对于位运算，请使用相应的中缀函数）|
| ==、!= | 相等操作符（对于非原生类型会翻译为调用  equals() ） |
| ===、!== | 引用相等操作符 |
| <、>、<=、>= | 比较操作符（对于非原生类型会翻译为调用  compareTo() ）[ 、  ]  —— 索引访问操作符（会翻译为调用  get  与  set ）|
|     !!  |  断言一个表达式非空 |
|     ?.  |  执行安全调用（如果接收者非空，就调用一个方法或访问一个属性） |
|     ?:  |  如果左侧的值为空，就取右侧的值（elvis 操作符） |
|     ::  |  创建一个成员引用或者一个类引用 |
|     ..  |  创建一个区间 |
|      :  |  分隔声明中的名称与类型 |
|      ?  |  将类型标记为可空 |
|     ->  |  分隔 lambda 表达式的参数与主体； 分隔在函数类型中的参数类型与返回类型声明； 分隔 when 表达式分支的条件与代码体 |
|      @  |  引入一个注解；引入或引用一个循环标签；引入或引用一个 lambda 表达式标签；引用一个来自外部作用域的 “this”表达式；引用一个外部超类 |
|      ;  |  分隔位于同一行的多个语句 |
|      $  |  在字符串模版中引用变量或者表达式 |
|      _  |  在 lambda 表达式中代替未使用的参数； 在解构声明中代替未使用的参数 |

### 基本类型

| 数字类型 | 位宽 |
|----|-----|
| Double | 64 |
| Float  | 32 |
| Long   | 64 |
| Int    | 32 |
| Short  | 16 |
| Byte   |  8 |

 * 十进制:  123； Long 类型用大写  L  标记:  123L
 * 十六进制:  0x0F
 * 二进制:  0b00001011
 * 注意: 不支持八进制
 * 类型不同的值是无法进行相等比较的（需要显示转换为相同的类型）

### 数字字面值中的下划线（使用下划线使数字常量更易读）

```kotlin
val oneMillion = 1_000_000
val creditCardNumber = 1234_5678_9012_3456L
val socialSecurityNumber = 999_99_9999L
val hexBytes = 0xFF_EC_DE_5E
val bytes = 0b11010010_01101001_10010100_10010010
```

### 字符串字面值

```kotlin
/*
* 原始字符串 使用三个引号（ """ ）分界符括起来，内部没有转义并且可以包含换行以及任何其他字符
* 通过  trimMargin()  函数去除前导空格
*/

val text = """
 |Tell me and I forget.
 |Teach me and I remember.  
 |Involve me and I learn.  
 |(Benjamin Franklin)  """.trimMargin()

/*
* 字符串模板 以美元符（ $ ）开头，由一个简单的名字构成或者用花括号括起来的任意表达式
*/

val i = 10
println("$s.length is ${s.length}")  //  输出“abc.length is 3”
```

## <a name="kotlin"></a>Kotlin 简单示例

  * 语法 while  与  do .. while  照常使用

  默认导入
  * kotlin.*
  * kotlin.annotation.*
  * kotlin.collections.*
  * kotlin.comparisons.* （自 1.1 起） 
  * kotlin.io.*
  * kotlin.ranges.*
  * kotlin.sequences.*
  * kotlin.text.*

  根据目标平台还会导入额外的包：
  * java.lang.*     （JVM）
  * kotlin.jvm.*    （JVM）
  * kotlin.js.*      (JS)


```kotlin
/*
* 定义包: 目录与包的结构无需匹配：源代码可以在文件系统的任意位置
*/
package my.demo 
import java.util.*  as jutil

/*
* 定义函数: 带有两个  Int  参数、返回  Int  的函数
*/
fun sum(a: Int, b: Int): Int {
  return a + b
}

fun main() {
   print("sum of 3 and 5 is ")
   println(sum(3, 5))
}

/*
* If 表达式
*/
val max = if (a > b) a else b

/*
* if 的分支可以是代码块，最后的表达式作为该块的值：
*/
val max = if (a > b) {
  print("Choose a")   
  a
} else {
  print("Choose b")  
  b
}

/*
* When 表达式
*/
when (x) {
  1 -> print("x == 1")  
  2, 3 -> print("x == 2or3")  
  else -> { // 注意这个块
    print("x is neither 1 nor 2")
  }
}

/*
* When 检测一个值在（ in ）或者不在（ !in ）一个区间或者集合中
*/
when (x) {
  in 1..10 -> print("x is in the range")  
  in validNumbers -> print("x is valid")
  !in 10..20 -> print("x is outside the range") 
  else -> print("none of the above")
}

/*
* 另一种可能性是检测一个值是（ is ）或者不是（ !is ）一个特定类型的值。
* 注意： 由于智能转换，你可以访问该类型的方法与属性而无需任何额外的检测。
*/
fun hasPrefix(x: Any) = when(x) {
  is String -> x.startsWith("prefix")  
  else -> false
}

/*
* when  也可以用来取代  if - else   if 链。 
* 如果不提供参数，所有的分支条件都是简单的布尔表达式，而当一个分支的条件为真时则执行该分支
*/
when {
  x.isOdd() -> print("x is odd")
  x.isEven() -> print("x is even") 
  else -> print("x is funny")
}

/**
* for  循环可以对任何提供迭代器（iterator）的对象进行遍历
*/
for (item: Int in ints) {
  // ...
}

for (i in 6 downTo 0 step 2) {
  println(i)
}

/**
* 可以用标签限制  break  或者 continue 
*/
loop@ for (i in 1..100) {
  for (j in 1..100) {
    if (……) break@loop
  }
}

```

## 类与对象

  类声明由类名、类头（指定其类型参数、主构造函数等）以及由花括号包围的类体构成。类头与类体都是可选的； 如果一个类没有类体，可以省略花括号。

  * 幕后字段 field 标识符只能用在属性的访问器内。
  * 接口与 Java 8 类似，既包含抽象方法的声明，也包含实现。与抽象类不同的是，接口无法保存状态。它可以有属性但必须声明为抽象或提供访问器实现。
  * 伴生对象，在 Kotlin 中类没有静态方法，你可以把它写成该类内对象声明中的一员。

#### 在包顶层声明，访问限制符

  * 如果你不指定任何可见性修饰符，默认为  public ，这意味着你的声明将随处可见；
  * 如果你声明为  private ，它只会在声明它的文件内可见；
  * 如果你声明为  internal ，它会在相同模块内随处可见；
  * protected  不适用于顶层声明。
  * 注意：要使用另一包中可见的顶层声明，仍需将其导入进来。
  

```kotlin
// 文件名：example.kt 
package foo

private fun foo() { …… }  // 在 example.kt 内可见
public var bar: Int = 5   // 该属性随处可见
private set               // setter 只在 example.kt 内可见
internal val baz = 6      // 相同模块内可见

```

#### 对于类内部声明的成员：

  * private  意味着只在这个类内部（包含其所有成员）可见；
  * protected —— 和  private 一样 + 在子类中可见。
  * internal  —— 能见到类声明的 本模块内 的任何客户端都可见其  internal  成员；
  * public  —— 能见到类声明的任何客户端都可见其  public  成员。

  注意 对于Java用户：Kotlin 中外部类不能访问内部类的 private 成员。
  如果你覆盖一个  protected  成员并且没有显式指定其可见性，该成员还会是  protected  可见性。

```kotlin
open class Outer {
  private val a = 1  
  protected open val b = 2  
  internal val c = 3
  val d = 4  // 默认 public

  protected class Nested {
    public val e: Int = 5
  }
}

class Subclass : Outer() {
  // a 不可见
  // b、c、d 可见
  // Nested 和 e 可见
  override val b = 5   // “b”为 protected
}

class Unrelated(o: Outer) {
  // o.a、o.b 不可见
  // o.c 和 o.d 可见（相同模块）
  // Outer.Nested 不可见，Nested::e 也不可见
}

```


### 构造函数

  * 在 Kotlin 中的一个类可以有一个主构造函数以及一个或多个次构造函数。主构造函数是类头的一部分：它跟在类名（与可选的类型参数）后。
  * 主构造函数不能包含任何的代码。初始化的代码可以放到以  init  关键字作为前缀的初始化块（initializer blocks）中。
  * 请注意，主构造的参数可以在初始化块中使用。它们也可以在类体内声明的属性初始化器中使用
  * 在 Kotlin 中所有类都有一个共同的超类  Any ，这对于没有超类型声明的类是默认超类; 注意： Any  并不是  java.lang.Object ；尤其是，它除了  equals() 、 hashCode()  与 toString()  外没有任何成员。

```kotlin

  class Person constructor(firstName: String) { ... }

  // 如果主构造函数没有任何注解或者可见性修饰符，可以省略这个  constructor  关键字。

  class Person(firstName: String) { ... }

  /**
  * 声明属性以及从主构造函数初始化属性, Kotlin 简洁语法
  */
  class Person(val firstName: String, val lastName: String, var age: Int) { …… }

  /**
   * 如果构造函数有注解或可见性修饰符，这个  constructor  关键字是必需的，并且这些修饰符在它前面
   */
  class Customer public @Inject constructor(name: String) { …… }

  /**
  * 类也可以声明前缀有  constructor 的次构造函数：
  */
  class Constructors {

    var street: String = ……

    init {
      println("Init block")
    }

    constructor(i: Int) {
      println("Constructor")
    }

    constructor(ctx: Context, attrs: AttributeSet) : super(ctx, attrs)

  }

  fun main() {
    Constructors(1)
  }

  /**
  * 在一个内部类中访问外部类的超类，可以通过由外部类名限定的  super  关键字来实现： super@Outer 
  */
  class Bar : Foo() {
    override fun f() { /* …… */ }
    override val x: Int get() = 0

    inner class Baz {
      fun g() {
        super@Bar.f() // 调用 Foo 实现的 f()
        println(super@Bar.x) // 使用 Foo 实现的 x 的 getter
      }
    }
  }

```

### 类覆盖规则

  * 如果一个类从它的直接超类继承相同成员的多个实现， 它必须覆盖这个成员并提供其自己的实现（也许用继承来的其中之一）

```kotlin
open class A {
  open fun f() { print("A") }  
  fun a() { print("a") }
}

interface B {
  fun f() { print("B") } // 接口成员默认就是“open”的  
  fun b() { print("b") }
}

class C() : A(), B {
  // 编译器要求覆盖 f()：  
  override fun f() {
    super<A>.f() // 调用 A.f()  
    super<B>.f() // 调用 B.f()
  }
}
```
 
### 扩展函数 (需要用一个 接收者类型 也就是被扩展的类型来作为他的前缀)

  * 扩展不能真正的修改他们所扩展的类。通过定义一个扩展，你并没有在一个类中插入新成员， 仅仅是可以通过该类型的变量用点表达式去调用这个新函数。

```kotlin
fun MutableList<Int>.swap(index1: Int, index2: Int) {
  val tmp = this[index1] // “this”对应该列表  
  this[index1] = this[index2]
  this[index2] = tmp
}

// 泛化写法
fun <T> MutableList<T>.swap(index1: Int, index2: Int) {
  val tmp = this[index1] // “this”对应该列表  
  this[index1] = this[index2]
  this[index2] = tmp
}

/*
* 注意可以为可空的接收者类型定义扩展。
* 这样的扩展可以在对象变量上调用， 即使其值为
null，并且可以在函数体内检测  this == null ，这能让你在没有检测 null 的时候调用 Kotlin中的toString()：检测发生在扩展函数的内部。
*/
fun Any?.toString(): String {
  if (this == null) return "null"
  // 空检测之后，“this”会自动转换为非空类型，所以下面的 toString()  
  // 解析为 Any 类的成员函数
  return toString()
}

/*
* 扩展属性
*/
val <T> List<T>.lastIndex: Int
  get() = size - 1

/*
* 注意：由于扩展没有实际的将成员插入类中，因此对扩展属性来说幕后字段是无效的。这就是为什么扩展属性不能有初始化器。他们的行为只能由显式提供的 getters/setters 定义。
*/
val Foo.bar = 1 // 错误：扩展属性不能有初始化器

/*
* 伴生对象的扩展
*/
class MyClass {
  companion object { } 
}

fun MyClass.Companion.foo() { …… }

```

## 数据类 （一些只保存数据的类）

  * 数据类可以扩展其他类

```kotlin

  data class User(val name: String = "", val age: Int = 0)

```
