# Spring Boot 2.x 笔记

### Spring Boot 的依赖和自动配置

	org.springframework.boot.autoconfigure.web.servlet

	@Component以及他的特殊化(@Controller, @Service 和 @Repository)允许在通过类路径扫描自动发现。

	@Bean却只能在配置类中明确的声明一个单例的bean。

|注解|说明|
|-----------|--------------------------------------|
|@Component  |  加到类路径自动扫描 |
|@Controller  |  一个web的控制层，在Spring MVC中使用 |
|@Repository  |  数据管理/存储,企业级应用使用(Dao, DDD) |
|@Service  |  提供一个商业逻辑 - 一个无状态的切面 |

	
### Spring 中是通过注解描述来创建IoC对象。

	Spring Boot并不建议使用XML，而是通过注解的描述生成对象。
	
	在Spring 中把每一个需要管理的对象称为Spring Bean（简称Bean ），而Spring 管理这些Bean 的容器
	被我们称为SpringIoC 容器（或者简称IoC 容器）。IoC 容器需要具备两个基本的功能：
	
		* 通过描述管理Bean ， 包括发布和获取Bean; 
		
		* 通过描述完成Bean 之间的依赖关系。
		
### 所有的IoC 容器都需要实现接口 BeanFactory ，它是一个顶级容器接口。
	
### ApplicationContext 是 BeanFactory 的子接口之一

	是最为重要的接口设计，大部分Spring IoC 容器是 ApplicationContext 接口的实现类。
	
### AnnotationConfigApplicationContext 是一个基于注解的IoC 容器。
	
	* @Configuration 和 @Bean 注解例子
	
```java
package com.xyz.config;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

// ＠Configuration 代表这是一个Java 配置文件
// Spring 的容器会根据它来生成IoC 容器去装配Bean

@Configuration
public class AppConfig {

	// @Bean 代表将initUser 方法返回的POJO 装配到IoC 容器中，而其属性name 定义这个Bean 的名称
	// 如果没有配置它，则将方法名称“initUser”作为Bean的名称保存到Spring IoC容器中。
	
	@Bean(name = "user")
	public User initUser () {
		User user= new User();
		user.setId (lL);
		user.setUserName ("user_name_l");
		user.setNote("note_l");
		return user;
	}
}
```	

	* 构建自己的IoC 容器
	
```java
package com.xyz.config ;
import org.apache.log4j.Logger;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;
import com.xyz.User;

public class IoCTest {

	private static Logger log = Logger.getLogger(IoCTest.class);

	public static void main (String [] args) {
		// 构建自己的 IoC 容器
		ApplicationContext ctx = new AnnotationConfigApplicationContext(AppConfig.class);
		User user= ctx.getBean(User.class);
		log.info (user.getid()) ;
	}
}

// AppConfig 传递给Annotat ionConfigApp li cationContext 的构造方法，这样它就能够读取配置了。
// 然后将配置里面的Bean 装配到IoC 容器中，于是可以使用getBean 方法获取对应的POJO
```
	
## 1. 装配你的Bean

	如果一个个的Bean 使用注解 @Bean 注入Spring loC 容器中，那将是一件很麻烦的事情。
	对于扫描装配而言使用的注解是 @Component 和 @ComponentScan。

	> @Bean 是标明“哪个方法”的返回值被扫描进入Spring IoC 容器

	> @Component 是标明“哪个类”被扫描进入Spring IoC 容器

	> @ComponentScan 是标明采用何种策略去扫描装配Bean。

```java
package com.xyz.config;

// 在config包里面Bean类，加上 @Component 注解
@Component ("user")
public class User {
	@Value ("1")
	private Long id;
	@Value ("user_name_1")
	private String userName ;
	@Value ("note_1")
	private String note;
	/**setter and getter **/
}
```
	其中配置的“user＂则是作为 Bean 的名称，当然你也可以不配置这个字符串
	那么IoC 容器就会把类名第一个字母作为小写，其他不变作为Bean 名称放入到IoC 容器中
	注解@Value 则是指定具体的值，使得Spring IoC 给予对应的属性注入对应的值。
	
### 为了让Spring IoC 容器装配这个 @Component 注解类， 需要改造配置类 AppConfig
	
```java
package com.xyz.config ;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration ;

// 它只会扫描类 AppConfig 所在的当前包和其子包
// 注解可以修改为其它扫描路径(和排除)：
// @ComponentScan("com.springboot.example.*")
// 或
// @ComponentScan(basePackages = {"com.springboot.example.pojo"})
// 或
// @ComponentScan(basePackageClasses = {User.class}, excludeFilters = {@Filter(classes = {Service.class})})

@Configuration
@ComponentScan
public class AppConfig {
	// ......
}
```
	
## 2. 自定义第三方Bean

	把第三方包的类对象也放入到Spring IoC 容器中，这时 @Bean 注解就可以发挥作用了。

```xml
<!-- 在 pom.xml 中定义DBCP 数据源 -->
<dependency>
	<groupid>org.apache.commons</groupid>
	<artifactid>commons-dbcp2</artifactid>
</dependency>
<dependency>
	<groupid>mysql</groupid>
	<artifactid>mysql-connector-java</artifactid>
</dependency>
```	

	* 使用DBCP 生成第三方数据源

```java
// 这里通过@Bean 定义了其配置项name 为“dataSource”，那么Spring 就会把它返回的对象用名称“dataSource” 保存在IoC 容器中。
// 当然， 你也可以不填写这个名称，那么它就会用你的方法名称作为Bean 名称保存到IoC 容器中。

@Bean(name = "dataSource")
public DataSource getDataSource () {

	Properties props= new Properties();
	props.setProperty("driver"，"com.mysql.jdbc.Driver");
	props.setProperty("url","jdbc:mysql://localhost:3306/chapter3");
	props.setProperty("username","root");
	props.setProperty("password","123456");
	DataSourcedataSource = null;

	try{
		dataSource = BasicDataSourceFactory.createDataSource(props);
	}catch(Exceptione){
		e.printStackTrace();
	}
	return dataSource;
}
```
	
## 3. 依赖注入： @Autowired

	例如，人类（Person）有时候利用 <---- 动物(Animal）去完成一些事情，
	比方说狗（Dog）是用来看门的，猫（Cat）是用来抓老鼠的，鹦鹉（Parrot)是用来迎客的…

```java
/* 人类接口 */
package com.xyz.pojo.definition ;

public interface Person {
	// 使用动物服务
	public void service();
	// 设置动物
	public void setAnimal(Animal animal);

}

/* 动物接口 */
package com.xyz.pojo.definition ;

public interface Animal {
	public void use();
}
```
### @Autowired 实现依赖注入

	它注入的机制最基本的一条是根据类型（by type），或者根据名称（by name）。

```java
/* 人类实现类 */
package com.xyz.pojo;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import com.xyz.pojo.definition.Animal;
import com.zyz.pojo.definition.Person;

@Component
public class BussinessPerson implements Person {

	@Autowired
	private Animal animal = null;

	@Override
	public void service() {
		this.animal.use();
	}

	@Override
	public void setAnimal(Animal animal) {
		this.animal = animal;
	}
}

/* 狗，动物的实现类 */
package com.xyz.pojo;
import org.springframework.stereotype.Component;
import com.xyz.po]o.definition.Animal;

@Component
public class Dog implements Animal {
	@Override
	public void use () {
		System.out.println("狗【"＋Dog.class.getSimpleName（）＋ "】是看门用的。");
	}
}
```

### 消除歧义性： @Primary 和 Quelifier

	BussinessPerson 类只是定义了一个动物属性（Animal） ，而我们却有两个动物， 一个狗， 一个猫？

```java
// @Primary，它是一个修改优先权的注解，当我们有猫有狗的时候，假设这次需要使用猫， 
// 那么只需要在猫类的定义上加入@Primarγ 就可以了。
@Component
@Primary
public class Cat implements Animal {
	@Override
	public void use () {
		System.out.println("猫【"＋Dog.class.getSimpleName（）＋ "】抓老鼠。");
	}
}

// @Quelifier(Bean标识) 与 @Autowired 组合使用，通过这个就可以消除歧义性了。
public class BussinessPerson implements Person {

	@Autowired
	@Quelifier("dog")
	private Animal animal = null;

	// ......
}

```

### 带有参数的构造方法类的装配

```java
public class BussinessPerson implements Person {

	private Animal animal = null;

	// 如果没有多种 Animal 实例Bean存在， @Qualifier 可以删除
	public BussinessPerson(@Autowired @Qualifier("dog") Animal animal) {
		this.animal = animal;
	}

	// ......
}

```

## 4. 生命周期

	大致分为Bean 定义、Bean 的初始化、Bean 的生存期和Bean 的销毁4 个部分。
	
	1. 资源定位： 通过@ComponentScan定义所有需要扫描的包
	2. Bean定义： 将Bean的定义保存到 BeanDefinition 的实例中
	3. 发布Bean定义： IoC容器装载Bean定义
	4. 实例化： 创建Bean的实例对象
	5. 依赖注入（DI)： 例如，使用@Autowired注入的各类资源

```java
// 在配置类AppConfig 的 @ComponentScan 中加入 lazyInit 配置，默认 false 不进行延迟初始化
// 它修改为了延迟初始化， Spring 并不会在发布Bean 定义后马上为我们完成实例化和依赖注入。
@ComponentScan (basePackages = "com.xyz.pojo.*", lazyInit = true)
```

	* 加入生命周期接口和自定义	
	
