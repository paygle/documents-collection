
# React.Component 导出

```js

React.Component = {

  constructor          // ES6语法

  props                // ES6语法

  defaultProps() {}    // ES6语法

  state

  context

  refs

  isReactComponent

  setState(){}

  // 客户端被实例化
  getDefaultProps(){}  // 在createReactClass方式使用， 不要在此调用 this.props

  getInitialState(){}  // 在createReactClass方式使用， 否则请用 constructor

  componentWillMount(){}  // 服务端可被实例化

  render(){}

  componentDidMount(){}


  updater(){}

  forceUpdate(){}

  // 存在期
  componentWillReceiveProps(nextProps){}

  shouldComponentUpdate(nextProps, nextState)

  componentWillUpdate(object nextProps, object nextState){} // 注意不要在此更新 props或者state

  componentDidUpdate(object prevProps, object prevState){}  // 在组件重新被渲染之后

  // 销毁时
  componentWillUnmount(){}
}

```


# ES6语法

```js

// 初始默认值设定
class Greeting extends React.Component { }
Greeting.defaultProps = { name: 'Mary' };



// Mixin(混入) 只有使用 create-react-class 才可用
var SetIntervalMixin = {
  componentWillUnmount: function() {
    this.intervals.forEach(clearInterval);
  }
};

var createReactClass = require('create-react-class');

var TickTock = createReactClass({
  mixins: [SetIntervalMixin], // 使用混入
  getInitialState: function() { return {seconds: 0}; },
  render: function() {
    return (
      <p>
        React has been running for {this.state.seconds} seconds.
      </p>
    );
  }
});

ReactDOM.render(
  <TickTock />,
  document.getElementById('example')
);

```


# React.js 导出

```js

const React = {

  Children: {
    map,
    forEach,
    count,
    toArray,
    only,
  },

  createRef,
  Component,
  PureComponent,

  createContext,
  forwardRef,

  Fragment
  StrictMode
  unstable_AsyncMode
  unstable_Profiler

  createElement
  cloneElement
  createFactory
  isValidElement: isValidElement,

  version: ReactVersion,

  __SECRET_INTERNALS_DO_NOT_USE_OR_YOU_WILL_BE_FIRED: ReactSharedInternals,
};

```


# ReactDOM 导出

```js

ReactDOM = {

  findDOMNode(any){}

  render(element, container, callback){}

  unmountComponentAtNode(container){}


  createPortal(children, container, null, key){}

  hydrate(element, container, callback){}

  flushSync(){}


  unstable_renderSubtreeIntoContainer(parentComponent, element, containerNode, callback ){}

  unstable_createPortal(){}

  unstable_batchedUpdates(){}

  unstable_deferredUpdates(){}

  unstable_interactiveUpdates(){}

  unstable_flushControlled(){}


  __SECRET_INTERNALS_DO_NOT_USE_OR_YOU_WILL_BE_FIRED: {
    // For TapEventPlugin which is popular in open source
    EventPluginHub,

    EventPluginRegistry,

    EventPropagators,

    ReactControlledComponent,

    ReactDOMComponentTree,

    ReactDOMEventListener,
  }
}


// ReactDOM/server 接口
{

  renderToString

  renderToStaticMarkup


  renderToNodeStream

  renderToStaticNodeStream

  version: ReactVersion
}

```