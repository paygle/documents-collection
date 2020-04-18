# web.xml 配置详细

## web.xml加载过程（步骤）

当启动一个WEB项目时，容器包括（JBoss、Tomcat等）首先会读取项目web.xml配置文件里的配置，当这一步骤没有出错并且完成之后，项目才能正常地被启动起来。

1. 启动WEB项目的时候，容器首先会去它的配置文件web.xml读取两个节点:&lt;listener&gt; 和 &lt;context-param&gt;。

2. 紧接着，容器创建一个ServletContext（Application），这个WEB项目所有部分都将共享这个上下文。

3. 容器以 &lt;context-param&gt; 的 name 作为键，value作为值，将其转化为键值对，存入ServletContext。

4. 容器创建 &lt;listener&gt; 中的类实例，根据配置的class类路径 &lt;listener-class&gt; 来创建监听，在监听中会有contextInitialized(ServletContextEvent args)初始化方法，启动Web应用时，系统调用Listener的该方法，在这个方法中获得：ServletContext application = ServletContextEvent.getServletContext(); context-param的值 = application.getInitParameter("context-param的键"); 得到这个context-param的值之后，你就可以做一些操作了。

5. 举例：你可能想在项目启动之前就打开数据库，那么这里就可以在 &lt;context-param&gt; 中设置数据库的连接方式（驱动、url、user、password），在监听类中初始化数据库的连接。这个监听是自己写的一个类，除了初始化方法，它还有销毁方法，用于关闭应用前释放资源。比如:说数据库连接的关闭，此时，调用contextDestroyed(ServletContextEvent args)，关闭Web应用时，系统调用Listener的该方法。
   
6. 接着，容器会读取&lt;filter&gt;，根据指定的类路径来实例化过滤器。
   
7. 以上都是在WEB项目还没有完全启动起来的时候就已经完成了的工作。如果系统中有Servlet，则Servlet是在第一次发起请求的时候被实例化的，而且一般不会被容器销毁，它可以服务于多个用户的请求。所以，Servlet的初始化都要比上面提到的那几个要迟。
   
8.  总的来说，web.xml的加载顺序是:&lt;context-param&gt;->&lt;listener&gt;->&lt;filter&gt;->&lt;servlet&gt;。其中，如果web.xml中出现了相同的元素，则按照在配置文件中出现的先后顺序来加载。
   
9.  对于某类元素而言，与它们出现的顺序是有关的。以&lt;filter&gt;为例，web.xml中当然可以定义多个&lt;filter&gt;，与&lt;filter&gt;相关的一个元素是&lt;filter-mapping&gt;，注意，对于拥有相同&lt;filter-name&gt;的&lt;filter&gt;和&lt;filter-mapping&gt;元素而言，&lt;filter-mapping&gt;必须出现在&lt;filter&gt;之后，否则当解析到&lt;filter-mapping&gt;时，它所对应的&lt;filter-name&gt;还未定义。web容器启动初始化每个&lt;filter&gt;时，按照&lt;filter&gt;出现的顺序来初始化的，当请求资源匹配多个&lt;filter-mapping&gt;时，&lt;filter&gt;拦截资源是按照&lt;filter-mapping&gt;元素出现的顺序来依次调用doFilter()方法的。&lt;servlet&gt;同&lt;filter&gt;类似，此处不再赘述。


