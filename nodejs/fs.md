### require('fs') 文件与文件操作(读写文件 删除 重命名)  

#### 删除文件

```js
    fs.unlink(path[,callback]) 或者 fs.unlinkSync(path)
```

#### 删除文件夹

* 方法1：使用递归

```js
  deleteFolderRecursive = function(path) {

      var files = [];

      if( fs.existsSync(path) ) {

          files = fs.readdirSync(path);

          files.forEach(function(file,index){

              var curPath = path + "/" + file;

              if(fs.statSync(curPath).isDirectory()) { // recurse

                  deleteFolderRecursive(curPath);

              } else { // delete file

                  fs.unlinkSync(curPath);

              }

          });

          fs.rmdirSync(path);

      }

  };
```

* 方法2：使用系统的命令

```js
    var exec = require('child_process').exec,child;

    child = exec('rm -rf test',function(err,out) {

      console.log(out); err && console.log(err);

    });
```

#### 读写文件 nodejs中操作相对就简单很多！来看看几个例子吧。
* 写文本文件

```js
    var fs = require("fs");

    var data = 'hello world';

    fs.writeFile('c:a.txt', data, 'ascii', function(err){

      if(err){

        console.log('写入文件失败');

      }else{

        console.log('保存成功, 赶紧去看看乱码吧');

      }
    })
```

* 注意：默认情况下，数据编码为utf8；mode=438 ；flag=w

* 读取文本文件

```js
    var fs = require("fs");

    var data = 'hello world';

    fs.readFile('c:a.txt','ascii', function(err, data){

      if(err){

        console.log('写入文件失败');

      }else{

        console.log(data);

      }

    })
```

* 注意： 如果没有特殊编码，那么就以二进制缓冲数据返回。

* 注意：二进制缓冲数据打印结果：<Buffer 68 65 6c 6c 6f 20 e8 97 5a a2>


#### 文件目录操作

    nodejs文件操作（fs）

    在操作文件时候，我们需要require(加载)File System包来获得文件操作功能。

    代码 var fs = require("fs");

    而这个fs如何使用呢，这个时候我们就需要查询nodejs官方的API

    基本文件操作包括：新建、重命名、删除等等，来看看几个例子

##### 新建文件夹

```js
    // 加载文件系统模块

    var fs = require("fs");

    // 在C盘创建一个名为a的文件夹

    fs.mkdir("c:a", function(err){

        if(!err){

            console.log("操作成功！");  

        }else{

            console.log("操作失败！");

        }

    });
```

* 注意：如果文件夹存在err就会有错误信息。

##### 删除文件夹

```js
  var fs = require("fs");

  // 删除C盘里的a文件夹

  fs.rmdir("c:a", function(err){

      if(err){

          console.log("删除失败！");

      }else{

          console.log("删除成功！");

      }

  });
```

* 注意：如果删除文件夹不存在，那么err就会有错误信息。

##### 重命名文件夹

```js
    var fs = require("fs");

    // 重命名a文件夹为b

    fs.rename("c:a","C:b",function(err){

        if(err){

            console.log("重命名失败！");

        }else{

            console.log("重命名成功！");

        }

    });
```

* 注意：文件夹不存在，那么err就会有错误信息。

##### 判断文件/文件夹是否存在

```js
    var fs = require("fs");

    // 判断a文件夹是否存在

    fs.exists("c:a", function(exists){

      if(exists){

          console.log("a文件夹存在");

      }else{

          console.log("a文件夹不存在")

      }

    });
```

* 注意：Then call the callback argument with either true or false
(这个回调函数参数值是true或者false)

##### 判断文件类型 fs.stat(), fs.lstat() and fs.fstat()

```js
    var fs = require("fs");

    // 获取a文件夹的类型

    fs.stat("C:a",function(err, stat){

        if(err){

            console.log("文件不存在！");

        }else{

            console.log("是否文件："+stat.isFile());

            console.log("是否文件夹："+stat.isDirectory());

        }
    });
```

* 其它状态函数：以下函数都有相同的异步函数追加 Async。如：isFileAsync()

```js
    stats.isFile()

    stats.isDirectory()

    stats.isBlockDevice()

    stats.isCharacterDevice()

    stats.isSymbolicLink()  (只针对 fs.lstat() 有效)

    stats.isFIFO()

    stats.isSocket()

    stats.isFIFO
```