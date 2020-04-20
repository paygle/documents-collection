# Mybatis 配置

## mybatis-config 配置

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE configuration PUBLIC "-//mybatis.org//DTD Config 3.0//EN" "http://mybatis.org/dtd/mybatis-3-config.dtd">
<!--配置一个全部参数的样例-->
<configuration>

    <!--1. 属性（properties） 这些属性可以在外部进行配置，并可以进行动态替换，也可以在 properties 元素的子元素中设置 -->
    <properties resource="configs.properties">
        <!--property里面的属性全局均可使用-->
        <property name="username" value="root"/>
        <property name="password" value="root"/>
    </properties>

    <!--2. 全局配置参数， 这是 MyBatis 中极为重要的调整设置，它们会改变 MyBatis 的运行时行为-->
    <settings>
        <!--设置是否启用缓存-->
        <setting name="cacheEnabled" value="true"/>
        <!--设置是否启用懒加载-->
        <setting name="lazyLoadingEndbled" value="true"/>
        <setting name="multipleResultSetsEnabled" value="true"/>
        <setting name="useColumnLabel" value="true"/>
        <setting name="useGeneratedKeys" value="false"/>
        <setting name="autoMappingBehavior" value="PARTIAL"/>
        <setting name="autoMappingUnknownColumnBehavior" value="WARNING"/>
        <setting name="defaultExecutorType" value="SIMPLE"/>
        <setting name="defaultStatementTimeout" value="25"/>
        <setting name="defaultFetchSize" value="100"/>
        <setting name="safeRowBoundsEnabled" value="false"/>
        <setting name="mapUnderscoreToCamelCase" value="false"/>
        <setting name="localCacheScope" value="SESSION"/>
        <setting name="jdbcTypeForNull" value="OTHER"/>
        <setting name="lazyLoadTriggerMethods" value="equals,clone,hashCode,toString"/>
    </settings>

    <!--3.别名设置，为Java 类型设置一个缩写名字，仅用于 XML 配置，意在降低冗余的全限定类名书写-->
    <typeAliases>
        <typeAlias alias="user" type="cn.com.mybatis.po.User"/>
        <typeAlias alias="integer" type="java.lang.Integer"/>
    </typeAliases>

    <!--4.类型转换器，处理任意继承了 Enum 的类-->
    <typeHandlers>
        <!--一个简单的类型转换器-->
        <typeHandler handler="org.apache.ibatis.type.EnumOrdinalTypeHandler" javaType="java.math.RoundingMode"/>
    </typeHandlers>

    <!--5.对象工厂， 每次创建结果对象的新实例时，它都会使用一个对象工厂（ObjectFactory）实例来完成实例化工作-->
    <objectFactory type="org.mybatis.example.ExampleObjectFactory">
        <!--对象工厂注入的参数-->
        <property name="someProperty" value="100"/>
    </objectFactory>

    <!--6.插件，在映射语句执行过程中的某一点进行拦截调用-->
    <plugins>
        <plugin interceptor="org.mybatis.example.ExamplePlugin">
            <property name="someProperty" value="100"/>
        </plugin>
    </plugins>

    <!--7.environments数据库环境配置-->
    <!--和Spring整合后environments配置将被废除-->
    <environments default="development">
        <environment id="development">
            <transactionManager type="JDBC"/>
            <!--数据库连接池 -->
            <dataSource type="POOLED">
                <property name="driver" value="${driver}"/>
                <property name="url" value="${url}"/>
                <property name="username" value="${username}"/>
                <property name="password" value="${password}"/>
            </dataSource>
        </environment>
    </environments>

    <!--加载映射器（mappers）文件-->
    <mappers>
        <!-- 使用相对于类路径的资源引用 -->
        <mapper resource="mapper/UserMapper.xml"/>

        <!-- 使用完全限定资源定位符（URL） -->
        <mapper url="file:///var/mappers/AuthorMapper.xml"/>

        <!-- 使用映射器接口实现类的完全限定类名 -->
        <mapper class="org.mybatis.builder.AuthorMapper"/>

        <!-- 将包内的映射器接口实现全部注册为映射器 -->
        <package name="org.mybatis.builder"/>
    </mappers>

    <!-- 数据库厂商标识（databaseIdProvider）-->
    <databaseIdProvider type="DB_VENDOR">
        <property name="SQL Server" value="sqlserver"/>
        <property name="DB2" value="db2"/>
        <property name="Oracle" value="oracle" />
    </databaseIdProvider>
</configuration>
```

## mybatis mapper 详解

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="org.apache.ibatis.submitted.rounding.Mapper">

	<resultMap type="org.apache.ibatis.submitted.rounding.User" id="usermap">
		<id column="id" property="id"/>
		<result column="name" property="name"/>
		<result column="funkyNumber" property="funkyNumber"/>
		<result column="roundingMode" property="roundingMode"/>
	</resultMap>

	<select id="getUser" resultMap="usermap">
		select * from users
	</select>

	<insert id="insert">
	    insert into users (id, name, funkyNumber, roundingMode) values (
	    	#{id}, #{name}, #{funkyNumber}, #{roundingMode}
	    )
	</insert>

	<resultMap type="org.apache.ibatis.submitted.rounding.User" id="usermap2">
		<id column="id" property="id"/>
		<result column="name" property="name"/>
		<result column="funkyNumber" property="funkyNumber"/>
		<result column="roundingMode" property="roundingMode" typeHandler="org.apache.ibatis.type.EnumTypeHandler"/>
	</resultMap>

	<select id="getUser2" resultMap="usermap2">
		select * from users2
	</select>

	<insert id="insert2">
	    insert into users2 (id, name, funkyNumber, roundingMode) values (
	    	#{id}, #{name}, #{funkyNumber}, #{roundingMode, typeHandler=org.apache.ibatis.type.EnumTypeHandler}
	    )
	</insert>

</mapper>
```