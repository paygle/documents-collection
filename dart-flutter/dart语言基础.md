# 帮助文档

* [Dart 关键字](#keywords)


## <a name="keywords"></a> Dart 关键字

|||||||||||
|---|---|---|---|---|---|---|---|---|---|
|abstract |continue	|false |new |this |as |default |final |null |throw	|
|assert |deferred |finally |operator |true |async |do	|for |part |try	|
|async*	|dynamic |get |rethrow |typedef |await |else |if |return |var	|
|break |enum |implements |set |void |case |export |import |static	|while |
|catch |external |in |super |with |class |extends |is |switch |yield |
|const |factory |library |sync*  |yield* |

 * Hello world
```dart
import 'dart:math';
import 'package:angular2/angular2.dart';

// 类型定义
class Spacecraft {
  String name;
  DateTime launchDate;
  int launchYear;

  // 构造函数
  Spacecraft(this.name, this.launchDate) {
    launchYear = launchDate?.year;
  }

  // Named constructor that forwards to the default one.
  Spacecraft.unlaunched(String name) : this(name, null);

  // Method.
  void describe() {
    print('Spacecraft: $name');
    if (launchDate != null) {
      int years = new DateTime.now().difference(launchDate).inDays ~/ 365;
      print('Launched: $launchYear ($years years ago)');
    } else {
      print('Unlaunched');
    }
  }
}

// 主程序入口
void main() {
  var year = 1977;
  print('Hello, World!');
}
```