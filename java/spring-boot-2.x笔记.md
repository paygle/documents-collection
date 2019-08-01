# Spring Boot 2.x 笔记

### Spring Boot 的依赖和自动配置

	org.springframework.boot.autoconfigure.web.servlet
	
### Spring 中是通过注解描述来创建IoC对象。

	Spring Boot并不建议使用XML，而是通过注解的描述生成对象。
	
	在Spring 中把每一个需要管理的对象称为Spring Bean（简称Bean ），而Spring 管理这些Bean 的容器，被我们称为SpringIoC 容器（或者简称IoC 容器）。IoC 容器需要具备两个基本的功能：
	
		* 通过描述管理Bean ， 包括发布和获取Bean; 
		
		* 通过描述完成Bean 之间的依赖关系。
		
### 所有的IoC 容器都需要实现接口 BeanFactory ，它是一个顶级容器接口。
	
### ApplicationContext 是 BeanFactory 的子接口之一

	是最为重要的接口设计，大部分Spring IoC 容器是 ApplicationContext 接口的实现类。
	
### org.springframework.context.annotation.AnnotationConfigApplicationContext 是一个基于注解的IoC 容器。
	
	* @Configuration 和 @Bean 注解例子
	
```java
package com.xyz.config;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

// ＠Configuration 代表这是一个Java 配置文件
// Spring 的容器会根据它来生成IoC 容器去装配Bean

@Configuration
public class AppConfig {

	// @Bean 代表将initUser 方法返回的POJO 装配到IoC 容器中，而其属性name 定义这个Bean 的名称，
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

	如果一个个的Bean 使用注解＠Bean 注入Spring loC 容器中，那将是一件很麻烦的事情。
	对于扫描装配而言使用的注解是 @Component 和 @ComponentScan。

	> @Component 是标明哪个类被扫描进入Spring IoC 容器

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
	其中配置的“user＂则是作为 Bean 的名称，当然你也可以不配置这个字符串，那么IoC 容器就会把类名第一个字母作为小写，其他不变作为Bean 名称放入到IoC 容器中；
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

	例如，人类（Person）有时候利用 <---- 动物(Animal）去完成一些事情，比方说狗（Dog）是用来看门的，猫（Cat）是用来抓老鼠的，鹦鹉（Parrot)是用来迎客的…

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
// 	@Primary，它是一个修改优先权的注解，当我们有猫有狗的时候，假设这次需要使用猫， 那么只需要在猫类的定义上加入@Primarγ 就可以了。
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
public class BussinessPerson implements Person, 
				BeanNameAware, BeanFactoryAware, ApplicationContextAware, InitializingBean, DisposableBean {

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
			System.out.println("BeanPostProcessor 调用 postProcessBeforeinitialization 方法，参数【" + bean.getClas().getSimpleName() + "】【" + beanName + "】");
			retur bean;
	}

	@Override
	public Object postProcessAfterInitializatiion(Object bean, String beanName) throw BeansException {
		System.out.println("BeanPostProcessor 调用 postProcessAfterInitializatiion 方法，参数【" + bean.getClass().getSimpleName() + "】【" + beanName + "】");
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
	可以通过@Value 注解， 使用 ${...} 这样的占位符读取配置在属性文件的内容。这里的 @Value 注解，既可以加载属性， 也可以加在方法上。
	使用 @ConfigurationProperties 通过它使得配置上有所减少，去掉 @Value 的配置。
	注解 @ConfigurationProperties 中配置的字符串 database ，将与POJO 的属性名称组成属性的全限定名去配置文件里查找，这样就能将对应的属性读入到POJO 当中。
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

## 8. 部署环境的切换 @Profile

62


























	
	
	
	
	
	