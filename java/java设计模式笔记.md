# Java 设计模式笔记


## idea 多线程调试：

	1. 在需要断点的行打上“断点”，在断点上右键弹出选项中选择： Suspend -> Thread
    2. 启动调试在 Debugger 窗口 -> Frames 左上侧的下拉选项中选择对应的断点项。
	3. 使用 synchronized 同步关键字修饰单例对象的生成方法，防止多线程创建多个单例实例。
	4. 在项目的pom文件位置处打开cmd执行命令： mvn idea:module 生成 *.iml 工程文件

## 模式应用
	
	1. 适配器模式
	不是软件设计阶段考虑的模式，是软件维护时对已经存在的类，这的方法和需求不匹配时（方法结果相同或相似）应用。
	
	
## 优化同步锁性能开销：
	
```java
public class LazySingeton {

	// 使用 volatile 禁止程序执行 赋值地址和初始操作随意排序
	private volatile static LazyInst = null
	
	private LazySingeton() {
		if (LazyInst != null) {
			throw new RuntimeException("单例构造器，禁止反射调用");
		}
	}
	
	public static LazySingeton getInstance() {
		if (LazyInst == null) {
			// 保证初始化开销最小
			synchronized (LazySingeton.class) {
				if (LazyInst == null) {
					LazyInst = new LazySingeton();
				}
			}
		}
		return LazyInst;
	}
	
	// 序列化、反序列化获取相同的对象
	private Object readResolve() {
		return LazyInst
	}

}

// 枚举类是天然序列化对象，处理单例反射攻击
public enum EnumInstance {

	INSTANCE;
	
	private Object data;
	
	public Object getData() {
		return data;
	}
	
	public volid setData(Object data){
		this.data = data;
	}
	
	public static EnumInstance getInstance(){
		return INSTANCE
	}
}
```

## 原型模式 

	即在内存中大量深层的复制一个已经创建好的类，而不是去重新申请这个类。
	
```java
// 可以被克隆的对象
public class Mail implements Cloneable {
	// ...
	
	@Override
	protected Object clone() throws CloneNotSupportedException {
	
		Mail mail = (Mail)super.clone();
		
		// 深克隆, 子对象也必须克隆一遍
		mail.sendTime = (Date) mail.sendTime.clone();
		
		return super.clone();
		
		// return getInstance();  // 阻止破坏单例对象
	}
}

// 调用测试
public class Test {
	public static void main(String[] args) throws CloneNotSupportedException {
		Mail mail = new Mail();
		mail.setContent("初始化模板");
		for(int i = 0; i < 10; i++){
			Mail mailTemp = (Mail) mail.clone();
			mailTemp.setName("name is " + i);
			mailTemp.setAddress("name" + i + "@qq.com");
			MailUtil.sendMail(mailTemp)
		}
		MailUtil.saveOriginMailRecord(mail);
	}
}

```

## 装饰者模式

	* 扩展一个类的功能或给一个类添加附加职责（增加复杂性）
	
	* 动态给一个对象添加功能，这些功能可以再动态的撤销
	
	* 是继承有力的补充，比继承灵活，不改变原有对象的情况下给一个对象扩展功能
	
```java

// 煎饼类, 被包装的对象
public abstract class ABattercake {
	protected abstract String getDesc();
	protected abstract int cost();
}

// 煎饼实现类
public class Battercake extends ABattercake {
	@Override
	protected abstract String getDesc(){
		return "煎饼";
	}
	
	@Override
	protected abstract int cost(){
		return 8;
	}
}

// 装饰类, 以参数的形式注入被装饰的煎饼类
public class AbstractDecorator extends ABattercake {

	private ABattercake aBattercake;
	
	public AbstractDecorator(ABattercake aBattercake) {
		this.aBattercake = aBattercake;
	}
	
	@Override
	protected String getDesc(){
		return this.aBattercake.getDesc();
	}
	
	@Override
	protected int cost(){
		return this.aBattercake.cost();
	}
}

// 加香肠装饰类
public class SausageDecorator extends AbstractDecorator {

	public SausageDecorator(ABattercake aBattercake) {
		supper(aBattercake);
	}
	
	@Override
	protected String getDesc(){
		return supper.getDesc() + " 加一根香肠";
	}
	
	@Override
	protected int cost(){
		return supper.cost() + 2;
	}
}

// 加一个鸡蛋装饰类
public class EggDecorator extends AbstractDecorator {

	public EggDecorator(ABattercake aBattercake) {
		supper(aBattercake);
	}
	
	@Override
	protected String getDesc(){
		return supper.getDesc() + " 加一个鸡蛋";
	}
	
	@Override
	protected int cost(){
		return supper.cost() + 1;
	}
}

/**
 * 测试类
 */
public class Test {
	public static void main(String[] args) {
		ABattercake aBattercake;
		aBattercake = new Battercake();
		aBattercake = new EggDecorator(aBattercake);
		aBattercake = new EggDecorator(aBattercake);
		aBattercake = new SausageDecorator(aBattercake);
		// 装饰多次后的结果
		System.out.println(aBattercake.getDesc() + aBattercake.cost());
	}
}

```
	