```java
package com.xyz,pojo;

// 实现生命周期中单个Bean 可以实现的所有接口，
@Component
public class BussinessPerson implements Person, BeanNameAware, 
		BeanFactoryAware, ApplicationContextAware, InitializingBean, DisposableBean {

	private Animal animal = null;

	@Override
	public void service () { this.animal.use(); }

	@Override
	@Autowired
	@Qualifier("dog")
	public void setAnimal(Animal animal) {
		System.out.println("延时依赖注入");
		this.animal = animal;
	}

	@Override
	public void setBeanName(String beanName) {
		System.out.println("【" + this.getClass().getSimpleName()+"】重写 BeanNameAware 的 setBeanName");
	}

	@Override
	public void setBeanFactory(BeanFactory beanFactory) throw BeanException {
		System.out.println("【" + this.getClass().getSimpleName()+"】重写 BeanFactoryAware 的 setBeanFactory");
	}

	@Override
	public void setApplicationContext(ApplicationContext applicationContext) throw BeanException {
		System.out.println("【" + this.getClass().getSimpleName()+"】重写 ApplicationContextAware 的 setApplicationContext");
	}

	@Override
	public void afterPropertiesSet() throw Exception {
		System.out.println("【" + this.getClass().getSimpleName()+"】重写 InitializingBean 的 afterPropertiesSet");
	}

	@PostConstruct
	public void init() {
		System.out.println("【" + this.getClass().getSimpleName()+"】注解@PostConstruct定义的自定义初始化方法");
	}

	@PreDestroy
	public void destroyCustom() {
		System.out.println("【" + this.getClass().getSimpleName()+"】注解@PreDestroy定义的自定义销毁化方法");
	}

	@Override
	public void destroy() throw Exception {
		System.out.println("【" + this.getClass().getSimpleName()+"】重写 DisposableBean 方法");
	}
}

```

	* 后置Bean 初始化器，对所有的Bean 有效

	* 可以使用注解＠Bean 来配置自定义初始化和销毁方法： @Bean(initMethod="init", destroyMethod="destroy")

```java
package com.xyz.life;

@Component
public class BeanPostProcessorExample implements BeanPostProcessor {
	@Override
	public Object postProcessBeforeInitialization(Object bean , String beanName) throws BeansException {
			System.out.println("BeanPostProcessor 调用 postProcessBeforeinitialization 方法，参数【" 
			                  + bean.getClas().getSimpleName() + "】【" + beanName + "】");
			retur bean;
	}

	@Override
	public Object postProcessAfterInitializatiion(Object bean, String beanName) throw BeansException {
		System.out.println("BeanPostProcessor 调用 postProcessAfterInitializatiion 方法，参数【" 
		                    + bean.getClass().getSimpleName() + "】【" + beanName + "】");
		retur bean ;
	}
}
```
	
## 5. 使用属性文件
	
	在Spring Boot 中使用属性文件，默认为我们准备的 application.properties ，也可以使用自定义的配置文件。

```xml
<!-- 在 POM 中加入属性文件依赖 ， 就可以直接使用 application.properties 文件 -->
<dependency>
	<groupId>org.springframework.boot</groupId>
	<artifactId>spring-boot-configuration-processor</artifactId>
	<optional>true</optional>
</dependency>
```

	* 使用属性配置
	

```java
package com.xyz.pojo ;

/*
	可以通过@Value 注解， 使用 ${...} 这样的占位符读取配置在属性文件的内容。
	这里的 @Value 注解，既可以加载属性， 也可以加在方法上。
	使用 @ConfigurationProperties 通过它使得配置上有所减少，去掉 @Value 的配置。
	注解 @ConfigurationProperties 中配置的字符串 database ，
	将与POJO 的属性名称组成属性的全限定名去配置文件里查找，这样就能将对应的属性读入到POJO 当中。
*/
@Component
@ConfigurationProperties("database")  // 使用它需要去掉 @Value 注解
@PropertySource(value={"classpath:jdbc.properties"}, ignoreResourceNotFound=true)  // 选择配置文件 jdbc.properties
public class DataBaseProperties {

	@Value("${database.driverName}")
	private String driverName = null;

	@Value("${database.url}")
	private String url = null;

	private String username = null;

	private String password = null ;

	public void setDriverName (String driverName) {
		this.driverName = driverName;
	}
	public void setUrl (String url) {
		this.url = url;
	}

	@Value("${database.username}")
	public void setUsername (String username) {
		this.username = username;
	}

	@Value("${database.password}")
	public void setPassword (String password) {
		this.password= password;
	}
	/** getters **/
}
```
	
## 6. 条件装配Bean， 注解 @Conditional
	
```java
// ......
	@Bean(name="dataSource", destroyMethod="close")
	@Conditional(DatabaseConditional.class)
	public DataSource getDataSource(
			@Value("${database.driverName}") String driver,
			@Value("${database.url}") String url,
			@Value("${database.username}") String username,
			@Value("${database.password}") String password,
		) {

		Properties props = new Properties();
		props.setProperty("driver", driver);
		props.setProperty("url", url);
		props.setProperty("username", username);
		props.setProperty("password", password);
		DataSource dataSource = null;
		try {
			dataSource = BasicDataSourceFactory.createDataSource(props);
		} catch (Exception e) {
			e.printStackTrace();
		}
		return dataSource;
	}

// ......
// 以上代码加入了 @Conditional 注解， 井且配置了类 DatabaseConditional ，那么这个类就必须实现Condition 接口。
// 定义初始化数据库的条件
public class DatabaseConditional implements Condition {
		/* 
		* 数据库装配条件
		* @param context 条件上下文
		* @param metadata 注释类型的元数据
		* @return true 装配Bean ，否则不装配
		*/
	@Override
	public boolean matches(ConditionContext context, AnnotatedTypeMetadata metadata) {
		// 取出环境配置
		Environment env = context.getEnvironment();
		// 判断属性文件是否存在对应的数据库配置
		return env.containsProperty("database.driverName")
				&& env.containsProperty("database.url")
				&& env.containsProperty("database.username")
				&& env.containsProperty("database.password");
	}
}
```

## 7. Bean 的作用域

| 作用域类型 | 使用范围 | 作用域描述 |
|-----------|-------------------|-------------------------------|
| singleton | 所有 Spring 应用 | 默认值， loC 容器只存在单例 |
| prototype | 所有 Spring 应用 | 每当从IoC 容器中取出一个Bean ，则创建一个新的Bean |
| session | Spring Web 应用 | HTTP 会话 |
| application | Spring Web 应用 | Web 工程生命周期 |
| request | Spring Web 应用 | Web 工程单次请求 (request) |
| globalSession | Spring Web 应用 | 在一个全局的HTTP Session 中， 一个Bean 定义对应一个实例。实践中基本不使用 |

```java

	// ConfigurableBeanFactory 只能提供单例(SCOPE_ SINGLETON)和原型(SCOPE_PROTOTYPE) 两种作用域供选择
  @Scope(ConfigurableBeanFactory.SCOPE_PROTOTYPE)
	public class ScopeBeanA { 
		// ......
	}

	// WebApplicationContext 能提供 请求(SCOPE REQUEST)、会话(SCOPE_SESSION) 和应用(SCOPEAPPLICATION)
	@Scope(WebApplicationContext.SCOPE_REQUEST)
	public class ScopeBeanB { 
		// ......
	}

```

## 8. 部署环境的切换， @Profile 定义配置环境

	启动Profile机制参数: JAVA_OPTS="-Dspring.profiles.active=dev"
	若不配置启动参数，被@Profile 标注的Bean 将不会被Spring 装配到 IoC 容器中。
	Spring 优先先判定 spring.profiles.active 配置后， 不存在再去查找spring.profiles.default 的配置

	按照 Spring Boot 的规则，假设把选项 -Dspring.profiles.active 配置的值记为 {profile}
	则它会用 application-{profile}.properties 文件去代替原来默认的 application.properties 文件

```java
// 配置开发环境
@Bean(name="dataSource", destroyMethod="close")
@Profile("dev")
public DataSource getDevDataSource() {
	Properties props = new Properties();
	props.setProperty("driver"，"com.mysql.jdbc.Driver");
	props.setProperty("url","jdbc:mysql://localhost:3306/dev_example");
	props.setProperty("username","root");
	props.setProperty("password","123456");
	DataSourcedataSource = null;

	try{
		dataSource = BasicDataSourceFactory.createDataSource(props);
	}catch(Exceptione){
		e.printStackTrace();
	}
	return dataSource;
}

// 配置测试环境
@Bean(name="dataSource", destroyMethod="close")
@Profile("test")
public DataSource getTestDataSource() {
	// ......
	return dataSource;
}
```

## 9. 引入 XML 配置Bean

	使用的是注解 @ImportResource ，通过它可以引入对应的XML 文件，用以加载Bean。

	使用XML配置的Bean，要求所在的包并不在 @ComponentScan 定义的扫描包范围内，而且没有标注 @Component

	* 被配置到 Bean XML 上的 Bean 类

```java
package com.my.pojo;
public class MyHelloWord { /* ...... */ }
```

	* 扫描 XML Beans 配置

```xml
<!-- spring-beans.xml -->
<?xml version="1.0" encoding="UTF-8"?>
<beans xmlns="http://www.springframework.org/schema/beans" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
       xsi:schemaLocation="http://www.springframework.org/schema/beans http://www.springframework.org/schema/beans/spring-beans.xsd">
    <!-- 自定义配置bean，首字母小写  -->
    <bean id="myHelloWord" class="com.my.pojo.MyHelloWord"/>
</beans>
```
	* 使用注解 @ImportResource 装配定义 Beans 的 XML文件

