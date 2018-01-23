# python3 语法基础

## Python标识符

Python标识符是用来标识变量，函数，类，模块或其他对象的名称。标识符是以字母A到Z开始或a〜z或后跟零个或多个字母下划线(_)，下划线和数字(0〜9)。

Python标识符范围内的不容许有如：@, $ 和 % 符号。Python是一种区分大小写的编程语言。因此，Manpower 和 manpower 在Python中是两种不同的标识符。

### Python 标识符命名的约定

* 类名称使用大写字母。所有其它标识符开始使用小写字母。
* 开头使用一个下划线的标识符表示该标识符是私有的。
* 开始以两个前导下划线的标识符表示强烈私有的标识符。
* 如果标识符使用两个下划线作为结束时，所述标识符是语言定义的特殊的名字。

## 行和缩进

* Python不使用大括号({})来表示的代码块类和函数定义或流程控制。代码块由行缩进，这是严格执行表示。因此，Python中所有连续不换行，同样数量的空格缩进将形成一个块。

* 在缩进位的数目是可变的，但该块内的所有语句的缩进量必须相同。

### 多行语句

```py
# 许使用续行字符(\)表示让行可以连续下去
total = item_one + \
        item_two + \
        item_three

# 语句中包含 [], {}, 或() 括号内不需要使用续行字符。 例如
days = ['Monday', 'Tuesday', 'Wednesday',
        'Thursday', 'Friday']

```

## 在Python的引号

* Python接受单引号(')，双引号(“)和三('''或”“”)引用来表示字符串，只要是同一类型的引号开始和结束。
* 三重引号可用于跨越多个行字符串。例如，下面所有的都是合法的

```py
word = 'word'
sentence = "This is a sentence."
paragraph = """This is a paragraph. It is
made up of multiple lines and sentences."""
```

## Python中的注释

哈希符号(＃)这是一个字符作为注释的开头。在#之后到行末的所有字符都是注释的一部分，Python解释器会忽略它们。Python没有多行注释功能

```py
#!/usr/bin/python3

# First comment
print ("Hello, Python!") # second comment
```

## 在一行多条语句

分号(;)允许给在单行有多条语句，而不管语句开始一个新的代码块。

```py
import sys; x = 'foo'; sys.stdout.write(x + '\n')
```

## 多重声明组为套件


头部行开始的语句(以关键字)，并用冒号终止(：)，后面跟一行或多行组成套件。例如

```py
if expression :
   suite
elif expression : 
   suite 
else : 
   suite
```

## Python保留字

|所有|Python|的关键字|仅包含|小写|字母|
|---|---|---|---|---|---|
|and	|exec	|Not |as	|finally	|or|
|assert	|for	|pass|break	|from	|print|
|class	|global	|raise |continue	|if	|return|
|def	|import	|try |del	|in	|while|
|elif	|is	|with |else	|lambda	|yield|
|except|

## Python3变量类型

* 赋值给变量

Python变量不需要显式声明保留内存空间。当赋值给一个变量这些声明自动发生。等号(=)是用来赋值给变量。

=运算符的左边是变量名称，而 =运算符右侧是存储在变量的值。

```py
#!/usr/bin/python3

counter = 100          # An integer assignment
miles   = 1000.0       # A floating point
name    = "John"       # A string

print (counter)
print (miles)
print (name)
```

* 多重赋值

```py
# 同时分配一个值给几个变量。
a = b = c = 1

# 将多个对象同时分配多个变量。
a, b, c = 1, 2, "john"
```

## 标准数据类型

### Python数字

```py
# 数字数据类型存储数值。当分配一个值给创建数值对象。
var1 = 1
var2 = 10

# 使用 del 语句删除引用的那一个数字对象
del var1[,var2[,var3[....,varN]]]]

# 使用del语句删除单个或多个对象
del var_a, var_b
```

* Python支持四种数字类型: int (有符号整数)、float (浮点实数值)、complex (复数)、long的整数类型

