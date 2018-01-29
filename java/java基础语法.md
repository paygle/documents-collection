# Java 基础语法

* [Java 包](#golang)
* [枚举 Enumeration](#enumeration)
* [位集合 BitSet](#bitset)
* [向量 Vector ](#vector)
* [栈 Stack ](#stack)
* [Map接口来获取键/值的存储功能](#map)
* [哈希表 Hashtable](#hashtable)
* [属性 Properties](#property)
* [泛型](#genericity)
* [序列化](#serialize)
* [Java关键字](#keywords)
* [标识符](#identifier)
* [修饰符](#modifier)
* [变量](#variable)
* [对象和类](#objclass)
* [接口 interface](#interface)
* [接口的实现 implements](#implements)
* [接口的继承 extends](#extends)
* [基本数据类型](#basetype)
* [变量类型](#vartype)
* [访问控制修饰符](#ctrlmodifier)
* [static 修饰符](#static)
* [final 修饰符](#final)
* [abstract 修饰符](#abstract)
* [synchronized 修饰符](#synchronized)
* [transient 修饰符](#transient)
* [volatile 修饰符](#volatile)
* [java运算符](#operator)
* [while/do...while/for 循环](#dowhilefor)
* [分支语句结构](#ifswitch)
* [类型包装器](#decorator)
* [数组](#array)
* [Java 异常处理](#exception)
* [throws/throw 关键字](#throws)
* [声明自定义异常](#customexc)
* [Socket 编程](#socket)
* [Java 多线程编程](#multithread)
* [Java 文档注释](#doccomment)
* [Java8 新增特性](#java8new)


* 一个Java程序可以认为是一系列对象的集合，而这些对象通过调用彼此的方法来协同工作。下面简要介绍下类、对象、方法和实例变量的概念。

* 对象：对象是类的一个实例，有状态和行为。例如，一条狗是一个对象，它的状态有：颜色、名字、品种；行为有：摇尾巴、叫、吃等。

* 类：类是一个模板，它描述一类对象的行为和状态。

* 方法：方法就是行为，一个类可以有很多方法。逻辑运算、数据修改以及所有动作都是在方法中完成的。

* 实例变量：每个对象都有独特的实例变量，对象的状态由这些实例变量的值决定。

> 第一个Java程序

```java
public class HelloWorld {
    /* 第一个Java程序
     * 它将打印字符串 Hello World
     */
    public static void main(String []args) {
        System.out.println("Hello World"); // 打印 Hello World
    }
}
```

## 编写Java程序时，应注意以下几点：

* 大小写敏感：__Java是大小写敏感的__，这就意味着标识符Hello与hello是不同的。

* 类名：对于所有的类来说，类名的首字母应该大写。如果类名由若干单词组成，那么每个单词的首字母应该大写，例如 MyFirstJavaClass 。

* 方法名：__所有的方法名都应该以小写字母开头__。如果方法名含有若干单词，则后面的每个单词首字母大写。

* 源文件名：__源文件名必须和类名相同__。当保存文件的时候，你应该使用类名作为文件名保存（切记Java是大小写敏感的），文件名的后缀为.java。（如果文件名和类名不相同则会导致编译错误）。

* 主方法入口：所有的Java 程序由以下方法开始执行。

```java
public static void main(String []args){}
```

## <a name="golang"></a>Java 包(package)

1. 把功能相似或相关的类或接口组织在同一个包中，方便类的查找和使用。

2. 如同文件夹一样，包也采用了树形目录的存储方式。同一个包中的类名字是不同的，不同的包中的类的名字是可以相同的，当同时调用两个不同包中相同类名的类时，应该加上包名加以区别。因此，包可以避免名字冲突。

3. 包也限定了访问权限，拥有包访问权限的类才能访问某个包中的类。


Java 使用包（package）这种机制是为了防止命名冲突，访问控制，提供搜索和定位类（class）、接口、枚举（enumerations）和注释（annotation）等。

```java
// package pkg1[．pkg2[．pkg3…]];

package net.java.util

public class Something{
   //...
}
```

## Java 数据结构

### <a name="enumeration"></a>枚举（Enumeration）

虽然它本身不属于数据结构,但它在其他数据结构的范畴里应用很广。 枚举（The Enumeration）接口定义了一种从数据结构中取回连续元素的方式。

```java
import java.util.Vector;
import java.util.Enumeration;
 
public class EnumerationTester {
 
   public static void main(String args[]) {
      Enumeration<String> days;
      Vector<String> dayNames = new Vector<String>();
      dayNames.add("Sunday");
      dayNames.add("Monday");
      dayNames.add("Tuesday");
      dayNames.add("Wednesday");
      dayNames.add("Thursday");
      dayNames.add("Friday");
      dayNames.add("Saturday");
      days = dayNames.elements();
      while (days.hasMoreElements()){
         System.out.println(days.nextElement()); 
      }
   }
}
```

### <a name="bitset"></a>位集合（BitSet）

Bitset类创建一种特殊类型的数组来保存位值。BitSet中数组大小会随需要增加。这和位向量（vector of bits）比较类似。

```java
import java.util.BitSet;
 
public class BitSetDemo {
 
  public static void main(String args[]) {
     BitSet bits1 = new BitSet(16);
     BitSet bits2 = new BitSet(16);
      
     // set some bits
     for(int i=0; i<16; i++) {
        if((i%2) == 0) bits1.set(i);
        if((i%5) != 0) bits2.set(i);
     }
     System.out.println("Initial pattern in bits1: ");
     System.out.println(bits1);
     System.out.println("\nInitial pattern in bits2: ");
     System.out.println(bits2);
 
     // AND bits
     bits2.and(bits1);
     System.out.println("\nbits2 AND bits1: ");
     System.out.println(bits2);
 
     // OR bits
     bits2.or(bits1);
     System.out.println("\nbits2 OR bits1: ");
     System.out.println(bits2);
 
     // XOR bits
     bits2.xor(bits1);
     System.out.println("\nbits2 XOR bits1: ");
     System.out.println(bits2);
  }
}
```

### <a name="vector"></a>向量（Vector）

向量（Vector）类和传统数组非常相似，但是Vector的大小能根据需要动态的变化。

和数组一样，Vector对象的元素也能通过索引访问。

使用Vector类最主要的好处就是在创建对象的时候不必给对象指定大小，它的大小会根据需要动态的变化。

```java
import java.util.*;

public class VectorDemo {

   public static void main(String args[]) {
      // initial size is 3, increment is 2
      Vector v = new Vector(3, 2);
      System.out.println("Initial size: " + v.size());
      System.out.println("Initial capacity: " +
      v.capacity());
      v.addElement(new Integer(1));
      v.addElement(new Integer(2));
      v.addElement(new Integer(3));
      v.addElement(new Integer(4));
      System.out.println("Capacity after four additions: " +
          v.capacity());

      v.addElement(new Double(5.45));
      System.out.println("Current capacity: " +
      v.capacity());
      v.addElement(new Double(6.08));
      v.addElement(new Integer(7));
      System.out.println("Current capacity: " +
      v.capacity());
      v.addElement(new Float(9.4));
      v.addElement(new Integer(10));
      System.out.println("Current capacity: " +
      v.capacity());
      v.addElement(new Integer(11));
      v.addElement(new Integer(12));
      System.out.println("First element: " +
         (Integer)v.firstElement());
      System.out.println("Last element: " +
         (Integer)v.lastElement());
      if(v.contains(new Integer(3)))
         System.out.println("Vector contains 3.");
      // enumerate the elements in the vector.
      Enumeration vEnum = v.elements();
      System.out.println("\nElements in vector:");
      while(vEnum.hasMoreElements())
         System.out.print(vEnum.nextElement() + " ");
      System.out.println();
   }
}
```

### <a name="stack"></a>栈（Stack）

实现了一个后进先出（LIFO）的数据结构。

你可以把栈理解为对象的垂直分布的栈，当你添加一个新元素时，就将新元素放在其他元素的顶部。

当你从栈中取元素的时候，就从栈顶取一个元素。换句话说，最后进栈的元素最先被取出。

```java
import java.util.*;
 
public class StackDemo {
 
    static void showpush(Stack<Integer> st, int a) {
        st.push(new Integer(a));
        System.out.println("push(" + a + ")");
        System.out.println("stack: " + st);
    }
 
    static void showpop(Stack<Integer> st) {
        System.out.print("pop -> ");
        Integer a = (Integer) st.pop();
        System.out.println(a);
        System.out.println("stack: " + st);
    }
 
    public static void main(String args[]) {
        Stack<Integer> st = new Stack<Integer>();
        System.out.println("stack: " + st);
        showpush(st, 42);
        showpush(st, 66);
        showpush(st, 99);
        showpop(st);
        showpop(st);
        showpop(st);
        try {
            showpop(st);
        } catch (EmptyStackException e) {
            System.out.println("empty stack");
        }
    }
}
```

### <a name="map"></a>Map接口来获取键/值的存储功能。

* 给定一个键和一个值，你可以将该值存储在一个Map对象. 之后，你可以通过键来访问对应的值。
* 当访问的值不存在的时候，方法就会抛出一个NoSuchElementException异常.
* 当对象的类型和Map里元素类型不兼容的时候，就会抛出一个 ClassCastException异常。
* 当在不允许使用Null对象的Map中使用Null对象，会抛出一个NullPointerException 异常。
* 当尝试修改一个只读的Map时，会抛出一个UnsupportedOperationException异常。

```java
import java.util.*;

public class CollectionsDemo {

   public static void main(String[] args) {
      Map m1 = new HashMap(); 
      m1.put("Zara", "8");
      m1.put("Mahnaz", "31");
      m1.put("Ayan", "12");
      m1.put("Daisy", "14");
      System.out.println();
      System.out.println(" Map Elements");
      System.out.print("\t" + m1);
   }
}
```

### <a name="hashtable"></a>哈希表（Hashtable）

Hashtable类提供了一种在用户定义键结构的基础上来组织数据的手段。

例如，在地址列表的哈希表中，你可以根据邮政编码作为键来存储和排序数据，而不是通过人名。

哈希表键的具体含义完全取决于哈希表的使用情景和它包含的数据。

```java
import java.util.*;

public class HashTableDemo {

   public static void main(String args[]) {
      // Create a hash map
      Hashtable balance = new Hashtable();
      Enumeration names;
      String str;
      double bal;

      balance.put("Zara", new Double(3434.34));
      balance.put("Mahnaz", new Double(123.22));
      balance.put("Ayan", new Double(1378.00));
      balance.put("Daisy", new Double(99.22));
      balance.put("Qadir", new Double(-19.08));

      // Show all balances in hash table.
      names = balance.keys();
      while(names.hasMoreElements()) {
         str = (String) names.nextElement();
         System.out.println(str + ": " +
         balance.get(str));
      }
      System.out.println();
      // Deposit 1,000 into Zara's account
      bal = ((Double)balance.get("Zara")).doubleValue();
      balance.put("Zara", new Double(bal+1000));
      System.out.println("Zara's new balance: " +
      balance.get("Zara"));
   }
}
```

### <a name="property"></a>属性（Properties）

Properties 继承于 Hashtable.Properties 类表示了一个持久的属性集.属性列表中每个键及其对应值都是一个字符串。

Properties 类被许多Java类使用。例如，在获取环境变量时它就作为System.getProperties()方法的返回值。

```java
import java.util.*;

public class PropDemo {

   public static void main(String args[]) {
      Properties capitals = new Properties();
      Set states;
      String str;
      
      capitals.put("Illinois", "Springfield");
      capitals.put("Missouri", "Jefferson City");
      capitals.put("Washington", "Olympia");
      capitals.put("California", "Sacramento");
      capitals.put("Indiana", "Indianapolis");

      // Show all states and capitals in hashtable.
      states = capitals.keySet(); // get set-view of keys
      Iterator itr = states.iterator();
      while(itr.hasNext()) {
         str = (String) itr.next();
         System.out.println("The capital of " +
            str + " is " + capitals.getProperty(str) + ".");
      }
      System.out.println();

      // look for state not in list -- specify default
      str = capitals.getProperty("Florida", "Not Found");
      System.out.println("The capital of Florida is "
          + str + ".");
   }
}
```

## <a name="genericity"></a>Java 泛型

泛型提供了编译时类型安全检测机制，该机制允许程序员在编译时检测到非法的类型。

泛型的本质是参数化类型，也就是说所操作的数据类型被指定为一个参数。

> 定义泛型方法的规则：

* 所有泛型方法声明都有一个类型参数声明部分（由尖括号分隔），该类型参数声明部分在方法返回类型之前（在下面例子中的<E>）。
* 每一个类型参数声明部分包含一个或多个类型参数，参数间用逗号隔开。一个泛型参数，也被称为一个类型变量，是用于指定一个泛型类型名称的标识符。
* 类型参数能被用来声明返回值类型，并且能作为泛型方法得到的实际参数类型的占位符。
* 泛型方法体的声明和其他方法一样。注意类型参数只能代表引用型类型，不能是原始类型（像int,double,char的等）。

```java
//可能有时候，你会想限制那些被允许传递到一个类型参数的类型种类范围。
//例如，一个操作数字的方法可能只希望接受Number或者Number子类的实例。这就是有界类型参数的目的。

//要声明一个有界的类型参数，首先列出类型参数的名称，后跟extends关键字，最后紧跟它的上界。

public class MaximumTest
{
   // 比较三个值并返回最大值
   public static <T extends Comparable<T>> T maximum(T x, T y, T z)
   {                     
      T max = x; // 假设x是初始最大值
      if ( y.compareTo( max ) > 0 ){
         max = y; //y 更大
      }
      if ( z.compareTo( max ) > 0 ){
         max = z; // 现在 z 更大           
      }
      return max; // 返回最大对象
   }
   public static void main( String args[] )
   {
      System.out.printf( "%d, %d 和 %d 中最大的数为 %d\n\n",
                   3, 4, 5, maximum( 3, 4, 5 ) );
 
      System.out.printf( "%.1f, %.1f 和 %.1f 中最大的数为 %.1f\n\n",
                   6.6, 8.8, 7.7, maximum( 6.6, 8.8, 7.7 ) );
 
      System.out.printf( "%s, %s 和 %s 中最大的数为 %s\n","pear",
         "apple", "orange", maximum( "pear", "apple", "orange" ) );
   }
}
```

```java
// 定义一个泛型类
public class Box<T> {
   
  private T t;
 
  public void add(T t) {
    this.t = t;
  }
 
  public T get() {
    return t;
  }
 
  public static void main(String[] args) {
    Box<Integer> integerBox = new Box<Integer>();
    Box<String> stringBox = new Box<String>();
 
    integerBox.add(new Integer(10));
    stringBox.add(new String("菜鸟教程"));
 
    System.out.printf("整型值为 :%d\n\n", integerBox.get());
    System.out.printf("字符串为 :%s\n", stringBox.get());
  }
}
```

## <a name="serialize"></a>Java 序列化

该机制中，一个对象可以被表示为一个字节序列，该字节序列包括该对象的数据、有关对象的类型的信息和存储在对象中数据的类型。

类 ObjectInputStream 和 ObjectOutputStream 是高层次的数据流，它们包含序列化和反序列化对象的方法。

```java
public class Employee implements java.io.Serializable
{
   public String name;
   public String address;
   public transient int SSN;
   public int number;
   public void mailCheck()
   {
      System.out.println("Mailing a check to " + name
                           + " " + address);
   }
}
```
* __请注意__，一个类的对象要想序列化成功，必须满足两个条件：

> 该类必须实现 __java.io.Serializable__ 对象。

> 该类的所有属性必须是可序列化的。如果有一个属性不是可序列化的，则该属性必须注明是短暂的。

> 如果你想知道一个 Java 标准类是否是可序列化的，请查看该类的文档。检验一个类的实例是否能序列化十分简单， 只需要查看该类有没有实现 java.io.Serializable接口。

```java
import java.io.*;
 
public class SerializeDemo
{
   public static void main(String [] args)
   {
      Employee e = new Employee();
      e.name = "Reyan Ali";
      e.address = "Phokka Kuan, Ambehta Peer";
      e.SSN = 11122333;
      e.number = 101;
      try
      {
         FileOutputStream fileOut =
         new FileOutputStream("/tmp/employee.ser");
         ObjectOutputStream out = new ObjectOutputStream(fileOut);
         out.writeObject(e);
         out.close();
         fileOut.close();
         System.out.printf("Serialized data is saved in /tmp/employee.ser");
      }catch(IOException i)
      {
          i.printStackTrace();
      }
   }
}
```

## <a name="keywords"></a>Java关键字

> 下面列出了Java保留字。这些保留字不能用于常量、变量、和任何标识符的名称。

|类别	|关键字	|说明|
----------|----------|----------------|
|访问控制	|private	|私有的|
|  |protected|	受保护的|
|  |public	|公共的|
|类、方法和变量修饰符	|abstract	|声明抽象|
|  |class	|类|
|  |extends	|扩允,继承|
|  |final	|最终值,不可改变的|
|  |implements	|实现（接口）|
|  |interface	|接口|
|  |native	|本地，原生方法（非Java实现）|
|  |new	|新,创建|
|  |static	|静态|
|  |strictfp	|严格,精准|
|  |synchronized	|线程,同步|
|  |transient	|短暂|
|  |volatile	|易失|
|程序控制语句|	break|	跳出循环|
|  |case	|定义一个值以供switch选择|
|  |continue	|继续|
|  |default	|默认|
|  |do	|运行|
|  |else	|否则|
|  |for	|循环|
|  |if	|如果|
|  |instanceof	|实例|
|  |return	|返回|
|  |switch	|根据值选择执行|
|  |while	|循环|
|错误处理	|assert	|断言表达式是否为真|
|  |catch	|捕捉异常|
|  |finally	|有没有异常都执行|
|  |throw	|抛出一个异常对象|
|  |throws	|声明一个异常可能被抛出|
|  |try	|捕获异常|
|包相关	|import	|引入|
|  |package	|包|
|基本类型	|boolean	|布尔型|
|  |byte	|字节型|
|  |char	|字符型|
|  |double	|双精度浮点|
|  |float	|单精度浮点|
|  |int	|整型|
|  |long	|长整型|
|  |short	|短整型|
|  |null	|空|
|变量引用	|super	|父类,超类|
|  |this	|本类|
|  |void	|无返回值|
|保留关键字	|goto	|是关键字，但不能使用|
|  |const	|是关键字，但不能使用|

### <a name="identifier"></a>Java标识符

Java所有的组成部分都需要名字。类名、变量名以及方法名都被称为标识符。

### 关于Java标识符，有以下几点需要注意：

* 所有的标识符都应该以字母（A-Z或者a-z）,美元符（$）、或者下划线（_）开始

* 首字符之后可以是字母（A-Z或者a-z）,美元符（$）、下划线（_）或数字的任何字符组合

* 关键字不能用作标识符

* 标识符是大小写敏感的

* 合法标识符举例：age、$salary、_value、__1_value

* 非法标识符举例：123abc、-salary

### <a name="modifier"></a>Java修饰符

> 像其他语言一样，Java可以使用修饰符来修饰类中方法和属性。主要有两类修饰符：

* 访问控制修饰符 : default, public , protected, private

* 非访问控制修饰符 : final, abstract, strictfp

### <a name="variable"></a>Java变量

> Java中主要有如下几种类型的变量

* 局部变量：在方法、构造方法、语句块中定义的变量。其声明和初始化在方法中实现，在方法结束后自动销毁

```java
public class  ClassName{
    public void printNumber（）{
        int a;   // 局部变量
    }
}
```

* 成员变量（非静态变量）: 定义在类中，方法体之外。变量在创建对象时实例化。成员变量可被类中的方法、构造方法以及特定类的语句块访问。

```java
public class  ClassName{
    int a;            // 成员变量
    public void printNumber（）{ }
}
```

* 类变量（静态变量）: 定义在类中，方法体之外，但必须要有 static 来声明变量类型。静态成员属于整个类，可通过对象名或类名来调用。

```java
public class  ClassName{
    static int a;        // 类变量
    public void printNumber（）{  }
}
```

## <a name="objclass"></a>Java 对象和类

* 相关概念：多态、继承、封装、抽象、类、对象、实例、方法、重载

> Employee.java 文件代码:

```java
import java.io.*;

public class Employee{
   String name;
   int age;
   String designation;
   double salary;
   // Employee 类的构造器
   public Employee(String name){
      this.name = name;
   }
   // 设置age的值
   public void empAge(int empAge){
      age =  empAge;
   }
   /* 设置designation的值*/
   public void empDesignation(String empDesig){
      designation = empDesig;
   }
   /* 设置salary的值*/
   public void empSalary(double empSalary){
      salary = empSalary;
   }
   /* 打印信息 */
   public void printEmployee(){
      System.out.println("名字:"+ name );
      System.out.println("年龄:" + age );
      System.out.println("职位:" + designation );
      System.out.println("薪水:" + salary);
   }
}
```

> EmployeeTest.java 文件代码：

```java
import java.io.*;
public class EmployeeTest{
 
   public static void main(String args[]){
      /* 使用构造器创建两个对象 */
      Employee empOne = new Employee("RUNOOB1");
      Employee empTwo = new Employee("RUNOOB2");
 
      // 调用这两个对象的成员方法
      empOne.empAge(26);
      empOne.empDesignation("高级程序员");
      empOne.empSalary(1000);
      empOne.printEmployee();
 
      empTwo.empAge(21);
      empTwo.empDesignation("菜鸟程序员");
      empTwo.empSalary(500);
      empTwo.printEmployee();
   }
}
```

* java因强制要求类名（唯一的public类）和文件名统一，因此在引用其它类时无需显式声明。在编译时，编译器会根据类名去寻找同名文件。

## <a name="interface"></a>Java 接口 interface

* 接口是隐式抽象的，当声明一个接口的时候，不必使用abstract关键字。
* 接口中每一个方法也是隐式抽象的，声明时同样不需要abstract关键字。
* 接口中的方法都是公有的。

```java
[可见度] interface 接口名称 [extends 其他的类名] {
        // 声明变量
        // 抽象方法
}
```

Interface关键字用来声明一个接口。下面是接口声明的一个简单例子。

```java
/* 文件名 : NameOfInterface.java */
import java.lang.*;
//引入包

public interface NameOfInterface
{
   //任何类型 final, static 字段
   //抽象方法
}
```

### <a name="implements"></a>接口的实现 implements

```java
// ...implements 接口名称[, 其他接口, 其他接口..., ...] ...

public class MammalInt implements Animal{
 
   public void eat(){
      System.out.println("Mammal eats");
   }
 
   public void travel(){
      System.out.println("Mammal travels");
   } 
 
   public int noOfLegs(){
      return 0;
   }
 
   public static void main(String args[]){
      MammalInt m = new MammalInt();
      m.eat();
      m.travel();
   }
}
```

* 重写接口中声明的方法时，需要注意以下规则：

```txt
类在实现接口的方法时，不能抛出强制性异常，只能在接口中，或者继承接口的抽象类中抛出该强制性异常。
类在重写方法时要保持一致的方法名，并且应该保持相同或者相兼容的返回值类型。
如果实现接口的类是抽象类，那么就没必要实现该接口的方法。
```

* 在实现接口的时候，也要注意一些规则：

```txt
一个类可以同时实现多个接口。
一个类只能继承一个类，但是能实现多个接口。
一个接口能继承另一个接口，这和类之间的继承比较相似。
```

### <a name="extends"></a>接口的继承 extends

```java
// public interface Hockey extends Sports, Event

public interface Sports
{
   public void setHomeTeam(String name);
   public void setVisitingTeam(String name);
}
 
// 文件名: Football.java
public interface Football extends Sports
{
   public void homeTeamScored(int points);
   public void visitingTeamScored(int points);
   public void endOfQuarter(int quarter);
}
 
// 文件名: Hockey.java
public interface Hockey extends Sports
{
   public void homeGoalScored();
   public void visitingGoalScored();
   public void endOfPeriod(int period);
   public void overtimePeriod(int ot);
}
```

## <a name="basetype"></a>Java 基本数据类型

* Java语言提供了八种基本类型。六种数字类型（四个整数型，两个浮点型），一种字符类型，还有一种布尔型。

* byte 类型 8 位

```txt
最小值是 -128（-2^7）；
最大值是 127（2^7-1）；
默认值是 0；
例子：byte a = 100，byte b = -50。
```

* short 类型 16 位

```txt
最小值是 -32768（-2^15）；
最大值是 32767（2^15 - 1）；
默认值是 0；
例子：short s = 1000，short r = -20000。
```

* int 类型 32 位, 一般默认类型

```txt
最小值是 -2,147,483,648（-2^31）；
最大值是 2,147,483,647（2^31 - 1）
默认值是 0；
例子：int a = 100000, int b = -200000。
```

* long 类型 64 位, 一般默认类型

```txt
最小值是 -9,223,372,036,854,775,808（-2^63）；
最大值是 9,223,372,036,854,775,807（2^63 -1）；
默认值是 0L；
例子： long a = 100000L，Long b = -200000L。
"L"理论上不分大小写，但是若写成"l"容易与数字"1"混淆，不容易分辩。所以最好大写
```

* long 类型 64 位, 一般默认类型

```txt
最小值是 -9,223,372,036,854,775,808（-2^63）；
最大值是 9,223,372,036,854,775,807（2^63 -1）；
默认值是 0L；
例子： long a = 100000L，Long b = -200000L。
"L"理论上不分大小写，但是若写成"l"容易与数字"1"混淆，不容易分辩。所以最好大写
```

* float 类型 32 位

```txt
float 在储存大型浮点数组的时候可节省内存空间；
默认值是 0.0f；
浮点数不能用来表示精确的值，如货币；
例子：float f1 = 234.5f。
```

* double 类型是双精度、64 位

```txt
浮点数的默认类型为double类型；
double类型同样不能表示精确的值，如货币；
默认值是 0.0d；
例子：double d1 = 123.4。
```

* boolean类型表示一位的信息；

```txt
只有两个取值：true 和 false；
这种类型只作为一种标志来记录 true/false 情况；
默认值是 false；
例子：boolean one = true。
```

* char类型 16 位 Unicode 字符；

```txt
最小值是 \u0000（即为0）；
最大值是 \uffff（即为65,535）；
char 数据类型可以储存任何字符；
例子：char letter = 'A';。
```

## <a name="vartype"></a>Java 变量类型

* 所有的变量在使用前必须声明。声明变量的基本格式如下：

```txt
type identifier [ = value][, identifier [= value] ...] ;
```

## Java 修饰符

### <a name="ctrlmodifier"></a>访问控制修饰符

* default (即缺省，什么也不写）: 在同一包内可见，不使用任何修饰符。使用对象：类、接口、变量、方法。

* private : 在同一类内可见。使用对象：变量、方法。 注意：不能修饰类（外部类）

* public : 对所有类可见。使用对象：类、接口、变量、方法

* protected : 对同一包内的类和所有子类可见。使用对象：变量、方法。 注意：不能修饰类（外部类）。

|修饰符|	当前类|	同一包内|	子孙类|	其他包|
|-----|------|---------|------|-------|
|public	|Y	|Y	|Y	|Y|
|protected	|Y	|Y	|Y	|N|
|default	|Y	|Y	|N	|N|
|private	|Y	|N	|N	|N|

### 非访问修饰符

#### <a name="static"></a>static 修饰符

* 静态变量：

static 关键字用来声明独立于对象的静态变量，无论一个类实例化多少对象，它的静态变量只有一份拷贝。 静态变量也被称为类变量。局部变量不能被声明为 static 变量。

* 静态方法：

static 关键字用来声明独立于对象的静态方法。静态方法不能使用类的非静态变量。静态方法从参数列表得到数据，然后计算这些数据。

```java
//对类变量和方法的访问可以直接使用以下方式访问
classname.variablename
classname.methodname
```

#### <a name="final"></a>final 修饰符

* final 变量：

final 变量能被显式地初始化并且只能初始化一次。被声明为 final 的对象的引用不能指向不同的对象。但是 final 对象里的数据可以被改变。也就是说 final 对象的引用不能改变，但是里面的值可以改变。

final 修饰符通常和 static 修饰符一起使用来创建类常量。

```java
public class Test{
  final int value = 10;
  // 下面是声明常量的实例
  public static final int BOXWIDTH = 6;
  static final String TITLE = "Manager";
 
  public void changeValue(){
     value = 12; //将输出一个错误
  }
}
```

* final 方法

类中的 final 方法可以被子类继承，但是不能被子类修改。
声明 final 方法的主要目的是防止该方法的内容被修改。

```java
public class Test{
    public final void changeName(){}
}
```

* final 类

final 类不能被继承，没有类能够继承 final 类的任何特性。

```java
public final class Test {
   // 类体
}
```

#### <a name="abstract"></a>abstract 修饰符

* 抽象类：

抽象类不能用来实例化对象，声明抽象类的唯一目的是为了将来对该类进行扩充。

一个类不能同时被 abstract 和 final 修饰。如果一个类包含抽象方法，那么该类一定要声明为抽象类，否则将出现编译错误。

抽象类可以包含抽象方法和非抽象方法。

```java
abstract class Caravan{
   private double price;
   private String model;
   private String year;
   public abstract void goFast(); //抽象方法
   public abstract void changeColor();
}
```

* 抽象方法

抽象方法是一种没有任何实现的方法，该方法的的具体实现由子类提供。

抽象方法不能被声明成 final 和 static。

任何继承抽象类的子类必须实现父类的所有抽象方法，除非该子类也是抽象类。

如果一个类包含若干个抽象方法，那么该类必须声明为抽象类。抽象类可以不包含抽象方法。

抽象方法的声明以分号结尾，例如：public abstract sample();。

```java
public abstract class SuperClass{
    abstract void m(); //抽象方法
}

class SubClass extends SuperClass{
     //实现抽象方法
      void m(){
          .........
      }
}
```

#### <a name="synchronized"></a>synchronized 修饰符

该关键字声明的方法同一时间只能被一个线程访问。synchronized 修饰符可以应用于四个访问修饰符。

```java
public synchronized void showDetails(){
.......
}
```

#### <a name="transient"></a>transient 修饰符

序列化的对象包含被 transient 修饰的实例变量时，java 虚拟机(JVM)跳过该特定的变量。

该修饰符包含在定义变量的语句中，用来预处理类和变量的数据类型。

```java
public transient int limit = 55;   // 不会持久化
public int b; // 持久化
```

#### <a name="volatile"></a>volatile 修饰符

volatile 修饰的成员变量在每次被线程访问时，都强制从共享内存中重新读取该成员变量的值。而且，当成员变量发生变化时，会强制线程将变化值回写到共享内存。这样在任何时刻，两个不同的线程总是看到某个成员变量的同一个值。

一个 volatile 对象引用可能是 null。

```java
public class MyRunnable implements Runnable
{
    private volatile boolean active;
    public void run()
    {
        active = true;
        while (active) // 第一行
        {
            // 代码
        }
    }
    public void stop()
    {
        active = false; // 第二行
    }
}
```

## <a name="operator"></a>Java 运算符

* 算术运算符

表格中的实例假设整数变量A的值为10，变量B的值为20：

|操作符	|描述	|例子|
|-------|-----------|-----------|
|+	|加法 - 相加运算符两侧的值	|A + B 等于 30|
|-	|减法 - 左操作数减去右操作数	|A – B 等于 -10|
|*	|乘法 - 相乘操作符两侧的值	|A * B等于200|
|/	|除法 - 左操作数除以右操作数	|B / A等于2|
|％	|取模 - 左操作数除以右操作数的余数	|B%A等于0|
|++	|自增: 操作数的值增加1	|B++ 或 ++B 等于 21|
|--	|自减: 操作数的值减少1	|B-- 或 --B 等于 19|

* 关系运算符

|操作符	|描述	|例子|
|-------|-----------|-----------|
|==	|检查如果两个操作数的值是否相等，如果相等则条件为真。	|（A == B）为假(非真)。|
|!=	|检查如果两个操作数的值是否相等，如果值不相等则条件为真。	|(A != B) 为真。|
|> 	|检查左操作数的值是否大于右操作数的值，如果是那么条件为真。	|（A> B）非真。|
|< 	|检查左操作数的值是否小于右操作数的值，如果是那么条件为真。	|（A <B）为真。|
|>=	|检查左操作数的值是否大于或等于右操作数的值，如果是那么条件为真。	|（A> = B）为假。|
|<=	|检查左操作数的值是否小于或等于右操作数的值，如果是那么条件为真。	|（A <= B）为真。|

* 位运算符

下表列出了位运算符的基本运算,假设整数变量A的值为60和变量B的值为13

|操作符	|描述	|例子|
|-------|-----------|-----------|
|＆	如果相对应位都是1，则结果为1，否则为0	|（A＆B），得到12，即0000 1100|
|	&brvbar; | 如果相对应位都是0，则结果为0，否则为1	|（A &brvbar; B）得到61，即 0011 1101|
| ^	|如果相对应位值相同，则结果为0，否则为1	|（A ^ B）得到49，即 0011 0001|
| 〜	|按位补运算符翻转操作数的每一位，即0变成1，1变成0。	|（〜A）得到-61，即1100 0011|
| << 	|按位左移运算符。左操作数按位左移右操作数指定的位数。|	A << 2得到240，即 1111 0000|
| >> 	|按位右移运算符。左操作数按位右移右操作数指定的位数。	|A >> 2得到15即 1111|
| >>> 	|按位右移补零操作符。左操作数的值按右操作数指定的位数右移，移动得到的空位以零填充。	|A>>>2得到15即0000 1111|

* 逻辑运算符

|操作符	|描述	|例子|
|-------|-----------|-----------|
| &&	|称为逻辑与运算符。当且仅当两个操作数都为真，条件才为真。	|（A && B）为假。|
|	&brvbar;&brvbar; |	称为逻辑或操作符。如果任何两个操作数任何一个为真，条件为真。	|（A &brvbar;&brvbar; B）为真。|
| ！	|称为逻辑非运算符。用来反转操作数的逻辑状态。如果条件为true，则逻辑非运算符将得到false。|	！（A && B）为真。|

+ 赋值运算符

|操作符	|描述	|例子|
|-------|-----------|-----------|
| =	|简单的赋值运算符，将右操作数的值赋给左侧操作数	| C = A + B将把A + B得到的值赋给C|
| +=	|加和赋值操作符，它把左操作数和右操作数相加赋值给左操作数	| C + = A等价于C = C + A|
| -=	|减和赋值操作符，它把左操作数和右操作数相减赋值给左操作数	| C - = A等价于C = C - A|
| *=	|乘和赋值操作符，它把左操作数和右操作数相乘赋值给左操作数	| C * = A等价于C = C * A|
| /=	|除和赋值操作符，它把左操作数和右操作数相除赋值给左操作数	| C / = A等价于C = C / A|
|（％）=	|取模和赋值操作符，它把左操作数和右操作数取模后赋值给左操作数	| C％= A等价于C = C％A|
| <<=	|左移位赋值运算符	| C << = 2等价于C = C << 2|
| >>=	|右移位赋值运算符	| C >> = 2等价于C = C >> 2|
| ＆=	|按位与赋值运算符	| C＆= 2等价于C = C＆2|
| ^=	|按位异或赋值操作符	| C ^ = 2等价于C = C ^ 2|
| &brvbar;=	|按位或赋值操作符	| C &brvbar;= 2 等价于C = C&brvbar;2 |

+ 条件运算符（?:）

* instanceof 运算符

该运算符用于操作对象实例，检查该对象是否是一个特定类型（类类型或接口类型）

```txt
( Object reference variable ) instanceof  (class/interface type)
```

## <a name="dowhilefor"></a>Java 循环结构

* while 循环

```java
while( 布尔表达式 ) {
  //循环内容
}
```

* do…while 循环

```java
do {
       //代码语句
}while(布尔表达式);
```

* for循环

```java
for(初始化; 布尔表达式; 更新) {
    //代码语句
}
```

* 主要用于数组的增强型 for 循环语法格式

```java
for(声明语句 : 表达式)
{
   //代码句子
}

//示例
public class Test {
   public static void main(String args[]){
      int [] numbers = {10, 20, 30, 40, 50};
 
      for(int x : numbers ){
         System.out.print( x );
         System.out.print(",");
      }
      System.out.print("\n");
      String [] names ={"James", "Larry", "Tom", "Lacy"};
      for( String name : names ) {
         System.out.print( name );
         System.out.print(",");
      }
   }
}
```

* break 关键字

主要用在循环语句或者 switch 语句中，用来跳出整个语句块。跳出最里层的循环，并且继续执行该循环下面的语句。

* continue 关键字

适用于任何循环控制结构中。作用是让程序立刻跳转到下一次循环的迭代。

在 for 循环中，continue 语句使程序立即跳转到更新语句。

在 while 或者 do…while 循环中，程序立即跳转到布尔表达式的判断语句。

## <a name="ifswitch"></a>Java 分支结构

* if...else...

```java
if(布尔表达式){
   //如果布尔表达式的值为true
}else{
   //如果布尔表达式的值为false
}
```

* switch 语句判断一个变量与一系列值中某个值是否相等，每个值称为一个分支。

```java
switch(expression){
    case value :
       //语句
       break; //可选
    case value :
       //语句
       break; //可选
    //你可以有任意数量的case语句
    default : //可选
       //语句
}
```

## <a name="decorator"></a>java 类型包装器

* 所有的包装类（Integer、Long、Byte、Double、Float、Short）都是抽象类 Number 的子类

* Math 类包含了用于执行基本数学运算的属性和方法，如初等指数、对数、平方根和三角函数。

* Character 类用于对单个字符进行操作。

* String 类有 11 种构造方法，这些方法提供不同的参数来初始化字符串，String 类是不可改变的，所以你一旦创建了 String 对象，那它的值就无法改变了

* StringBuffer 和 StringBuilder 类的对象能够被多次的修改，并且不产生新的未使用对象。然而在应用程序要求线程安全的情况下，则必须使用 StringBuffer 类。

* Scanner 类 来获取用户的输入

## <a name="array"></a>Java 数组

* 注意：建议使用 dataType[] arrayRefVar 的声明风格声明数组变量。 dataType arrayRefVar[] 风格是来自 C/C++ 语言 ，在Java中采用是为了让 C/C++ 程序员能够快速理解java语言。

```java
// dataType[] arrayRefVar = new dataType[arraySize];       // 首选的方法
// dataType[] arrayRefVar = {value0, value1, ..., valuek};
double[] myList;

//或

//dataType arrayRefVar[];  // 效果相同，但不是首选方法
double myList[];
```

## <a name="exception"></a>Java 异常处理

所有的异常类是从 java.lang.Exception 类继承的子类。

* 检查性异常：最具代表的检查性异常是用户错误或问题引起的异常，这是程序员无法预见的。例如要打开一个不存在文件时，一个异常就发生了，这些异常在编译时不能被简单地忽略。

* 运行时异常： 运行时异常是可能被程序员避免的异常。与检查性异常相反，运行时异常可以在编译时被忽略。

* 错误： 错误不是异常，而是脱离程序员控制的问题。错误在代码中通常被忽略。例如，当栈溢出时，一个错误就发生了，它们在编译也检查不到的。

> 多重捕获块

```java
try {
   // 程序代码
} catch(异常类型1 异常的变量名1) {
  // 程序代码
} catch(异常类型2 异常的变量名2) {
  // 程序代码
} catch(异常类型2 异常的变量名2) {
  // 程序代码
} finally {
  // 程序代码
}
```

### <a name="throws"></a>throws/throw 关键字

如果一个方法没有捕获一个检查性异常，那么该方法必须使用 throws 关键字来声明。throws 关键字放在方法签名的尾部。

也可以使用 throw 关键字抛出一个异常，无论它是新实例化的还是刚捕获到的。

```java
import java.io.*;
public class className
{
   public void withdraw(double amount) throws RemoteException,
                              InsufficientFundsException
   {
      // Method implementation
      throw new RemoteException();
   }
   //Remainder of class definition
}
```

#### <a name="customexc"></a>声明自定义异常

* 所有异常都必须是 Throwable 的子类。
* 如果希望写一个检查性异常类，则需要继承 Exception 类。
* 如果你想写一个运行时异常类，那么需要继承 RuntimeException 类。

```java
import java.io.*;
 
//自定义异常类，继承Exception类
public class InsufficientFundsException extends Exception
{
  //此处的amount用来储存当出现异常（取出钱多于余额时）所缺乏的钱
  private double amount;
  public InsufficientFundsException(double amount)
  {
    this.amount = amount;
  } 
  public double getAmount()
  {
    return amount;
  }
}
```

## Java 网络编程

java.net 包中提供了两种常见的网络协议的支持：

* TCP：TCP 是传输控制协议的缩写，它保障了两个应用程序之间的可靠通信。通常用于互联网协议，被称 TCP / IP。

* UDP：UDP 是用户数据报协议的缩写，一个无连接的协议。提供了应用程序之间要发送的数据的数据包。

### <a name="socket"></a>Socket 编程

java.net.Socket 类代表一个套接字，并且 java.net.ServerSocket 类为服务器程序提供了一种来监听客户端，并与他们建立连接的机制。

> 以下步骤在两台计算机之间使用套接字建立TCP连接时会出现：

+ 服务器实例化一个 ServerSocket 对象，表示通过服务器上的端口通信。

+ 服务器调用 ServerSocket 类的 accept() 方法，该方法将一直等待，直到客户端连接到服务器上给定的端口。

+ 服务器正在等待时，一个客户端实例化一个 Socket 对象，指定服务器名称和端口号来请求连接。

+ Socket 类的构造函数试图将客户端连接到指定的服务器和端口号。如果通信被建立，则在客户端创建一个 Socket 对象能够与服务器进行通信。

+ 在服务器端，accept() 方法返回服务器上一个新的 socket 引用，该 socket 连接到客户端的 socket。

连接建立后，通过使用 I/O 流在进行通信，每一个socket都有一个输出流和一个输入流，客户端的输出流连接到服务器端的输入流，而客户端的输入流连接到服务器端的输出流。

TCP 是一个双向的通信协议，因此数据可以通过两个数据流在同一时间发送.以下是一些类提供的一套完整的有用的方法来实现 socket。

## <a name="multithread"></a>Java 多线程编程

> 创建一个线程，Java 提供了三种创建线程的方法

* 通过实现 Runnable 接口；

```java
class RunnableDemo implements Runnable {
   private Thread t;
   private String threadName;
   
   RunnableDemo( String name) {
      threadName = name;
      System.out.println("Creating " +  threadName );
   }
   
   public void run() {
      System.out.println("Running " +  threadName );
      try {
         for(int i = 4; i > 0; i--) {
            System.out.println("Thread: " + threadName + ", " + i);
            // 让线程睡眠一会
            Thread.sleep(50);
         }
      }catch (InterruptedException e) {
         System.out.println("Thread " +  threadName + " interrupted.");
      }
      System.out.println("Thread " +  threadName + " exiting.");
   }
   
   public void start () {
      System.out.println("Starting " +  threadName );
      if (t == null) {
         t = new Thread (this, threadName);
         t.start ();
      }
   }
}
 
public class TestThread {
 
   public static void main(String args[]) {
      RunnableDemo R1 = new RunnableDemo( "Thread-1");
      R1.start();
      
      RunnableDemo R2 = new RunnableDemo( "Thread-2");
      R2.start();
   }   
}
```

* 通过继承 Thread 类本身；

```java
class ThreadDemo extends Thread {
   private Thread t;
   private String threadName;
   
   ThreadDemo( String name) {
      threadName = name;
      System.out.println("Creating " +  threadName );
   }
   
   public void run() {
      System.out.println("Running " +  threadName );
      try {
         for(int i = 4; i > 0; i--) {
            System.out.println("Thread: " + threadName + ", " + i);
            // 让线程睡眠一会
            Thread.sleep(50);
         }
      }catch (InterruptedException e) {
         System.out.println("Thread " +  threadName + " interrupted.");
      }
      System.out.println("Thread " +  threadName + " exiting.");
   }
   
   public void start () {
      System.out.println("Starting " +  threadName );
      if (t == null) {
         t = new Thread (this, threadName);
         t.start ();
      }
   }
}
 
public class TestThread {
 
   public static void main(String args[]) {
      ThreadDemo T1 = new ThreadDemo( "Thread-1");
      T1.start();
      
      ThreadDemo T2 = new ThreadDemo( "Thread-2");
      T2.start();
   }   
}
```

* 通过 Callable 和 Future 创建线程。

```java
public class CallableThreadTest implements Callable<Integer> {
    public static void main(String[] args)  
    {  
        CallableThreadTest ctt = new CallableThreadTest();  
        FutureTask<Integer> ft = new FutureTask<>(ctt);  
        for(int i = 0;i < 100;i++)  
        {  
            System.out.println(Thread.currentThread().getName()+" 的循环变量i的值"+i);  
            if(i==20)  
            {  
                new Thread(ft,"有返回值的线程").start();  
            }  
        }  
        try  
        {  
            System.out.println("子线程的返回值："+ft.get());  
        } catch (InterruptedException e)  
        {  
            e.printStackTrace();  
        } catch (ExecutionException e)  
        {  
            e.printStackTrace();  
        }  
  
    }
    @Override  
    public Integer call() throws Exception  
    {  
        int i = 0;  
        for(;i<100;i++)  
        {  
            System.out.println(Thread.currentThread().getName()+" "+i);  
        }  
        return i;  
    }  
}
```

创建线程的三种方式的对比

1. 采用实现 Runnable、Callable 接口的方式创见多线程时，线程类只是实现了 Runnable 接口或 Callable 接口，还可以继承其他类。

2. 使用继承 Thread 类的方式创建多线程时，编写简单，如果需要访问当前线程，则无需使用 Thread.currentThread() 方法，直接使用 this 即可获得当前线程。

* 线程的几个主要概念： 线程同步；线程间通信；线程死锁；线程控制（挂起、停止和恢复）。

## <a name="doccomment"></a>Java 文档注释

|标签	|描述	|示例|
|-------|----------------|------------------|
|@author	|标识一个类的作者|	@author description|
|@deprecated	|指名一个过期的类或成员|	@deprecated description|
|{@docRoot}	|指明当前文档根目录的路径	|Directory Path|
|@exception	|标志一个类抛出的异常	|@exception exception-name explanation|
|{@inheritDoc}	|从直接父类继承的注释|	Inherits a comment from the immediate surperclass.|
|{@link}	|插入一个到另一个主题的链接|	{@link name text}|
|{@linkplain}	|插入一个到另一个主题的链接，但是该链接显示纯文本字体|	Inserts an in-line link to another topic.|
|@param	|说明一个方法的参数|	@param parameter-name explanation|
|@return	|说明返回值类型|	@return explanation|
|@see	|指定一个到另一个主题的链接|	@see anchor|
|@serial	|说明一个序列化属性|	@serial description|
|@serialData	|说明通过writeObject( ) 和 writeExternal( )方法写的数据|	@serialData description|
|@serialField	|说明一个ObjectStreamField组件|	@serialField name type description|
|@since	|标记当引入一个特定的变化时|	@since release|
|@throws	|和 @exception标签一样.|	The @throws tag has the same meaning as the @exception tag.|
|{@value}	|显示常量的值，该常量必须是static属性。|	Displays the value of a constant, which must be a static field.|
|@version	|指定类的版本	|@version info|

## <a name="java8new"></a>Java8 新增了非常多的特性，我们主要讨论以下几个：

* __Lambda 表达式__ − Lambda允许把函数作为一个方法的参数（函数作为参数传递进方法中。

* __方法引用__ − 方法引用提供了非常有用的语法，可以直接引用已有Java类或对象（实例）的方法或构造器。与lambda联合使用，方法引用可以使语言的构造更紧凑简洁，减少冗余代码。

* __默认方法__ − 默认方法就是一个在接口里面有了一个实现的方法。

* __新工具__ − 新的编译工具，如：Nashorn引擎 jjs、 类依赖分析器jdeps。

* __Stream API__ −新添加的Stream API（java.util.stream） 把真正的函数式编程风格引入到Java中。

* __Date Time API__ − 加强对日期与时间的处理。

* __Optional 类__ − Optional 类已经成为 Java 8 类库的一部分，用来解决空指针异常。

* __Nashorn, JavaScript 引擎__ − Java 8提供了一个新的Nashorn javascript引擎，它允许我们在JVM上运行特定的javascript应用。