```java
package com.my;

@Configuration
@ComponentScan(basePackages="com.my.*")
@ImportResource(value={"classpath:spring-beans.xml"})
public class AppConfig {
	// ......
}
```

## 10. 使用 Spring EL 表达式 #{...} 内计算

  数字型的可以使用 == 比较符，如果是字符串型的可以使用 eq 比较符

```java
// ${...} 代表占位符，会读取上下文的属性值装配到属性中
@Value("${database.driverName}")

// beanName 是Spring IoC 容器Bean 的名称
@Value ("#{beanName.str}" )

// 采用#{...}代表启用Spring 表达式，它将具有运算的功能
// T(...）代表的是引入类，System 是默认加载的包，可以不必写全名，如果是其他包，则需要写出全名才能引用类
@Value("#{T(System).currentTimeMillis()}")

// ? 表达式，判断这个属性是否为空。不为空才会执行toUppercase方法，并把属性转换为大写，赋予当前属性。
@Value("#{beanName.str?.toUpperCase()}" )

```

## 11. Spring AOP 约定编程

	任何AOP编程，首先要确定的是在什么地方需要AOP ，也就是需要确定连接点（在Spring 中就是什么类的什么方法）的问题。

```java
// 简易接口 HelloService
package com.mytest.service;
public interface HelloService {
	public void sayHello(String name);
}

// HelloService 实现类 HelloServiceImpl
package com.mytest.service.impl;
public class HelloServiceImpl implements HelloService {
	@Override
	public void sayHello(String name){
		if (name == null || name.trim() == "") {
			throw new RuntimeException("parameter is null");
		}
		System.out.println("hello " + name);
	}
}
```

	* 拦截器接口

```java
package com.mytest.intercept;
import java.lang.reflect.InvocationTargetException;
import com.mytest.invoke.Invocation;

public interface Interceptor {
	// 事前方法
	public boolean before() ;
	// 事后方法
	public void after ();
	/**
	* 取代原有事件方法
	* @param invocation 回调参数，可以通过它的 proceed 方法， 回调原有事件
	* @return 原有事件返回对象
	* @throws InvocationTargetException
	* @throws IllegalAccessException
	*/
	public Object around(Invocation invocation) throws InvocationTargetException, IllegalAccessException;
	// 是否返回方法。事件没有发生异常执行
	public void afterReturning() ;
	// 事后异常方法， 当事件发生异常后执行
	public void afterThrowing ();
	// 是否使用 around 方法取代原有方法
	boolean useAround() ;
}
```
	* Invocation 类

```java
package com.mytest.invoke;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;

public class Invocation {
	private Object[] params;
	private Method method ;
	private Object target ;

	public Invocation(Object target, Method method,Object[] params) {
		this.target=target;
		this.method=method;
		this.params=params;
	}
	// 反射方法
	public Object proceed() throws InvocationTargetException , IllegalAccessException {
		return method.invoke(target, params);
	}
	/**** setter and getter ****/
}
```

	* 开发自己的拦截器

```java
package com.mytest.intercept;
import java.lang.reflect.InvocationTargetException;
import com.mytest.invoke.Invocation;

public class MyInterceptor implements Interceptor {
	@Override // 事前方法
	public boolean before() {
		System.out.println("before ......");
		return true;
	}

	@Override // 事后方法
	public void after () {
		System.out.println("after ......");
	}

	@Override 
	public Object around(Invocation invocation) throws 
			InvocationTargetException, IllegalAccessException {
		System.out.println("around before ......");
		Object obj = invocation.proceed();
		System.out.println("around after ......");
		return obj;
	}

	@Override 
	public void afterReturning() {
		System.out.println("afterReturning ......");
	}

	@Override 
	public void afterThrowing () {
		System.out.println("afterThrowing ......");
	}

	@Override 
	public boolean useAround() {
		return true;
	}
}

```

### ProxyBean 的实现

	如何将服务类和拦截方法织入对应的流程，是ProxyBean 要实现的功能。首先要理解动态代理模式。

```java
// 在JDK 中，提供了类Proxy 的静态方法 newProxyInstance，其内容具体如下：
public static Object newProxyInstance(ClassLoader classLoader,
		Class <?>[] interfaces, InvocationHandler invocationHandler) throws IllegalArguπ1entException;

/*
	生成一个代理对象（proxy），它有3 个参数：
	• classLoader  我加载器;
	• interfaces 绑定的接口，也就是把代理对象绑定到哪些接口下， 可以是多个；
	• invocationHandler  绑定代理对象逻辑实现。
*/
```
```java
// ProxyBean 的实现
package com.mytest.proxy;

public class ProxyBean implements invocationHandler {

	private Object target = null;
	private Interceptor interceptor = null

	/**
	 * 绑定代理对象
	 * @param target 被代理的对象
	 * @param interceptor 拦截器
	 * @return 代理的对象
	 */
	 public static Object getProxyBean(Object target, Interceptor interceptor) {
		 ProxyBean proxyBean = new ProxyBean();
		 // 保存被代理对象
		 proxyBean.target = target;
		 // 保存拦截器
		 proxyBean.interceptor = interceptor;
		 // 生成代理对象
		 Object proxy = Proxy.newProxyInstance(
						target.getClass().getClassLoader(), 
						target.getClass().getInterfaces(),
						ProxyBean
					);
		// 返回代理对象
		return proxy;
	 }

	/**
	 * 处理代理对象方法逻辑
	 * @param proxy 代理对象
	 * @param method 当前方法
	 * @param args 运行参数
	 * @return 方法调用结果
	 * @throws Throwable 异常
	 */
	 @Override
	 public Object invoke(Object proxy, Method method, Object[] args) {
		 // 异常标识
		 boolean exceptionFlag = false;
		 Invocation invocation = new Invocation(target, method, args);
		 Object retObj = null;
		 try {
			 if (this.interceptor.before()) {
				 retObj = this.interceptor.around(invocation);
			 } else {
				 retObj = method.invoke(target, args);
			 }
		 } catch(Exception ex) {
			 // 产生异常
			 exceptionFlag = true;
		 }
		 this.interceptor.after();
		 if (exceptionFlag) {
			 this.interceptor.afterThrowing();
		 } else {
			 this.interceptor.afterReturning();
			 return retObj;
		 }
		 return null;
	 }
}

```

	* 测试约定流程

```java
private static void testProxy() {
	HelloService helloService = new HelloServiceImpl();
	// 按约定获取 proxy
	HellowService proxy = (HelloService) ProxyBean.getProxyBean(helloService, new MyInterceptor());
	prox.sayHello("ZhangSan");
}

```

### AOP 的概念，使用 @AspectJ 注解，只能对方法进行拦截

	使用Spring AOP 可以处理一些无法使用OOP 实现的业务逻辑。
	其次，通过约定，可以将一些业务逻辑织入流程中，并且可以将一些通用的逻辑抽取出来，
	然后给予默认实现，这样你只需要完成部分的功能就可以了,
	这样做可以使得开发者的代码更加简短，同时可维护性也得到提高

	AOP 最为典型的应用实际就是数据库事务的管控。例如， 当我们需要保存一个用户时，可能要连同它的角色信息一并保存到数据库中。

	@Transactional 注解，表明该方法需要事务运行，实现了数据库资源的打开和关闭、事务的回漆和提交。

#### AOP 术语和流程
	
	1. 连接点（joinpoint）：对应的是具体被拦截的对象，因为Spring只能支持方法，所以被拦截的对象往往就是指特定的方法。
		例如，我们前面提到的HelloServiceimpl的sayHello方法就是一个连接点，AOP将通过动态代理技术把它织入对应的流程中。

	2. 切点（point cut）：有时候，我们的切面不单单应用于单个方法，也可能是多个类的不同方法，这时，可以通过正则式和指示器的规则去定义，从而适配连接点。切点就是提供这样一个功能的概念。

	3. 通知（advice）：就是按照约定的流程下的方法，分为前置通知（before advice）、后置通知（after advice）、环绕通知（around advice）、事后返回通知（afterRetuming advice）和异常通知（afterThrowing advice），它会根据约定织入流程中，需要弄明白它们在流程中的顺序和运行的条件。

	4. 目标对象（target）：即被代理对象，例如，约定编程中的 HelloServiceImpl 实例就是一个目标对象，它被代理了。

	5. 引入（introduction）：是指引入新的类和其方法，增强现有Bean的功能。

	6. 织入（weaving）：它是一个通过动态代理技术，为原有服务对象生成代理对象，然后将与切点定义匹配的连接点拦截，并按约定将各类通知织入约定流程的过程。

	7. 切面（aspect）：是一个可以定义切点、各类通知和引入的内容，Spring AOP将通过它的信息来增强Bean的功能或者将对应的方法织入流程。
	
### AOP 开发详解

	1. 确定连接点

```java
// 用户服务接口
package com.xyz.aspect.service;
import com.xyz.pojo.User;
public interface UserService{
	public void printUser(User user);
}

// 用户服务接口实现类
package com.xyz.aspect.service.impl;
@Service
public class UserServiceImpl implements UserService{
	@Override
	public void printUser(User user){    // 以此方法为连接点
		if (user == null){
			throw new RuntimeException("检查用户参数是否为空......");
		}
		System.out.println("Id = " + user.getId());
		System.out.println("\tusername = " + user.getUsername());
	}
}
```
	
	2. 开发切面

