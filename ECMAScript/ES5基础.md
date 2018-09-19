# ECMAScript 5 基础语法

* [复用和继承模式](#inherit-reuse)


## <a name="inherit-reuse"></a>复用和继承模式

  * 对于构造函数的一般经验法则是：应该将可复用的成员添加到原型中


### 继承一： 默认类继承模式

  * 其继承自身属性的同时也包括原型属性和方法

  * **缺点：**在于同时继承了两个对象的属性，即添加到this的属性和原型属性。效率低下！

```js

function Parent(name) { this.name = name || 'Adam'; }

Parent.prototype.say = function() { return this.name; }

function Child(name) {}

// 默认继承方式
function inherit(Child, Parent) {

  // 注意, 原型属性应该指向一个对象，面不是一个函数
  Child.prototype = new Parent();

}

var kid = new Child();
kid.name = '输出“pack”';
kid.say();  // 输出“pack”

```


### 继承二： 构造函数继承

  * 只能继承在父构造函数中添加到this的属性，并不能继承prototype原型中的成员

  * 解决从子构造函数到父构造函数的参数传递问题，可获得父对象自身成员的真实副本并且不会有覆父类风险

  * **缺点：**无法继承原型链接任何东西，并且原型也仅是添加可重用方法及属性，它并不会为每个实例重新创建原型

```js

// 构造函数继承
function Child(a, b, c, d) {

  // 使用该模式时，子对象获得了继承成员的副本
  Parent.apply(this, arguments);

}


// 多重继承，实例：
function Cat() {
  this.legs = 4;
  this.say = function() { return "meaowww"; }
}

function Bird() {
  this.wings = 2;
  this.fly = true;
}

function CatWings() {
  Cat.apply(this);
  Bird.apply(this);
}

var jane = new CatWings();

```


### 继承三： 借用和设置原型继承

  * 主要是结合前两种模式，即先借用构造函数，然后设置子构造函数的原型使其指向一个构造函数创的建实例

  * **缺点：**父构造函数被调用了两次，因此导致了效率低下的问题

```js

// 借用和设置原型继承
function Child(a, b, c, d) {

  // 使用该模式时，子对象获得了继承成员的副本
  Parent.apply(this, arguments);

}

Child.prototype = new Parent();

```


### 继承四： 共享原型链接

  * 可复用成员应该转移到原型中而不是放置在this中，任何值得继承的东西都应该放置在原型中实现

  * **缺点：**父构造函数被调用了两次，因此导致了效率低下的问题

```js

// 共享原型链接继承
function inherit(Child, Parent) {

  // 所有的对象实际是共享了同一个原型
  Child.prototype = Parent.prototype;

}

```


### 继承五： 临时函数函数

  * 解决共享同一个原型所带来的问题，同时还能够继续受益于原型链接带来的好处

  这种模式了被称之为使用代理函数或代理构造函数的模式，而不是使用临时构造函数的模式，这是因为临时构造函数实际上是一个用于获得父对象的原型链接代理

```js

// 临时函数函数继承
function inherit(Child, Parent) {

  // 临时函数， 断开父对象与子对象的原型之间的直接链接关系
  var F = function() {}

  F.prototype = Parent.prototype;

  Child.prototype = new F();  // 临时构造函数

  // 存储超类， 看应用场景是否需要
  Child._SUPUER_PROTO_ = Parent.prototype;

  // 需要重置构造函数，默认指向父类
  Child.prototype.constructor = Child;
}

```


### 继承六： 原型继承

  * 现代无类继承模式

```js

function object(o) {
  var F = function() {}
  F.prototype = o;
  return new F();
}

// 要继承的对象
var parent = { name: 'Papa' };

// 新对象
var child = object(parent);

// 使用ES5中的 Object.create 方法代替上面的 object()
var child = Object.create(parent)


// 通过(深/浅)复制属性实现继承
function extend(parent, child) {
  var i;
  child = child || {};
  for(i in parent) {
    if (parent.hasOwnProperty(i)) {
      child[i] = parent[i];
    }
  }
  return child;
}


// 借用某个方式，不需要父子继承关系
OtherObject.doAction.call(myObject, param1, param1, ...param);
OtherObject.doAction.apply(myObject, [param1, param1, ...param]);

function myFun() {
  return [].slice.call(arguments, 1, 3);
}

// 借用和绑定 bind(), 把函数永远的绑定到给定的对象上，无法改变
var newFunc = obj.someFunc.bind(myobj, 1, 2, 3);


```