```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE html>
<web-app version="4.0" 
    xmlns="http://xmlns.jcp.org/xml/ns/javaee" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:schemaLocation="http://xmlns.jcp.org/xml/ns/javaee http://xmlns.jcp.org/xml/ns/javaee/web-app_4_0.xsd">

    <!-- web项目的名字，提供GUI工具可能会用来标记这个特定的Web应用的一个名称 -->
    <display-name>serTest</display-name>
    <!-- 对Web应用进行相关描述说明 -->
    <disciption>doing some Servlet Test!</disciption> 

    <!-- 指出IDE和GUI工具用来表示Web应用的大图标和小图标 -->
    <icon>
        <small-icon>/images/app_small.gif</small-icon>
        <large-icon>/images/app_large.gif</large-icon>
    </icon>

    <!-- 声明初始化参数，这里初始化对应的xml文件 -->
    <context-param>
        <!-- 参数名称  -->
        <param-name>contextConfigLocation</param-name>
        <!-- 参数值  -->
        <param-value>
            classpath*:/config/spring/applicationContext.xml,
            classpath*:/config/spring/applicationContext-repository.xml,
            classpath*:/config/spring/applicationContext-dubbo.xml
        </param-value>
        <description>参数描述</description>
    </context-param>


    <!-- 声明一个servlet的数据  -->
    <servlet>
        <!--- 指定servlet的名称 -->
        <servlet-name>myServlet</servlet-name>
        <!--- Servlet类，须包含完整路径 -->
        <servlet-class>*.myservlet</servlet-class>
        <display-name>为Servlet提供一个简短的名字被某些工具显示</display-name>
        <description>Servlet相关描述说明</description>
        <!-- 设置Servlet加载的次序 -->
        <load-on-startup>1</load-on-startup>
        <!--- 定义初始化参数，可有多个init-param -->
        <init-param>
            <param-name>driver</param-name>
            <param-value>com.mysql.jdbc.Driver</param-value>
        </init-param>
        <init-param>
            <param-name>url</param-name>
            <param-value>jdbc:mysql://localhost:3306/myDatabase</param-value>
        </init-param>
        <init-param>
            <param-name>username</param-name>
            <param-value>tang</param-value>
        </init-param>
        <init-param>
            <param-name>passwd</param-name>
            <param-value>whu</param-value>
        </init-param>
        <!-- 它会重写用于调用Web应用中servlet所设定的Enterprise JavaBean(EJB)的安全身份 -->
        <run-as>
            <description>Security role for anonymous access</description>
            <role-name>tomcat</role-name>
        </run-as>
        <jsp-file>指定web站台中的某个JSP网页的完整路径,必须由/开始</jsp-file>
    </servlet>
    <!-- 配置对应servlet 拦截url与应用中contrller的url的映射规则 -->
    <servlet-mapping>
        <servlet-name>myServlet</servlet-name>
        <!--
            容器无法识别同时拥有两种匹配规则的pattern, 优先级从高到低排列：
            1. 精确匹配：类似于/myServlet的精确路径
            2. 通配符匹配：/*
            3. 扩展名匹配：*.html，*.jpg ，.do ，.action之类的
            4. 默认匹配（/）——当之前匹配都不成功时
        -->
        <url-pattern>/myServlet</url-pattern>
    </servlet-mapping>

    <!-- 用于设定web应用的过滤器，可以过滤url或servlet请求 -->
    <filter>
        <filter-name>MyFilter</filter-name>
        <filter-class>com.demo.MyFilter</filter-class>
        <description>过滤器描述说明</description>
        <init-param>
            <param-name>test-param</param-name>
            <param-value>Initialization Parameter</param-value>
        </init-param>
    </filter>
    <!-- 命名一个过滤器，就要利用filter-mapping元素把它与一个或多个servlet或JSP页面相关联 -->
    <filter-mapping>
        <filter-name>MyFilterr</filter-name>
        <url-pattern>/*</url-pattern>
        <servlet-name>被过滤的servlet名称</servlet-name>
        <!--
        <dispatcher> 子元素可以设置的值及其意义
        REQUEST：(默认值)当用户直接访问页面时，Web容器将会调用过滤器。如果目标资源是通过RequestDispatcher的include()或forward()方法访问时，那么该过滤器就不会被调用。

        INCLUDE：如果目标资源是通过RequestDispatcher的include()方法访问时，那么该过滤器将被调用。除此之外，该过滤器不会被调用。

        FORWARD：如果目标资源是通过RequestDispatcher的forward()方法访问时，那么该过滤器将被调用，除此之外，该过滤器不会被调用。

        ERROR：如果目标资源是通过声明式异常处理机制调用时，那么该过滤器将被调用。除此之外，过滤器不会被调用。
        -->
        <dispatcher>指定过滤器所拦截的资源被 Servlet 容器调用的方式</dispatcher>
    </filter-mapping>

    <!-- 用于定义注册一个监听器类，可以收到事件什么时候发生以及用什么作为响应的通知 -->
    <listener>
        <listener-class>com.demo.MyListener</listenerclass>
    </listener>

    <!-- 如果某个会话在一定时间内未被访问，服务器可以抛弃它以节省内存 -->
    <session-config>
        <!-- 配置会话超时，单位是分钟 -->
        <session-timeout>120</session-timeout>
    </session-config>

    <!-- welcome-file-list 定义了首页文件，也就是用户直接输入域名时跳转的页面 -->
    <welcome-file-list>
        <welcome-file>index.html</welcome-file>
        <welcome-file>default.html</welcome-file>
        <welcome-file>default.jsp</welcome-file>
    </welcome-file-list>

    <!-- 在返回特定HTTP状态代码时，或者特定类型的异常被抛出时，能够制定将要显示的页面 -->
    <error-page> 
        <error-code>404</error-code> 
        <location>/NotFound.jsp</location> 
    </error-page>
    <error-page> 
        <exception-type>java.lang.NullException</exception-type> 
        <location>/error.jsp</location> 
    </error-page> 

    <!-- 分配特定的MIME类型 -->
    <mime-mapping>
        <extension>htm</extension> 
        <mime-type>text/html</mime-type> 
    </mime-mapping>


    <!-- 声明与资源相关的一个管理对象 -->
    <resource-env-ref>
        <resource-env-ref-name>jms/StockQueue</resource-env-ref-name>
    </resource-env-ref>
    <!-- 声明一个资源工厂使用的外部资源 -->
    <resource-ref>
        <description>JNDI JDBC DataSource ofshop</description>
        <res-ref-name>jdbc/sample_db</res-ref-name>
        <res-type>javax.sql.DataSource</res-type>
        <res-auth>Container</res-auth>
    </resource-ref>
    <!-- 安全限制配置,制定应该保护的URL。它与login-config元素联合使用 -->
    <security-constraint>
        <display-name>Example SecurityConstraint</display-name>
        <web-resource-collection>
            <web-resource-name>Protected Area</web-resource-name>
            <url-pattern>/jsp/security/protected/*</url-pattern>
            <http-method>DELETE</http-method>
            <http-method>GET</http-method>
            <http-method>POST</http-method>
            <http-method>PUT</http-method>
        </web-resource-collection>
        <auth-constraint>
            <role-name>tomcat</role-name>
            <role-name>role1</role-name>
        </auth-constraint>
   </security-constraint>
    <!-- 登陆验证配置 它与sercurity-constraint元素联合使用 -->
    <login-config>
        <auth-method>FORM</auth-method>
        <realm-name>Example-Based AuthentiationArea</realm-name>
        <form-login-config>
            <form-login-page>/jsp/security/protected/login.jsp</form-login-page>
            <form-error-page>/jsp/security/protected/error.jsp</form-error-page>
        </form-login-config>
    </login-config>
    <!--
        给出安全角色的一个列表，这些角色将出现在servlet元素内的security-role-ref元素的role-name子元素中。分别地声明角色可使高级IDE处理安全信息更为容易。
    -->
    <security-role>
        <role-name>tomcat</role-name>
    </security-role>
    <!-- Web环境参数：env-entry元素声明Web应用的环境项 -->
    <env-entry>
        <env-entry-name>minExemptions</env-entry-name>
        <env-entry-value>1</env-entry-value>
        <env-entry-type>java.lang.Integer</env-entry-type>
    </env-entry>

</web-app>
```