|int	|float	|complex|
|-----|-------|------------|
|10	|0.0	|3.14j|
|100	|15.20	|45.j|
|-786|	-21.9	|9.322e-36j|
|080|	32.3+e18|	.876j|
|-0490|	-90.|	-.6545+0J|
|-0x260	|-32.54e100|	3e+26J|
|0x69	|70.2-E12|	4.53e-7j|

> 复数由一对有序组成，通过 x + yj 来表示实浮点数， 其中 x 和 y 是实数并且 j 是虚数单位

### Python字符串

字符串子集可以用切片操作符 ([ ] and [:] ) ：字符串的索引从0开始，并以-1结束。加号(+)号是字符串连接运算符和星号(*)是重复操作符。

```py
#!/usr/bin/python3

str = 'Hello World!'

print (str)          # Prints complete string
print (str[0])       # Prints first character of the string
print (str[2:5])     # Prints characters starting from 3rd to 5th
print (str[2:])      # Prints string starting from 3rd character
print (str * 2)      # Prints string two times
print (str + "TEST") # Prints concatenated string
```

### Python列表

列表是最通用的Python复合数据类型。列表中包含用逗号分隔并使用方括号[]来包含项目。从某种程度上讲，列表类似于C语言中的数组。一个较大的区别是，所有在一个列表中的项目可以是不同的数据类型。

存储在一个列表中的值可以使用切片操作符([]和[:])进行访问：列表的0索引位置为起点位置，并在以-1 结束。 加号(+)号是列表中连接运算，星号(*)是重复操作符。

```py
#!/usr/bin/python3

list = [ 'abcd', 786 , 2.23, 'john', 70.2 ]
tinylist = [123, 'john']

print (list)          # Prints complete list
print (list[0])       # Prints first element of the list
print (list[1:3])     # Prints elements starting from 2nd till 3rd 
print (list[2:])      # Prints elements starting from 3rd element
print (tinylist * 2)  # Prints list two times
print (list + tinylist) # Prints concatenated lists
```

### Python元组

元组是另一个序列数据类型，它类似于列表。元组中使用单个逗号来分隔每个值。不像列表，元组的值是放列在圆括号中。

列表和元组之间的主要区别是：列表是包含在方括号[]中，并且它们的元素和大小是可以改变的，而元组元素是括在括号()中，不能进行更新。元组可以被认为是只读的列表。

```py
#!/usr/bin/python3

tuple = ( 'abcd', 786 , 2.23, 'john', 70.2  )
tinytuple = (123, 'john')

print (tuple)           # Prints complete tuple
print (tuple[0])        # Prints first element of the tuple
print (tuple[1:3])      # Prints elements starting from 2nd till 3rd 
print (tuple[2:])       # Prints elements starting from 3rd element
print (tinytuple * 2)   # Prints tuple two times
print (tuple + tinytuple) # Prints concatenated tuple
```

### Python字典

字典是一种哈希表类型。它们工作的方式就类似在Perl中关联数组或哈希、键-值对。字典的键可以是几乎任何Python类型，但通常是数字或字符串。另一方面，它的值可以是任意Python对象。

字典是用大括号括起来({})，并且值可分配并使用方括号([])来访问。

```py
#!/usr/bin/python3

dict = {}
dict['one'] = "This is one"
dict[2]     = "This is two"

tinydict = {'name': 'john','code':6734, 'dept': 'sales'}


print (dict['one'])       # Prints value for 'one' key
print (dict[2])           # Prints value for 2 key
print (tinydict)          # Prints complete dictionary
print (tinydict.keys())   # Prints all the keys
print (tinydict.values()) # Prints all the values
```

### 数据类型转换

