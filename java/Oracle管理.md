# Oracle 管理设置

## 查看用户具有的权限和角色
```sql
/* 1.查看所有用户 */
select * from dba_users;
select * from all_users;
select * from user_users;

/* 2.查看用户或角色系统权限(直接赋值给用户或角色的系统权限) */
select * from dba_sys_privs;
select * from user_sys_privs;

/* 3.查看角色(只能查看登陆用户拥有的角色)所包含的权限 */
sql>select * from role_sys_privs;

/* .查看用户对象权限 */
select * from dba_tab_privs;
select * from all_tab_privs;
select * from user_tab_privs;

/* 5.查看所有角色 */
select * from dba_roles;

/* 6.查看用户或角色所拥有的角色 */
select * from dba_role_privs;
select * from user_role_privs;

/* 7.查看哪些用户有sysdba或sysoper系统权限(查询时需要相应权限)
select * from V$PWFILE_USERS

/* 查看某一用户 wkk的拥有的权限 */
SQL> select * from dba_sys_privs where grantee='Wkk'

/* 比如我要查看用户 wzsb的拥有的角色 */
SQL> select * from dba_role_privs where grantee='WZSB'

/* 查看一个用户所有的权限及角色 */
select privilege from dba_sys_privs where grantee='WZSB' union
select privilege from dba_sys_privs where grantee in (select granted_role from dba_role_privs where grantee='WZSB' );

/* 8.当前用户被激活的全部角色 */
select * from session_roles;

/* 9.查看某个用户所拥有的角色 */

select * from dba_role_privs where grantee='用户名';

/* 10.查看某个角色所拥有的权限 */

select * from dba_sys_privs where grantee='CONNECT';

select * from session_privs; --当前用户所拥有的全部权限

/* 11.查看哪些用户具有DBA的角色 */

select grantee from dba_role_privs where granted_role='DBA';

/* 12.查看Oracle提供的系统权限 */

select name from sys.system_privilege_map;

/* 13. 查询当前用户可以访问的所有数据字典视图 */
select * from dict where comments like '%grant%';

/* 14.显示当前数据库的全称 */
select * from global_name;

/* 15.查询sytem用户的所属表空间 */

select username,default_tablespace from dba_users  where username=“sytem”

```

## 创建表空间

  CREATE TABLESPACE tablespaceName
  DATAFILE 'filename'
  [SIZE integer [K | M]]  
  [AUTOEXTEND [OFF | ON]];

```sql
/* 例如 */
  CREATE TABLESPACE geeksss
  DATAFILE 'D:\ORACLE\DATA\GEEKSSS.DBF'
  SIZE 10M 
  AUTOEXTEND ON;

/* 删除表空间 */
  DROP TABLESPACE tablespaceName;
```
  tablespaceName 是需创建的表空间名称；
  DATAFILE 指定组成表空间的一个或多个数据文件，当有多个数据文件时使用逗号分隔；
  filename 是数据文件的路径和名称；
  SIZE指定文件的大小，用K指定千字节大小，用M指定兆字节大小；
  AUTOEXTEND子句用来启用或禁用数据文件的自动扩展，设置为ON则空间使用完毕会自动扩展，设置为OFF则很容易出现表空间剩余容量为0的情况，使数据不能存储到数据库中。

## 自定义用户管理

  CREATE USER username
  IDENTIFIED BY password
  [DEFAULT TABLESPACE tablespaceName]
  [TEMPORARY TABLESPACE tablespaceName]

```sql
  CREATE USER martin
  IDENTIFIED BY martinpwd
  DEFAULT TABLESPACE geeksss
  TEMPORARY TABLESPACE temp;

  /* 将martin的口令修改为mpwd */
  ALTER USER martin IDENTIFIED BY mpwd;

  /* 删除用户martin */
  DROP USER martin CASCADE;

  /* 解锁用户martin */
  alter user test account unlock;

  /* 修改密码 */
  alter user system identified by password;
```
  username 是用户名，用户名必须是一个标识符；
  password 是用户口令，口令必须是一个标识符，且不区分大小写；
  DEFAULT或TEMPORARY TABLESPACE为用户确定默认表空间或临时表空间。