```java
// 定义切面
package com.xyz.aspect;
@Aspect
@Order(1)  // 定义多切面执行顺序，或者实现 Ordered 接口
public class MyAspect implements Ordered {

	// 使用 @Order 或者 Ordered 实现，任选一种来确定多切面执行顺序
	@Override
	public int getOrder() {
		return 1;
	}

	/*
		注解 @Pointcut 来定义切点，它标注在方法pointCut 上，
		则在后面的通知注解中就可以使用方法名称来定义了。
	*/
	@Pointcut("execution(* com.xyz.aspect.service.impl.UserServiceImpl.printUser(..))")
	public void pointCut() { }

	@Before("pointCut() && args(user)")  // 获取通知参数
	public boolean before(JoinPoint point, User user) {
		Object[] args = point.getArgs();
		System.out.println("before ......");
		return true;
	}

	@After("pointCut()")
	public void after () {
		System.out.println("after ......");
	}

	@Around("pointCut()")  // 环境通知
	public void around(ProceedingJoinPoint jp) throws Throwable {
		// ProceedingJoinPoint 对象有一个proceed 方法，可以回调原有目标对象的方法。
		System.out.println("around before ......");
		// 回调目标对象的原有方法
		jp.proceed();
		System.out.println("around after ......");
	}

	@AfterReturning("pointCut()") 
	public void afterReturning() {
		System.out.println("afterReturning ......");
	}

	@AfterThrowing("pointCut()")
	public void afterThrowing () {
		System.out.println("afterThrowing ......");
	}
}
/**
• execution 表示在执行的时候，拦截里面的正则匹配的方法;
• * 号, 表示任意返回类型的方法;
• com.xyz.aspect.service.impl.UserServiceImpl 指定目标对象的全限定名称;
• printUser 指定目标对象的方法;
• (..) 表示任意参数进行匹配;
*/
```


	* 环绕通知（Around） 是所有通知中最为强大的通知，强大也意味着难以控制。

	* 使用它的场景是在你需要大幅度修改原有目标对象的服务逻辑时， 否则都尽量使用其他的通知。

	* 环绕通知是一个取代原有目标对象方法的通知， 当然它也提供了回调原有目标对象方法的能力。

	
#### AspectJ 关于Spring AOP 切点的指示器

|项目类型|描述|
|-------|------------------------------------|
|arg()  |  限定连接点方法参数 |
|@args()  |  通过连接点方法参数上的注解进行限定|
|execution()  |  用于匹配是连接点的执行方法|
|this()  |  限制连接点匹配AOP 代理Bean 引用为指定的类型|
|target  |  目标对象（即被代理对象）|
|@target()  |  限制目标对象的配置了指定的注解|
|within  |  限制连接点匹配指定的类型|
|@within()  |  限定连接点带有匹配注解类型|
|@annotation()  |  限定带有指定注解的连接点|

表达式： execution(* com.*.UserServiceImpl.printUser(..) && bean('userServiceImpl'))

表达式中的&&代表“并且”的意思，而bean中定义代表对Spring Bean 容器名称的限定，这样就限定具体的类了


	3. AOP使用

```java
// 用户控制器
package com.xyz.aspect.controller;

@Controller  // 定义控制器
@RequestMapping("/user") // 定义请求路径
public class UserController {
	
	@Autowired  // 注入用户服务
	private UserService userService = null;

	@RequestMapping("/print") // 定义请求路径
	@ResponseBody  // 转换为JSON
	public User printUser(Long id, String userName, String note) {
		User user = new User();
		user.setId(id);
		user.setUserName(userName);
		user.setNote(note);
		userService.printUser(user); // 若 user == null, 则执行 afterthrowing 方法
		return user; // 加入断点
	}
}

// Spring Boot 配置启动文件
package com.xyz.main;

// 指定扫描包
@SpringBootApplication(scanBasePackage="com.xyz.aspect")
public class MainApplication {
	// 定义切面
	@Bean(name="myAspect")
	public MyAspect initMyAspect() {
		return new MyAspect();
	}

	// 同一个点的多切面引入
	@Bean(name="myAspect2")
	public MyAspect2 initMyAspect() {
		return new MyAspect2();
	}

	// 启动切面
	public static void main(String[] args){
		SpringApplication.run(MainApplication.class, args);
	}
}
```

	4. AOP 引入, Spring 允许增强接口的功能

```java
// 用户检测的接口UserValidator
package com.xyz.aspect.validator;
import com.xyz.pojo.User;

public interface UserValidator {
	// 检测用户对象是否为空
	public boolean validate(User user);
}

/*
  UserValidator 的实现类
*/
package com.xyz.aspect.validator.impl;

public class UserValidatorImpl implements UserValidator {
	@Override
	public boolean validate(User user) {
		System.out.println("引入新的接口" + UserValidator.class.getsSmpleName());
		return user != null;
	}
}
```
```java
/*
	在JDK 动态代理中下挂的两个接口， 于是我们可以将这个代理对象通过这两个接口相互转换， 
	然后调度其对应的方法， 这就是引入的原理。

	注解 @DeclareParents， 它的作用是引入新的类来增强服务， 
	它有两个必须配置的属性value 和defaultlmpl 。
*/
@Aspect  // 在 MyAspect引入新的接口
public class MyAspect {
	@DeclareParents(value= "com.xyz.aspect.service.impl.UserServiceImpl+",
		        defaultImpl=UserValidatorImpl.class)
	public UserValidator userValidator;
	// ......
}
/*
• value：指向你要增强功能的目标对象， 这里是要增强UserServicelmpl 对象， 
	因此可以看到配置为com.xyz.aspect.service.impl.UserServiceImpl+。

• defaultImpl：引入增强功能的类， 这里配置为UserValid ator lmpl ，
	用来提供校验用户是否为空的功能。
*/
```
```java
// 测试引入的验证器
@RequestMapping("/vp") // 定义请求
@ResponseBody    // 返回JSON
public User validateAndPrint(Long id, String userName, String note) {
	User user = new User();
	user.setId(id);
	user.setUserName(userName);
	user.setNote(note);
	// 强制转换
	UserValidator userValidator = (UserValidator)userService;
	// 验证用户是否为空
	if (userValidator.validate(user)) {
		userService.printUser(user);
	}
	return user;
}
```

	5. AOP 织入是一个生成动态代理对象并且将切面和目标对象方法编织成为约定流程的过程。
		
	我们都是采用接口＋实现类的模式， 这是Spring 推荐的方式。
	但是对于是否拥有接口则不是Spring AOP 的强制要求， 对于动态代理的也有多种实现方式， 我业界比较流行的还有CGLIB 、Javassist 、ASM 等。Spring 采用了JDK 和CGLIB ， 对于JDK 而言，它是要求被代理的目标对象必须拥有接口， 而对于CGLIB 则不做要求。因此在默认的情况下， Spring 会按照这样的一条规则处理，即当你需要使用AOP 的类拥有接口时， 它会以JDK 动态代理运行，否则以CGLIB 运行。


## 访问数据库

	* POM.XML中添加 DBCP 数据源依赖，commons-dbcp2

```conf
# application.properties 文件中数据源连接参数
spring.datasource.url=jdbc:mysql://localhost:3306/mydatabase
spring.datasource.username=root
spring.datasource.password=l23456
#spring .datasource.driver-class-name=com.mysql . jdbc Driver
# 指定数据库连接池的类型
spring.datasource.type=org.apache.commonsdbcp2.BasicDataSource
```

	* 监测数据库连接池类型

```java
package com.xyz.db;

@Component
// 实现 Sprint Bean 生命周期接口 ApplicationContextAware 
public class DataSourceShow implements ApplicationContextAware {
	ApplicationContext applicationContext = null;

	// Sprint 容器会自动调用这个方法，注入 Spring IoC 容器
	@Override
	public void setApplicationContext(ApplicationContext applicationContext)
			throws BeanException {
		this.applicationContext = applicationContext;
		DataSource dataSource = applicationContext.getBean(DataSource.class);
		System.out.println(dataSource.getClass().getName());		
	}
}
```

### JPA (Java Persistence API) 用 (Hibernate）操作数据

```java
package com.xyz.pojo;

@Entity(name="user")  // 标明一个实体类
@Table(name="t_user")  // 定义映射的表
public class User {
	@Id   // 标明主键
	@GenerateValue(strategy=GenerationType.IDENTITY) // 主键策略，递增
	private Long id = null;

	@Column(name="user_name") // 定义属性和表的映射关系
	private String userName = null;

	private Strint note = null;

	@Convert(converter=SexConverter.class) // 定义转换器
	private SexEnum sex = null;
	/** setter and getter **/
}

/*------------- 性别转换器 ---------------*/
public class SexConverter implements AttributeConverter<SexEnum, Integer> {
	// 将枚举转换为数据库列
	@Override
	public Integer converterToDatabaseColumn(SexEnum sex) {
		return sex.getId();
	}
	// 将数据库列转换为枚举
	@Override
	public SexEnum converterToEntityAttibute(Integer id){
		return SexEnum.getEnumById(id);
	}
}

/* --------- Spring Boot 启动文件配置 -------- */
// 定义 Spring boot 扫描路径
@SpringBootApplication(scanBasePackage="com.xyz.example")
// 定义JPA 接口扫描包路径
@EnableJpaRepositories(basePackage="com.xyz.example.dao")
// 定义实体 Bean 扫描包路径
@EntityScan(basePackages="com.xyz.example.pojo")
public class MainApplication {
	public static void main(String[] args) {
		SpringApplication.run(MainApplication.class, args);
	}
}
```