## 享元模式，复用对象 （Hashmap实现共享池）

	提供了减少对象数量从而改善应用所需的对象结构的方式
	运用共享技术有效地支持大量细粒度的对象
	
	常常应用于系统底层的开发，以便解决系统的性能问题
	系统有大量的相似对象，需要缓冲池的场景
	
	要关注内部（不变）/外部（可变）状态，关注线程安全问题，系统逻辑复杂化
	
```java
public interface Employee {
	void report();
}

public class Manager implements Employee {
	@Override
	void report(){
		System.out.println(reportContent);
	};
	
	// 不变的内部状态
	private String title = "部门经理";
	
	// 外部传入的状态
	private String department;
	private String reportContent;
	
	public void setReportContent(String reportContent) {
		this.reportContent = reportContent;
	}
	
	public Manager(String department) {
		this.department = department;
	}
}

public class EmployeeFactory {
	private static final Map<String, Employee> EMPLOYE_MAP = new HashMap<>();
	
	public static Employee getManager(String department) {
		Manager manager = (Manager) EMPLOYE_MAP.get(department);
		
		if (manager == null) {
			manager = new Manager(department);
			System.out.println("创建部门经理: " + department);
			String reportContent = department + "部门汇报: 此次报告内容......";
			manager.setReportContent(reportContent);
			System.out.println("创建报告: " + reportContent);
			EMPLOYE_MAP.put(department, manager);
		}
		return manager;
	}
}

public class Test {

	private static final String departments[] = {"RD", "QA", "PM"};
	
	public static void main(String[] args) {
		for(int i=0; i<10; i++) {
			String department = departments[(int)(Math.random)];
			Manager manager = (Manager) EmployeeFactory.getManager(department);
			manager.report();
		}
	}

}


```
	
## 组合模式（结构型）

- 定义：将对象组合成树形结构以表示“部分-整体”的层次结构

- 组合模式使用客户端对单个对象和组合对象保持一致的方式处理

- 缺点：一、限制类型时会比较复杂；二、使设计变得更加抽象


## 桥接模式（结构型）

- 定义：将抽象部分与它的具体实现部分分离，使它们都可以独立地变化

- 通过组合的方式建立两个类之间的联系，而不是继承

- 适用场景
	1. 抽象和具体实现之间增加更多的灵活性
	2. 一个类存在两个（或多个）独立变化的维度，且这两个（或多个）维度都需要进行扩展
	3. 不希望使用继承，或因为多层继承导致系统类的个数剧增

- 缺点： 增加了系统的理解与设计难度；需要正确地识别出系统中两个独立变化的维度

```java
/**
* 账户接口
*/
public interface Account {
	Account openAccount();
	void showAccountType();
}
/**
* 定期账户实现层
*/
public class DepositAccount implements Account {
	@Override
	Account openAccount() {
		System.out.println("打开定期账号");
		return new DepositAccount();
	}
	@Override
	void showAccountType() {
		System.out.println("这是一个定期账号");
	}
}
/**
* 活期账户实现层
*/
public class SavingAccount implements Account {
	@Override
	Account openAccount() {
		System.out.println("打开活期账号");
		return new SavingAccount();
	}
	@Override
	void showAccountType() {
		System.out.println("这是一个活期账号");
	}
}
/**
* 定义一个抽象银行
*/
public abstract class Bank {

	protected Account account; // 桥接抽象层

	public Bank(Account account) {
		this.account = account;
	}
	// 委托
	abstract Account openAccount();
}
/**
* 定义一个农业银行账号
*/
public class ABCBank extends Bank {

	public ABCBank(Account account) {
		super(account);
	}
	@Override
	Account openAccount() {
		System.out.println("打开中国农业银行账号");
		// 桥接-具体行为委托调用
		account.openAccount();
		return account;
	}
}
/**
* 定义一个工商银行账号
*/
public class ICBCBank extends Bank {

	public ABCBank(Account account) {
		super(account);
	}
	@Override
	Account openAccount() {
		System.out.println("打开中国工商银行账号");
		// 桥接-具体行为委托调用
		account.openAccount();
		return account;
	}
}

/**
* 测试类
*/
public class Test {
	public static void main(String[] args) {
		Bank icbcBank = new ICBCBank(new DepositAccount());
		Account icbcAccount = icbcBank.openAccount();
		icbcAccount.showAccountType();
	}
}
```