## 设置Oracle用户永不被锁

```sql
/*  查看RESOURCE_NAME 的 FAILED_LOGIN_ATTEMPTS 的值  */
select * from dba_profiles;

/*  修改为30次  */
alter profile default limit FAILED_LOGIN_ATTEMPTS 30;
/* 修改为无限次 */
alter profile default limit FAILED_LOGIN_ATTEMPTS unlimited;
/* 查看当前系统用户profile情况 */
select username,account_status,default_tablespace,profile from dba_users;
/* 分配新的profile */
alter user FS8004 profile fsl_profile;

/* 创建profile */
ALTER SYSTEM SET RESOURCE_LIMIT=TRUE;

create profile fsl_profile limit
 FAILED_LOGIN_ATTEMPTS unlimited
 PASSWORD_LOCK_TIME 1/24
 PASSWORD_REUSE_MAX 5
 PASSWORD_REUSE_TIME 1800
 PASSWORD_LIFE_TIME 360
 PASSWORD_GRACE_TIME 30
 SESSIONS_PER_USER UNLIMITED 
 CPU_PER_SESSION UNLIMITED
 CPU_PER_CALL UNLIMITED
 IDLE_TIME 360
 CONNECT_TIME UNLIMITED
 LOGICAL_READS_PER_SESSION UNLIMITED 
 LOGICAL_READS_PER_CALL UNLIMITED 
 PRIVATE_SGA UNLIMITED
 COMPOSITE_LIMIT UNLIMITED;

```

## 权限

　　create session  允许用户登录数据库权限

　　create table   允许用户创建表权限

　　unlimited tablespace  允许用户在其他表空间随意建表

## 角色

　　Oracle中常用的系统预定义角色如下：

    CONNECT ：需要连接上数据库的用户，特别是那些不需要创建表的用户，通常授予该角色。
    RESOURCE ：更为可靠和正式的数据库用户可以授予该角色，可以创建表、触发器、过程等。
    DBA ：数据库管理员角色，拥有管理数据库的最高权限，一个具有DBA角色的用户可以撤销任何别的用户甚至别的DBA权限，这是很危险的，所以不要轻易授予该角色。

　　CONNECT 角色： --是授予最终用户的典型权利，最基本的权力，能够连接到ORACLE数据库中，并在对其他用户的表有访问权限时，做SELECT、UPDATE、INSERTT等操作。
    ALTER SESSION --修改会话
    CREATE CLUSTER --建立聚簇
    CREATE DATABASE LINK --建立数据库链接
    CREATE SEQUENCE --建立序列
    CREATE SESSION --建立会话
    CREATE SYNONYM --建立同义词
    CREATE VIEW --建立视图
    RESOURCE 角色： --是授予开发人员的，能在自己的方案中创建表、序列、视图等。
    CREATE CLUSTER --建立聚簇
    CREATE PROCEDURE --建立过程
    CREATE SEQUENCE --建立序列
    CREATE TABLE --建表
    CREATE TRIGGER --建立触发器
    CREATE TYPE --建立类型

　　DBA角色，是授予系统管理员的，拥有该角色的用户就能成为系统管理员了，它拥有所有的系统权限

　　例：

```sql
#sqlplus /nolog

SQL> conn / as sysdba;

SQL> create user username identified by password  --username/password都是用户自定义

SQL> grant dba to username;

SQL> conn username/password

SQL> select * from user_sys_privs;
```
　　我们将从创建Oracle用户权限表开始谈起，然后讲解登陆等一般性动作，使大家对Oracle用户权限表有个深入的了解。

### 一、创建

