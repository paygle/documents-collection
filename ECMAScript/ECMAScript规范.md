## ECMAScript 6 基础

### 1. ECMAScript 语法提案的批准流程

  + Stage 0 - Strawman（展示阶段）
  + Stage 1 - Proposal（征求意见阶段）
  + Stage 2 - Draft（草案阶段）
  + Stage 3 - Candidate（候选人阶段）
  + Stage 4 - Finished（定案阶段）

### 2. let 和 const

#### let 声明

  + ES6 新增了let命令，用来声明变量。它的用法类似于var，但是所声明的变量，只在let命令所在的代码块内有效。**“for循环”** 的计数器，就很合适使用let命令。
  + var命令会发生**“变量提升”**现象，即变量可以在声明之前使用，值为undefined。这种现象多多少少是有些奇怪的，按照一般的逻辑，变量应该在声明语句之后才可以使用。为了纠正这种现象，let命令改变了语法行为，它所声明的变量一定要在声明后使用，否则报错。
  + **暂时性死区**，只要块级作用域内存在let命令，它所声明的变量就“绑定”（binding）这个区域，不再受外部的影响。
  + let不允许在相同作用域内，重复声明同一个变量。
  + ES6 引入了块级作用域，明确允许在块级作用域之中声明函数。ES6 规定，块级作用域之中，函数声明语句的行为类似于let，在块级作用域之外不可引用。ES5 规定，函数只能在顶层作用域和函数作用域之中声明，不能在块级作用域声明。但是，浏览器没有遵守这个规定，为了兼容以前的旧代码，还是支持在块级作用域之中声明函数，因此上面两种情况实际都能运行，不会报错。

#### const 声明

  + const声明一个只读的常量。一旦声明，常量的值就不能改变。对于const来说，只声明不赋值，就会报错。
  + const的作用域与let命令相同：只在声明所在的块级作用域内有效。
  + const实际上保证的，并不是变量的值不得改动，而是变量指向的那个内存地址不得改动。const只能保证这个指针是固定的，至于它指向的数据结构是不是可变的，就完全不能控制了。因此，将一个对象声明为常量必须非常小心。

#### ES6 声明变量的六种方法

  + ES5 只有两种声明变量的方法：var命令和function命令。ES6 除了添加let和const命令，后面章节还会提到，另外两种声明变量的方法：import命令和class命令。
  + 顶层对象，在浏览器环境指的是window对象，在 Node 指的是global对象。ES5 之中，顶层对象的属性与全局变量是等价的。从 ES6 开始，全局变量将逐步与顶层对象的属性脱钩。即 var a !== window.a 。

```js
// 获取顶层对象的两种方法

// 方法一
(typeof window !== 'undefined'
   ? window
   : (typeof process === 'object' &&
      typeof require === 'function' &&
      typeof global === 'object')
     ? global
     : this);

// 方法二
var getGlobal = function () {
  if (typeof self !== 'undefined') { return self; }
  if (typeof window !== 'undefined') { return window; }
  if (typeof global !== 'undefined') { return global; }
  throw new Error('unable to locate global object');
};

// 现在有一个提案，在语言标准的层面，引入global作为顶层对象。也就是说，在所有环境下，global都是存在的，都可以从它拿到顶层对象。
// ES6 模块的写法
import shim from 'system.global/shim';
shim();

```

### 3. 变量的解构赋值

  + ES6 允许按照一定模式，从数组和对象中提取值，对变量进行赋值，这被称为解构。本质上，这种写法属于“模式匹配”，只要等号两边的模式相同，左边的变量就会被赋予对应的值。如果解构不成功，变量的值就等于undefined。
  + 解构赋值允许指定默认值。注意，ES6 内部使用严格相等运算符（===），判断一个位置是否有值。所以，只有当一个数组成员严格等于undefined，默认值才会生效。

```js
let [foo, [[bar], baz]] = [1, [[2], 3]];
foo // 1
bar // 2
baz // 3

let [ , , third] = ["foo", "bar", "baz"];
third // "baz"

let [x, , y] = [1, 2, 3];
x // 1
y // 3

let [head, ...tail] = [1, 2, 3, 4];
head // 1
tail // [2, 3, 4]

let [x, y, ...z] = ['a'];
x // "a"
y // undefined
z // []

let [x, y = 'b'] = ['a']; // x='a', y='b'
let [x, y = 'b'] = ['a', undefined]; // x='a', y='b'

```

  + 对象的解构赋值，解构不仅可以用于数组，还可以用于对象。对象的解构与数组有一个重要的不同。数组的元素是按次序排列的，变量的取值由它的位置决定；而对象的属性没有次序，变量必须与属性同名，才能取到正确的值。

