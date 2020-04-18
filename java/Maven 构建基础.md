### 安装文件到本地仓库（配置完成Maven环境变量之后）：

mvn install:install-file -DgroupId=com.oracle -DartifactId=ojdbc8 -Dversion=12.2.0.1 -Dpackaging=jar -Dfile=ojdbc8.jar 

### mvn help用法
在命令控制台，进入 /project 目录，执行以下命令：
```shell
# 检查当前Maven环境启用的文件
C:\project>mvn help:effective-settings

# 查看当前项目的pom配置，包括所有依赖
C:\project>mvn help:effective-pom

# 查看当前处于激活状态的profile
C:\project>mvn help:active-profiles

# 指定使用某个配置文件执行Maven命令
C:\project>mvn -s <filepath> <goal>
C:\project>mvn -s ~/.m2/settings_local.xml clean deploy

# 检查环境
C:\project>mvn -X

# 打印所有可用的环境变量和Java系统属性
C:\project>mvn help:system
```

### Maven 构建生命周期

|阶段	 | 处理 | 描述|
|------|----|----------|
|validate | 验证项目 | 验证项目是否正确且所有必须信息是可用的
|compile  | 执行编译 | 源代码编译在此阶段完成
|Test	  | 测试 | 使用适当的单元测试框架（例如JUnit）运行测试。
|package  | 打包 | 创建JAR/WAR包如在 pom.xml 中定义提及的包
|verify	  | 检查 | 对集成测试的结果进行检查，以保证质量达标
|install  | 安装 | 安装打包的项目到本地仓库，以供其他项目使用
|deploy	  | 部署 | 拷贝最终的工程包到远程仓库中，以共享给其他开发人员和工程

一个插件目标代表一个特定的任务（比构建阶段更为精细），这有助于项目的构建和管理。这些目标可能被绑定到多个阶段或者无绑定。不绑定到任何构建阶段的目标可以在构建生命周期之外通过直接调用执行。这些目标的执行顺序取决于调用目标和构建阶段的顺序。


### Default (Build) 生命周期
|生命周期阶段	|描述
|------------|--------
|validate（校验）	| 	校验项目是否正确并且所有必要的信息可以完成项目的构建过程。
|initialize（初始化）	| 	初始化构建状态，比如设置属性值。
|generate-sources（生成源代码）	| 	生成包含在编译阶段中的任何源代码。
|process-sources（处理源代码）	| 	处理源代码，比如说，过滤任意值。
|generate-resources（生成资源文件）	| 生成将会包含在项目包中的资源文件。
|process-resources （处理资源文件）	| 复制和处理资源到目标目录，为打包阶段最好准备。
|compile（编译）|	编译项目的源代码。
|process-classes（处理类文件）|	处理编译生成的文件，比如说对Java class文件做字节码改善优化。
|generate-test-sources（生成测试源代码）|	生成包含在编译阶段中的任何测试源代码。
|process-test-sources（处理测试源代码）|	处理测试源代码，比如说，过滤任意值。
|generate-test-resources（生成测试资源文件）|	为测试创建资源文件。
|process-test-resources（处理测试资源文件）|	复制和处理测试资源到目标目录。
|test-compile（编译测试源码）|	编译测试源代码到测试目标目录.
|process-test-classes（处理测试类文件）|	处理测试源码编译生成的文件。
|test（测试）|	使用合适的单元测试框架运行测试（Juint是其中之一）。
|prepare-package（准备打包）|	在实际打包之前，执行任何的必要的操作为打包做准备。
|package（打包）|	将编译后的代码打包成可分发格式的文件，比如JAR、WAR或者EAR文件。
|pre-integration-test（集成测试前）|	在执行集成测试前进行必要的动作。比如说，搭建需要的环境。
|integration-test|（集成测试）	处理和部署项目到可以运行集成测试环境中。
|post-integration-test（集成测试后）|	在执行集成测试完成后进行必要的动作。比如说，清理集成测试环境。
|verify （验证）|	运行任意的检查来验证项目包有效且达到质量标准。
|install（安装）|	安装项目包到本地仓库，这样项目包可以用作其他本地项目的依赖。
|deploy（部署）|	将最终的项目包复制到远程仓库中与其他开发者和项目共享。