### 整合 MyBatis 框架

	My Batis 的配置文件包括两个大的部分， 一是基础配置文件， 一个是映射文件。
	
	My Batis 是一个基于 SqlSessionFactory 构建的框架。在MyBatis 应用的生命周期中理当只存在一个 SqlSessionFactory 对象，并且往往会使用单例模式。而构建 SqlSessionFactory 是通过配置类（Configuration）来完成的，因此对于mybatis-spring-boot-starter ，它会给予我们在配置文件（application.properties）进行 Configuration 配置的相关内容。

```xml
<!-- 引入 MyBatis 的 starter -->
<dependency>
	<groupId>org.mybatis.spring.boot</groupId>
	<artifactId>mybatis-spring-boot-starter</artifactId>
	<version>2.1.0</version>
</dependency>
```

#### MyBatis 可配置的内容

	•  properties （属性）： 属性文件在实际应用中一般采用Spring 进行配置，而不是MyBatis

	• settings（设置）：它的配置将改变MyBatis 的底层行为，可以配置映射规则，如自动映射和驼峰映射、执行器（Executor ）类型、缓存等内容，比较复杂，具体配置项可看 MyBatis在线参考。

	• typeAliases（类型别名）：使用类全限定名会比较长，所以MyBatis 会对常用的类提供默认的别名，此外还允许我们通过typeAliases 配置自定义的别名。

	• typeHandlers（类型处理器）： 这是MyBatis 的重要配置之一，在MyBatis 写入和读取数据库的过程中对于不同类型的数据（对于Java 是JavaType，对于数据库则是JdbcType ）进行自定义转换，在大部分的情况下我们不需要使用自定义的typeHandler ，因为在MyBatis 自身就已经定义了比较多的typeHandler, MyBatis 会自动识别javaTyp巳和jdbcType ，从而实现各种类型的转换。一般来说， typeHandler的使用集中在枚举类型上。

	• objectFactory（对象工厂）：这是一个在MyBatis 生成返回的POJO 时会调用的工厂类。一般我们使用MyBatis 默认提供的对象工厂类（DefaultObjectFactory）就可以了，而不需要任何配置。

	• mappers（映射器）： 最核心的组件，它提供SQL和POJO 映射关系， 这是MyBatis开发的核心。

	• plugins（插件）：有时候也称为拦截器， 是MyBatis 最强大也是最危险的组件，它通过动态代理和责任链模式来完成，可以修改MyBatis 底层的实现功能。掌握它需要比较多的MyBatis知识。

	• environments（数据库环境）： 可以配置数据库连接内容和事务。一般这些交由Spring托管。

	• databaseIdProvider（数据库厂商标识）：允许MyBatis 配置多类型数据库支持， 不常用。

```java
/* ---------   在用户类使用MyBatis 别名  ---------- */
package com.xyz.example.pojo;

@Alias(value="user")  // Mybatis 指定别名
public class User {
	private Long id = null;
	private String userName = null;
	private String note = null;
	// 性别枚举，这里需要使用 typeHandler 进行转换
	private SexEnum sex = null;

	public User() { }
	/*** setter and getter ***/
}
```
```java
/* ---------   性别 typeHandler  ---------- */
package com.xyz.example.typehandler;
// 声明 JdbcType　为整数
@MappedJdbcTypes(JdbcType.INTEGER)
// 声明 JavaType 为 SexEnum
@MappedType(value=SexEnum.class)
public class SexTypeHandler extends BaseTypeHandler<SexEnum> {
	// 通过列名读取性别
	@Overrde
	public SexEnum getNullableResult(ResultSet rs, String col) throws SQLException {
		int sex = rs.getInt(col);
		if (sex != 1 && sex != 2) {
			return null;
		}
		return SexEnum.getEnumById(sex);
	}
	// 通过下标读取性别
	@Overrde
	public SexEnum getNullableResult(ResultSet rs, int idx) throws SQLException {
		int sex = rs.getInt(idx);
		if (sex != 1 && sex != 2) {
			return null;
		}
		return SexEnum.getEnumById(idx);
	}
	// 通过存储过程读取性别
	@Overrde
	public SexEnum getNullableResult(CallableStatement cs, int idx) throws SQLException {
		int sex = cs.getInt(idx);
		if (sex != 1 && sex != 2) {
			return null;
		}
		return SexEnum.getEnumById(idx);
	}
	// 设置非空性别参数
	@Overrde
	public void setNonNullParameter(PreparedStatement ps, int idx, SexEnum sex, JdbcType jdbcType) 
			throws SQLException {
		ps.setInt(idx, sex.getId());
	}

}
```
```xml
<!-- 用户映射文件 userMapper.xml -->
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="com.xyz.example.dao.MyBatisUserDao">
	<select id="getUser " parameterType="long" resultType="user" >
		SELECT id, user name AS userName, sex, note FROM t_user where id = #{id}
	</select>
</mapper>
```
```java
/* --------- 定义 MyBatis 操作接口 -------- */
package com.xyz.example.dao;
// @Repository 注解，扫描在加载 MyBatis 接口Bean 时需要，也可以使用注解 @Mapper
@Repository
public interface MyBatisUserDao {
	public User getUser(Long id);
}
```
```conf
# 配置 application.properties 中映射文件和扫描别名
# MyBatis 映射文件通配
mybatis.mapper-locations=classpath:com/xyz/example/mapper/*.xml
# MyBatis 扫描别名包，和主解 @Alias 联用
mybatis.type-aliases-package=com.xyz.example.pojo
# 配置 typeHandler 的扫描包
mybatis.type-handlers-package=com.xyz.example.typehandler
```

### Spring Boot 整合 MyBatis

	• MapperFactoryBean 是针对一个接口配置
	• MapperScannerConfigurer 提供扫描装配 MyBatis 的接口到 Spring IoC 容器中。
	• MyBatis 提供了注解 @MapperScan ，也能够所需的对应接口扫描装配到Spring IoC 容器中

```java
/* --------- 便用 MapperFactoryBean 装配MyBatis 接口 -------- */
@Autowired
SqlSessinFactory sqlSessinFactory = null;
// 定义一个 Mybatis 的 Mapper 接口
public MapperFactoryBean<MyBatisUserDao> initMyBatisUserDao() {
	MapperFactoryBean<MyBatisUserDao> bean = new MapperFactoryBean<>();
	bean.setMapperInterface(MyBatisUserDao.class);
	bean.setSqlSessionFactory(sqlSessionFactory);
	return bean;
}
```

```java
/* ------------ 使用MyBatis 接口  ---------- */
package com.xyz.example.service;
import com.xyz.example.pojo.User;
public interface MyBatisUserService {
	public User getUser(Long id);
}

/* ------------ 使用MyBatis 接口实现  ---------- */
package com.xyz.example.service.impl;
@Service
public class MyBatisUserServiceImpl implements MyBatisUserService {
	@Autowired
	private MyBatisUserDao myBatisUserDao = null;
	@Override
	public User getUser(Long id) {
		return myBatisUserDao.getUser(id);
	}
}

/* ------------ 使用控制器测试MyBatis 接口 ---------- */
package com.xyz.example.controller;
@Controller
@RequestMapping("/mybatis")
public class MyBatisController {
	@Autowired
	private MyBatisUserService myBatisUserService = null;

	@RequestMapping("/getUser")
	@ResponseBody
	public User getUser(Long id) {
		return myBatisUserService.getUser(id);
	}

}
```
```java
/* ------- 使用MapperScannerConfigurer 扫描装配MyBatis 接口 ----- */
// 配置 Mybatis 接口扫描，返回扫描器
@Bean
public MapperScannerConfigurer mapperScannerConfig() {
	// 定义扫描器实例
	MapperScannerConfigurer mapperScannerConfigurer = new MapperScannerConfigurer();
	// 加载 SqlSessionFactory, Spring boot 会自动生产， SqlSessionFactory 实例
	mapperScannerConfigurer.setSqlSessionFactoryBeanName("sqlSessionFactory");
	// 定义扫描的包
	mapperScannerConfigurer.setBasePackage("com.xyz.example.*");
	// 限定被标注 @Repository 的接口才被扫描
	mapperScannerConfigurer.setAnnotationClass(Repository.class);
	// 通过继承某个接口限制扫描，一般使用不多
	// mapperScannerConfigurer.setMarkerinterface( .. .... );
	return mapperScannerConfigurer;
}
```

#### 使用简单的 @MapperScan 定义扫描

```java
/* ----------- 定义 MyBatis 插件 ---------- */
package com.xyz.example.plugin;
// 定义拦截器签名
@Intercepts({
	@Singature(type = StatementHandler.class,
	method = "prepare",
	args = { Connection.class, Integer.class})})
public class MyPlugin implements Interceptor {
	Properties properties = null;
	// 拦截方法逻辑
	@Override
	public Object intercept(Invocation invocation) throws Throwable {
		System.out.println("插件拦截方法......");
		return invocation.proceed();
	}
	// 生成 Mybatis拦截器代理对象
	@Override
	public Object plugin(Object target) {
		return Plugin.wrap(target, this);
	}
	// 设置插件属性
	@Override
	public void setProperties(Properties properties) {
		return this.properties = properties;
	}
}

```
	在 application.properties 文件中配置
	mybatis.config location=classpath:mybatis/mybatis-config xml