## 代理模式（结构型）

- 定义： 为其他对象提供一种代理，以控制对这个对象的访问

- 代理对象在客户端和目标对象之间起到中介作用

- 使用场景： 保护目标对象；增强目标对象

- 优点

	1. 能将代理对象与真实被调用的目标对象分离
	2. 一定程度上降低了系统的耦合度，扩展性好
	3. 保护目标对象
	4. 增强目标对象

- 静态代理

```java
public class Order {
	private Object orderInfo;
	private Integer userId;
	/** gettter and setter **/
}
public interface IOrderService {
	int saveOrder(Order order);
}
public interface IOrderDao {
	int insert(Order order);
}
public class OrderDaoImpl implements IOrderDao {
	@Override
	public int insert(Order order) {
		System.out.println("Dao层添加Order成功");
		return 1;
	}
}
public class OrderServiceImpl implements IOrderService {
	private IOrderDao iOrderDao;
	@Override
	public int saveOrder(Order order) {
		// Spring 会自动注入
		iOrderDao = new OrderDaoImpl();
		System.out.println("Service层调用Dao层添加Order");
		return iOrderDao.insert(order);
	}
}

/**
* 静态代理类
*/
public class OrderServiceStaticProxy {
	private IOrderService iOrderService;
	public int saveOrder(Order order) {
		beforeMethod(order);
		iOrderService = new OrderServiceImpl();
		int result = iOrderService.saveOrder(order);
		afterMethod();
		return result;
	}
	private void beforeMethod(Order order) {
		int userId = order.getUserId();
		int dbRouter = userId % 2;
		System.out.println("静态代理分配到【db"+dbRouter+"】处理数据");
		// 设置 dataSource
		DataSourceContextHolder.setDBType("db"+String.ValueOf(dbRouter));
		System.out.println("静态代理 before code");
	}
	private void afterMethod() {
		System.out.println("静态代理 after code");
	}
}
```

- JDK动态代理-接口代理

```java
public class OrderServiceDynamicProxy implements InvocationHandler {
	private Object target;

	public OrderServiceDynamicProxy(Object target) {
		this.target = target;
	}

	public Object bind() {
		Class cls = target.getClass();
		return Proxy.newProxyInstance(cls.getClassLoader(), cls.getInterfaces(), this);
	}

	@Override
	public Object invoke(Object proxy, Method method, Object[] args) throws Exception {
		Object argObject = args[0];
		beforeMethod(argObject);
		Object object = method.invoke(target, args);
		afterMethod();
		return object;
	}

	private void beforeMethod(Object obj) {

		int userId = 0;
		System.out.println("动态代理 before code");
		if (obj instanceof Order) {
			Order order = (Order)obj;
			userId = order.getUserId();
		}

		int dbRouter = userId % 2;
		System.out.println("动态代理分配到【db"+dbRouter+"】处理数据");
		// 设置 dataSource
		DataSourceContextHolder.setDBType("db"+String.ValueOf(dbRouter));
		
	}

	private void afterMethod() {
		System.out.println("动态代理 after code");
	}
}
```

- CGLib代理-类代理（生成子类）


## 模板方法模式（行为型）

- 定义了一个算法的骨架，并允许子类为一个或多个步骤提供实现

- 模板方法使得子类可以在不改变算法结构的情况下，重新定义算法某些步骤

- 缺点
	1. 类数目增加
	2. 增加了系统实现的复杂度
	3. 继承关系自身缺点，如果父类添加新的抽象方法，所有子类都要改一遍

- 工厂方法是模板方法的特殊实现

```java
public abstract class ACourse {

	// 模板方法核心方法，不希望被子类修改
	protected final void makeCourse() {
		this.makePPT();
		this.makePPT();
		if (this.needWriteArticle()) {
			this.writeArticle();
		}
		this.packageCource();
	}

	final void makePPT() {
		System.out.println("制作PPT");
	}
	final void makeVideo() {
		System.out.println("制作视频");
	}
	final void writeArticle() {
		System.out.println("编写手记");
	}
	// 钩子方法
	protected boolean needWriteArticle() {
		return false;
	}
	// 子类实现
	abstract void packageCource();
}
/**
* 实现课程
*/
public class DesignPatternCourse extends ACourse {
	@Override
	void packageCource() {
		System.out.println("提供课程Java源代码");
	}
	@Override
	boolean needWriteArticle() {
		return true;
	}
}
```




	
	
	
	
	
	
	