|函数	|描述|
|----------------------|--------------------------|
|int(x [,base])|转换x为整数。x是字符串则 base 为指定的基数|
|float(x)|转换x为一个浮点数|
|complex(real [,imag])|创建一个复数|
|str(x)|转换对象x为字符串表示|
|repr(x)|转换对象x为表达式字符串|
|eval(str)|计算一个字符串，并返回一个对象|
|tuple(s)|转换s为一个元组|
|list(s)|转换s为一个列表|
|set(s)|转换s为一个集合|
|dict(d)|创建一个字典。 d必须是(键，值)元组序列|
|frozenset(s)|转换s为冷冻集|
|chr(x)|将一个字符转换为整数|
|unichr(x)|Unicode字符转换为整数|
|ord(x)|单个字符其转换为整数值|
|hex(x)|十六进制字符串转换为整数|
|oct(x)|转换整数成为八进制字符串|

## Python算术运算符

|操作符	|描述（变量 a=10 和 变量b=20）	|示例|
|-----|-------------------------|--------------------|
|+ 	|将操作符的两侧数值增加(加法运算)|a + b = 30|
|- 	| 左操作数减去右操作数|a – b = -10|
|*|操作符两侧数据相乘(乘法运算)|a * b = 200|
|/|操作符左侧操作数除以右操作数|b / a = 2|
|%|左操作数除以右侧操作数的余数|b % a = 0|
| ** |执行指数(幂)计算 | a ** b 就是10 的20 次幂|
|//| 地板除 - 除法不管操作数为何种数值类型，总是会舍去小数部分，返回数字序列中比真正的商小的最接近的数字 |9//2 = 4 以及 9.0//2.0 = 4.0|

## Python3比较运算符

|操作符	|描述（变量 a=10 和 变量b=20）	|示例|
|-----|-------------------------|--------------------|
|==|如果两个操作数的值相等，则条件变为真|(a == b) 不为 true.|
|!=|如果两个操作数的值不相等，则条件变为真||
|<>|如果两个操作数的值不相等，则条件变为真|(a <> b) 为 true. 这类似于 != 运算符|
|>|如果左操作数的值大于右操作数的值，则条件为真|(a > b) 不为true.|
|<|如果左操作数的值小于右操作数的值，则条件为真。|(a < b) 为 true.|
|>=|如果左操作数的值大于或等于右操作数的值，则条件为真|(a >= b) 不为 true.|
|<=|如果左操作数的值小于或等于右操作数的值，则条件为真|(a <= b) 为 true.|


## Python3赋值运算符

|操作符	|描述（变量 a=10 和 变量b=20）	|示例|
|-----|-------------------------|--------------------|
|=|将操作符的右侧操作数赋值给左侧的操作数|c = a + b 就是将 a + b 的值赋给 c|
|+=|它将右操作数和左操作数相加并分配结果给左操作数|c += a 相当于 c = c + a|
|-=|左操作数减去右操作数，并把结果赋给左操作数|c -= a 相当于 c = c - a|
|*=|左操作数乘以右操作数并分配结果给左操作数|c *= a 相当于 c = c * a|
|/=|把左操作数除以右操作数，并把结果赋给左操作数|c /= a 相当于 c = c / a；c /= a相当于 c = c / a|
|%=|两个操作数取模，并把结果赋给左操作数|c %= a 相当于 c = c % a
|**=|执行运算符指数(幂)计算并分配值给左操作数|c **= a 相当于c = c ** a
|//=|这对操作符两侧的操作数进行地板除并赋值给左操作数|c //= a 相当于 c = c // a|


## Python3位运算符

|操作符	|描 	|示例|
|-----|-------------------------|--------------------|
|&|二进制和|(a & b) (二进制表示为 0000 1100)|
| &brvbar;|二进制或|(a &brvbar; b) = 61 (二进制表示为 0011 1101)|
|^|二进制异或运算|(a ^ b) = 49 (二进制表示为 0011 0001)|
|~|二进制的补|(~a ) = -61 (二进制表示为1100 0011以2的补码形式，由一个带符号二进制数)|
|<<|二进制左移|a << = 240 (二进制表示为 1111 0000)|
|>>|二进制右移|a >> = 15 (二进制表示为 0000 1111)|