```js
let { bar, foo } = { foo: "aaa", bar: "bbb" };
foo // "aaa"
bar // "bbb"

// 如果变量名与属性名不一致，必须写成下面这样
let { foo: baz } = { foo: 'aaa', bar: 'bbb' };
baz // "aaa"

// 对象的解构赋值是下面形式的简写
// 先找到同名属性，然后再赋给对应的变量。真正被赋值的是后者，而不是前者。
// foo是匹配的模式，baz才是变量。真正被赋值的是变量baz，而不是模式foo。
let { foo: foo, bar: bar } = { foo: "aaa", bar: "bbb" };

// 与数组一样，解构也可以用于嵌套结构的对象。
// 注意，这时p键是模式，不是变量，因此不会被赋值。如果p也要作为变量赋值，需要前面加下p。
let obj = {
  p: [
    'Hello',
    { y: 'World' }
  ]
};

let { p, p: [x, { y }] } = obj;
x // "Hello"
y // "World"
p // ["Hello", {y: "World"}]

```

  + 解构赋值虽然很方便，但是解析起来并不容易。可以使用圆括号的情况只有一种：赋值语句的非模式部分，可以使用圆括号。，因为它们都是赋值语句，而不是声明语句；其次它们的圆括号都不属于模式的一部分。

```js

[(b)] = [3]; // 正确
({ p: (d) } = {}); // 正确
[(parseInt.prop)] = [3]; // 正确

```

  + 变量的解构赋值用途

```js
// 1. 交换变量的值
let x = 1, y = 2;
[x, y] = [y, x];

// 2. 从函数返回多个值,函数只能返回一个值，如果要返回多个值，只能将它们放在数组或对象里返回。
function example() { return [1, 2, 3]; }
let [a, b, c] = example();

// 3. 函数参数的定义

// 参数是一组有次序的值
function f([x, y, z]) { ... }
f([1, 2, 3]);

// 参数是一组无次序的值
function f({x, y, z}) { ... }
f({z: 3, y: 2, x: 1});

// 4. 提取 JSON 数据
let jsonData = {
  id: 42,
  status: "OK",
  data: [867, 5309]
};
let { id, status, data: number } = jsonData;

// 5. 函数参数的默认值
jQuery.ajax = function (url, {
  async = true,
  beforeSend = function () {},
  cache = true,
  complete = function () {},
  crossDomain = false,
  global = true,
  // ... more config
} = {}) {
  // ... do stuff
};

// 6. 遍历 Map 结构
const map = new Map();
map.set('first', 'hello');
map.set('second', 'world');

for (let [key, value] of map) {
  console.log(key + " is " + value);
}
// first is hello
// second is world
// 获取键名
for (let [key] of map) { /* ... */}

// 获取键值
for (let [,value] of map) { /* ... */}

// 7. 输入模块的指定方法
const { SourceMapConsumer, SourceNode } = require("source-map");

```

### 4. 字符串的扩展


  + 字符的 Unicode 表示法
  + codePointAt()
  + String.fromCodePoint()
  + 字符串的遍历器接口
  + at()
  + normalize()
  + includes(), startsWith(), endsWith()
  + repeat()
  + padStart()，padEnd()
  + matchAll()
  + 模板字符串
  + 实例：模板编译
  + 标签模板
  + String.raw()
  + 模板字符串的限制

  + 字符的 Unicode 表示法，采用\uxxxx形式表示一个字符，其中xxxx表示字符的 Unicode 码点。但是，这种表示法只限于码点在\u0000~\uFFFF之间的字符。超出这个范围的字符，必须用两个双字节的形式表示。"\uD842\uDFB7"。
  + ES6 对这一点做出了改进，只要将码点放入大括号，就能正确解读该字符。大括号表示法与四字节的 UTF-16 编码是等价的。

```js
"\u{20BB7}"
// "𠮷"

"\u{41}\u{42}\u{43}"
// "ABC"

let hello = 123;
hell\u{6F} // 123

'\u{1F680}' === '\uD83D\uDE80'
// true

// 有了这种表示法之后，JavaScript 共有 6 种方法可以表示一个字符。
'\z' === 'z'  // true
'\172' === 'z' // true
'\x7A' === 'z' // true
'\u007A' === 'z' // true
'\u{7A}' === 'z' // true
```

  + ES6 提供了codePointAt方法，能够正确处理 4 个字节储存的字符，返回一个字符的码点。

