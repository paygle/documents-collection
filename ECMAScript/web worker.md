# Web Worker Usage

#### 创建内联函数 worker
```js
function createWorker(f) {
  var blob = new Blob(['(' + f.toString() +')()'], {type: 'text/javascript'});
  var url = window.URL.createObjectURL(blob);
  var worker = new Worker(url);
  return worker;
}

function workerjs() {
  self.addEventListener('message', function(e) {  
    self.postMessage('you say' + e.data)
  })
}

var myworker = createWorker(workerjs)

```