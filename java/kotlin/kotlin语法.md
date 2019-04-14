# Kotlin 语言

* [关键字](#kotlin)


## Kotlin 基本语法

## <a name="kotlin"></a>Kotlin 关键字

### 简单示例

```kotlin
/*
* 定义包: 目录与包的结构无需匹配：源代码可以在文件系统的任意位置
*/
package my.demo 
import java.util.*

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

```

