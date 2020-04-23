#  maven自动部署到tomcat9

  maven的自动部署功能可以很方便的将maven工程自动部署到远程tomcat服务器

## 第一步：配置 Tomcat 访问权限

  首先，我们需要先打开 Tomcat 的 manager 功能，找到 conf 文件夹下的 tomcat-users.xml文件中的&lt;tomcat-users&gt;标签，
  然后添加如下内容（可以直接在其文档注释部分找到对应的模版，然后进行修改）：

```xml
<role rolename="manager-gui"/> 
<role rolename="manager-script"/>
<role rolename="manager-jmx"/>
<role rolename="manager-status"/>
<user username="admin" password="1234" roles="manager-gui,manager-script,manager-jmx,manager-status" />
```

  配置好之后，ctrl+s 保存文件。紧接着，双击 tomcat 解压包中 bin 目录下的 startup.bat 命令进行启动Tomcat服务器。在浏览器地址来中进行访问http://localhost:8080/manager，按下 Enter 回车键，即可看到弹窗，需要我们输入上面配置好的用户名和密码，才能进行登录，如果顺利则请进入下一步。

## 第二步：配置maven的settings.xml

  在 conf/settings.xml 文件中的标签 <servers> 添加子标签。通过标签名字，我们知道这主要是为了让 maven 去关联我们的 Tomcat 服务器。
  注意，这里配置的 username 和 password 一定要和 tomcat 中的 tomcat_user.xml 中一致，否则关联不起来。

```xml
<server> 
    <id>tomcat9</id>
    <username>admin</username>
    <password>1234</password>
</server>
```

## 第三步：回到我们的 Eclipse 中，然后在 pom.xml 文件中，在原来 tomcat7 插件的基础上，往 <project> 下添加 <configuration> 子标签进行配置即可。

```xml
<build>
  <plugins>
    <plugin>
      <groupId>org.apache.maven.plugins</groupId>
      <artifactId>maven-compiler-plugin</artifactId>
      <version>3.8.1</version>
      <configuration>
        <source>1.8</source>
        <target>1.8</target>
      </configuration>
    </plugin>

    <plugin>
      <groupId>org.apache.tomcat.maven</groupId>
      <artifactId>tomcat7-maven-plugin</artifactId>
      <version>2.2</version>
      <configuration>
        <!-- 直接访问 Tomcat 服务器的 manager -->
        <url>http://localhost:8080/manager</url>
        <server>tomcat9</server>
        <username>admin</username>
        <password>1234</password>
        <update>true</update>
        <path>/webapp</path>
      </configuration>
    </plugin>

  </plugins>
</build>
```

## 执行命令

  (执行过程中，tomcat9要先启动)

  1. Run as → clean install
  2. Run as → tomcat7:deploy 注：第1次部署执行
  3. Run as → tomcat7:redeploy 注：第2次或以后需要重新发布执行
  4. Run as → tomcat7:run 注：部署到 tomcat 中启动（执行之前先关闭tomcat9，如果部署到本地，防止端口占用启动不了）


## eclipse内使用tomcat项目究竟被部署到了哪里

  eclipse并不像MyEclipse默认将项目部署到tomcat安装目录下的webapps中，而默认部署到工作目录(workspace)下的.metadata/.plugins/org.eclipse.wst.server.core/tmp0/wtpwebapps中。（tmp0、tmp1的不同是目前这个server容器的顺序）

- 如何修改
  
  为了使项目默认部署到tomcat安装目录下的webapps中，show view-> servers-> 找到需要修改的tomcat-> 右击+open或者双击

  1. 停止eclipse内的tomcat服务器（stop）
  2. 删除该容器中部署的项目(add and remove)
  3. 清除该容器相关数据(clean), 不清除是无法修改 Servers location的
  4. 打开tomcat的修改界面(open)
  5. 找到Servers location, 选择第二个(User tomcat Installation)
  6. 修改deploy path为webapps
  7. 保存关闭
   
