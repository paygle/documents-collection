# DOM 常用API

* [Element.classList](#classList)

## <a name="classList"></a>Element.classList

  Element.classList 是一个只读属性，返回一个元素的类属性的实时DOMTokenList 集合。
  使用 classList 是替代element.className作为空格分隔的字符串访问元素的类列表的一种方便的方法。
  如果类属性未设置或为空，那么 elementClasses.length 返回 0。虽然element.classList 本身是只读的，但是你可以使用 add() 和 remove() 方法修改它。

  * classList 方法

```js

add( String [, String] )
// 添加指定的类值。如果这些类已经存在于元素的属性中，那么它们将被忽略。

remove( String [,String] )
// 删除指定的类值。

item ( Number )
// 按集合中的索引返回类值。

toggle ( String [, force] )
// 当只有一个参数时：切换 class value; 即如果类存在，则删除它并返回false，如果不存在，则添加它并返回true。
// 当存在第二个参数时：如果第二个参数的计算结果为true，则添加指定的类值，如果计算结果为false，则删除它

contains( String )
// 检查元素的类属性中是否存在指定的类值。

replace( oldClass, newClass )
// 用一个新类替换已有类。

```


### requestAnimationFrame 兼容

```js

(function() {
    var lastTime = 0;
    var vendors = ['ms', 'moz', 'webkit', 'o'];
    for(var x = 0; x < vendors.length && !window.requestAnimationFrame; ++x) {
        window.requestAnimationFrame = window[vendors[x]+'RequestAnimationFrame'];
        window.cancelAnimationFrame = window[vendors[x]+'CancelAnimationFrame'] 
                                   || window[vendors[x]+'CancelRequestAnimationFrame'];
    }
 
    if (!window.requestAnimationFrame)
        window.requestAnimationFrame = function(callback, element) {
            var currTime = new Date().getTime();
            var timeToCall = Math.max(0, 16 - (currTime - lastTime));
            var id = window.setTimeout(function() { callback(currTime + timeToCall); }, timeToCall);
            lastTime = currTime + timeToCall;
            return id;
        };
 
    if (!window.cancelAnimationFrame)
        window.cancelAnimationFrame = function(id) { clearTimeout(id); };
}());

```
