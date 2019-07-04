# Groovy 快速入门

* [注释](#docs)
* [枚举 Enumeration](#enumeration)

Groovy是一门基于JVM的动态语言，很多语法和Java类似。大部分Java代码也同时是合法的Groovy代码。本文是快速入门，所以针对语法并不会做非常详细的介绍。如果需要详细语法，请直接查看Groovy官方文档。另外为了省事，本文中的大部分代码例子直接引用了Groovy文档。

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




