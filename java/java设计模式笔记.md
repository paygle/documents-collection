# Java 设计模式笔记


## idea 多线程调试：

	1. 在需要断点的行打上“断点”，在断点上右键弹出选项中选择： Suspend -> Thread
    2. 启动调试在 Debugger 窗口 -> Frames 左上侧的下拉选项中选择对应的断点项。
	3. 使用 synchronized 同步关键字修饰单例对象的生成方法，防止多线程创建多个单例实例。

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
	
	
	
	