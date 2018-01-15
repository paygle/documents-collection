
# Proxy 代理对象

> ES6规范定义了一个全新的全局构造函数：代理（Proxy）。接受两个参数：目标对象（target）与句柄对象（handler）。
> 代理对象在IE浏览器中无法使用，而且还没有支持这一特性polyfill。


```js
    var target = {}, handler = {};
    var proxy = new Proxy(target, handler);
```

```js
var obj = new Proxy({}, {
  get: function (target, key, receiver) {
    console.log(`getting ${key}!`);
    return Reflect.get(target, key, receiver);
  },
  set: function (target, key, value, receiver) {
    console.log(`setting ${key}!`);
    return Reflect.set(target, key, value, receiver);
  }
});
```

* 所有对象都会共享一些基础功能：

> 对象都有属性。你可以get、set或删除它们或做更多操作。
> 对象都有原型。这也是JS中继承特性的实现方式。
> 有一些对象是可以被调用的函数或构造函数。

* 代理到底好在哪里？

代理可以帮助你观察或记录对象访问，当调试代码时助你一臂之力，测试框架也可以用代理来创建模拟对象（mock object）。

代理可以帮助你强化普通对象的能力，例如：惰性属性填充。

我不太想提到这一点，但是如果要想了解代理在代码中的运行方式，将代理的句柄对象包裹在另一个代理中是一个非常不错的办法，每当句柄方法被访问时就可以将你想要的信息输出到控制台中。

正如上文中只读视图的示例readOnlyView，我们可以用代理来限制对象的访问。当然在应用代码中很少遇到这种用例，但是Firefox在内部使用代理来实现不同域名之间的安全边界，是我们的安全模型的关键组成部分。

* 与WeakMap深度结合。在我们的readOnlyView示例中，每当对象被访问的时候创建一个新的代理。这种做法可以帮助我们节省在WeakMap中创建代理时的缓存内存，所以无论传递多少次对象给readOnlyView，只会创建一个代理。

这也是一个动人的WeakMap用例。

* 代理可解除。ES6规范中还定义了另外一个函数：Proxy.revocable(target, handler)。这个函数可以像new Proxy(target, handler)一样创建代理，但是创建好的代理后续可被解除。（Proxy.revocable方法返回一个对象，该对象有一个.proxy属性和一个.revoke方法。）一旦代理被解除，它即刻停止运行并抛出所有内部方法。

* 对象不变性。在某些情况下，ES6需要代理的句柄方法来报告与目标对象状态一致的结果，以此来保证所有对象甚至是代理的不变性。举个例子，除非目标不可扩展（inextensible），否则代理不能被声明为不可扩展的。

不变性的规则非常复杂，如果你看到类似“proxy can't report a non-existent property as non-configurable”这样的错误信息，就可以考虑从不变性的角度解决问题，最可能的补救方法是改变代理报告本身，或者在运行时改变目标对象来反射代理的报告指向。