```xml
<!-- MyBatis 配置文件C mybatis/mybatis-config .xml ) -->
<configuration>
	<plugins>
		<plugin interceptor= "com.xyz.example.plugin.MyPlugin" >
		<property name="keyl" value="valuel" />
		<property name="key2" value="value2" />
		<property name="key3" value="value3" />
		</plugiη 〉
	</plugins>
</configuration>
```

```java
package com.xyz.example;
// 定义 Spring Boot 扫描包路径
@SpringBootApplication(scanBasePackages = {"com.xyz.example"})
// 定义JPA 接口扫描包路径
@EnableJpaRepositories(basePackages = "com.xyz.example.dao")
// 定义实体Bean 扫描包路径
@EntityScan(basePackages = "com.xyz.example.pojo")
// 定义 MyBatis 的扫描
@MapperScan(
	// 指定扫描包
	basePackages = "com.xyz.example.*",
	// 指定 SqlSessionFactory，如果 sqlSessionTemplate 被指定,则作废
	sqlSessionFactoryRef = "sqlSessionFactory",
	// 指定 sqlSessionTemplate ，将忽略 sqlSessionFactory 的配置
	sqlSessionTemplateRef = "sqlSessionTemplate",
	// markerinterface = Class.class,    // 限定扫描接口，不常用
	annotationClass = Repository.class
)
public class MainApplication {
	// SqlSessionFactory 对象由 Spring Boot 自动配置生成
	@Autowired
	SqlSessionFactory sqlSessionFactory = null;

	// 启用Spring Bean 生命周期执行方法， 加入插件
	@PostConstruct
	public void initMyBatis() {
		// 插件实例
		Interceptor plugin =new MyPlugin();
		// 设置插件属性
		Properties properties = new Properties();
		properties.setProperty("keyl", "valuel");
		properties.setProperty("key2", "value2");
		properties.setProperty("key3", "value3");
		plugin.setProperties(properties);
		// 在sqlSessionFactory 中添加插件
		sqlSessionFactory . getConfiguration() .addinterceptor(plugin);
	}
	// ......
}
```

## 数据库事务处理

	在Spring 中，数据库事务是通过AOP 技术来提供服务的。

	如一个批处理，它将处理多个交易，但是在一些交易中发生了异常， 这个时候则不能将所有的交易都回滚。如果所有的交易都回渎，那么那些本能够正常处理的业务也无端地被回滚了，这显然不是我们所期待的结果。通过Spring 的数据库事务传播行为，可以很方便地处理这样的场景。

### JDBC 的数据库事务

```java
package com.xyz.example.service.impl;
@Service
public class JdbcServiceImpl implements JdbcService {
	@Autowired
	private DataSource dataSource = null;
	@Override
	public int insertUser(String userName, String note) {
		Connection conn = null;
		int result = 0;
		try {
			// 获取连接
			conn = dataSource.getConnection();
			// 开启事务
			conn.setAutoCommit(false);
			// 设置隔离级别
			conn.setTransactionIsolation(TransactionIsolationLevel.READ_COMMITTED.getLevel());
			// 执行业务SQL代码，其他都属性JDBC功能代码
			PreparedStatement ps = conn.prepareStatement(
				"INSERT INTO t_user(user_name, note) VALUES(?, ?)"
			);
			ps.setString(1, userName);
			ps.setString(2, note);
			result = ps.executeUpdate();
			// 提交事务
			conn.commit();
		} catch (Exception e) {
			// 回滚事务
			if (conn != null) {
				try {
					conn.rollback();
				} catch (SQLException ex) {
					ex.printStackTrace();
				}
			}
			e.printStackTrace();
		} finally {
			// 关闭数据库连接
			try {
				if (conn != null && !conn.isClosed()) {
					conn.close();
				}
			} catch (SQLException ex) {
				ex.printStackTrace();
			}
		}
		return result;
	}
}
```

### Spring 声明式数据库事务约定

	@Transactional 注解可以标注在类或者方法上，当它标注在类上时，代表这个类所有公共（public）非静态的方法都将启用事务功能。在 @Transactional 中，还允许配置许多的属性，如事务的隔离级别和传播行为。

	PlatformTransactionManager 事务管理器类，getTransaction 方法的参数是一个事务定义器，它是依赖于我们配置的 @ Transactional 的配置项生成，且能够设置事务的属性，而提交和回滚事务也就可以通过commit和rollback方法来执行。

```sql
/* 创建用户表 */
CREATE TABLE t_user (
	id int(12) auto_increment,
	user_name VARCHAR(60) NOT NULL,
	note VARCHAR(512),
	PRIMARY KEY(id)
);
```
```java
/* ------- 用户POJO ------ */
package com.xyz.example.pojo;
@Alias("user")    // Mybatis 别名注解
public class User {
	private Long id;
	private String userName;
	private String note;
	/* setter and getter */
}
```
```java
/* ------- MyBatis 接口文件 ------ */
package com.xyz.example.dao;
@Repository
public interface UserDao {
	User getUser(Long id);
	int insertUser(User user);
}
```
```xml
<!--  用户POJO -->
<mapper namespace="com.xyz.example.dao.UserDao" 〉
	<select id="getUser" parameterType="long" resultType="user">
		SELECT id, user_name AS userName, note FROM t_user where id=#{id}
	</select>
	<insert id="insertUser" useGeneratedKeys="true" keyProperty="id" >
		INSERT INTO t_user(user_name, note) value(#{userName}, #{note})
	</insert>
</mapper>
```
```java
/* ------- 用户服务接口 ------ */
package com.xyz.example.service;
public interface UserService {
	// 获取用户信息
	public User getUser(Long id);
	// 新增用户
	public int insertUser(User user);
}
```
```java
/* ------- 用户服务接口实现类 ------ */
package com.xyz.example.service.impl;
@Service
public interface UserService {
	@Autowired
	private UserDao userDao = null;
	// 事务定义
	@Override
	@Transactional(isolation=Isolation.READ_COMMITTED, timeout=1)
	public User getUser(Long id) {
		return userDao.getUser(id);
	}
	// 事务定义
	@Override
	@Transactional(isolation=Isolation.READ_COMMITTED, timeout=1)
	public User insertUser(User user) {
		return userDao.insertUser(user);
	}
}
```
```java
/* ------- 测试数据库事务 ------ */
package com.xyz.example.controller;
@Controller
@RequestMapping("/user")
public class UserController {
	// 注入Service
	@Autowired
	private UserService userService = null;
	// 测试获取用户
	@RequestMapping("/getUser")
	@ResponseBody
	public User getUser(Long id) {
		return userService.getUser(id);
	}
	// 测试插入用户
	@RequestMapping("/insertUser")
	@ResponseBody
	public Map<String, Object> insertUser(String userName, String note) {
		User user = new User();
		user.setUserName(userName);
		user.setNote(note);
		// 结果会回填，返回插入条数
		int update = userService.insertUser(user);
		Map<String, Object> result = new HashMap<>();
		result.put("success", update == 1);
		result.put("user", user);
		return result;
	}
}
```
	配置My Batis
	mybatis.mapper-locations=classpath:com/xyz/example/mapper/*.xml
	mybatis.type-aliases-package=com.xyz.example.pojo

```java
/* ------- Spring Boot 启动文件 ------ */
package com.xyz.example;

@MapperScan(basePackage="com.xyz.example", annotationClass=Repository.class)
@SpringBootApplication(scanBasePackages="com.xyz.example")
public class MainApplication {
	public static void main(String[] args) throws Exception {
		SpringApplication.run(MainApplication.class, args);
	}
	// 注入事务管理器， 它由Spring Boot 自动生成
	@Autowired
	PlatformTransactionManager transactionMananger = null;
	// 使用后初始化方法，观察自动生成的事务管理器
	@PostConstruct
	public void viewTransactionManager() {
		// 启动前加入断点观测
		System.out.println(transactionMananger.getClass().getName());
	}
}
```

	多个事务都提交引发的丢失更新称为第二类丢失更新。为了克服这些问题， 数据库提出了事务之间的隔离级别的概念。
	数据库现有的技术完全可以避免丢失更新，但是这样做的代价， 就是付出锁的代价，在互联网中， 系统不单单要考虑数据的一致性，还要考虑系统的性能。

	1. 未提交读（read uncommitted）是最低的隔离级别，其含义是允许一个事务读取另外一个事务没有提交的数据。
	未提交读是一种危险的隔离级别，所以一般在我们实际的开发中应用不广， 但是它的优点在于并发能力高，适合那些对数据一致性没有要求而追求高并发的场景，它的最大坏处是出现脏读。

	2. 写提交（read committed）隔离级别， 是指一个事务只能读取另外一个事务已经提交的数据，不能读取未提交的数据。

	3. 可重复读，等待已有事务提交才允许读取数据库。

	4. 串行化（Serializable）是数据库最高的隔离级别，它会要求所有的SQL 都会按照顺序执行，这样就可以克服上述隔离级别出现的各种问题，所以它能够完全保证数据的一致性。


	* 隔离级别和可能发生的现象

	Oracle 只能支持读写提交和串行化，默认的隔离级别为读写提交.
	MySQL 则能够支持4 种，，默认的隔离级别为可重复读。

|隔离级别|脏读|不可重复读|幻读|
|---------|---|---|---|
| 未提交读 | √ | √ | √ |
| 读写提交 | × | √ | √ | 
| 可重复读 | × | × | √ |  
| 串行化   | × | × | × | 

```ini
# application.properties 设置隔离级别数字配置的含义
#-1 数据库默认隔离级别
# 1 未提交读
# 2 读写提交
# 4 可重复读
# 8 串行化
# tomcat 数据源默认隔离级别
spring.datasource.tomcat.default-transaction-isolation=2
# dbcp2 数据库连接池默认隔离级别
#spring.datasource.dbcp2.default-transaction-isolation=2
```

#### 传播行为， 是方法之间调用事务采取的策略问题

	org.springframework.transaction.annotation.Propagation
	其中， REQUIRED, REQUIRES_NEW 和 NESTED 这3 种最常用的传播行为。

	在绝大部分的情况下，我们会认为数据库事务要么全部成功， 要么全部失败。但现实中也许会有特殊的情况。

	@Transactional 自调用失效问题, AOP 的原理是动态代理， 在自调用的过程中， 是类自身的调用，而不是代理对象去调用， 那么就不会产生AOP，这样Spring就不能把你的代码织入到约定的流程中， 于是就产生了现在看到的失败场景。 用一个Service 去调用另一个Service ， 这样就是代理对象的调用， Spring才会将你的代码织入事务流程。当然也可以从Spring IoC 容器中获取代理对象去启用AOP。

```java
package com.xyz.example.service;
public interface UserBatchService {
	public int insertUsers(List<User> userList);
}

