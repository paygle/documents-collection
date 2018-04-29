# Matlab 语法基础

* [基础语法](#basic)
* [矩阵和数组](#matrix)
* [二维图和三维图](#plot23)
* [数据类型](#datatype)
* [if运算](#ifelse)
* [循环语句](#while)
* [函数](#funcs)


## <a name="basic"></a>基础语法
  * MATLAB为某些数学符号提供了一些特殊表达式，如pi为π，Inf为∞，i(和j)为√-1，.Nan代表“非数字”等。
  * 在MATLAB中使用分号(;)
  * 百分比符号(%)用于指示注释行
  * 变量名称由一个字母组成，后跟任意数字的字母，数字或下划线。

  * 冒号(:)是MATLAB中最有用的操作符之一。它用于创建向量，下标数组，并为迭代指定。

```m

    A(:,j)	        % 是A的第j列
    A(i,:)	        % 是A的第i行
    A(:,:)	        % 是等效的二维数组。对于矩阵，这与A相同。
    A(j:k)	        % 是A(j), A(j+1),...,A(k)
    A(:,j:k)	    % 是A(:,j), A(:,j+1),...,A(:,k)
    A(:,:,k)	    % 是三维数组A的第k页
    A(i,j,k,:)	    % 是四维数组A中的向量。向量包括A(i，j，k，1)，A(i，j，k，2)，A(i，j，k，3)等。
    A(:)	        % 是A的所有要素，被视为单列。在赋值语句的左侧，A(:)填充A，从之前保留其形状。在这种情况下，右侧必须包含与A相同数量的元素。

```

  * 常用的运算符和特殊字符

```m
    +	    % 相加; 加法运算符。
    -	    % 相减; 减法运算符。
    *	    % 标量和矩阵乘法运算符。
    .*	    % 阵列乘法运算符。
    ^	    % 标量和矩阵求幂运算符。
    .^	    % 阵列求幂运算符。
    \	    % 左除法运算符。
    /	    % 右除法运算符。
    .\	    % 阵列左除法运算符。
    ./	    % 右除法运算符。
    :	    % 冒号; 生成规则间隔的元素，并表示整个行或列。
    ( )	    % 括号; 包含函数参数和数组索引; 覆盖优先级。
    [ ]	    % 括号; 罩住阵列元素。
    .	    % 小数点。
    …	    % 省略号; 行连续运算符
    ,	    % 逗号; 分隔一行中的语句和元素
    ;	    % 分号; 分隔列并抑制输出显示。
    %	    % 百分号;指定一个注释并指定格式。
    _	    % 引用符号和转置运算符。
    ._	    % 非共轭转置运算符。
    =	    % 赋值运算符。
    '	    % 矩阵转置。 '是A的线性代数转置。对于复数矩阵，这是复共轭转置。
    .'	    % 数组转置。.'是A的数组转置。对于复数矩阵，这不涉及共轭。

    <	    % 小于
    <=	    % 小于或等于
    >	    % 大于
    >=	    % 大于或等于
    ==	    % 等于
    ~=	    % 不等于

    &&      % 条件与，或用AND
    OR      % 条件或，或用OR
    NOT     % 非
    &       % 位运算
    |       % 位运算
    ~       % 位运算
    ^       % 位运算
```

  * 特殊变量和常数

```m
    ans	    % 最近的回应/回答。
    eps	    % 浮点精度精度。
    i,j	    % 虚构单位√-1。
    Inf	    % 无穷
    NaN	    % 未定义的数值结果(非数字)。
    pi	    % 数字π
```

## <a name="datatype"></a>数据类型

```m

    int8	    % 8位有符号整数
    uint8	    % 8位无符号整数
    int16	    % 16位有符号整数
    uint16	    % 16位无符号整数
    int32	    % 32位有符号整数
    uint32	    % 32位无符号整数
    int64	    % 64位有符号整数
    uint64	    % 64位无符号整数
    single	    % 单精度数值数据
    double	    % 双精度数值数据
    logical	    % 逻辑值为1或0，分别代表true和false
    char	    % 字符数据(字符串作为字符向量存储)
    单元格阵列	  % 索引单元阵列，每个都能够存储不同维数和数据类型的数组
    结构体	     % C型结构，每个结构具有能够存储不同维数和数据类型的数组的命名字段
    函数处理     % 指向一个函数的指针
    用户类	     % 用户定义的类构造的对象
    Java类	    % 从Java类构造的对象

    % 示例
    str = 'Hello World!'
    n = 2345
    d = double(n)
    un = uint32(789.50)
    rn = 5678.92347
    c = int32(rn)
```

  * 数据类型转换

```m

    char	        % 转换为字符数组(字符串)
    int2str	        % 将整数数据转换为字符串
    mat2str	        % 将矩阵转换为字符串
    num2str	        % 将数字转换为字符串
    str2double      % 将字符串转换为双精度值
    str2num	        % 将字符串转换为数字
    native2unicode	% 将数字字节转换为Unicode字符
    unicode2native	% 将Unicode字符转换为数字字节
    base2dec        % 将基数N字符串转换为十进制数
    bin2dec	        % 将二进制数字串转换为十进制数
    dec2base        % 将十进制转换为字符串中的N数字
    dec2bin	        % 将十进制转换为字符串中的二进制数
    dec2hex	        % 将十进制转换为十六进制数字
    hex2dec	        % 将十六进制数字字符串转换为十进制数
    hex2num	        % 将十六进制数字字符串转换为双精度数字
    num2hex	        % 将单数转换为IEEE十六进制字符串
    cell2mat        % 将单元格数组转换为数组
    cell2struct     % 将单元格数组转换为结构数组
    cellstr	        % 从字符数组创建字符串数组
    mat2cell        % 将数组转换为具有潜在不同大小的单元格的单元阵列
    num2cell        % 将数组转换为具有一致大小的单元格的单元阵列
    struct2cell     % 将结构转换为单元格数组

```

  * 数据类型确定

```m

    is	            % 检测状态
    isa	            % 确定输入是否是指定类的对象
    iscell          % 确定输入是单元格数组
    iscellstr       % 确定输入是字符串的单元格数组
    ischar	        % 确定项目是否是字符数组
    isfield	        % 确定输入是否是结构数组字段
    isfloat	        % 确定输入是否为浮点数组
    ishghandle      % 确定是否用于处理图形对象句柄
    isinteger       % 确定输入是否为整数数组
    isjava	        % 确定输入是否为Java对象
    islogical       % 确定输入是否为逻辑数组
    isnumeric       % 确定输入是否是数字数组
    isobject        % 确定输入是否为MATLAB对象
    isreal	        % 检查输入是否为实数数组
    isscalar        % 确定输入是否为标量
    isstr           % 确定输入是否是字符数组
    isstruct        % 确定输入是否是结构数组
    isvector        % 确定输入是否为向量
    class           % 确定对象的类
    validateattributes	  % 检查数组的有效性
    whos            % 在工作区中列出变量，其大小和类型

```



## <a name="matrix"></a>矩阵和数组

  * 数组创建，要创建每行包含四个元素的数组，请使用逗号 (,) 或空格分隔各元素。这种数组为行向量。

```m
  a = [1 2 3 4]

  % 要创建包含多行的矩阵，请使用分号分隔各行
  a = [1 2 3; 4 5 6; 7 8 10]

  % 要转置矩阵，请使用单引号 (')：
  a'

  % 计算 a 的各个元素的三次方
  a.^3

```

  * 串联，连接数组以便形成更大数组的过程。

```m
  % 列数增加，如果各数组具有相同的列数，则可以使用分号垂直串联
  A = [a; a]

  % 使用逗号将彼此相邻的数组串联起来称为水平串联。每个数组必须具有相同的行数。
  A = [a, a]
```

  * 复数，包含实部和虚部，虚数单位是 -1 的平方根。

```m
  sqrt(-1)
  ans = 0.0000 + 1.0000i

  % 要表示复数的虚部，请使用 i 或 j。
  c = [3+4i, 4+3j; -i, 10j]
```

  * 引用数组中的特定元素有两种方法。最常见的方法是指定行和列下标

```m

  A(4, 2)

  % 一种方法不太常用，但有时非常有用，即使用单一下标按顺序向下遍历每一列
  A(12)

  % 可以在赋值语句左侧指定当前维外部的元素。数组大小会增大以便容纳新元素
  A(4, 8) = 19

  % 要引用多个数组元素，请使用冒号运算符，这使您可以指定一个格式为 start:end 的范围。
  A(1:3, 2)

  % 单独的冒号（没有起始值或结束值）指定该维中的所有元素
  A(3,:)

  % 冒号运算符还允许您使用较通用的格式 start:step:end 创建等距向量值
  B = 0:10:100

```

## <a name="plot23"></a>二维图和三维图

  * 线图，要创建二维线图，请使用 plot 函数。

```m
  x = 0:pi/100:2*pi;
  y = sin(x);
  plot(x,y)

  % 可以标记轴并添加标题。
  xlabel('x')
  ylabel('sin(x)')
  title('Plot of the Sine Function')

```

## <a name="ifelse"></a>if运算

```m

  % if…end语句	if ... end语句包含一个布尔表达式，后跟一个或多个语句。

    if <expression>
    % statement(s) will execute if the boolean expression is true
    <statements>
    end

    % 示例
    a = 10;
    % check the condition using if statement
    if a < 20
    % if condition is true then print the following
        fprintf('a is less than 20\n' );
    end
    fprintf('value of a is : %d\n', a);
```

```m

  % if…else…end语句	if语句可以跟随一个可选的else语句，当布尔表达式为false时，else语句块将执行。
  % if…elseif…elseif…else…end语句	if语句后面可以有一个(或多个)可选elseif ...和一个else语句，这对于测试各种条件非常有用。

    if <expression>
    % statement(s) will execute if the boolean expression is true
    <statement(s)>
    else
    <statement(s)>
    % statement(s) will execute if the boolean expression is false
    end

    % 示例
    a = 100;
    % check the boolean condition
    if a < 20 
        % if condition is true then print the following
        fprintf('a is less than 20\n' );
    else
        % if condition is false then print the following
        fprintf('a is not less than 20\n' );
    end
        fprintf('value of a is : %d\n', a);
```

```m

  % 嵌套if语句	可以在一个if或elseif语句中使用另一个if或elseif语句。

    if <expression 1>
    % Executes when the boolean expression 1 is true
        if <expression 2>
            % Executes when the boolean expression 2 is true
        end
    end

    % 示例
    a = 100;
    b = 200;
    % check the boolean condition 
    if( a == 100 )
        % if condition is true then check the following 
        if( b == 200 )
            % if condition is true then print the following 
            fprintf('Value of a is 100 and b is 200\n' );
        end

    end
    fprintf('Exact value of a is : %d\n', a );
    fprintf('Exact value of b is : %d\n', b );
```

```m
  % switch语句	switch语句用来测试一个变量与值列表的相等性。
  % 嵌套switch语句	可以在一个switch语句中使用一个switch语句。

    switch <switch_expression>
        case <case_expression>
            <statements>
        case <case_expression>
            <statements>
            ...
            ...
        otherwise
            <statements>
    end

    % 示例
    grade = 'B';
    switch(grade)
        case 'A' 
            fprintf('Excellent!\n' );
        case 'B' 
            fprintf('Well done\n' );
        case 'C' 
            fprintf('Well done\n' );
        case 'D'
            fprintf('You passed\n' );

        case 'F' 
            fprintf('Better try again\n' );

        otherwise
            fprintf('Invalid grade\n' );
    end
```

## <a name="while"></a>循环语句

  * break语句	终止循环语句，并将执行转移到循环之后的语句。
  * continue语句	导致循环跳过主体的剩余部分，并在重申之前立即重新测试其状态。

  * while 语句

```m

    while <expression>
        <statements>
    end

    % 示例

    a = 10;
    % while loop execution
    while( a < 20 )
        fprintf('value of a: %d\n', a);
        a = a + 1;
    end

```

  * for循环

```m

    for index = values
        <program statements>
                    ...
    end

    % 示例

    for a = 1.0: -0.1: 0.0
        disp(a)
    end

    for a = [24,18,17,23,28]
        disp(a)
    end

    % 值(values)具有以下格式
    initval:endval	% index变量从initval到endval每次递增1，并重复程序语句的执行，直到index大于endval。

    initval:step:endval	 % 通过每次迭代值步长(step)增加索引(index)的值，或者当step为负时递减。

    valArray	% 在每个迭代中从数组valArray的后续列创建列向量索引。 例如，在第一次迭代中，index = valArray(:，1)。 循环最多执行n次，其中n是由numel(valArray，1，:)给出的valArray的列数。valArray可以是任何MATLAB数据类型，包括字符串，单元格数组或结构体。

```

## <a name="funcs"></a>函数

  * MATLAB® 程序文件可以包含用于多个函数的代码。在函数文件中，第一个函数称为主函数(必须与文件同名）。此函数对其他文件中的函数可见，或者您也可以从命令行调用它。文件中的其他函数称为局部函数，它们可以任意顺序出现在主函数后面。局部函数仅对同一文件中的其他函数可见。它们等效于其他编程语言的子例程，有时被称为子函数。

  * function [out1,out2, ..., outN] = myfun(in1,in2,in3, ..., inN)

```m

    function max = mymax(n1, n2, n3, n4, n5)
    % This function calculates the maximum of the
    % five numbers given as input
    max =  n1;
    if(n2 > max)
        max = n2;
    end
    if(n3 > max)
        max = n3;
    end
    if(n4 > max)
        max = n4;
    end
    if(n5 > max)
        max = n5;
    end

```

* 匿名函数

```m
    sqr = @(x) x.^2;
```

    变量 sqr 是一个函数句柄。@ 运算符创建句柄，@ 运算符后面的圆括号 () 包括函数的输入参数。该匿名函数接受单个输入 x，并显式返回单个输出，即大小与包含平方值的 x 相同的数组。

