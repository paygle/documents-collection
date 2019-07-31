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

	// @Bean 代表将initUser 方法返回的POJO 装配到IoC 容器中，而其属性name 定义这个Bean 的名称，
	// 如果没有配置它，则将方法名称“initUser”作为Bean的名称保存到Spring IoC容器中。
	
	@Bean(name = ”user”)
	public User initUser () {
		User user= new User();
		user.setId (lL);
		user.setUserName (”user_name_l”);
		user.setNote( ”note_l” );
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
	
## 装配你的Bean

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
	
## 自定义第三方Bean

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
	
## 依赖注人 

46
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	