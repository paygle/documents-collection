
## XMLHttpRequest封装在HttpClient.js文件内，代码如下：

```js

function HttpClient(){}
    
HttpClient.prototype={
    //传给open方法的GET或者POST类型
    requestType: 'GET',
        
    //设置为TRUE时将发生异步调用  ，false时为同步调用  
    isAsync: false,

    //指定XMLHttpRequest返回的是文本对象还是XML文档对象  
    isGetXML: false,  

    //保存XMLHttpRequest实例的地方    
    xmlhttp: false,
        
    //发送了一个成功的异步调用后将调用的内容    
    callback: false,    
        
    //当调用XMLHttpRequest的send方法时将调用的内容    
    //自定义  ,不用的话注释掉即可。  
    onSend: function(){    
        document.getElementById("HttpClientStatus").style.display='block';
    },    
        
    //当readyState的值为4时将调用的内容，在回调函数之前进行    
    //自定义    ,不用的话注释掉即可。  
    onLoad: function(){    
        document.getElementById("HttpClientStatus").style.display='none';    
    },    
        
    //当出现http错误时将调用的内容    
    onError: function(error){    
        alert(error);    
    },    
        
    //实例化一个xmlhttpclient    
    init: function(){    
        try{  //对于Firefox和Opera等浏览器  
              this.xmlhttp=new XMLHttpRequest();  
              //有此版本的Mozilla浏览器在处理服务器返回的包含XML mime-type头部信息内容时会出错。    
              //所以，为了确保返回内容是text/xml信息，需要包含下面的语句。    
              if(xmlrequest.overrideMimeType)    
              {    
                    xmlrequest.overrideMimeType("text/xml");    
              }    
        }catch(e){  //对于IE浏览器  
            var XMLHTTP_IDS=new Array('Msxml2.XMLHTTP.7.0',
                                      'Msxml2.XMLHTTP.6.0',
                                      'Msxml2.XMLHTTP.5.0',
                                      'Msxml2.XMLHTTP.4.0',    
                                      'Msxml2.XMLHTTP.3.0',    
                                      'Msxml2.XMLHTTP',    
                                      'Micrsoft.XMLHTTP');
            var success=false;    
            for(var i=0;i<XMLHTTP_IDS.length && !success;i++){    
                try{    
                    this.xmlhttp=new ActiveXObject(XMLHTTP_IDS[i]);    
                    success=true;    
                    break;  
                }catch(e){    
                        
                }    
                    
            }    
            if(!success){    
                this.onError('Unable to create XMLHttpRequest.');    
            }    
        }    
    },    
        
    //用来处理readyState改变的内部方法    
    _readyStateChangeCallback: function(){
            
        switch(this.xmlhttp.readyState){    
        case 2: this.onSend();break;
        case 4:     
            this.onLoad();    
            if(this.xmlhttp.status==200){    
                if(this.isGetXML)  
                    this.callback( this.xmlhttp.responseXML);   
                else  
                    this.callback(this.xmlhttp.responseText);    
            }else{    
                this.onError('HTTP ERROR MAKING REQUEST: '+    
                              '['+this.xmlhttp.status+']'+    
                              this.xmlhttp.statusText);    
            }    
            break;    
        }    
    },    
        
    //发起页面请求的方法    
    //@ url,字符串型，为请求的页面    
    //@ playload ,字符串型，如果是post请求就需要    
    makeRequest: function(url,payload){
            
        if(!this.xmlhttp){    
            this.init();
        }    
        try{    
        this.xmlhttp.open(this.requestType,url,this.isAsync);
        }catch(e){    
            alert('error');    
        }    
        var self=this;    
            
        this.xmlhttp.onreadystatechange=function(){    
            //回调函数  
            self._readyStateChangeCallback();    
        }    
            
        this.xmlhttp.send(payload);
            
        if(!this.isAsync){    
            if(this.isGetXML)  
                return this.xmlhttp.responseXML;    
            else  
                return this.xmlhttp.responseText;    
        }
    }
}

```

### 要读取的xml文件，a.xml如下：

```html

<html>  
<head>  
    <title>Hello, Ajax!</title>  
</head>  
<body>I am there.</body>  
</html>

```

### HTML页面实现功能：可以以返回XML文档对象或者返回文件对象的形式，读取a.xml的内容并显示在当前页面内；代码如下：

```html

<!DOCTYPE html>  
<html>    
  <head>    
    <title>HttpClientTest</title>       
    <meta http-equiv="content-type" content="text/html; charset=UTF-8">  
  </head>    
  <script type="text/javascript" src="HttpClient.js" mce_src="HttpClient.js"></script>    
  <script type="text/javascript" >  
  var client=new HttpClient();    
  client.isAsync=true;    
  client.isGetXML=false;  
      
  function test(){    
      
      client.callback=function(result){    
  
    /*      //isGetXML设置为true时,返回的是XML文档对象  
          var title = result.getElementsByTagName("title")[0].childNodes[0].nodeValue;    
          alert(title);   
          document.getElementById('target').innerHTML=title;  */  
            
          //isGetXML设置为false时,返回的是文本  
          alert(result);   
          document.getElementById('target').innerHTML=result;   
      }    
          
      client.makeRequest('a.xml',null);    
          
  }    
 </script>    
  <body>    
    <div id="HttpClientStatus" style="display:block" mce_style="display:none">Loading....</div>    
    <a href="javascript:test()" mce_href="javascript:test()"> Make an Async Test call</a>    
    <div id='target'></div>    
</body>    
</html>  

```