　　sys;       // 系统管理员，拥有最高权限

　　system;    // 本地管理员，次高权限

　　scott;     // 普通用户，密码默认为tiger,默认未解锁

　　oracle有三个默认的用户名和密码~
　　1.用户名:sys密码:change_on_install
　　2.用户名:system密码:manager
　　3.用户名:scott密码:tiger

### 二、登陆

```sql
　　sqlplus / as sysdba;//登陆sys帐户

　　sqlplus sys as sysdba;//同上

　　sqlplus scott/tiger;//登陆普通用户scott
```

### 三、管理用户

```sql
　　create user zhangsan;//在管理员帐户下，创建用户zhangsan

　　alert user scott identified by tiger;//修改密码
```

### 四，授予权限

　　1、默认的普通用户scott默认未解锁，不能进行那个使用，新建的用户也没有任何权限，必须授予权限

```sql
　　grant create session to zhangsan;    /* 授予zhangsan用户创建session的权限，即登陆权限，允许用户登录数据库 */

　　grant unlimited tablespace to zhangsan;    /* 授予zhangsan用户使用表空间的权限 */

　　grant create table to zhangsan;    /* 授予创建表的权限 */

　　grant drop table to zhangsan;    /* 授予删除表的权限 */

　　grant insert table to zhangsan;    /* 插入表的权限 */

　　grant update table to zhangsan;    /* 修改表的权限 */

　　grant all to public;    /* 这条比较重要，授予所有权限(all)给所有用户(public) */
```

　　2、oralce对权限管理比较严谨，普通用户之间也是默认不能互相访问的，需要互相授权

```sql
　　grant select on tablename to zhangsan;    /* 授予zhangsan用户查看指定表的权限 */

　　grant drop on tablename to zhangsan;    /* 授予删除表的权限 */

　　grant insert on tablename to zhangsan;    /* 授予插入的权限 */

　　grant update on tablename to zhangsan;    /* 授予修改表的权限 */

　　grant insert(id) on tablename to zhangsan;

　　grant update(id) on tablename to zhangsan;    /* 授予对指定表特定字段的插入和修改权限，注意，只能是insert和update */

　　grant alert all table to zhangsan;    /* 授予zhangsan用户alert任意表的权限 */
```

### 五、撤销权限

　　基本语法同 grant, 关键字为 revoke

    GRANT 权限|角色 TO 用户名;

    REVOKE 权限|角色 FROM 用户名;

```sql
/* 实例：如何授予和撤销martin用户的CONNECT和RESOURCE两个角色： */

 GRANT connect,resource to martin;       -- 授予CONNECT和RESOURCE两个角色

 REVOKE connect, resource FROM martin;   -- 撤销CONNECT和RESOURCE两个角色

 GRANT SELECT ON SCOTT.emp TO martin;        -- 允许用户查看emp表中的记录

 GRANT UPDATE ON SCOTT.emp TO martin;        -- 允许用户更新emp表中的记录
```

### 六、查看权限

```sql

　　select * from user_sys_privs;    /* 查看当前用户所有权限 */

　　select * from user_tab_privs;    /* 查看所用用户对表的权限 */

```

### 七、操作表的用户的表

```sql
　　select * from zhangsan.tablename
```

### 八、权限传递

　　即用户A将权限授予B，B可以将操作的权限再授予C，命令如下：

```sql

　　grant alert table on tablename to zhangsan with admin option;    /* 关键字 with admin option */

　　grant alert table on tablename to zhangsan with grant option;    /* 关键字 with grant option效果和admin类似 */

```

### 九、角色

　　角色即权限的集合，可以把一个角色授予给用户

```sql
　　create role myrole;    /* 创建角色 */

　　grant create session to myrole;  /* 将创建session的权限授予myrole */

　　grant myrole to zhangsan;  /* 授予zhangsan用户myrole的角色 */

　　drop role myrole;  /* 删除角色 */
```