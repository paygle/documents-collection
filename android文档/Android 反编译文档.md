# Android 反编译文档

	Bandican 录像软件

	smali 语法, 

 1. 安装 SDK Tool 所有版本 和 需要开发的目标平台 SDK（ 版本 > 5)
	包含：adb（在platform-tools目录）、DDMS(Sdk\tools目录内存监视工具，执行 monitor.bat文件） 等工具。
 
 2. 反编译工具： AndroidKiller ApkIDE  IDA DDMS 等等
 
    AndroidKiller，ApkIDE 静态代码注入，debug.smali 文件加入需要注入的被反编译的APK中
	
	JEB（Android 5 以上版本，动态调试，搜索中文清空先）
 
 3. 使用 ps 命令（安卓9不可行） 获取进程列表，通过程序名称截取 pid 数值，使用 killProcess 干掉应用。
 
	◆ 系统proc文件夹方法
	
	原理：Android系统在管理进程时，通过low memory killer机制来定时找出oom_score评分高出阈值的进程，进行回收，那么反过来考虑，oom_score值最低的，且oom_adj值为0的进程(0为前台进程的adj值)，就是很可能是当前的前台进程了。
	这种方式可以静默运行，但是也有两个问题：
	1.获取的其实是进程名，默认进程名为App的包名，但是开发者可以自定义进程名，所以你可能拿不到App的包名。
	2.获取的结果并不精准，因为评分最低的不一定是当前显示的App，还有可能是某个进程优先级很高的后台服务，我们可能需要维护一个黑名单，在代码中屏蔽掉这个名单上的所有后台服务。
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	

 
 
 