/** -------- 批量用户实现类 --------- **/
package com.xyz.example.service.impl;
@Service
public class UserBatchServiceImpl implements UserBatchService /*, ApplicationContextAware */ {
	@Autowired
	private UserService userService = null;

	/* 
	private ApplicationContext applicationContext = null;

	// 实现生命周期方法，设置IoC 容器
	@Override
	public void setApplicationContext(ApplicationContext applicationContext) 
		throws BeanException {
		this.applicationContext = applicationContext;
	}
	*/

	@Override
	@Transactional(isolation=Isolation.READ_COMMITTED, propagation=Propagation.REQUIRED)
	public int insertUsers(List<User> userList) {
		int count = 0;
		// 从 IoC 容器中取出代理对象
		// UserService userService = applicationContext.getBean(UserService.class);
		for (User user : userList) {
			// 调用子方法，将使用 @Transactional 定义的传播行为
			count+= userService.insertUser(user);
		}
		return count;
	}
}
```

## 使用性能利器 Redis

	Redis 支持Lua 语言，而且在Redis 中Lua 语言的执行是原子性的，也就是在Redis 执行Lua 时， 不会被其他命令所打断，这样就能够保证在高并发场景下的一致性。Redis 除了操作那些数据类型的功能外， 还能支持事务、流水线、发布订阅等功能。

	使用 Spring 缓存注解操作 Redis

	• @CachePut 表放示将方法结果返回存到缓存中。
	• @Cacheable 表示先从缓存中通过定义的键查询，如果可以查询到数据，则返回，否则执行该方法，返回数据，并且将返回结果保存到缓存中。
	• @CacheEvict 通过定义的键移除缓存，它有一个Boolean 类型的配置项beforeInvocation，表示在方法之前或者之后移除缓存。因为其默认值为false，所以默认为方法之后将缓存移除。

```xml
<dependency>
	<groupId>org.springframework.boot</groupId>
	<artifactId>spring-boot-starter-data-redis</artifactId>
	<exclusions>
		<!-- 不依赖Redis的异步客户端lettuce -->
		<exclusion>
			<groupId>io.lettuce</groupId>
			<artifactId>lettuce-core</artifactId>
		</exclusion>
	</exclusions>
</dependency>
<!-- 引入Redis的客户端驱动jedis -->
<dependency>
	<groupId>redis.clients</groupId>
	<artifactId>jedis</artifactId>
</dependency>
```
```java
/* -------- 通过一个连接池的配置创建 RedisConnectionFactory 对象------- */
package com.xyz.example.config;
@configuration
public class RedisConfig {
	private RedisConnectionFactory connectionFactory = null;
	@Bean(name = "RedisConnectionFactory")
	public RedisConnectionFactory initRedisConnectionFactory() {
		if (this.connectionFactory != null) {
			return this.connectionFactory;
		}
		JedisPoolConfig poolConfig = new JedisPoolConfig();
		// 最大空闲数
		poolConfig.setMaxIds(30);
		// 最大连接数
		poolConfig.setMaxTotal(50);
		// 最大等待毫秒数
		poolConfig.setMaxWaitMillis(2000);
		// 创建 Jedis 连接工厂
		JedisConnectionFactory connectionFactory = new JedisConnectionFactory(poolConfig);
		// 获取单机的Redis配置
		RedisStandaloneConfiguration rsCfg = connectionFactory.getStandaloneConfiguration();
		connectionFactory.setHostName("192.168.11.131");
		connectionFactory.setPort(6379);
		connectionFactory.setPassword("123456");
		this.connectionFactory = connectionFactory;
		return connectionFactory;
	}
	/* ...... */
}
// Spring 为了简化开发，提供了 RedisTemplate 类管理 Redis
```

	RedisTemplate 是一个强大的类，它会自动从RedisConnectionFactory 工厂中获取连接，然后执行对应的Redis命令，在最后还会关闭Redis的连接。

	* RedisTemplate 中的序列化器属性

|属性|描述|备注|
|-----|----------|----------------------|
|defaultSerializer | 默认序列化器 | 如果没有设置，则使用JdkSerializationRedisSeralizer |
|keySerializer | Redis 键序列化器 | 如果没有设置，则使用默认序列化器 |
|valueSerializer | Redis 值序列化器 | 如果没有设置，则使用默认序列化器 |
|hashKeySerializer | Redis 散列结构field序列化器 | 如果没有设置，则使用默认序列化器 |
|hashSerializer | Redis 散列结构value序列化器 | 如果没有设置，则使用默认序列化器 |
|stringSerializer | 字符串序列化器 | RedisTemplate 自动赋值为StringRedisSerializer 对象 |


```java
/* ------- 创建 RedisTemplate -------- */
@Bean(name="redisTemplate")
public RedisTemplate<Object, Object> initRedisTemplate() {
	RedisTemplate<Object, Object> redisTemplate = new RedisTemplate<>();
	// RedisTemplate 会自动初始化 StringRedisSerializer ，所以这里直接获取
	RedisSerializer stringRedisSerializer = redisTemplate.getStringSerializer();
	// 设置字符串序列化器，这样Spring 就会把Redis 的key 当作字符串处理了
	redisTemplate.setKeySerializer(stringRedisSerializer);
	redisTemplate.setHashKeySerializer(stringRedisSerializer);
	redisTemplate.setHashValueSerializer(stringRedisSerializer);
	redisTemplate.setConnectionFactory(initConnectionFactory());
	return redisTemplate;
}
```
```java
/* ------- 测试 RedisTemplate -------- */
package com.xyz.example.main;
public class ExampleMain {
	public static void main(String[] args) {
		ApplicationContext ctx = new AnnotationConfigApplicationContext(RedisConfig.class);
		RedisTemplate redisTemplate = ctx.getBean(RedisTemplate.class);
		redisTemplate.opsForValue().set("Key1", "value1");
		redisTemplate.opsForHash().put("hash", "field", "hvalue");
	}
}
```

	* Spring 对Redis 数据类型操作的封装

	Redis 能够支持7 种类型的数据结构，这7 种类型是字符串、散列、列表（链表） 、集合、有序集合、基数和地理位置。

```java
// 获取地理位置操作接口 GeoOperations
redisTemplate.opsForGeo();
// 获取散列操作接口 HashOperations
redisTemplate.opsForHash();
// 获取基数操作接口 HyperLogLogOperations
redisTemplate.opsForHyperLogLog();
// 获取列表操作接口 ListOperations
redisTemplate.opsForList();
// 获取集合操作接口 SetOperations
redisTemplate.opsForSet();
// 获取字符串操作接口 ValueOperations
redisTemplate.opsForValue();
// 获取有序集合操作接口 ZSertOperations
redisTemplate.opsForZSet();
```

	* SessionCallback 和 RedisCallback 接口

	它们的作用是让 RedisTemplate 进行回调，通过它们可以在同一条连接下执行多个Redis 命令。

```java
/* RedisCallback 接口比较底层， 需要处理的内容也比较多，可读性较差，所以在非必要的时候尽量不选择使用它 */
public void useRedisCallback(RedisTemplate redisTemplate) {
	redisTemplate.execute(new RedisCallback() {
		@Override
		public Object doInRedis(RedisConnection rc) throws DataAccessException {
			rc.set("key1".getBytes(), "value1".getBytes());
			rc.hSet("hash".getBytes(), "field".getBytes(), "hvalue".getBytes());
			return null;
		}
	});
}
// 使用 Lambda 表达式
public void useRedisCallback(RedisTemplate redisTemplate) {
	redisTemplate.execute((RedisConnection rc) -> {
			rc.set("key1".getBytes(), "value1".getBytes());
			rc.hSet("hash".getBytes(), "field".getBytes(), "hvalue".getBytes());
			return null;
	});
}

