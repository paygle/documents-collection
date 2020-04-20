# Mybatis 配置

## mybatis-config 配置

[点击获取 mybatis-config 配置帮助](https://mybatis.org/mybatis-3/zh/configuration.html)

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
            <!--
               JDBC 事务管理器被用作当应用程序负责管理数据库连接的生命周期（提交、回退等等）的时候。内部将使用 JdbcTransactionFactory 类创建事务管理器
               MANAGED 事务管理器是当由应用服务器负责管理数据库连接生命周期的时候使用。内部使用 ManagedTransactionFactory 类创建事务管理器
            -->
            <transactionManager type="JDBC"/>

            <!-- mybatis提供了3种数据源类型，分别是：POOLED, UNPOOLED, JNDI -->
            <!-- POOLED – 这种数据源的实现利用“池”的概念将 JDBC 连接对象组织起来，避免了创建新的连接实例时所必需的初始化和认证时间 -->
            <!-- UNPOOLED – 这个数据源的实现只是每次被请求时打开和关闭连接 -->
            <!-- JNDI – 这个数据源的实现是为了能在如 EJB 或应用服务器这类容器中使用，容器可以集中或在外部配置数据源，然后放置一个 JNDI 上下文的引用。-->
            <dataSource type="POOLED">
                <property name="driver" value="${driver}"/>
                <property name="url" value="${url}"/>
                <property name="username" value="${username}"/>
                <property name="password" value="${password}"/>
            </dataSource>
            <!-- 
              JNDI – 这种数据源配置只需要两个属性:  
                initial_context – 属性用来在 InitialContext 中寻找上下文，是可选属性，默认直接从 InitialContext 中寻找。
                data_source – 这是引用数据源实例位置的上下文的路径，没有提供时则直接在 InitialContext 中查找。
            -->
            <dataSource type="JNDI">
                <property name="data_source" value="PUBUSERDATA" />
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

## Mapper SQL 映射文件详解

[点击获取 Mapper 映射帮助](https://mybatis.org/mybatis-3/zh/sqlmap-xml.html)

SQL 映射文件只有很少的几个顶级元素（按照应被定义的顺序列出）：

  - cache – 该命名空间的缓存配置。
  - cache-ref – 引用其它命名空间的缓存配置。
  - resultMap – 描述如何从数据库结果集中加载对象，是最复杂也是最强大的元素。
  - sql – 可被其它语句引用的可重用语句块。
  - insert – 映射插入语句。
  - update – 映射更新语句。
  - delete – 映射删除语句。
  - select – 映射查询语句。

### select 

```xml
<select
  id="selectPerson"
  parameterType="int"
  parameterMap="deprecated"
  resultType="hashmap"
  resultMap="personResultMap"
  flushCache="false"
  useCache="true"
  timeout="10"
  fetchSize="256"
  statementType="PREPARED"
  resultSetType="FORWARD_ONLY">
  SELECT * FROM PERSON WHERE ID = #{id}
</select>
```

| 属性 |	描述
|--------|-------------------------|
| id |	在命名空间中唯一的标识符，可以被用来引用这条语句。
| parameterType |	将会传入这条语句的参数的类全限定名或别名。这个属性是可选的，因为 MyBatis 可以通过类型处理器（TypeHandler）推断出具体传入语句的参数，默认值为未设置（unset）。
| resultType |	期望从这条语句中返回结果的类全限定名或别名。 注意，如果返回的是集合，那应该设置为集合包含的类型，而不是集合本身的类型。 resultType 和 resultMap 之间只能同时使用一个。
| resultMap |	对外部 resultMap 的命名引用。结果映射是 MyBatis 最强大的特性，如果你对其理解透彻，许多复杂的映射问题都能迎刃而解。 resultType 和 resultMap 之间只能同时使用一个。
| flushCache |	将其设置为 true 后，只要语句被调用，都会导致本地缓存和二级缓存被清空，默认值：false。
| useCache |	将其设置为 true 后，将会导致本条语句的结果被二级缓存缓存起来，默认值：对 select 元素为 true。
| timeout |	这个设置是在抛出异常之前，驱动程序等待数据库返回请求结果的秒数。默认值为未设置（unset）（依赖数据库驱动）。
| fetchSize |	这是一个给驱动的建议值，尝试让驱动程序每次批量返回的结果行数等于这个设置值。 默认值为未设置（unset）（依赖驱动）。
| statementType |	可选 STATEMENT，PREPARED 或 CALLABLE。这会让 MyBatis 分别使用 Statement，PreparedStatement 或 CallableStatement，默认值：PREPARED。
| resultSetType |	FORWARD_ONLY，SCROLL_SENSITIVE, SCROLL_INSENSITIVE 或 DEFAULT（等价于 unset） 中的一个，默认值为 unset （依赖数据库驱动）。
| databaseId |	如果配置了数据库厂商标识（databaseIdProvider），MyBatis 会加载所有不带 databaseId 或匹配当前 databaseId 的语句；如果带和不带的语句都有，则不带的会被忽略。
| resultOrdered |	这个设置仅针对嵌套结果 select 语句：如果为 true，将会假设包含了嵌套结果集或是分组，当返回一个主结果行时，就不会产生对前面结果集的引用。 这就使得在获取嵌套结果集的时候不至于内存不够用。默认值：false。
| resultSets |	这个设置仅适用于多结果集的情况。它将列出语句执行后返回的结果集并赋予每个结果集一个名称，多个名称之间以逗号分隔。

### insert, update 和 delete 实现非常接近

```xml
<insert
  id="insertAuthor"
  parameterType="domain.blog.Author"
  flushCache="true"
  statementType="PREPARED"
  keyProperty=""
  keyColumn=""
  useGeneratedKeys=""
  timeout="20">
  insert into Author (id,username,password,email,bio)
  values (#{id},#{username},#{password},#{email},#{bio})
</insert>

<update
  id="updateAuthor"
  parameterType="domain.blog.Author"
  flushCache="true"
  statementType="PREPARED"
  timeout="20">
  update Author set
    username = #{username},
    password = #{password},
    email = #{email},
    bio = #{bio}
  where id = #{id}
</update>

<delete
  id="deleteAuthor"
  parameterType="domain.blog.Author"
  flushCache="true"
  statementType="PREPARED"
  timeout="20">
    delete from Author where id = #{id}
</delete>
```

| 属性 |	描述
|-----------|---------------------|
| id |	在命名空间中唯一的标识符，可以被用来引用这条语句。
| parameterType |	将会传入这条语句的参数的类全限定名或别名。这个属性是可选的，因为 MyBatis 可以通过类型处理器（TypeHandler）推断出具体传入语句的参数，默认值为未设置（unset）。
| flushCache |	将其设置为 true 后，只要语句被调用，都会导致本地缓存和二级缓存被清空，默认值：（对 insert、update 和 delete 语句）true。
| timeout |	这个设置是在抛出异常之前，驱动程序等待数据库返回请求结果的秒数。默认值为未设置（unset）（依赖数据库驱动）。
| statementType |	可选 STATEMENT，PREPARED 或 CALLABLE。这会让 MyBatis 分别使用 Statement，PreparedStatement 或 CallableStatement，默认值：PREPARED。
| useGeneratedKeys |	（仅适用于 insert 和 update）这会令 MyBatis 使用 JDBC 的 getGeneratedKeys 方法来取出由数据库内部生成的主键（比如：像 MySQL 和 SQL Server 这样的关系型数据库管理系统的自动递增字段），默认值：false。
| keyProperty |	（仅适用于 insert 和 update）指定能够唯一识别对象的属性，MyBatis 会使用 getGeneratedKeys 的返回值或 insert 语句的 selectKey 子元素设置它的值，默认值：未设置（unset）。如果生成列不止一个，可以用逗号分隔多个属性名称。
| keyColumn |	（仅适用于 insert 和 update）设置生成键值在表中的列名，在某些数据库（像 PostgreSQL）中，当主键列不是表中的第一列的时候，是必须设置的。如果生成列不止一个，可以用逗号分隔多个属性名称。
| databaseId |	如果配置了数据库厂商标识（databaseIdProvider），MyBatis 会加载所有不带 databaseId 或匹配当前 databaseId 的语句；如果带和不带的语句都有，则不带的会被忽略。

### sql 代码片段，以便在其它语句中使用

```xml
<sql id="userColumns"> ${alias}.id,${alias}.username,${alias}.password </sql>

<select id="selectUsers" resultType="map">
  select
    <include refid="userColumns"><property name="alias" value="t1"/></include>,
    <include refid="userColumns"><property name="alias" value="t2"/></include>
  from some_table t1
    cross join some_table t2
</select>
```

### 结果映射（resultMap）

| 属性 |	描述
|-----|-----------------------|
| id |	当前命名空间中的一个唯一标识，用于标识一个结果映射。
| type |	类的完全限定名, 或者一个类型别名（关于内置的类型别名，可以参考上面的表格）。
| autoMapping |	如果设置这个属性，MyBatis 将会为本结果映射开启或者关闭自动映射。 这个属性会覆盖全局的属性 autoMappingBehavior。默认值：未设置（unset）。

```xml
<!-- 非常复杂的结果映射 -->
<resultMap id="detailedBlogResultMap" type="Blog">
  <constructor>
    <idArg column="blog_id" javaType="int"/>
  </constructor>
  <result property="title" column="blog_title"/>
  <association property="author" javaType="Author">
    <id property="id" column="author_id"/>
    <result property="username" column="author_username"/>
    <result property="password" column="author_password"/>
    <result property="email" column="author_email"/>
    <result property="bio" column="author_bio"/>
    <result property="favouriteSection" column="author_favourite_section"/>
  </association>
  <collection property="posts" ofType="Post">
    <id property="id" column="post_id"/>
    <result property="subject" column="post_subject"/>
    <association property="author" javaType="Author"/>
    <collection property="comments" ofType="Comment">
      <id property="id" column="comment_id"/>
    </collection>
    <collection property="tags" ofType="Tag" >
      <id property="id" column="tag_id"/>
    </collection>
    <discriminator javaType="int" column="draft">
      <case value="1" resultType="DraftPost"/>
    </discriminator>
  </collection>
</resultMap>
```
 1. constructor - 用于在实例化类时，注入结果到构造方法中
    - idArg - ID 参数；标记出作为 ID 的结果可以帮助提高整体性能
    - arg - 将被注入到构造方法的一个普通结果
 2. id – 一个 ID 结果；标记出作为 ID 的结果可以帮助提高整体性能
 3. result – 注入到字段或 JavaBean 属性的普通结果
 4. association – 一个复杂类型的关联；许多结果将包装成这种类型
    - 嵌套结果映射 – 关联可以是 resultMap 元素，或是对其它结果映射的引用
 5. collection – 一个复杂类型的集合
    - 嵌套结果映射 – 集合可以是 resultMap 元素，或是对其它结果映射的引用
 6. discriminator – 使用结果值来决定使用哪个 resultMap
    - case – 基于某些值的结果映射
      嵌套结果映射 – case 也是一个结果映射，因此具有相同的结构和元素；或者引用其它的结果映射

#### Id 和 Result 的属性

| 属性 |	描述
|----------|-----------------------------|
| property |	映射到列结果的字段或属性。如果 JavaBean 有这个名字的属性（property），会先使用该属性。否则 MyBatis 将会寻找给定名称的字段（field）。 无论是哪一种情形，你都可以使用常见的点式分隔形式进行复杂属性导航。 比如，你可以这样映射一些简单的东西：“username”，或者映射到一些复杂的东西上：“address.street.number”。
| column |	数据库中的列名，或者是列的别名。一般情况下，这和传递给 resultSet.getString(columnName) 方法的参数一样。
| javaType |	一个 Java 类的全限定名，或一个类型别名（关于内置的类型别名，可以参考上面的表格）。 如果你映射到一个 JavaBean，MyBatis 通常可以推断类型。然而，如果你映射到的是 HashMap，那么你应该明确地指定 javaType 来保证行为与期望的相一致。
| jdbcType |	JDBC 类型，所支持的 JDBC 类型参见这个表格之后的“支持的 JDBC 类型”。 只需要在可能执行插入、更新和删除的且允许空值的列上指定 JDBC 类型。这是 JDBC 的要求而非 MyBatis 的要求。如果你直接面向 JDBC 编程，你需要对可以为空值的列指定这个类型。
| typeHandler |	我们在前面讨论过默认的类型处理器。使用这个属性，你可以覆盖默认的类型处理器。 这个属性值是一个类型处理器实现类的全限定名，或者是类型别名。


### mapper 实例

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