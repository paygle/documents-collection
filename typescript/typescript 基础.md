# Typescript 基础语法

* [基础类型](#base-type)
* [接口](#interface)
* [类](#class)
* [函数](#functions)
* [泛型](#generics)
* [枚举](#enums)
* [Symbol 类型](#symbols)
* [可迭代性](#iterators)
* [模块定义](#module-def)
* [命名空间](#namespace)
* [声明合并](#declare-merge)



## <a name="base-type"></a>基础类型

#### 布尔值

```ts
let isDone: boolean = false;
```

#### 数字

```ts
let decLiteral: number = 6;
let hexLiteral: number = 0xf00d;
let binaryLiteral: number = 0b1010;
let octalLiteral: number = 0o744;
```

#### 字符串

```ts
let name: string = "bob";
name = "smith";
```

```ts
let name: string = `Gene`;
let age: number = 37;
let sentence: string = `Hello, my name is ${ name }.

I'll be ${ age + 1 } years old next month.`;
```

#### 数组

```ts
let list: number[] = [1, 2, 3];
```

第二种方式是使用数组泛型，`Array<元素类型>`：

```ts
let list: Array<number> = [1, 2, 3];
```

#### 元组 Tuple

元组类型允许表示一个已知元素数量和类型的数组，各元素的类型不必相同。

```ts
// Declare a tuple type
let x: [string, number];
// Initialize it
x = ['hello', 10]; // OK
// Initialize it incorrectly
x = [10, 'hello']; // Error
```

当访问一个已知索引的元素，会得到正确的类型
当访问一个越界的元素，会使用联合类型替代：

```ts
x[3] = 'world'; // OK, 字符串可以赋值给(string | number)类型

console.log(x[5].toString()); // OK, 'string' 和 'number' 都有 toString

x[6] = true; // Error, 布尔不是(string | number)类型
```

#### 枚举

默认情况下，从`0`开始为元素编号。

```ts
enum Color {Red, Green, Blue}
let c: Color = Color.Green;
```

```ts
enum Color {Red = 1, Green, Blue}
let c: Color = Color.Green;
```

或者，全部都采用手动赋值：

```ts
enum Color {Red = 1, Green = 2, Blue = 4}
let c: Color = Color.Green;
```

枚举类型提供的一个便利是你可以由枚举的值得到它的名字。

```ts
enum Color {Red = 1, Green, Blue}
let colorName: string = Color[2];

alert(colorName);  // 显示'Green'因为上面代码里它的值是2
```

#### 任意值

可以使用`any`类型来标记这些变量：

```ts
let notSure: any = 4;
notSure = "maybe a string instead";
notSure = false; // okay, definitely a boolean
```

#### 空值

```ts
function warnUser(): void {
    alert("This is my warning message");
}
```

声明一个`void`类型的变量没有什么大用，因为你只能为它赋予`undefined`和`null`：

```ts
let unusable: void = undefined;
```

#### Null 和 Undefined

TypeScript里，`undefined`和`null`两者各自有自己的类型分别叫做`undefined`和`null`。
和`void`相似，它们的本身的类型用处不是很大：

```ts
// Not much else we can assign to these variables!
let u: undefined = undefined;
let n: null = null;
```

默认情况下`null`和`undefined`是所有类型的子类型。
就是说你可以把`null`和`undefined`赋值给`number`类型的变量。

然而，当你指定了`--strictNullChecks`标记，`null`和`undefined`只能赋值给`void`和它们各自。
这能避免*很多*常见的问题。

> 注意：我们鼓励尽可能地使用`--strictNullChecks`，但在本手册里我们假设这个标记是关闭的。

#### Never

`never`类型表示的是那些永不存在的值的类型。

`never`类型是任何类型的子类型，也可以赋值给任何类型；然而，*没有*类型是`never`的子类型或可以赋值给`never`类型（除了`never`本身之外）。
即使`any`也不可以赋值给`never`。


```ts
// 返回never的函数必须存在无法达到的终点
function error(message: string): never {
    throw new Error(message);
}

// 推断的返回值类型为never
function fail() {
    return error("Something failed");
}

// 返回never的函数必须存在无法达到的终点
function infiniteLoop(): never {
    while (true) {
    }
}
```

## 类型断言

> 其一是“尖括号”语法：

```ts
let someValue: any = "this is a string";

let strLength: number = (<string>someValue).length;
```

> 另一个为`as`语法：

```ts
let someValue: any = "this is a string";

let strLength: number = (someValue as string).length;
```

两种形式是等价的。当你在TypeScript里使用JSX时，只有`as`语法断言是被允许的。

## `let` 声明

> 块作用域

当用`let`声明一个变量，它使用的是*词法作用域*或*块作用域*。
拥有块级作用域的变量的另一个特点是，它们不能在被声明之前读或写。
不同于使用`var`声明的变量那样可以在包含它们的函数外访问，块作用域变量在包含它们的块或`for`循环之外是不能访问的。

```ts
function f(input: boolean) {
    let a = 100;

    if (input) {
        // Still okay to reference 'a'
        let b = a + 1;
        return b;
    }

    // Error: 'b' doesn't exist here
    return b;
}
```

> 重定义及屏蔽

我们提过使用`var`声明时，它不在乎你声明多少次；你只会得到1个。

在一个嵌套作用域里引入一个新名字的行为称做*屏蔽*。
它是一把双刃剑，它可能会不小心地引入新问题，同时也可能会解决一些错误。

*通常*来讲应该避免使用屏蔽，因为我们需要写出清晰的代码。
同时也有些场景适合利用它，你需要好好打算一下。

## `const` 声明

> 它们与`let`声明相似，但是就像它的名字所表达的，它们被赋值后不能再改变。
> 换句话说，它们拥有与`let`相同的作用域规则，但是不能对它们重新赋值。

除非你使用特殊的方法去避免，实际上`const`变量的内部状态是可修改的。
幸运的是，TypeScript允许你将对象的成员设置成只读的。

这很好理解，它们引用的值是*不可变的*。

```ts
const numLivesForCat = 9;
const kitty = {
    name: "Aurora",
    numLives: numLivesForCat,
}

// Error
kitty = {
    name: "Danielle",
    numLives: numLivesForCat
};

// all "okay"
kitty.name = "Rory";
kitty.name = "Kitty";
kitty.name = "Cat";
kitty.numLives--;
```

## 解构

### 解构数组

> 最简单的解构莫过于数组的解构赋值了：

```ts
let input = [1, 2];
let [first, second] = input;
console.log(first); // outputs 1
console.log(second); // outputs 2
```

这创建了2个命名变量 `first` 和 `second`。
相当于使用了索引，但更为方便：

```ts
first = input[0];
second = input[1];
```

> 解构作用于已声明的变量会更好：

```ts
// swap variables
[first, second] = [second, first];
```

> 作用于函数参数：

```ts
function f([first, second]: [number, number]) {
    console.log(first);
    console.log(second);
}
f(input);
```

> 你可以在数组里使用`...`语法创建剩余变量：

```ts
let [first, ...rest] = [1, 2, 3, 4];
console.log(first); // outputs 1
console.log(rest); // outputs [ 2, 3, 4 ]
```

当然，由于是JavaScript, 你可以忽略你不关心的尾随元素：

```ts
let [first] = [1, 2, 3, 4];
console.log(first); // outputs 1
```

或其它元素：

```ts
let [, second, , fourth] = [1, 2, 3, 4];
```

### 对象解构

```ts
let o = {
    a: "foo",
    b: 12,
    c: "bar"
};
let { a, b } = o;
```

就像数组解构，你可以用没有声明的赋值：

```ts
({ a, b } = { a: "baz", b: 101 });
```

> 注意，我们需要用括号将它括起来，因为Javascript通常会将以 `{` 起始的语句解析为一个块。

> 你可以在对象里使用`...`语法创建剩余变量：

```ts
let { a, ...passthrough } = o;
let total = passthrough.b + passthrough.c.length;

```

### 属性重命名

```ts
let { a: newName1, b: newName2 } = o;
```

这里的语法开始变得混乱。
你可以将 `a: newName1` 读做 "`a` 作为 `newName1`"。
方向是从左到右，好像你写成了以下样子：

```ts
let newName1 = o.a;
let newName2 = o.b;
```

令人困惑的是，这里的冒号*不是*指示类型的。
如果你想指定它的类型， 仍然需要在其后写上完整的模式。

```ts
let {a, b}: {a: string, b: number} = o;
```

### 默认值

> 默认值可以让你在属性为 undefined 时使用缺省值：

```ts
function keepWholeObject(wholeObject: { a: string, b?: number }) {
    let { a, b = 1001 } = wholeObject;
}
```

现在，即使 `b` 为 undefined ， `keepWholeObject` 函数的变量 `wholeObject` 的属性 `a` 和 `b` 都会有值。

## 函数声明

解构也能用于函数声明。

```ts
type C = { a: string, b?: number }
function f({ a, b }: C): void {
    // ...
}
```

但是，通常情况下更多的是指定默认值，解构默认值有些棘手。
首先，你需要在默认值之前设置其格式。

```ts
function f({ a, b } = { a: "", b: 0 }): void {
    // ...
}
f(); // ok, default to { a: "", b: 0 }
```

> 要小心使用解构。解构表达式要尽量保持小而简单。

## 展开

> 展开操作符正与解构相反。
它允许你将一个数组展开为另一个数组，或将一个对象展开为另一个对象。

```ts
let first = [1, 2];
let second = [3, 4];
let bothPlus = [0, ...first, ...second, 5];
```

这会令`bothPlus`的值为`[0, 1, 2, 3, 4, 5]`。
展开操作创建了`first`和`second`的一份浅拷贝。
它们不会被展开操作所改变。

你还可以展开对象：

```ts
let defaults = { food: "spicy", price: "$$", ambiance: "noisy" };
let search = { ...defaults, food: "rich" };
```

`search`的值为`{ food: "rich", price: "$$", ambiance: "noisy" }`。
对象的展开比数组的展开要复杂的多。
像数组展开一样，它是从左至右进行处理，但结果仍为对象。
这就意味着出现在展开对象后面的属性会覆盖前面的属性。
因此，如果我们修改上面的例子，在结尾处进行展开的话：

```ts
let defaults = { food: "spicy", price: "$$", ambiance: "noisy" };
let search = { food: "rich", ...defaults };
```

那么，`defaults`里的`food`属性会重写`food: "rich"`，在这里这并不是我们想要的结果。

> 对象展开还有其它一些意想不到的限制。

# <a name="interface"></a>接口

> 它有时被称做“鸭式辨型法”或“结构性子类型化”。口的作用就是为这些类型命名和为你的代码或第三方代码定义契约。

```ts
interface LabelledValue {
  label: string;
}

function printLabel(labelledObj: LabelledValue) {
  console.log(labelledObj.label);
}

let myObj = {size: 10, label: "Size 10 Object"};
printLabel(myObj);
```

`LabelledValue`接口就好比一个名字，用来描述上面例子里的要求。
它代表了有一个`label`属性且类型为`string`的对象。
需要注意的是，我们在这里并不能像在其它语言里一样，说传给`printLabel`的对象实现了这个接口。我们只会去关注值的外形。
只要传入的对象满足上面提到的必要条件，那么它就是被允许的。

还有一点值得提的是，类型检查器不会去检查属性的顺序，只要相应的属性存在并且类型也是对的就可以。

## 可选属性

> 带有可选属性的接口与普通的接口定义差不多，只是在可选属性名字定义的后面加一个`?`符号。

```ts
interface SquareConfig {
  color?: string;
  width?: number;
}

function createSquare(config: SquareConfig): {color: string; area: number} {
  let newSquare = {color: "white", area: 100};
  if (config.color) {
    newSquare.color = config.color;
  }
  if (config.width) {
    newSquare.area = config.width * config.width;
  }
  return newSquare;
}

let mySquare = createSquare({color: "black"});
```

## 只读属性

> 一些对象属性只能在对象刚刚创建的时候修改其值。
> 你可以在属性名前用`readonly`来指定只读属性:

```ts
interface Point {
    readonly x: number;
    readonly y: number;
}
```

## `readonly` vs `const`

最简单判断该用`readonly`还是`const`的方法是看要把它做为变量使用还是做为一个属性。
做为变量使用的话用`const`，若做为属性则使用`readonly`。

## 额外的属性检查

> 对象字面量会被特殊对待而且会经过*额外属性检查*，当将它们赋值给变量或作为参数传递的时候。
> 如果一个对象字面量存在任何“目标类型”不包含的属性时，你会得到一个错误。

```ts
// error: 'colour' not expected in type 'SquareConfig'
let mySquare = createSquare({ colour: "red", width: 100 });
```

> 绕开这些检查非常简单。使用类型断言：

```ts
let mySquare = createSquare({ width: 100, opacity: 0.5 } as SquareConfig);
```

> 如果`SquareConfig`带有上面定义的类型的`color`和`width`属性，并且*还会*带有任意数量的其它属性，那么我们可以这样定义它：

```ts
interface SquareConfig {
    color?: string;
    width?: number;
    [propName: string]: any;
}
```

> 还有最后一种跳过这些检查的方式，这可能会让你感到惊讶，它就是将这个对象赋值给一个另一个变量：
> 因为`squareOptions`不会经过额外属性检查，所以编译器不会报错。

```ts
let squareOptions = { colour: "red", width: 100 };
let mySquare = createSquare(squareOptions);
```

> 要留意，在像上面一样的简单代码里，你可能不应该去绕开这些检查。
> 对于包含方法和内部状态的复杂对象字面量来讲，你可能需要使用这些技巧，但是大部额外属性检查错误是真正的bug。

## 函数类型

> 为了使用接口表示函数类型，我们需要给接口定义一个调用签名。
> 它就像是一个只有参数列表和返回值类型的函数定义。参数列表里的每个参数都需要名字和类型。

```ts
interface SearchFunc {
  (source: string, subString: string): boolean;
}
```

这样定义后，我们可以像使用其它接口一样使用这个函数类型的接口。

```ts
let mySearch: SearchFunc;
mySearch = function(source: string, subString: string) {
  let result = source.search(subString);
  return result > -1;
}
```

对于函数类型的类型检查来说，函数的参数名不需要与接口里定义的名字相匹配。

```ts
let mySearch: SearchFunc;
mySearch = function(src: string, sub: string): boolean {
  let result = src.search(sub);
  return result > -1;
}
```

## 可索引的类型

> 与使用接口描述函数类型差不多，我们也可以描述那些能够“通过索引得到”的类型，比如`a[10]`或`ageMap["daniel"]`。

```ts
interface StringArray {
  [index: number]: string;
}

let myArray: StringArray;
myArray = ["Bob", "Fred"];

let myStr: string = myArray[0];
```

## 类类型

### 实现接口

> TypeScript也能够用它来明确的强制一个类去符合某种契约。

```ts
interface ClockInterface {
    currentTime: Date;
}

class Clock implements ClockInterface {
    currentTime: Date;
    constructor(h: number, m: number) { }
}
```

> 你也可以在接口中描述一个方法，在类里实现它，如同下面的`setTime`方法一样：

```ts
interface ClockInterface {
    currentTime: Date;
    setTime(d: Date);
}

class Clock implements ClockInterface {
    currentTime: Date;
    setTime(d: Date) {
        this.currentTime = d;
    }
    constructor(h: number, m: number) { }
}
```

> 接口描述了类的公共部分，而不是公共和私有两部分。

## 类静态部分与实例部分的区别

当你操作类和接口的时候，你要知道类是具有两个类型的：静态部分的类型和实例的类型。
你会注意到，当你用构造器签名去定义一个接口并试图定义一个类去实现这个接口时会得到一个错误：

```ts
interface ClockConstructor {
    new (hour: number, minute: number): ClockInterface;
}
interface ClockInterface {
    tick();
}

function createClock(ctor: ClockConstructor, hour: number, minute: number): ClockInterface {
    return new ctor(hour, minute);
}

class DigitalClock implements ClockInterface {
    constructor(h: number, m: number) { }
    tick() {
        console.log("beep beep");
    }
}
class AnalogClock implements ClockInterface {
    constructor(h: number, m: number) { }
    tick() {
        console.log("tick tock");
    }
}

let digital = createClock(DigitalClock, 12, 17);
let analog = createClock(AnalogClock, 7, 32);
```

> 因为`createClock`的第一个参数是`ClockConstructor`类型，在`createClock(AnalogClock, 7, 32)`里，会检查`AnalogClock`是否符合构造函数签名。

## 继承接口

```ts
interface Shape {
    color: string;
}

interface Square extends Shape {
    sideLength: number;
}

let square = <Square>{};
square.color = "blue";
square.sideLength = 10;
```

> 一个接口可以继承多个接口，创建出多个接口的合成接口。

```ts
interface Shape {
    color: string;
}

interface PenStroke {
    penWidth: number;
}

interface Square extends Shape, PenStroke {
    sideLength: number;
}

let square = <Square>{};
square.color = "blue";
square.sideLength = 10;
square.penWidth = 5.0;
```

## 混合类型

> 一个对象可以同时做为函数和对象使用，并带有额外的属性。

```ts
interface Counter {
    (start: number): string;
    interval: number;
    reset(): void;
}

function getCounter(): Counter {
    let counter = <Counter>function (start: number) { };
    counter.interval = 123;
    counter.reset = function () { };
    return counter;
}

let c = getCounter();
c(10);
c.reset();
c.interval = 5.0;
```

在使用JavaScript第三方库的时候，你可能需要像上面那样去完整地定义类型。

# 接口继承类

> 当接口继承了一个类类型时，它会继承类的成员但不包括其实现。
> 接口同样会继承到类的private和protected成员。
> 这意味着当你创建了一个接口继承了一个拥有私有或受保护的成员的类时，这个接口类型只能被这个类或其子类所实现（implement）。

当你有一个庞大的继承结构时这很有用，但要指出的是你的代码只在子类拥有特定属性时起作用。
这个子类除了继承至基类外与基类没有任何关系。

```ts
class Control {
    private state: any;
}

interface SelectableControl extends Control {
    select(): void;
}

class Button extends Control implements SelectableControl {
    select() { }
}

class TextBox extends Control {
    select() { }
}

// Error: Property 'state' is missing in type 'Image'.
class Image implements SelectableControl {
    select() { }
}

class Location {

}
```

在上面的例子里，`SelectableControl`包含了`Control`的所有成员，包括私有成员`state`。
因为`state`是私有成员，所以只能够是`Control`的子类们才能实现`SelectableControl`接口。
因为只有`Control`的子类才能够拥有一个声明于`Control`的私有成员`state`，这对私有成员的兼容性是必需的。

在`Control`类内部，是允许通过`SelectableControl`的实例来访问私有成员`state`的。
实际上，`SelectableControl`就像`Control`一样，并拥有一个`select`方法。
`Button`和`TextBox`类是`SelectableControl`的子类（因为它们都继承自`Control`并有`select`方法），但`Image`和`Location`类并不是这样的。

# <a name="class"></a>类

下面看一个使用类的例子：

```ts
class Greeter {
    greeting: string;
    constructor(message: string) {
        this.greeting = message;
    }
    greet() {
        return "Hello, " + this.greeting;
    }
}

let greeter = new Greeter("world");
```

## 继承

```ts
class Animal {
    name: string;
    constructor(theName: string) { this.name = theName; }
    move(distanceInMeters: number = 0) {
        console.log(`${this.name} moved ${distanceInMeters}m.`);
    }
}

class Snake extends Animal {
    constructor(name: string) { super(name); }
    move(distanceInMeters = 5) {
        console.log("Slithering...");
        super.move(distanceInMeters);
    }
}

class Horse extends Animal {
    constructor(name: string) { super(name); }
    move(distanceInMeters = 45) {
        console.log("Galloping...");
        super.move(distanceInMeters);
    }
}

let sam = new Snake("Sammy the Python");
let tom: Animal = new Horse("Tommy the Palomino");

sam.move();
tom.move(34);
```

```text
Slithering...
Sammy the Python moved 5m.
Galloping...
Tommy the Palomino moved 34m.
```

### 公共，私有与受保护的修饰符

> 默认为`public`
> 当成员被标记成`private`时，它就不能在声明它的类的外部访问。

```ts
class Animal {
    private name: string;
    constructor(theName: string) { this.name = theName; }
}

new Animal("Cat").name; // 错误: 'name' 是私有的.
```

> 如果其中一个类型里包含一个`private`成员，那么只有当另外一个类型中也存在这样一个`private`成员， 并且它们都是来自同一处声明时，我们才认为这两个类型是兼容的。对于`protected`成员也使用这个规则。

```ts
class Animal {
    private name: string;
    constructor(theName: string) { this.name = theName; }
}

class Rhino extends Animal {
    constructor() { super("Rhino"); }
}

class Employee {
    private name: string;
    constructor(theName: string) { this.name = theName; }
}

let animal = new Animal("Goat");
let rhino = new Rhino();
let employee = new Employee("Bob");

animal = rhino;
animal = employee; // 错误: Animal 与 Employee 不兼容.
```

> 理解`protected`

`protected`修饰符与`private`修饰符的行为很相似，但有一点不同，`protected`成员在派生类中仍然可以访问。

```ts
class Person {
    protected name: string;
    constructor(name: string) { this.name = name; }
}

class Employee extends Person {
    private department: string;

    constructor(name: string, department: string) {
        super(name)
        this.department = department;
    }

    public getElevatorPitch() {
        return `Hello, my name is ${this.name} and I work in ${this.department}.`;
    }
}

let howard = new Employee("Howard", "Sales");
console.log(howard.getElevatorPitch());
console.log(howard.name); // 错误
```

## readonly修饰符

> 你可以使用`readonly`关键字将属性设置为只读的。
> 只读属性必须在声明时或构造函数里被初始化。

```ts
class Octopus {
    readonly name: string;
    readonly numberOfLegs: number = 8;
    constructor (theName: string) {
        this.name = theName;
    }
}
let dad = new Octopus("Man with the 8 strong legs");
dad.name = "Man with the 3-piece suit"; // 错误! name 是只读的.
```

## 存取器

> TypeScript支持通过getters/setters来截取对对象成员的访问。
> 下面来看如何把一个简单的类改写成使用`get`和`set`。
> 存取器要求你将编译器设置为输出ECMAScript 5或更高。
> 其次，只带有`get`不带有`set`的存取器自动被推断为`readonly`。

```ts
let passcode = "secret passcode";

class Employee {
    private _fullName: string;

    get fullName(): string {
        return this._fullName;
    }

    set fullName(newName: string) {
        if (passcode && passcode == "secret passcode") {
            this._fullName = newName;
        }
        else {
            console.log("Error: Unauthorized update of employee!");
        }
    }
}

let employee = new Employee();
employee.fullName = "Bob Smith";
if (employee.fullName) {
    alert(employee.fullName);
}
```

## 静态属性

> 如同在实例属性上使用`this.`前缀来访问属性一样，这里我们使用`Grid.`来访问静态属性。

```ts
class Grid {
    static origin = {x: 0, y: 0};
    calculateDistanceFromOrigin(point: {x: number; y: number;}) {
        let xDist = (point.x - Grid.origin.x);
        let yDist = (point.y - Grid.origin.y);
        return Math.sqrt(xDist * xDist + yDist * yDist) / this.scale;
    }
    constructor (public scale: number) { }
}

let grid1 = new Grid(1.0);  // 1x scale
let grid2 = new Grid(5.0);  // 5x scale

console.log(grid1.calculateDistanceFromOrigin({x: 10, y: 10}));
console.log(grid2.calculateDistanceFromOrigin({x: 10, y: 10}));
```

## 抽象类

> 抽象类做为其它派生类的基类使用。它们一般不会直接被实例化。不同于接口，抽象类可以包含成员的实现细节。
> `abstract`关键字是用于定义抽象类和在抽象类内部定义抽象方法。
> 抽象类中的抽象方法不包含具体实现并且必须在派生类中实现。

```ts
abstract class Department {

    constructor(public name: string) {
    }

    printName(): void {
        console.log('Department name: ' + this.name);
    }

    abstract printMeeting(): void; // 必须在派生类中实现
}

class AccountingDepartment extends Department {

    constructor() {
        super('Accounting and Auditing'); // 在派生类的构造函数中必须调用 super()
    }

    printMeeting(): void {
        console.log('The Accounting Department meets each Monday at 10am.');
    }

    generateReports(): void {
        console.log('Generating accounting reports...');
    }
}

let department: Department; // 允许创建一个对抽象类型的引用
department = new Department(); // 错误: 不能创建一个抽象类的实例
department = new AccountingDepartment(); // 允许对一个抽象子类进行实例化和赋值
department.printName();
department.printMeeting();
department.generateReports(); // 错误: 方法在声明的抽象类中不存在
```

## 高级技巧

> 构造函数

```ts
class Greeter {
    greeting: string;
    constructor(message: string) {
        this.greeting = message;
    }
    greet() {
        return "Hello, " + this.greeting;
    }
}

let greeter: Greeter;
greeter = new Greeter("world");
console.log(greeter.greet());
```

## 把类当做接口使用

> 因为类可以创建出类型，所以你能够在允许使用接口的地方使用类。

```ts
class Point {
    x: number;
    y: number;
}

interface Point3d extends Point {
    z: number;
}

let point3d: Point3d = {x: 1, y: 2, z: 3};
```

# <a name="functions"></a>函数

## 为函数定义类型

让我们为上面那个函数添加类型：

```ts
function add(x: number, y: number): number {
    return x + y;
}

let myAdd = function(x: number, y: number): number { return x + y; };
```

我们可以给每个参数添加类型之后再为函数本身添加返回值类型。
TypeScript能够根据返回语句自动推断出返回值类型，因此我们通常省略它。

## 书写完整函数类型

```ts
let myAdd: (baseValue: number, increment: number) => number =
    function(x: number, y: number): number { return x + y; };
```

只要参数类型是匹配的，那么就认为它是有效的函数类型，而不在乎参数名是否正确。

第二部分是返回值类型。
对于返回值，我们在函数和返回值类型之前使用(`=>`)符号，使之清晰明了。
如之前提到的，返回值类型是函数类型的必要部分，如果函数没有返回任何值，你也必须指定返回值类型为`void`而不能留空。

函数的类型只是由参数类型和返回值组成的。
函数中使用的捕获变量不会体现在类型里。
实际上，这些变量是函数的隐藏状态并不是组成API的一部分。

## 推断类型

如果你在赋值语句的一边指定了类型但是另一边没有类型的话，TypeScript编译器会自动识别出类型：

```ts
// myAdd has the full function type
let myAdd = function(x: number, y: number): number { return x + y; };

// The parameters `x` and `y` have the type number
let myAdd: (baseValue: number, increment: number) => number =
    function(x, y) { return x + y; };
```

> 可选参数和默认参数, 可选参数必须跟在必须参数后面（除非设置默认值）。

## 剩余参数

> 编译器创建参数数组，名字是你在省略号（`...`）后面给定的名字，你可以在函数体内使用这个数组。
在JavaScript里，你可以使用`arguments`来访问所有传入的参数。
在TypeScript里，你可以把所有参数收集到一个变量里：

```ts
function buildName(firstName: string, ...restOfName: string[]) {
  return firstName + " " + restOfName.join(" ");
}

let employeeName = buildName("Joseph", "Samuel", "Lucas", "MacKinzie");
```

## `this`和箭头函数

> JavaScript里，`this`的值在函数被调用的时候才会指定。
> 箭头函数能保存函数创建时的`this`值，而不是调用时的值：

```ts
interface Card {
    suit: string;
    card: number;
}
interface Deck {
    suits: string[];
    cards: number[];
    createCardPicker(this: Deck): () => Card;
}
let deck: Deck = {
    suits: ["hearts", "spades", "clubs", "diamonds"],
    cards: Array(52),
    // NOTE: The function now explicitly specifies that its callee must be of type Deck
    createCardPicker: function(this: Deck) {
        return () => {
            let pickedCard = Math.floor(Math.random() * 52);
            let pickedSuit = Math.floor(pickedCard / 13);

            return {suit: this.suits[pickedSuit], card: pickedCard % 13};
        }
    }
}

let cardPicker = deck.createCardPicker();
let pickedCard = cardPicker();

alert("card: " + pickedCard.card + " of " + pickedCard.suit);
```

现在TypeScript知道`createCardPicker`期望在某个`Deck`对象上调用。
也就是说`this`是`Deck`类型的，而非`any`，因此`--noImplicitThis`不会报错了。

## 重载

> 它查找重载列表，尝试使用第一个重载定义。如果匹配的话就使用这个。
> 在定义重载的时候，一定要把最精确的定义放在最前面。

```ts
let suits = ["hearts", "spades", "clubs", "diamonds"];

function pickCard(x: {suit: string; card: number; }[]): number;
function pickCard(x: number): {suit: string; card: number; };
function pickCard(x): any {
    // Check to see if we're working with an object/array
    // if so, they gave us the deck and we'll pick the card
    if (typeof x == "object") {
        let pickedCard = Math.floor(Math.random() * x.length);
        return pickedCard;
    }
    // Otherwise just let them pick the card
    else if (typeof x == "number") {
        let pickedSuit = Math.floor(x / 13);
        return { suit: suits[pickedSuit], card: x % 13 };
    }
}

let myDeck = [{ suit: "diamonds", card: 2 }, { suit: "spades", card: 10 }, { suit: "hearts", card: 4 }];
let pickedCard1 = myDeck[pickCard(myDeck)];
alert("card: " + pickedCard1.card + " of " + pickedCard1.suit);

let pickedCard2 = pickCard(15);
alert("card: " + pickedCard2.card + " of " + pickedCard2.suit);
```

> 注意，`function pickCard(x): any`并不是重载列表的一部分，因此这里只有两个重载：一个是接收对象另一个接收数字。以其它参数调用`pickCard`会产生错误。

# <a name="generics"></a>泛型

下面来创建第一个使用泛型的例子：identity函数。
这个函数会返回任何传入它的值。
你可以把这个函数当成是`echo`命令。

不用泛型的话，这个函数可能是下面这样：

```ts
function identity(arg: number): number {
    return arg;
}
```

或者，我们使用`any`类型来定义函数：

```ts
function identity(arg: any): any {
    return arg;
}
```

使用`any`类型会导致这个函数可以接收任何类型的`arg`参数，这样就丢失了一些信息：传入的类型与返回的类型应该是相同的。
如果我们传入一个数字，我们只知道任何类型的值都有可能被返回。

因此，我们需要一种方法使返回值的类型与传入参数的类型是相同的。
这里，我们使用了*类型变量*，它是一种特殊的变量，只用于表示类型而不是值。

```ts
function identity<T>(arg: T): T {
    return arg;
}
```

我们给identity添加了类型变量`T`。
`T`帮助我们捕获用户传入的类型（比如：`number`），之后我们就可以使用这个类型。
之后我们再次使用了`T`当做返回值类型。现在我们可以知道参数类型与返回值类型是相同的了。
这允许我们跟踪函数里使用的类型的信息。

我们把这个版本的`identity`函数叫做泛型，因为它可以适用于多个类型。
不同于使用`any`，它不会丢失信息，像第一个例子那像保持准确性，传入数值类型并返回数值类型。

我们定义了泛型函数后，可以用两种方法使用。
第一种是，传入所有的参数，包含类型参数：

```ts
let output = identity<string>("myString");  // type of output will be 'string'
```

这里我们明确的指定了`T`是`string`类型，并做为一个参数传给函数，使用了`<>`括起来而不是`()`。

第二种方法更普遍。利用了*类型推论* -- 即编译器会根据传入的参数自动地帮助我们确定T的类型：

```ts
let output = identity("myString");  // type of output will be 'string'
```

注意我们没必要使用尖括号（`<>`）来明确地传入类型；编译器可以查看`myString`的值，然后把`T`设置为它的类型。
类型推论帮助我们保持代码精简和高可读性。如果编译器不能够自动地推断出类型的话，只能像上面那样明确的传入T的类型，在一些复杂的情况下，这是可能出现的。

### 使用泛型变量

使用泛型创建像`identity`这样的泛型函数时，编译器要求你在函数体必须正确的使用这个通用的类型。
换句话说，你必须把这些参数当做是任意或所有类型。

看下之前`identity`例子：

```ts
function identity<T>(arg: T): T {
    return arg;
}
```

如果我们想同时打印出`arg`的长度。
我们很可能会这样做：

```ts
function loggingIdentity<T>(arg: T): T {
    console.log(arg.length);  // Error: T doesn't have .length
    return arg;
}
```

如果这么做，编译器会报错说我们使用了`arg`的`.length`属性，但是没有地方指明`arg`具有这个属性。
记住，这些类型变量代表的是任意类型，所以使用这个函数的人可能传入的是个数字，而数字是没有`.length`属性的。

现在假设我们想操作`T`类型的数组而不直接是`T`。由于我们操作的是数组，所以`.length`属性是应该存在的。
我们可以像创建其它数组一样创建这个数组：

```ts
function loggingIdentity<T>(arg: T[]): T[] {
    console.log(arg.length);  // Array has a .length, so no more error
    return arg;
}
```

你可以这样理解`loggingIdentity`的类型：泛型函数`loggingIdentity`，接收类型参数`T`和参数`arg`，它是个元素类型是`T`的数组，并返回元素类型是`T`的数组。
如果我们传入数字数组，将返回一个数字数组，因为此时`T`的的类型为`number`。
这可以让我们把泛型变量T当做类型的一部分使用，而不是整个类型，增加了灵活性。

我们也可以这样实现上面的例子：

```ts
function loggingIdentity<T>(arg: Array<T>): Array<T> {
    console.log(arg.length);  // Array has a .length, so no more error
    return arg;
}
```

### 泛型类

> 泛型类看上去与泛型接口差不多。
> 泛型类使用（`<>`）括起泛型类型，跟在类名后面。

```ts
class GenericNumber<T> {
    zeroValue: T;
    add: (x: T, y: T) => T;
}

let myGenericNumber = new GenericNumber<number>();
myGenericNumber.zeroValue = 0;
myGenericNumber.add = function(x, y) { return x + y; };
```

`GenericNumber`类的使用没有什么去限制它只能使用`number`类型。也可以使用字符串或其它更复杂的类型。

### 泛型约束

> 创建一个包含`.length`属性的接口，使用这个接口和`extends`关键字来实现约束：

```ts
interface Lengthwise {
    length: number;
}

function loggingIdentity<T extends Lengthwise>(arg: T): T {
    console.log(arg.length);  // Now we know it has a .length property, so no more error
    return arg;
}
```

现在这个泛型函数被定义了约束，因此它不再是适用于任意类型：

```ts
loggingIdentity(3);  // Error, number doesn't have a .length property
```

我们需要传入符合约束类型的值，必须包含必须的属性：

```ts
loggingIdentity({length: 10, value: 3});
```

### 在泛型约束中使用类型参数

我们想要确保这个属性存在于对象`obj`上，因此我们需要在这两个类型之间使用约束。

```ts
function getProperty<T, K extends keyof T>(obj: T, key: K) {
    return obj[key];
}

let x = { a: 1, b: 2, c: 3, d: 4 };

getProperty(x, "a"); // okay
getProperty(x, "m"); // error: Argument of type 'm' isn't assignable to 'a' | 'b' | 'c' | 'd'.
```

### 在泛型里使用类类型

在TypeScript使用泛型创建工厂函数时，需要引用构造函数的类类型。比如，

```ts
function create<T>(c: {new(): T; }): T {
    return new c();
}
```

一个更高级的例子，使用原型属性推断并约束构造函数与类实例的关系。

```ts
class BeeKeeper {
    hasMask: boolean;
}

class ZooKeeper {
    nametag: string;
}

class Animal {
    numLegs: number;
}

class Bee extends Animal {
    keeper: BeeKeeper;
}

class Lion extends Animal {
    keeper: ZooKeeper;
}

function createInstance<A extends Animal>(c: new () => A): A {
    return new c();
}

createInstance(Lion).keeper.nametag;  // typechecks!
createInstance(Bee).keeper.hasMask;   // typechecks!
```

# <a name="enums"></a>枚举

TypeScript支持数字的和基于字符串的枚举。

### 数字枚举

> 数字枚举可以被混入到计算过的和常量成员

```ts
enum Direction {
    Up = 1,
    Down,
    Left,
    Right
}

console.log("Direction:", Direction.Down)
```

如上，我们定义了一个数字枚举，`Up`使用初始化为`1`。其余的成员会从`1`开始自动增长。
换句话说，`Direction.Up`的值为`1`，`Down`为`2`，`Left`为`3`，`Right`为`4`。

我们还可以完全不使用初始化器：`Up`的值为`0`，`Down`的值为`1`等等。
当我们不在乎成员的值的时候，这种自增长的行为是很有用处的，但是要注意每个枚举成员的值都是不同的。

### 字符串枚举

> 在一个字符串枚举里，每个成员都必须用字符串字面量，或另外一个字符串枚举成员进行初始化。

```ts
enum Direction {
    Up = "UP",
    Down = "DOWN",
    Left = "LEFT",
    Right = "RIGHT",
}
```

### 异构枚举（Heterogeneous enums）

从技术的角度来说，枚举可以混合字符串和数字成员，但是似乎你并不会这么做：

```ts
enum BooleanLikeHeterogeneousEnum {
    No = 0,
    Yes = "YES",
}
```

除非你真的想要利用JavaScript运行时的行为，否则我们不建议这样做。

> 计算的和常量成员。每个枚举成员都带有一个值，它可以是*常量*或*计算出来的*。

### 反向映射

> 数字枚举成员还具有了*反向映射*，要注意的是*不会*为字符串枚举成员生成反向映射。

```ts
enum Enum {
    A
}
let a = Enum.A;
let nameOfA = Enum[a]; // "A"
```

TypeScript可能会将这段代码编译为下面的JavaScript：

```js
var Enum;
(function (Enum) {
    Enum[Enum["A"] = 0] = "A";
})(Enum || (Enum = {}));
var a = Enum.A;
var nameOfA = Enum[a]; // "A"
```

### `const`枚举

> 常量枚举通过在枚举上使用`const`修饰符来定义，常量枚举不允许包含计算成员。

```ts
const enum Directions {
    Up,
    Down,
    Left,
    Right
}
```

### 外部枚举

> 外部枚举用来描述已经存在的枚举类型的形状。

```ts
declare enum Enum {
    A = 1,
    B,
    C = 2
}
```

# <a name="symbols"></a>Symbol 类型

除了用户定义的symbols，还有一些已经众所周知的内置symbols。
内置symbols用来表示语言内部的行为。

以下为这些symbols的列表：

### `Symbol.hasInstance`

方法，会被`instanceof`运算符调用。构造器对象用来识别一个对象是否是其实例。

### `Symbol.isConcatSpreadable`

布尔值，表示当在一个对象上调用`Array.prototype.concat`时，这个对象的数组元素是否可展开。

### `Symbol.iterator`

方法，被`for-of`语句调用。返回对象的默认迭代器。

### `Symbol.match`

方法，被`String.prototype.match`调用。正则表达式用来匹配字符串。

### `Symbol.replace`

方法，被`String.prototype.replace`调用。正则表达式用来替换字符串中匹配的子串。

### `Symbol.search`

方法，被`String.prototype.search`调用。正则表达式返回被匹配部分在字符串中的索引。

### `Symbol.species`

函数值，为一个构造函数。用来创建派生对象。

### `Symbol.split`

方法，被`String.prototype.split`调用。正则表达式来用分割字符串。

### `Symbol.toPrimitive`

方法，被`ToPrimitive`抽象操作调用。把对象转换为相应的原始值。

### `Symbol.toStringTag`

方法，被内置方法`Object.prototype.toString`调用。返回创建对象时默认的字符串描述。

### `Symbol.unscopables`

对象，它自己拥有的属性会被`with`作用域排除在外。

# <a name="iterators"></a>可迭代性

一些内置的类型如`Array`，`Map`，`Set`，`String`，`Int32Array`，`Uint32Array`等都已经实现了各自的`Symbol.iterator`。
对象上的`Symbol.iterator`函数负责返回供迭代的值。

## `for..of` 语句

`for..of`会遍历可迭代的对象，调用对象上的`Symbol.iterator`方法。

```ts
let someArray = [1, "string", false];

for (let entry of someArray) {
    console.log(entry); // 1, "string", false
}
```

### `for..of` vs. `for..in` 语句

> `for..of`和`for..in`均可迭代一个列表；但是用于迭代的值却不同
> `for..in`迭代的是对象的 *键* 的列表，可以操作任何对象
> `for..of`则迭代对象的键对应的值

下面的例子展示了两者之间的区别：

```ts
let list = [4, 5, 6];

for (let i in list) {
    console.log(i); // "0", "1", "2",
}

for (let i of list) {
    console.log(i); // "4", "5", "6"
}
```

> 目标代码生成为 ES5 和 ES3

当生成目标为ES5或ES3，迭代器只允许在`Array`类型上使用。
在非数组值上使用`for..of`语句会得到一个错误，就算这些非数组值已经实现了`Symbol.iterator`属性。

> 目标代码生成为 ECMAScript 2015 或更高

当目标为兼容ECMAScipt 2015的引擎时，编译器会生成相应引擎的`for..of`内置迭代器实现方式。

# <a name="module-def"></a>模块定义

> 模块是自声明的；两个模块之间的关系是通过在文件级别上使用imports和exports建立的。
> TypeScript与ECMAScript 2015一样，任何包含顶级`import`或者`export`的文件都被当成一个模块。
> 相反地，如果一个文件不带有顶级的`import`或者`export`声明，那么它的内容被视为全局可见的（因此对模块也是可见的）。

## export 导出声明

> 任何声明（比如变量，函数，类，类型别名或接口）都能够通过添加`export`关键字来导出。

```ts
class ZipCodeValidator implements StringValidator {
    isAcceptable(s: string) {
        return s.length === 5 && numberRegexp.test(s);
    }
}
export { ZipCodeValidator };
export { ZipCodeValidator as mainValidator };
```

### 重新导出

```ts
export class ParseIntBasedZipCodeValidator {
    isAcceptable(s: string) {
        return s.length === 5 && parseInt(s).toString() === s;
    }
}

// 导出原先的验证器但做了重命名
export {ZipCodeValidator as RegExpBasedZipCodeValidator} from "./ZipCodeValidator";
```

> 或者一个模块可以包裹多个模块，并把他们导出的内容联合在一起通过语法：`export * from "module"`。

```ts
export * from "./StringValidator"; // exports interface StringValidator
export * from "./ZipCodeValidator";  // exports class ZipCodeValidator
```

## import 导入声明

```ts
import { ZipCodeValidator } from "./ZipCodeValidator";

let myValidator = new ZipCodeValidator();
```

可以对导入内容重命名

```ts
import { ZipCodeValidator as ZCV } from "./ZipCodeValidator";
let myValidator = new ZCV();
```

## 将整个模块导入到一个变量，并通过它来访问模块的导出部分

```ts
import * as validator from "./ZipCodeValidator";
let myValidator = new validator.ZipCodeValidator();
```

> 具有副作用的导入模块 (尽管不推荐这么做，一些模块会设置一些全局状态供其它模块使用)

```ts
import "./my-module.js";
```

### 默认导出

每个模块都可以有一个`default`导出。
默认导出使用`default`关键字标记；并且一个模块只能够有一个`default`导出。
需要使用一种特殊的导入形式来导入`default`导出。

##### JQuery.d.ts

```ts
declare let $: JQuery;
export default $;
```

## `export =` 和 `import = require()`

> TypeScript模块支持`export =`语法以支持传统的CommonJS和AMD的工作流模型。
> `export =`语法定义一个模块的导出对象。它可以是类，接口，命名空间，函数或枚举。
> 若要导入一个使用了`export =`的模块时，必须使用TypeScript提供的特定语法`import module = require("module")`。

##### ZipCodeValidator.ts

```ts
let numberRegexp = /^[0-9]+$/;
class ZipCodeValidator {
    isAcceptable(s: string) {
        return s.length === 5 && numberRegexp.test(s);
    }
}
export = ZipCodeValidator;
```

##### Test.ts

```ts
import zip = require("./ZipCodeValidator");

// Some samples to try
let strings = ["Hello", "98052", "101"];

// Validators to use
let validator = new zip();

// Show whether each string passed each validator
strings.forEach(s => {
  console.log(`"${ s }" - ${ validator.isAcceptable(s) ? "matches" : "does not match" }`);
});
```

#### 可选的模块加载和其它高级加载场景

有时候，你只想在某种条件下才加载某个模块。
在TypeScript里，使用下面的方式来实现它和其它的高级加载场景，我们可以直接调用模块加载器并且可以保证类型完全。

编译器会检测是否每个模块都会在生成的JavaScript中用到。
如果一个模块标识符只在类型注解部分使用，并且完全没有在表达式中使用时，就不会生成`require`这个模块的代码。
省略掉没有用到的引用对性能提升是很有益的，并同时提供了选择性加载模块的能力。

这种模式的核心是`import id = require("...")`语句可以让我们访问模块导出的类型。
模块加载器会被动态调用（通过`require`），就像下面`if`代码块里那样。
它利用了省略引用的优化，所以模块只在被需要时加载。
为了让这个模块工作，一定要注意`import`定义的标识符只能在表示类型处使用（不能在会转换成JavaScript的地方）。

为了确保类型安全性，我们可以使用`typeof`关键字。
`typeof`关键字，当在表示类型的地方使用时，会得出一个类型值，这里就表示模块的类型。

##### 示例：Node.js里的动态模块加载

```ts
declare function require(moduleName: string): any;

import { ZipCodeValidator as Zip } from "./ZipCodeValidator";

if (needZipValidation) {
    let ZipCodeValidator: typeof Zip = require("./ZipCodeValidator");
    let validator = new ZipCodeValidator();
    if (validator.isAcceptable("...")) { /* ... */ }
}
```

##### 示例：require.js里的动态模块加载

```ts
declare function require(moduleNames: string[], onLoad: (...args: any[]) => void): void;

import * as Zip from "./ZipCodeValidator";

if (needZipValidation) {
    require(["./ZipCodeValidator"], (ZipCodeValidator: typeof Zip) => {
        let validator = new ZipCodeValidator.ZipCodeValidator();
        if (validator.isAcceptable("...")) { /* ... */ }
    });
}
```

##### 示例：System.js里的动态模块加载

```ts
declare const System: any;

import { ZipCodeValidator as Zip } from "./ZipCodeValidator";

if (needZipValidation) {
    System.import("./ZipCodeValidator").then((ZipCodeValidator: typeof Zip) => {
        var x = new ZipCodeValidator();
        if (x.isAcceptable("...")) { /* ... */ }
    });
}
```

##### 使用其它的JavaScript库，它们通常是在`.d.ts`文件里定义的。

要想描述非TypeScript编写的类库的类型，我们需要声明类库所暴露出的API。
我们可以使用顶级的`export`声明来为每个模块都定义一个`.d.ts`文件，但最好还是写在一个大的`.d.ts`文件里。
我们使用与构造一个外部命名空间相似的方法，但是这里使用`module`关键字并且把名字用引号括起来，方便之后`import`。
例如：

##### node.d.ts (simplified excerpt)

```ts
declare module "url" {
    export interface Url {
        protocol?: string;
        hostname?: string;
        pathname?: string;
    }

    export function parse(urlStr: string, parseQueryString?, slashesDenoteHost?): Url;
}

declare module "path" {
    export function normalize(p: string): string;
    export function join(...paths: any[]): string;
    export let sep: string;
}
```

现在我们可以`/// <reference>` `node.d.ts`并且使用`import url = require("url");`或`import * as URL from "url"`加载模块。

```ts
/// <reference path="node.d.ts"/>
import * as URL from "url";
let myUrl = URL.parse("http://www.typescriptlang.org");
```

#### 外部模块简写

假如你不想在使用一个新模块之前花时间去编写声明，你可以采用声明的简写形式以便能够快速使用它。

##### declarations.d.ts

```ts
declare module "hot-new-module";
```

简写模块里所有导出的类型将是`any`。

```ts
import x, {y} from "hot-new-module";
x(y);
```

#### 模块声明通配符

```ts
declare module "*!text" {
    const content: string;
    export default content;
}
// Some do it the other way around.
declare module "json!*" {
    const value: any;
    export default value;
}
```

现在你可以就导入匹配`"*!text"`或`"json!*"`的内容了。

```ts
import fileContent from "./xyz.txt!text";
import data from "json!http://example.com/data.json";
console.log(data, fileContent);
```

> 模块里不要使用命名空间

#### 危险信号

以下均为模块结构上的危险信号。重新检查以确保你没有在对模块使用命名空间：

* 文件的顶层声明是`export namespace Foo { ... }` （删除`Foo`并把所有内容向上层移动一层）
* 文件只有一个`export class`或`export function` （考虑使用`export default`）
* 多个文件的顶层具有同样的`export namespace Foo {` （不要以为这些会合并到一个`Foo`中！）

# <a name="namespace"></a>命名空间

> 命名空间是位于全局命名空间下的一个普通的带有名字的JavaScript对象。
> 像命名空间一样，模块可以包含代码和声明。不同的是模块可以*声明*它的依赖。
> 对于Node.js应用来说，模块是默认并推荐的组织代码的方式。
> 不应该对模块使用命名空间，使用命名空间是为了提供逻辑分组和避免命名冲突。
> 模块文件本身已经是一个逻辑分组，并且它的名字是由导入这个模块的代码指定，所以没有必要为导出的对象增加额外的模块层。
> 使用命名空间的验证器

```ts
namespace Validation {
    export interface StringValidator {
        isAcceptable(s: string): boolean;
    }

    const lettersRegexp = /^[A-Za-z]+$/;
    const numberRegexp = /^[0-9]+$/;

    export class LettersOnlyValidator implements StringValidator {
        isAcceptable(s: string) {
            return lettersRegexp.test(s);
        }
    }

    export class ZipCodeValidator implements StringValidator {
        isAcceptable(s: string) {
            return s.length === 5 && numberRegexp.test(s);
        }
    }
}

// Some samples to try
let strings = ["Hello", "98052", "101"];

// Validators to use
let validators: { [s: string]: Validation.StringValidator; } = {};
validators["ZIP code"] = new Validation.ZipCodeValidator();
validators["Letters only"] = new Validation.LettersOnlyValidator();

// Show whether each string passed each validator
for (let s of strings) {
    for (let name in validators) {
        console.log(`"${ s }" - ${ validators[name].isAcceptable(s) ? "matches" : "does not match" } ${ name }`);
    }
}
```

> 多文件中的命名空间, 尽管是不同的文件，它们仍是同一个命名空间，并且在使用的时候就如同它们在一个文件中定义的一样。
> 当涉及到多文件时，我们必须确保所有编译后的代码都被加载了。

我们有两种方式。

第一种方式，把所有的输入文件编译为一个输出文件，需要使用`--outFile`标记：

```Shell
tsc --outFile sample.js Test.ts
```

编译器会根据源码里的引用标签自动地对输出进行排序。你也可以单独地指定每个文件。

```Shell
tsc --outFile sample.js Validation.ts LettersOnlyValidator.ts ZipCodeValidator.ts Test.ts
```

第二种方式，我们可以编译每一个文件（默认方式），那么每个源文件都会对应生成一个JavaScript文件。
然后，在页面上通过`<script>`标签把所有生成的JavaScript文件按正确的顺序引进来，比如：

##### MyTestPage.html (excerpt)

```html
    <script src="Validation.js" type="text/javascript" />
    <script src="LettersOnlyValidator.js" type="text/javascript" />
    <script src="ZipCodeValidator.js" type="text/javascript" />
    <script src="Test.js" type="text/javascript" />
```

## 别名

另一种简化命名空间操作的方法是使用`import q = x.y.z`给常用的对象起一个短的名字。
不要与用来加载模块的`import x = require('name')`语法弄混了，这里的语法是为指定的符号创建一个别名。
你可以用这种方法为任意标识符创建别名，也包括导入的模块中的对象。

```ts
namespace Shapes {
    export namespace Polygons {
        export class Triangle { }
        export class Square { }
    }
}

import polygons = Shapes.Polygons;
let sq = new polygons.Square(); // Same as "new Shapes.Polygons.Square()"
```

注意，我们并没有使用`require`关键字，而是直接使用导入符号的限定名赋值。
这与使用`var`相似，但它还适用于类型和导入的具有命名空间含义的符号。
重要的是，对于值来讲，`import`会生成与原始符号不同的引用，所以改变别名的`var`值并不会影响原始变量的值。

### 外部命名空间

流行的程序库D3在全局对象`d3`里定义它的功能。
因为这个库通过一个`<script>`标签加载（不是通过模块加载器），它的声明文件使用内部模块来定义它的类型。
为了让TypeScript编译器识别它的类型，我们使用外部命名空间声明。
比如，我们可以像下面这样写：

##### D3.d.ts (部分摘录)

```ts
declare namespace D3 {
    export interface Selectors {
        select: {
            (selector: string): Selection;
            (element: EventTarget): Selection;
        };
    }

    export interface Event {
        x: number;
        y: number;
    }

    export interface Base extends Selectors {
        event: Event;
    }
}

declare var d3: D3.Base;
```

# <a name="declare-merge"></a>声明合并

> “声明合并”是指编译器将针对同一个名字的两个独立声明合并为单一声明。
> 合并后的声明同时拥有原先两个声明的特性。
> 任何数量的声明都可被合并；不局限于两个声明。

### 基础概念

TypeScript中的声明会创建以下三种实体之一：命名空间，类型或值。

| Declaration Type | Namespace | Type | Value |
|------------------|:---------:|:----:|:-----:|
| Namespace        |     X     |      |   X   |
| Class            |           |   X  |   X   |
| Enum             |           |   X  |   X   |
| Interface        |           |   X  |       |
| Type Alias       |           |   X  |       |
| Function         |           |      |   X   |
| Variable         |           |      |   X   |

理解每个声明创建了什么，有助于理解当声明合并时有哪些东西被合并了。

#### 合并接口

从根本上说，合并的机制是把双方的成员放到一个同名的接口里。

```ts
interface Box {
    height: number;
    width: number;
}

interface Box {
    scale: number;
}

let box: Box = {height: 5, width: 6, scale: 10};
```

接口的非函数的成员应该是唯一的。
如果它们不是唯一的，那么它们必须是相同的类型。
如果两个接口中同时声明了同名的非函数成员且它们的类型不同，则编译器会报错。
对于函数成员，每个同名函数声明都会被当成这个函数的一个重载。

* 同时需要注意，当接口`A`与后来的接口`A`合并时，后面的接口具有更高的优先级。
* 注意每组接口里的声明顺序保持不变，但各组接口之间的顺序是后来的接口重载出现在靠前位置。
* 合并命名空间，非导出成员仅在其原有的（合并前的）命名空间内可见。合并之后，从其它命名空间合并进来的成员无法访问非导出成员。

# <a name="Triple-Slash"></a> 三斜线指令

> 三斜线指令是包含单个XML标签的单行注释。
> 注释的内容会做为编译器指令使用。

三斜线指令*仅*可放在包含它的文件的最顶端。
一个三斜线指令的前面只能出现单行或多行注释，这包括其它的三斜线指令。
如果它们出现在一个语句或声明之后，那么它们会被当做普通的单行注释，并且不具有特殊的涵义。

> 用于声明文件间的*依赖*。

```ts
/// <reference path="..." />
```

> 声明了对某个包的依赖，仅当在你需要写一个`d.ts`文件时才使用这个指令。

```ts
/// <reference types="..." />
```

> 把一个文件标记成*默认库*。你会在`lib.d.ts`文件和它不同的变体的顶端看到这个注释。

```ts
/// <reference no-default-lib="true"/>
```