```js
// codePointAt方法是测试一个字符由两个字节还是由四个字节组成的最简单方法。
function is32Bit(c) {
  return c.codePointAt(0) > 0xFFFF;
}

is32Bit("𠮷") // true
is32Bit("a") // false

```

### 5. 正则的扩展

  + RegExp 构造函数
  + 字符串的正则方法
  + u 修饰符
  + y 修饰符
  + sticky 属性
  + flags 属性
  + s 修饰符：dotAll 模式
  + 后行断言
  + Unicode 属性类
  + 具名组匹配
  + String.prototype.matchAll


### 6. 数值的扩展

  + 二进制和八进制表示法
  + Number.isFinite(), Number.isNaN()
  + Number.parseInt(), Number.parseFloat()
  + Number.isInteger()
  + Number.EPSILON
  + 安全整数和 Number.isSafeInteger()
  + Math 对象的扩展
  + 指数运算符


### 7. 函数的扩展

  + 函数参数的默认值
  + rest 参数
  + 严格模式
  + name 属性
  + 箭头函数
  + 双冒号运算符
  + 尾调用优化
  + 函数参数的尾逗号


### 8. 数组的扩展

  + 扩展运算符
  + Array.from()
  + Array.of()
  + 数组实例的 copyWithin()
  + 数组实例的 find() 和 findIndex()
  + 数组实例的 fill()
  + 数组实例的 entries()，keys() 和 values()
  + 数组实例的 includes()
  + 数组的空位


### 9. 对象的扩展

  + 属性的简洁表示法
  + 属性名表达式
  + 方法的 name 属性
  + Object.is()
  + Object.assign()
  + 属性的可枚举性和遍历
  + Object.getOwnPropertyDescriptors()
  + __proto_属性，Object.setPrototypeOf()，Object.getPrototypeOf()
  + super 关键字
  + Object.keys()，Object.values()，Object.entries()
  + 对象的扩展运算符


### 10. Symbol

  + 作为属性名的 Symbol
  + 实例：消除魔术字符串
  + 属性名的遍历
  + Symbol.for()，Symbol.keyFor()
  + 实例：模块的 Singleton 模式
  + 内置的 Symbol 值


### 11. Set 和 Map 数据结构

  + Set
  + WeakSet
  + Map
  + WeakMap

### 12. Proxy

  + Proxy 实例的方法
  + Proxy.revocable()
  + this 问题
  + 实例：Web 服务的客户端


### 13. Reflect

  + 静态方法
  + 实例：使用 Proxy 实现观察者模式


### 14. Promise 对象

  + 基本用法
  + Promise.prototype.then()
  + Promise.prototype.catch()
  + Promise.prototype.finally()
  + Promise.all()
  + Promise.race()
  + Promise.resolve()
  + Promise.reject()
  + 应用
  + Promise.try()


### 15. Iterator 和 for...of 循环

  + Iterator（遍历器）的概念
  + 默认 Iterator 接口
  + 调用 Iterator 接口的场合
  + 字符串的 Iterator 接口
  + Iterator 接口与 Generator 函数
  + 遍历器对象的 return()，throw()
  + for...of 循环


### 16. Generator 函数的语法

  + next 方法的参数
  + for...of 循环
  + Generator.prototype.throw()
  + Generator.prototype.return()
  + next()、throw()、return() 的共同点
  + yield* 表达式
  + 作为对象属性的 Generator 函数
  + Generator 函数的this


### 17. async 函数

  + 基本用法
  + 语法
  + async 函数的实现原理
  + 与其他异步处理方法的比较
  + 实例：按顺序完成异步操作
  + 异步遍历器


### 18. Class 语法

  + 严格模式
  + constructor 方法
  + 类的实例对象
  + Class 表达式
  + 不存在变量提升
  + 私有方法和私有属性
  + this 的指向
  + name 属性
  + Class 的取值函数（getter）和存值函数（setter）
  + Class 的 Generator 方法
  + Class 的静态方法
  + Class 的静态属性和实例属性
  + new.target 属性



### 19. 修饰器

  + 类的修饰
  + 方法的修饰
  + 为什么修饰器不能用于函数？
  + core-decorators.js
  + 使用修饰器实现自动发布事件
  + Mixin
  + Trait
  + Babel 转码器的支持


### 20. Module 的语法

  + 严格模式
  + export 命令
  + import 命令
  + 模块的整体加载
  + export default 命令
  + export 与 import 的复合写法
  + 模块的继承
  + 跨模块常量
  + import()


### 21. ArrayBuffer

  + ArrayBuffer 对象
  + TypedArray 视图
  + 复合视图
  + DataView 视图
  + 二进制数组的应用
  + SharedArrayBuffer
  + Atomics 对象