## Python3逻辑运算符

|操作符	|描述（变量a=True和变量b=False）|	示例|
|-----|-------------------------|--------------------|
| and |如果两个操作数为真，则条件为true|(a and b) 结果 False.|
| or |如果两个操作数为非零，条件变为true|(a or b) 结果 True.|
| not |用来扭转操作数的逻辑状态|not(a 且b) 结果 True.|

## Python3成员运算符

|操作符	|描述）|	示例|
|-----|-------------------------|--------------------|
|in	|如果在指定的顺序找到则计算结果为true，否则变量结果值为false| x in y，这里在一个1的结果，如果 x 是序列 y 的成员|
|not in|如果不能在指定的顺序找到则计算结果为true，否则变量为 false|x not in y，这里在一个1的结果，如果 x 不是序列 y 成员|

## Python3标识运算符

|操作符	|描述）|	示例|
|-----|-------------------------|--------------------|
|is|如果操作符两侧的变量是相同的对象计算结果为true，否则返回 false|x is y,如果 id(x) 等于 id(y) ，则结果为1|
|is not	|如果操作符两侧的变量不是相同的对象计算结果为true，否则返回 false	|x is y,如果 id(x) 不等于 id(y) ，则结果为1|

## Python3运算符优先级

|操作符|	描述(从高到低）|
|-----|-------------------------|
| ** | 幂运算
|~ + -|补，一元加号和减号(方法名的最后两个 +@ 和 -@)|
|* / % //|乘，除，模运算和地板除|
|+ -|加法和减法|
|>> <<|左，右按位移动|
|&|位元“与”|
|^ &brvbar;|按位'异或'和常规 '或'|
|<= < > >=|比较运算符|
|<> == !=|操作符相等比较|
|= %= /= //= -= += *= **=|赋值运算符|
|is is not|标识操作符|
|in not in|	成员操作符|
|not or and	|逻辑运算符|

## if...elseif...else语句

```py
if expression:
   statement(s)
else:
   statement(s)
```

```py
# 多层嵌套if
if expression1:
   statement(s)
   if expression2:
      statement(s)
   elif expression3:
      statement(s)
   else
      statement(s)
elif expression4:
   statement(s)
else:
   statement(s)
```

```py
#!/usr/bin/python3

amount=int(input("Enter amount: "))

if amount<1000:
    discount=amount*0.05
    print ("Discount",discount)
else:
    discount=amount*0.10
    print ("Discount",discount)
    
print ("Net payable:",amount-discount)
```

## while循环语句

```py
while expression:
   statement(s)
```

```py
#!/usr/bin/python3

count = 0
while (count < 9):
   print ('The count is:', count)
   count = count + 1

print ("Good bye!")
```

## for循环语句

```py
for iterating_var in sequence:
   statements(s)
```

```py
for letter in 'Python':     # traversal of a string sequence
   print ('Current Letter :', letter)
print()

fruits = ['banana', 'apple',  'mango']

for fruit in fruits:        # traversal of List sequence
   print ('Current fruit :', fruit)

# 通过序列索引进行遍历
for index in range(len(fruits)):
   print ('Current fruit :', fruits[index])

print ("Good bye!")
```

### for循环使用else语句

如果else语句和for循环语句一起使用，else块只在 for 循环正常终止时执行(而不是遇到break语句)。

如果else语句用在 while循环中，当条件变为 False 时，则执行else语句。

```py
#!/usr/bin/python3

numbers=[11,33,55,39,55,75,37,21,23,41,13]

for num in numbers:
    if num%2==0:
        print ('the list contains an even number')
        break
else:
    print ('the list doesnot contain even number')
```

## break语句

用于提前终止当前循环。 抛弃循环后，重新开始执行下一个语句。break语句可以在 while 和for 这两个循环使用。

如果您使用嵌套循环，break语句停止内部循环的执行，并开始执行块之后下一行代码段。

## continue语句

返回控制到当前循环的开始。遇到 continue 时，循环不会再执行当前迭代剩余的语句，而是开始下一次迭代。continue语句可以在while和for循环中使用。

## pass语句

它在语法上是用来作必须声明，但又不希望执行任何命令或代码。pass语句是个空操作;在执行时没有任何反应。

```py
#!/usr/bin/python3

for letter in 'Python': 
   if letter == 'h':
      pass
      print ('This is pass block')
   print ('Current Letter :', letter)

print ("Good bye!")
```

## Python3函数

* 函数模块使用 def 关键字开头，后跟函数名以及括号( ( ) ).
* 任何输入参数或参数都应该放在这些括号内。 还可以定义这些括号内的参数。
* 函数的第一个语句可以是一个可选的声明 - 函数或文档说明的字符串。
* 每个函数中的代码块使用冒号(:)开始和缩进。
* 语句返回[expression]存在于函数中，一个表达式可选地传递给调用者。不带参数的return语句返回None。

```py
def functionname( parameters ):
   "function_docstring"
   function_suite
   return [expression]
```

* 可变长度参数

```py
def functionname([formal_args,] *var_args_tuple ):
   "function_docstring"
   function_suite
   return [expression]
```

星号(*)放在持有的所有非关键字变量参数值的变量名之前。如果函数调用期间没有指定任何其他参数此元组是空的。

```py
#!/usr/bin/python3

# Function definition is here
def printinfo( arg1, *vartuple ):
   "This prints a variable passed arguments"
   print ("Output is: ")
   print (arg1)
   for var in vartuple:
      print (var)
   return

# Now you can call printinfo function
printinfo( 10 )
printinfo( 70, 60, 50 )
```

### 匿名函数

* lambda形式可以使用任何数量的参数，但在表现形式上只返回一个值。 它们不能包含命令或多个表达式。
* 匿名函数不能直接调用打印，因为lambda需要表达式。
* lambda函数都有自己的命名空间，并且不能访问除了在其参数列表和在全局命名空间中的其他变量。
* 尽管似乎 lambda 是一个函数的单行版，它们不等同于C或C++的内联声明，它的目的是调用出于性能考虑，在传递函数由堆栈分配。

```py
lambda [arg1 [,arg2,.....argn]]:expression
```

```py
#!/usr/bin/python3

# Function definition is here
sum = lambda arg1, arg2: arg1 + arg2

# Now you can call sum as a function
print ("Value of total : ", sum( 10, 20 ))
print ("Value of total : ", sum( 20, 20 ))
```

## Python3模块

一个模块是 Python 代码的文件。 一个模块可以定义函数，类和变量。一个模块也可以包括可运行的代码。

### import 语句

```py
import module1[, module2[,... moduleN]
```

### from...import 语句

```py
# 从一个模块中导入特定的属性到当前的命名空间
from modname import name1[, name2[, ... nameN]]

# 将一个模块的所有名称导入到当前的命名空间
from modname import *
```

### 执行模块作为脚本

在一个模块，模块名(做为一个字符串)可以作为全局变量__name__的值。该模块中的代码会被执行，就好像导入它一样，但设定为__main__的__name__。这意味着，通过在模块的末尾添加以下代码：

```py
#!/usr/bin/python3

# Fibonacci numbers module

def fib(n): # return Fibonacci series up to n
    result = []
    a, b = 0, 1
    while b < n:
        result.append(b)
        a, b = b, a+b
    return result
if __name__ == "__main__":
    f=fib(100)
    print(f)
```

#### PYTHONPATH 变量

下面是 Windows系统中一个典型的 PYTHONPATH ：
```cmd
set PYTHONPATH=c:\python34\lib;
```

这里是 UNIX 系统的典型 PYTHONPATH ：
```sh
set PYTHONPATH=/usr/local/lib/python
```

#### dir( ) 函数

使用 dir()内置函数返回一个包含由模块定义的名称字符串的排序列表。
该列表包含一个模块中定义的所有的模块，变量和函数的名称。

```py
#!/usr/bin/python3

# Import built-in module math
import math

content = dir(math)

print (content)
```

#### globals() 和 locals() 函数

globals() 和 locals()函数可用于在全局和局部名字空间返回名称，取决于从哪里调用它们。
如果 locals() 从函数中调用， 它会返回所有可以从函数访问的名字。

如果 globals() 可以在一个函数中调用，它将返回所有可以在全局范围内，可从函数访问的名字。
两种这些函数的返回类型是字典。因此，名称可以使用 keys() 函数来提取。

#### reload() 函数

当模块被导入到一个脚本，在模块的顶层部的代码只执行一次。
因此，如果你希望模块重新执行的顶层代码， 可以使用 reload()函数。在reload()函数会再次导入先前导入模块。

```py
reload(module_name)
```

### Python包

包是一个分层文件目录结构，定义由模块和子包和子子包等一个单一的Python应用环境。

考虑在 Phone 目录下找到的文件Pots.py。此文件的源代码如下面的行

```py
#!/usr/bin/python3

def Pots():
   print ("I'm Pots Phone")  
```

类似的方式，这里有不同功能的两个相同名称文件如下

* Phone/Isdn.py有一个函数 Isdn()

* Phone/G3.py 有一个函数 G3()

现在，在 Phone 目录中创建一个文件__init__.py

```txt
Phone/__init__.py
```

为了让所有的功能可用，当导入Phone，需要把明确

```py
#  import 语句在 __init__.py 中如下
from Pots import Pots
from Isdn import Isdn
from G3 import G3
```

在添加这些行到 __init__.py 后，当导入 Phone 包所有的这些类可用。

```py
#!/usr/bin/python3

# Now import your Phone Package.
import Phone

Phone.Pots()
Phone.Isdn()
Phone.G3()
```

## Python3断言

当它遇到一个断言语句，Python评估计算之后的表达式，希望是 true 值。如果表达式为 false，Python 触发 AssertionError 异常。

```py
assert Expression[, Arguments]
```

```py
#!/usr/bin/python3
def KelvinToFahrenheit(Temperature):
    assert (Temperature >= 0),"Colder than absolute zero!"
    return ((Temperature-273)*1.8)+32

print (KelvinToFahrenheit(273))
print (int(KelvinToFahrenheit(505.78)))
print (KelvinToFahrenheit(-5))
```

## Python3异常处理

|标准异常名称|描述|
|-------------|-------------------------------------------|
|Exception|所有异常的基类|
|StopIteration|当一个迭代器的 next()方法不指向任何对象时引发|
|SystemExit|由 sys.exit()函数引发|
|StandardError|除了StopIteration异常和SystemExit，所有内置异常的基类|
|ArithmeticError|数值计算所发生的所有错误的基类|
|OverflowError|当数字类型计算超过最高限额引发|
|FloatingPointError|当一个浮点运算失败时触发|
|ZeroDivisonError|当除运算或模零在所有数值类型运算时引发|
|AssertionError|断言语句失败的情况下引发|
|AttributeError|属性引用或赋值失败的情况下引发|
|EOFError|当从 raw_input() 与 input() 函数输入，到达文件末尾时触发|
|ImportError|当一个 import 语句失败时触发|
|KeyboardInterrupt|当用户中断程序执行，通常是通过按 Ctrl+c 引发|
|LookupError|所有查找错误基类|
|IndexError|当在一个序列中没有找到一个索引时引发|
| KeyError|当指定的键没有在字典中找到引发|
|NameError|当在局部或全局命名空间中找不到的标识引发|
|UnboundLocalError|试图访问在函数或方法的局部变量时引发，但没有值分配给它。|
|EnvironmentError|Python环境之外发生的所有异常的基类。|
|IOError|当一个输入/输出操作失败，如打印语句或 open()函数试图打开不存在的文件时引发操作系统相关的错误时引发|
|SyntaxError|当在Python语法错误引发；|
|IndentationError|没有正确指定缩进引发。|
|SystemError|当解释器发现一个内部问题，但遇到此错误时，Python解释器不退出引发|
|SystemExit	|当Python解释器不使用sys.exit()函数引发。如果代码没有被处理，解释器会退出。当操作或函数在指定数据类型无效时引发|
|ValueError	|在内置函数对于数据类型，参数的有效类型时引发，但是参数指定了无效值|
|RuntimeError|当生成的错误不属于任何类别时引发|
|NotImplementedError|当要在继承的类来实现，抽象方法实际上没有实现时引发此异常|

### 处理异常

```py
try:
   You do your operations here
   ......................
except ExceptionI:
   If there is ExceptionI, then execute this block.
except ExceptionII:
   If there is ExceptionII, then execute this block.
   ......................
else:
   If there is no exception then execute this block.
```

实例

```py
#!/usr/bin/python3

try:
   fh = open("testfile", "r")
   fh.write("This is my test file for exception handling!!")
except IOError:
   print ("Error: can\'t find file or read data")
else:
   print ("Written content in the file successfully")
```

### except子句与多个异常

```py
try:
   You do your operations here
   ......................
except(Exception1[, Exception2[,...ExceptionN]]]):
   If there is any exception from the given exception list, 
   then execute this block.
   ......................
else:
   If there is no exception then execute this block.
```

### try-finally子句

```py
try:
   You do your operations here
   ......................
   Due to any exception, this may be skipped.
finally:
   This would always be executed.
   ......................
```

实例

```py
#!/usr/bin/python3

try:
   fh = open("testfile", "w")
   try:
      fh.write("This is my test file for exception handling!!")
   finally:
      print ("Going to close the file")
      fh.close()
except IOError:
   print ("Error: can\'t find file or read data")
```

### 异常的参数

如果写代码来处理一个异常，可以使用一个变量按照异常的名称在 except 语句中。 如果要捕捉多个异常，可以使用一个变量后跟一个异常的元组。

该变量接收大多含有异常的原因各种值。变量可以在一个元组的形式以接收一个或多个值。这个元组通常包含错误字符串，错误编号，以及错误位置。

```py
try:
   You do your operations here
   ......................
except ExceptionType as Argument:
   You can print value of Argument here...
```

实例

```py
#!/usr/bin/python3

# Define a function here.
def temp_convert(var):
   try:
      return int(var)
   except ValueError, as Argument:
      print ("The argument does not contain numbers\n", Argument)

# Call above function here.
temp_convert("xyz")
```

### 引发异常

```py
raise [Exception [, args [, traceback]]]
```

这里，Exception 是异常的类型(例如，NameError)argument 为异常的参数值。该参数是可选的;如果没有提供，异常的参数是None。
最后一个参数 traceback，也可选的(并在实践中很少使用)，并且如果存在的话，是用于异常的回溯对象。

```py
def functionName( level ):
    if level <1:
        raise Exception(level)
        # The code below to this would not be executed
        # if we raise the exception
    return level
```

实例

```py
#!/usr/bin/python3
def functionName( level ):
    if level <1:
        raise Exception(level)
        # The code below to this would not be executed
        # if we raise the exception
    return level

try:
    l=functionName(-10)
    print ("level=",l)
except Exception as e:
    print ("error in level argument",e.args[0])
```

### 用户定义的异常

在try块，用户定义的异常将引发，并夹在 except 块中。 变量e是用来创建网络错误 Networkerror 类的实例。

```py
class Networkerror(RuntimeError):
   def __init__(self, arg):
      self.args = arg
```

所以上面的类定义后，可以引发异常如下

```py
try:
   raise Networkerror("Bad hostname")
except Networkerror,e:
   print e.args
```