/* SessionCallback 提供了良好的封装，对于开发者比较友好，因此在实际的开发中应该优先选择使用它 */
public void useSessionCallback(RedisTemplate redisTemplate) {
	redisTemplate.execute(new SessionCallback() {
		@Override
		public Object execute(RedisConnection ro) throws DataAccessException {
			ro.opsForValue().set("key1", "value1");
			ro.opsForHash("hash", "field", "hvalue");
			return null;
		}
	});
}
// 使用 Lambda 表达式
public void useSessionCallback(RedisTemplate redisTemplate) {
	redisTemplate.execute((RedisConnection ro) -> {
			ro.opsForValue().set("key1", "value1");
			ro.opsForHash("hash", "field", "hvalue");
			return null;
	});
}
```

	* 使用Spring 操作列表(链表)

```java
@RequestMapping("/list")
@ResponseBody
public Map<String, Object> testList() {
	// 插入两个列表， 注意它们在链表的顺序
	// 链表从左到右顺序为vl0 , v8 ,v6 , v4,v2
	stringRedisTemplate.opsForList().leftPushAll("list1","v2","v4","v6","v8","vl0");
	// 链表从左到右顺序为 v1, v2, v 3, v 4 ,v5, v6
	stringRedisTemplate.opsForList().rightPushAll("list2","v1","v2","v3","v4","v5","v6");
	// 绑定list2 链表操作
	BoundListOperations listOps = stringRedisTemplate.boundListOps("list2");
	//  从右边弹出一个成员
	Object resultl = listOps.rightPop();
	// 获取定位元素, Redis 从0 开始计算, 这里值为v2
	Object result2 = listOps.index(1) ;
	// 从左边插入链表
	listOps.leftPush("v0");
	// 求链表长度
	Long size = listOps.size();
	// 求链表下标区间成员,整个链表下标范围为0 到size-1 ,这里不取最后一个元素
	List elements = listOps.range(0,size-2);
	Map<String , Object> map= new HashMap<String , Object>();
	map.put("success", true);
	return map;
}
```

#### 使用Redis 事务

	Redis 是支持一定事务能力的NoSQL ， 在Redis 中使用事务，通常的命令组合是watch...multi .. . exec，也就是要在一个Redis 连接中执行多个命令，这时我们可以考虑使用S巳ssionCallback 接口来达到这个目的。

	watch 命令是可以监控Redis 的一些键; 

	multi 命令是开始事务，开始事务后， 该客户端的命令不会马上被执行，而是存放在一个队列里，这点是需要注意的地方，也就是在这时我们执行一些返回数据的命令， Redis 也是不会马上执行的，而是把命令放到一个队列里，所以此时调用Redis 的命令，结果都是返回null ，这是初学者容易犯的错误；

	exec 命令的意义在于执行事务，只是它在队列命令执行前会判断被watch 监控的Redis 的键的数据是否发生过变化（即使赋予与之前相同的值也会被认为是变化过)，如果它认为发生了变化，那么Redis 就会取消事务， 否则就会执行事务， Redis 在执行事务时，要么全部执行， 要么全部不执行，而且不会被其他客户端打断，这样就保证了Redis 事务下数据的一致性。

```java
/* ------ 通过Spring 使用Redis 事务机制 ------ */
@RequestMapping("/multi")
@ResponseBody
public Map<String, Object> testMulti() {
	redisTemplate.opsForValue().set("key1", "value1");
	List list = (List) redisTemplate.execute((RedisOperations operations) -> {
		// 设置要监控key1
		operations.watch("key1");
		// 开启事务，在exec 命令执行前，全部都只是进入队列
		operations.multi();
		operations.opsForValue().set("key2","value2");
		// operations.opsForValue().increment("key1", 1);  // 第1步
		// 获取值将为null ， 因为自由s 只是把命令放入队列
		Object value2 = operations.opsForValue().get("key2");
		System.out.println("命令在队列，所以value 为null 【"+ value2 +"】");
		operations.opsForValue().set("key3","value3");
		Object value3 = operations.opsForValue().get("key3");
		System.out.println("命令在队列，所以value 为null 【" + value3 +"】");
		// 执行exec 命令，将先判别key1 是否在监控后被修改过，如果是则不执行事务，否则就执行事务
		return operations exec(); // 第2步
	});
	System.out.println(list);
	Map<String, Object> map = new HashMap<String, Object>();
	map.put("success", true);
	return map;
}
```

	* 使用Redis 流水线

	在默认的情况下， Redis 客户端是一条条命令发送给Redis 服务器的，这样显然性能不高。在关系数据库中我们可以使用批量，也就是只有需要执行SQL 时，才一次性地发送所有的SQL 去执行，这样性能就提高了许多。

```java
@RequestMapping("/pipeline")
@ResponseBody
public Map<String, Object> testPipeline() {
	Long start = System.currentTimeMillis();
	List list = (List)redisTemplate.executePipelined((RedisOperations operations) -> {
		for (int i = 1; i <= 10000; i++) {
			operations.opsForValue().set("pipeline_" + i, "value_" + i);
			String value = (String)operations.opsForValue().get("pipelin_" + i);
			if (i == 10000) {
				System.out.println("命令只是进入队列，所以值为空【" ＋ value +"】");
			}
		}
		return null;
	});
	Long end = System.currentTimeMillis();
	System.out.println("耗时：" + (end - start) + "毫秒。");
	Map<String, Object> map = new HashMap<String, Object>();
	map.put("success", true);
	return map;
}
```

	* Redis 发布订阅

	Redis 提供一个渠道，让消息能够发送到这个渠道上，而多个系统可以监听这个渠道， 如短信、微信和邮件系统都可以监昕这个渠道，当一条消息发送到渠道，渠道就会通知它的监昕者，这样短信、微信和邮件系统就能够得到这个渠道给它们的消息了，这些监听者会根据自己的需要去处理这个消息，于是我们就可以得到各种各样的通知了。

```java
/* -------- Redis 消息监听器  ------- */
package com.xyz.example.listener;
@Component
public class RedisMessageListener implements MessageListener {
	@Override
	public void onMessage(Message message, byte[] pattern) {
		// 消息体
		String body = new String(message.getBody());
		// 渠道名称
		String topic = new String(pattern);
		System.out.println(body);
		System.out.println(topic);
	}
}
```
```java
/* -------- 监听Redis 发布的消息  ------- */
package com.xyz.example.main;
@SpringBootApplication(scanBasePackages="com.xyz.example")
@MapperScan(basePackages="com.xyz.example", annotationClass=Repository.class)
public class TestApplication {
	/* ...... */
	@Autowired
	private RedisTemplate redisTemplate = null;
	// Redis 连接工厂
	@Autowired
	private RedisConnectionFactory connectionFactory = null;
	// Redis 消息监听器
	@Autowired
	private MessageListener redisMsgListener = null;
	// 任务池
	private ThreadPoolTaskScheduler taskScheduler = null;
	// 创建任务池，运行线程等待处理 Redis 的消息
	@Bean
	public ThreadPoolTaskScheduler initTaskScheduler() {
		if (taskScheduler != null) {
			return taskScheduler;
		}
		taskScheduler = new TheadPoolTaskScheduler();
		taskScheduler.setPoolSize(20);
		return taskScheduler;
	}
	// 定义Redis 的监听容器
	@Bean
	public RedisMessageListenerContainer initRedisContainer() {
		RedisMessageListenerContainer container = new RedisMessageListenerContainer();
		// Redis 连接工厂
		container.setConnectionFactory(connectionFactory);
		// 设置运行任务池
		container.setTaskExecutor(initTaskScheduler());
		// 定义监听渠道，名称为 topic1
		Topic topic = new ChannelTopic("topic1");
		// 使用监听器监听 Redis 消息
		container.addMessageListener(redisMsgListener, topic);
		return container;
	}
}
```

	* Redis 使用 Lua 脚本

	一种是直接发送Lua 到Redis 服务器去执行，另一种是先把Lua 发送给Redis, Redis 会对Lua 脚本进行缓存，然后返回一个SHA1 的32 位编码回来，之后只需要发送SHA1 和相关参数给Redis 便可以执行了。

```java
/* ------ 执行简易Lua 脚本 ------- */
@RequestMapping("/lua")
@ResponseBody
public Map<String, Object> testLua() {
	DefaultRedisScript<String> rs = new DefaultRedisScript<String>();
	// 设置脚本
	rs.setScriptText("return 'Hello Redis'");
	// 定义返回类型。注意如果没有这个定义， Spring 不会返回结果
	rs.setResultType(String.class);
	RedisSerializer<String> stringSerializer = redisTemplate.getStringSerializer();
	// 执行Lua脚本
	String str = (String)redisTemplate.execute(rs, stringSerializer, stringSerializer, null);
	Map<String, Object> map = new HashMap<String, Object>();
	map.put("str", str);
	return map;
}
```

	* 自定义缓存管理器

```java
// 注入连接工厂，由Spring Boot 自动配置生成
@Autowired
private RedisConnectionFactory connectionFactory = null;
	// 自定义Redis 缓存管理器
	@Bean(name ="redisCacheManager")
	public RedisCacheManager initRedisCacheManager() {
	// Redis 加锁的写入器
	RedisCacheWriter writer= RedisCacheWriter.lockiηgRedisCacheWriter(connectionfactory);
	// 启动Redis 缓存的默认设置
	RedisCacheConfiguration config = RedisCacheConfiguration.defaultCacheConfig();
	// 设置JDK 序列化器
	config = config.serializeValuesWith(SerializationPair.fromSerializer(new JdkSerializationRedisSerializer()));
	// 禁用前缀
	config = config.disableKeyPrefix();
	// 设置 10 min 超时
	config = config.entryTtl(Duration.ofMinutes(10)) ;
	// 创建缓Redis 存管理器
	RedisCacheManager redisCacheManager = new RedisCacheManager(writer , config) ;
	return redisCacheManager;
}
```


182