### 例如使用 <init-param> 来初始化数据库连接参数
```java
public void init(ServletConfig config) throws SevletException{
	super(config);
	String driver = config.getInitParameter("driver");
	String url = config.getInitParameter("url");
	String username = config.getInitParameter("username");
	String passwd = config.getInitParameter("passwd");
	try{
		Class.forName(driver).newInstance();
		this.conn = DriverManager.getConnection(url, username, passwd);
		System.out.println("Connection successful...");
	} catch(SQLExceprion se){
		System.out.println("se");
	} catch(Exception e){
		e.printStackTrace():
	}
	
}
```

# web.xml 各版本的 Schema 头部声明

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE html>
<web-app version="4.0" 
    xmlns="http://xmlns.jcp.org/xml/ns/javaee" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:schemaLocation="http://xmlns.jcp.org/xml/ns/javaee http://xmlns.jcp.org/xml/ns/javaee/web-app_4_0.xsd">
  
</web-app>
```

```xml
<?xml version="1.0" encoding="UTF-8"?>
<web-app version="3.1" 
    xmlns="http://xmlns.jcp.org/xml/ns/javaee" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:schemaLocation="http://xmlns.jcp.org/xml/ns/javaee http://xmlns.jcp.org/xml/ns/javaee/web-app_3_1.xsd">

</web-app>
```
    version 3.0 该版本已开始推荐使用注解进行web项目配置，但还是可以使用web.xml进行配置
```xml
<?xml version="1.0" encoding="UTF-8"?>
<web-app version="3.0" 
    xmlns="http://java.sun.com/xml/ns/javaee" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:schemaLocation="http://java.sun.com/xml/ns/javaee http://java.sun.com/xml/ns/javaee/web-app_3_0.xsd">

</web-app>
```

```xml
<?xml version="1.0" encoding="UTF-8"?>  
<web-app version="2.5" 
    xmlns="http://java.sun.com/xml/ns/javaee" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"  
    xsi:schemaLocation="http://java.sun.com/xml/ns/javaee http://java.sun.com/xml/ns/javaee/web-app_2_5.xsd">  

</web-app>
```

```xml
<?xml version="1.0" encoding="UTF-8"?>  
<web-app version="2.4" 
    xmlns="http://java.sun.com/xml/ns/j2ee" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"  
    xsi:schemaLocation="http://java.sun.com/xml/ns/j2ee http://java.sun.com/xml/ns/j2ee/web-app_2_4.xsd">  

</web-app>
```

```xml
<?xml version="1.0" encoding="UTF-8"?>  
<!DOCTYPE web-app PUBLIC "-//Sun Microsystems, Inc.//DTD Web Application 2.3//EN" "http://java.sun.com/dtd/web-app_2_3.dtd">  
<web-app>  
  <display-name>Servlet 2.3 Web Application</display-name>  
</web-app>
```
