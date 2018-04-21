# Matlab 的命令

* [管理会话](#conversation)
* [使用系统](#systemuse)
* [输入和输出](#inoutput)
* [fscanf和fprintf](#fprintf)
* [格式化format函数](#format)
* [向量，矩阵和数组命令](#vector)
* [绘图命令](#plot)


## <a name="conversation"></a>管理会话命令

```m
    clc	    % 清除命令窗口
    clear	% 从内存中删除变量
    exist	% 检查文件或变量是否存在
    global	% 声明变量为全局变量
    help	% 搜索帮助主题
    lookfor	% 搜索帮助关键字的条目。
    quit	% 停止MATLAB
    who	    % 列出当前变量
    whos	% 列出当前变量(长显示详细信息)
```

## <a name="systemuse"></a>使用系统命令

```m
    cd	          % 更改当前目录(进入指定目录)
    date	      % 显示当前日期
    delete	      % 删除文件
    diary	      % 打开/关闭日记文件记录
    dir	          % 列出当前目录下的所有文件
    load	      % 从文件加载工作区变量
    path	      % 显示搜索路径
    pwd	          % 显示当前目录
    save	      % 将工作空间变量保存在文件中
    type	      % 显示文件的内容
    what	      % 列出当前目录中的所有MATLAB文件
    wklread	      % 读取.wk1电子表格文件
```

## <a name="inoutput"></a>输入和输出命令

```m
    disp	      % 显示数组或字符串的内容
    fscanf	      % 从文件读取格式化数据
    format	      % 控制屏幕显示格式
    fprintf	      % 对屏幕或文件执行格式化的写入
    input	      % 显示提示并等待输入
    ;	          % 禁止打印显示
```

## <a name="fprintf"></a>fscanf和fprintf命令

```m
    %s	      % 格式化为字符串
    %d	      % 格式化为字符串整数值
    %f	      % 格式化为浮点值
    %e	      % 格式化为科学记数法的浮点值
    %g	      % 格式最紧凑的形式：%f或%e。
    \n	      % 在输出字符串中插入一个换行
    \t	      % 在输出字符串中插入制表符
```

## <a name="format"></a>格式化(format)函数具有以下用于数字显示的形式

```m
    format short	  % 4位十进制数(默认)
    format long		  % 16位数字
    format short e	  % 5位数加上指数
    format long e	  % 16位加上指数
    format bank		  % 两位十进制数字
    format +		  % 正，负或零
    format rat		  % 有理近似
    format compact	  % 禁止一些换行符
    format loose	  % 重置为较不紧凑的显示模式
```

## <a name="vector"></a>向量，矩阵和数组命令

```m
    cat	          % 连接数组
    find	      % 查找非零元素的索引
    length	      % 计算元素数量
    linspace      % 创建规则间隔的向量
    logspace      % 创建对数间隔向量
    max	          % 返回最大的元素
    min	          % 返回最小的元素
    prod	      % 产生的每列
    reshape	      % 改变大小
    size	      % 计算数组大小
    sort	      % 对每列进行排序
    sum	          % 对每列进行求和
    eye	          % 创建一个单位矩阵
    ones	      % 创建一个数组
    zeros	      % 创建一个零的数组
    cross	      % 计算矩阵交叉积
    dot	          % 计算矩阵点积
    det	          % 计算数组的行列式
    inv	          % 计算矩阵的倒数
    pinv	      % 计算矩阵的伪逆
    rank	      % 计算矩阵的秩
    rref	      % 计算简化行阶梯形式
    cell	      % 创建单元格数组
    celldisp      % 显示单元格数组
    cellplot      % 显示单元格阵列的图形表示
    num2cell      % 将数组转换为单元格数组
    deal	      % 匹配输入和输出列表
    iscell	      % 识别单元格数组
```

## <a name="plot"></a>绘图命令

```m
    axis	      % 设置轴限制
    fplot	      % 智能绘图功能
    grid	      % 显示网格线
    plot	      % 生成xy坐标图
    print	      % 打印或绘图到文件
    title	      % 在文字的顶部放置文字
    xlabel	      % 将文本标签添加到x轴
    ylabel	      % 将文本标签添加到y轴
    axes	      % 创建轴对象
    close	      % 关闭当前坐标图
    close all     % 关闭所有坐标图
    figure	      % 打开一个新的图形窗口
    gtext	      % 通过鼠标启用标签放置
    hold	      % 冻结当前坐标图
    legend	      % 通过鼠标图例位置
    refresh	      % 重新绘制当前图形窗口
    set	          % 指定诸如轴的对象的属性
    subplot	      % 在子窗口中创建图
    text	      % 在图开放置字符串
    bar	          % 创建条形图
    loglog	      % 创建日志记录图
    polar	      % 创建极坐标图
    semilogx      % 创建半标记图(对数横坐标)
    semilogy      % 创建半标记图(对数纵坐标)
    stairs	      % 创建梯形图
    stem	      % 创建茎图
```