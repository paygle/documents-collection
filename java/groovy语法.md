# Groovy 快速入门

* [注释](#docs)
* [标识符](#flag)
* [字符串](#string)
* [字符](#char)
* [布尔类型](#bool)
* [数字类型](#number)
* [数学计算](#math)
* [列表](#list)
* [数组](#array)
* [Map](#map)
* [运算符（和Java类似的运算符）](#operator)
* [可空运算符](#empty)
* [安全导航运算符](#safe)
* [字段访问运算符](#visit)
* [方法指针运算符](#methodp)
* [展开运算符](#expand)
* [范围运算符](#scope)
* [比较运算符](#compare)
* [成员运算符](#member)
* [相等运算符](#equal)
* [转换运算符](#convert)
* [声明变量](#define)
* [循环语句](#for)
* [断言语句](#assert)
* [面向对象-构造器](#constructor)
* [面向对象-方法](#methods)
* [面向对象-字段](#field)
* [面向对象-属性](#attr)
* [面向对象-特征类](#sclass)
* [闭包](#cloze)
* [调用闭包](#clozecall)
* [使用闭包](#clozeuse)

Groovy是一门基于JVM的动态语言，很多语法和Java类似。大部分Java代码也同时是合法的Groovy代码。本文是快速入门，所以针对语法并不会做非常详细的介绍。如果需要详细语法，请直接查看Groovy官方文档。另外为了省事，本文中的大部分代码例子直接引用了Groovy文档。由于其运行在 JVM 上的特性，Groovy 可以使用其他 Java 语言编写的库。



## Goovy 关键字(在Java的基础上多四个)

| 关键字    | as        | def        | in           | trait      |          |
| --------- | --------- | ---------- | ------------ | ---------- | -------- |
| abstract  | boolean   | break      | byte         | case       | catch    |
| char      | class     | const      | continue     | default    | do       |
| double    | else      | extends    | final        | finally    | float    |
| for       | if        | implements | import       | instanceof | int      |
| interface | long      | goto       | new          | native     | package  |
| private   | protected | public     | return       | short      | strictfp |
| super     | switch    | static     | synchronized | this       | throw    |
| throws    | transient | try        | void         | volatile   | while    |


* Groovy 会默认导入下面这些包、类，不需要使用import语句显式导入。

```groovy
import java.io.*
import java.lang.*
import java.math.BigDecimal
import java.math.BigInteger
import java.net.*
import java.util.*
import groovy.lang.*
import groovy.util.*
```

## Groovy类和java类比较

    1. 不需public修饰符， 默认访问修饰符就是public
    2. 不需要类型说明， 不关心变量和方法参数的具体类型
    3. 不需要getter/setter方法， 
    4. 不需要构造函数， 因为实际上只需要两个构造函数（1个不带参数的默认构造函数，1个只带一个map参数的构造函数--由于是map类型，通过这个参数可以构造对象时任意初始化它的成员变量）。
    5. 不需要return 来返回值。
    6. 方法调用可以省略（）（构造函数除外）。
    7. 额外的关键字: def、as、in、trait都是关键字


## 基本内容

### <a name="docs"></a>GroovyDoc 注释

单行注释，以//开头。 多行注释，/* */。

```groovy
/**
* Creates a greeting method for a certain person.
*
* @param otherPerson the person to greet
* @return a greeting message
*/
```

Shebang注释，和Linux其他注释类似，用于指定脚本解释器的位置。

```groovy
#!/usr/bin/env groovy
```

### <a name="flag"></a>标识符

大体上，Groovy标识符的命名规则和Java差不多。如果某个标识符在Groovy中合法，在Java中不合法，我们可以使用单引号或双引号将标识符包括起来。

### <a name="string"></a>字符串

字符串可以使用单引号'或双引号"包括起来。

```groovy
def str="1234"
def str2='1234'
```

多行字符串可以使用三个连续的单引号或双引号包括。不论是单行还是多行字符串， 都可以使用反斜杠转义字符。

```groovy
def multiline="""line1
line2
line3
"""
```

我们还可以将变量直接插入到字符串中，这叫做内插字符串（String interpolation）。语法和EL表达式类似。编译器会把美元和花括号中的内容替换成实际的值，内插字符串中还可以进行表达式计算。

```groovy
def name = 'Guillaume' // a plain string
def greeting = "Hello ${name}"
```

当内插字符串可以由前后的符号区分出的时候，花括号可以省略

```groovy
def person = [name: 'Guillaume', age: 36]
assert "$person.name is $person.age years old" == 'Guillaume is 36 years old'
```

当使用内插字符串的时候，字符串字面值是Groovy的字符串类型GString。这一点需要注意。普通的Java字符串是不变的，而GString是可变的。另外它们的哈希值也不同。因此在使用Map等数据类型的时候需要格外注意，避免使用GString作为Map的键。

### <a name="char"></a>字符

Groovy没有字符字面量。如果需要向Java方法传入单个字符的话，可以使用下面的方法进行转换。

```groovy
def c2 = 'B' as char 
assert c2 instanceof Character

def c3 = (char)'C' 
assert c3 instanceof Character
```

### <a name="bool"></a>布尔类型

Groovy的布尔类型和Java类似，也有true和false两个值。不过Groovy的布尔语义更丰富。未到结尾的迭代器、非空对象引用、非零数字都认为是真；空集合、空字符串等认为是假。详情参见Groovy文档

### <a name="number"></a>数字类型

Groovy支持byte、char 、short、 int 、long和 BigInteger等几种数字类型。如果使用普通方式声明，它们和Java中的变量很相似。

```groovy
int   i = 4

BigInteger bi =  6
```

如果使用def关键字声明，那么这些数字会自动选择可以容纳它们的类型。

```groovy
def a = 1
assert a instanceof Integer

// Integer.MAX_VALUE
def b = 2147483647
assert b instanceof Integer

// Integer.MAX_VALUE + 1
def c = 2147483648
assert c instanceof Long

// Long.MAX_VALUE
def d = 9223372036854775807
assert d instanceof Long

// Long.MAX_VALUE + 1
def e = 9223372036854775808
assert e instanceof BigInteger
```

这些整数还可以添加0b、0和0x前缀，分别代表8进制数，8进制数和16进制数。

另外Groovy还支持float、double和BigDecimal三种浮点数类型。原理同上。还可以使用科学计数法1.123E10这样的形式代表浮点数。

Groovy的数字常量同样支持后缀区分字面值类型，这几种类型和Java中的类似。唯一不同的是Groovy还支持G和g后缀，代表BigInteger或BigDecimal类型，根据字面值是否含有小数点来区分。

### <a name="math"></a>数学计算

数字的计算结果和Java规则类似：小于int的整数类型会被提升为int类型，计算结果也是int类型；小于long的整数类型和long计算，结果是long类型；BigInteger和其它整数类型计算，结果是BigInteger类型；BigDecimal和其它整数类型计算，结果是BigDecimal类型；BigDecimal和float、double等类型计算，结果是double类型。

### <a name="list"></a>列表
Groovy中的列表比较灵活，有点像Python中的列表。使用[....]语法可以声明列表，默认情况下列表是ArrayList实现。我们也可以使用as运算符自己选择合适的列表底层类型。

```groovy
def arrayList = [1, 2, 3]
assert arrayList instanceof java.util.ArrayList

def linkedList = [2, 3, 4] as LinkedList    
assert linkedList instanceof java.util.LinkedList
```

有了列表之后，就可以使用它了。使用方法和Python差不多。我们使用[索引]引用和修改列表元素。如果索引是负的，则从后往前计数。要在列表末尾添加元素，可以使用左移运算符<<。如果在方括号中指定了多个索引，会返回由这些索引对应元素组成的新列表。使用两个点加首位索引..可以选择一个子列表。

```groovy
def letters = ['a', 'b', 'c', 'd']

assert letters[0] == 'a'     
assert letters[1] == 'b'

assert letters[-1] == 'd'    
assert letters[-2] == 'c'

letters[2] = 'C'             
assert letters[2] == 'C'

letters << 'e'               
assert letters[ 4] == 'e'
assert letters[-1] == 'e'

assert letters[1, 3] == ['b', 'd']         
assert letters[2..4] == ['C', 'd', 'e'] 

// 列表还可以组合成复合列表。

def multi = [[0, 1], [2, 3]]     
assert multi[1][0] == 2   
```

### <a name="array"></a>数组

声明数组的方式和列表一样，只不过需要显示指定数组类型。数组的使用方法也和列表类似，只不过由于数组是不可变的，所以不能像数组末尾添加元素。

```groovy
int[] intArray = [1, 2, 3, 4, 5]

def intArray2 = [1, 2, 3, 4, 5, 6] as int[]
```

### <a name="map"></a>Map

创建Map同样使用方括号，不过这次需要同时指定键和值了。Map创建好之后，我们可以使用[键]或.键来访问对应的值。默认情况下创建的Map是java.util.LinkedHashMap，我们可以声明变量类型或者使用as关键字改变Map的实际类型。

```groovy
def colors = [red: '#FF0000', green: '#00FF00', blue: '#0000FF']   

assert colors['red'] == '#FF0000'    
assert colors.green  == '#00FF00' 
```

关于Map有一点需要注意。如果将一个变量直接作为Map的键的话，其实Groovy会用该变量的名称作为键，而不是实际的值。如果需要讲变量的值作为键的话，需要在变量上添加小括号。

```groovy
def key = 'name'
def person = [key: 'Guillaume']      //键是key而不是name

assert !person.containsKey('name')   
assert person.containsKey('key') 

//这次才正确的将key变量的值作为Map的键
person = [(key): 'Guillaume']        

assert person.containsKey('name')    
assert !person.containsKey('key')  
```

### <a name="operator"></a>运算符（和Java类似的运算符）

Groovy的数学运算符和Java类似，只不过多了一个乘方运算 \*\* 和乘方赋值 \*\*=

Groovy的关系运算符（大于、小于等于这些）和Java类似。

Groovy的逻辑运算符（与或非这些）和Java类似，也支持短路计算。

Groovy的位运算符合Java类似。

Groovy的三元运算符条件?值1:值2和Java类似。

### <a name="empty"></a>可空运算符

Groovy支持Elvis操作符，当对象非空的时候结果是值1，为空时结果是值2。或者更直接，对象非空是使用对象本身，为空时给另一个值，常用于给定某个可空变量的默认值。

```groovy
displayName = user.name ? user.name : 'Anonymous'   
displayName = user.name ?: 'Anonymous'  
```

### <a name="safe"></a>安全导航运算符

当调用一个对象上的方法或属性时，如果该对象为空，就会抛出空指针异常。这时候可以使用?.运算符，当对象为空时表达式的值也是空，不会抛出空指针异常。

```groovy
def person = Person.find { it.id == 123 }    
def name = person?.name                      
assert name == null  
```

### <a name="visit"></a>字段访问运算符

在Groovy中默认情况下使用点运算符.会引用属性的Getter或Setter。如果希望直接访问字段，需要使用.@运算符。

```groovy
class User {
    public final String name                 
    User(String name) { this.name = name}
    String getName() { "Name: $name" }       
}
def user = new User('Bob')
assert user.name == 'Name: Bob'   
assert user.@name == 'Bob'   
```

### <a name="methodp"></a>方法指针运算符

我们可以将方法赋给变量，这需要使用.&运算符。然后我们就可以像调用方法那样使用变量。方法引用的实际类型是Groovy的闭包Closure。这种运算符可以将方法作为参数，让Groovy语言非常灵活。

```groovy
def str = 'example of method reference'            
def fun = str.&toUpperCase                         
def upper = fun()                                  
assert upper == str.toUpperCase()   
```

### <a name="expand"></a>展开运算符

展开运算符*.会调用一个列表上所有元素的相应方法或属性，然后将结果再组合成一个列表。展开运算符还可以用于展开方法参数、列表和Map。

```groovy
class Car {
    String make
    String model
}
def cars = [
       new Car(make: 'Peugeot', model: '508'),
       new Car(make: 'Renault', model: 'Clio')
    ]       
def makes = cars*.make 
assert makes == ['Peugeot', 'Renault'] 

// 展开运算符是空值安全的，如果遇到了null值，不会抛出空指针异常，而是返回空值。
cars = [
   new Car(make: 'Peugeot', model: '508'),
   null,                                              
   new Car(make: 'Renault', model: 'Clio')
]
assert cars*.make == ['Peugeot', null, 'Renault']     
assert null*.make == null  
```

### <a name="scope"></a>范围运算符

使用..创建范围。默认情况下范围是闭区间，如果需要开闭区间可以在结束范围上添加<符号。范围的类型是groovy.lang.Range，它继承了List接口，也就是说我们可以将范围当做List使用。

```groovy
def range = 0..5                                    
assert (0..5).collect() == [0, 1, 2, 3, 4, 5]       
assert (0..<5).collect() == [0, 1, 2, 3, 4]         
assert (0..5) instanceof List                       
assert (0..5).size() == 6   
```

### <a name="compare"></a>比较运算符

```groovy
// <=>运算符相当于调用compareTo方法。
assert (1 <=> 1) == 0
```

### <a name="member"></a>成员运算符

```groovy
// 成员运算符in相当于调用contains或isCase方法。
def list = ['Grace','Rob','Emmy']
assert ('Emmy' in list)
```

### <a name="equal"></a>相等运算符

```groovy
// ==运算符和Java中的不同。在Groovy中它相当于调用equals方法。如果需要比较引用，使用is。
def list1 = ['Groovy 1.8','Groovy 2.0','Groovy 2.3']        
def list2 = ['Groovy 1.8','Groovy 2.0','Groovy 2.3']        
assert list1 == list2    //比较内容相等                                   
assert !list1.is(list2)   //比较引用相等
```

### <a name="convert"></a>转换运算符

我们可以使用Java形式的(String) i来转换类型。但是假如类型不匹配的话，就会抛出ClassCastException。而使用as运算符就会避免这种情况。
如果希望自己的类也支持as运算符的话，需要实现asType方法。

```groovy
Integer x = 123
String s = x as String 
```

##表达式和语句

### <a name="define"></a>声明变量

Groovy支持以传统方式使用变量类型 变量名的方式声明变量，也可以使用def关键字声明变量。使用def关键字的时候，变量类型由编译器自动推断，无法推断时就是Object类型。

```groovy
// Groovy可以同时声明多个变量。
def (a, b, c) = [10, 20, 'foo']
// 如果左边的变量数比右面的值多，那么剩余的变量就是null。
def (a, b, c) = [1, 2]
assert a == 1 && b == 2 && c == null
// 如果等号右面比左面多，那么多余的值会被忽略。
def (a, b) = [1, 2, 3]
assert a == 1 && b == 2
// 自定义对象也可以用多重赋值进行对象解构。该对象必须有getAt方法。
class CustomDestruction {
    int a
    int b
    int c
    //解构需要实现getAt方法
    def getAt(int i) {
        switch (i) {
            case 0: a; break
            case 1: b; break
            case 2: c; break
            default: throw new IllegalArgumentException()
        }
    }
    static void main(String[] args) {
        //对象解构
        def obj = new CustomDestruction(a: 3, b: 4, c: 5)
        def (x, y, z) = obj
        println("x=$x,y=$y,z=$z")
    }
}
```

##条件语句

Groovy的if语句和Java的类似。不过在Groovy中布尔值的真假不仅看条件比较的结果，还可以以其他情况判断。前面已经介绍过了。switch语句同理，真值判断非常自由。详情可参见Groovy文档 真值判断。

### <a name="for"></a>循环语句

Groovy支持传统的Java的for(int i=0;i<N;i++)和for(int i :array)两种形式。另外还支持for in loop形式，支持迭代范围、列表、Map、数组等多种形式的集合。
while语句的形式和Java相同。

```groovy
// 迭代范围
def x = 0
for ( i in 0..9 ) {
    x += i
}
assert x == 45

// 迭代列表
x = 0
for ( i in [0, 1, 2, 3, 4] ) {
    x += i
}
```

### <a name="assert"></a>断言语句

前面我们看到了很多Groovy断言的例子。Groovy断言语句的功能很强大，以至于文档中写的是强力断言（Power assertion）。

Groovy断言的形式如下。Groovy断言和Java断言完全不同。Groovy断言是一项语言功能，一直处于开启状态，和JVM的断言功能-ea完全无关。所以它是我们进行单元测试的首选方式。

```groovy
assert [left expression] == [right expression] : (optional message)

// 比如我们要断言1+1=3。结果应该类似这样。越复杂的表达式，断言效果越清晰。有兴趣的同学可以试试。
Caught: Assertion failed:

assert 1+1 == 3
        |  |
        2  false
```

## 面向对象编程

Groovy的面向对象编程和Java类似，但是提供了一系列功能简化面向对象开发。当然如果你想使用传统的Java语法来声明所有成员也可以，Groovy设计目的之一就是让Java程序员能够以低成本的方式切换到Groovy上。

    * 字段默认是私有的，Groovy会自动实现Getter和Setter方法。
    * 方法和属性默认是公有的。
    * 类不必和文件名相同，
    * 一个文件可以有多个类，如果一个类也没有，该文件就会被看做是脚本。

### <a name="constructor"></a>构造器

Groovy的构造器非常灵活，我们可以使用传统的Java方式声明和使用构造器，也可以完全不声明构造器。有时候不声明反而更简单。如果没有声明构造器的话，我们可以在构造对象的时候使用命名参数方式传递参数，这种方式非常方便，因为我们不需要声明所有参数，只要声明所需的参数即可。

如果希望对构造器进行限制，可以手动声明构造器，这样这种自动构造就不会进行了。

```groovy
class Product {
    String name
    double price

    String toString() {
        return "Product(name:$name,price:$price)"
    }

    static void main(String[] args) {
        def product = new Product(name: 'AMD Ryzen 1700', price: 2499)
        println("隐式构造器:$product")
    }
}
```

### <a name="methods"></a>方法

Groovy方法和Java方法基本相同。不过Groovy方法更方便：支持命名参数和默认参数。另外Groovy方法可以使用def关键字声明，这时候方法返回类型是Object。在Groovy中方法的返回语句可以省略，这时候编译器会使用方法的最后一个语句的值作为返回值。在前面我们还看到了def关键字定义变量，这时候变量的类型需要从代码中推断。

在使用命名参数的时候需要注意一点，方法参数需要声明为Map类型（不需要详细指定键和值的类型），在调用方法的时候使用命名参数方式传入参数。

```groovy
def foo(Map args) { "${args.name}: ${args.age}" }

//调用方法
foo(name: 'Marie', age: 1)

// 另外方法的括号是可选的，我们可以省略括号直接像这样调用方法。
methodWithDefaultParam '555', 42
```

### <a name="field"></a>字段

Groovy中字段和Java中的概念类似。不过Groovy更加方便：默认情况下字段是私有的，Groovy自动生成字段的Getter和Setter。如果需要更精细的控制，把它当成Java字段用就行了。不过如果自定义字段的话，Groovy不会自动生成对应的属性了。

### <a name="attr"></a>属性

如果字段上面没有声明访问修饰符（private、public这些），Groovy就会自动生成Gettter和Setter。如果字段是final的，那么只会生成Getter。这就是Groovy方便的属性功能。

当然Groovy的方便不止于此，我们的所有类似Java访问字段的语法，实际上都会调用字段对应的Getter和Setter。这样显著减少了代码量。如果在类内部的话，.字段语法会直接访问字段，这样做是为了防止无限递归调用属性。

下面的例子中，第一次调用p.name = 'Marge'如果在类内部，就直接写入字段，如果调用在类外部，就会使用Setter写入。第二次调用使用了方法语法，直接使用Setter写入，所以不管在类内还是类外，写入的值都是"Wonder$name"。

```groovy
class Person {
    String name
    void name(String name) {
        this.name = "Wonder$name"       
    }
    String wonder() {
        this.name                       
    }
}
def p = new Person()
p.name = 'Marge'                        
assert p.name == 'Marge'                
p.name('Marge')                         
assert p.wonder() == 'WonderMarge' 
```

### <a name="sclass"></a>特征类

Groovy和Scala一样，支持特征类（trait）。特征类就好像自带实现的接口。在Java中只能继承一个类和多个接口。在Groovy中，我们可以继承多个特征类。特征类和普通的Groovy类一样，可以包括属性、字段、方法等，特征类也可以是抽象的。

使用特征类，我们可以在Groovy中实现类似C++的多重继承。另外，特征类支持运行时动态绑定，在某些情况下非常有用。

```groovy
trait Readable {
    void read() {
        println("read...")
    }
}

trait Writable {
    void write(String text) {
        println("write $text")
    }
}

class Notebook implements Readable, Writable {
    static void main(String[] args) {
        //使用特性类
        def notebook = new Notebook()
        notebook.read()
        notebook.write("something")
    }
}
```

### <a name="cloze"></a>闭包

闭包是Groovy非常重要的一个功能，也是我们介绍的最后一个功能。要了解闭包，最好先知道Java的Lambda表达式、匿名内部类等概念。Groovy闭包和Lambda表达式概念相近，但是功能更强大。

```groovy
// 声明闭包
{ [closureParameters -> ] statements }
```

以下都是合法的Groovy闭包。所有闭包都是groovy.lang.Closure类型的实例。闭包的参数类型是可选的。如果闭包只有单个参数，参数名也是可选的。Groovy会隐式指定it作为参数名。Kotlin语言也是类似的做法，有助于我们先出可读性很好的闭包。

```groovy
{ item++ }                                          

{ -> item++ }                                       

{ println it }                                      

{ it -> println it }                                

{ name -> println name }                            

{ String x, int y ->                                
    println "hey ${x} the value is ${y}"
}

{ reader ->                                         
    def line = reader.readLine()
    line.trim()
}

```

### <a name="clozecall"></a>调用闭包

```groovy
// 闭包既可以当做方法来调用，也可以显示调用call方法。
def code = { 123 }
assert code() == 123
assert code.call() == 123

// 调用有参数的闭包也是类似的。
def isOdd = { int i-> i%2 == 1 }                            
assert isOdd(3) == true                                     
assert isOdd.call(2) == false  
```

### <a name="clozeuse"></a>使用闭包

Groovy闭包类似Java的Lambda表达式和匿名内部类，不过使用更方便，能让我们减少不少代码量。闭包还可以作为方法参数传递到其他地方，这让闭包更加灵活。

```groovy
static void funWithClosure(Closure closure) {
    closure()
}

//在其他地方调用该方法
funWithClosure({ println("Hello yitian") })

//括号还可以省略，更加简练
funWithClosure { println("Hello yitian") }
```

关于闭包，还有几个精彩的例子，就是Gradle脚本和Groovy模板引擎，它们都利用了Groovy强大的闭包功能，构建出简练而强大的DSL，让我们用很少的代码就可以实现强大的功能（虽然学起来稍微复杂点）。

