# 综述

本教程为精弘网络 2022 年暑期授课系列教程之讲义.

本文为Golang后端实战. 

作者 [FinleyGe](https://github.com/FinleyGe)

[GitHub上的代码仓库](https://github.com/zjutjh/Summer-Lessons-2022/tree/day6)

由于前面课程已经讲过了 Golang 的基础内容, 此处略去.

# 技术栈 & 名词解释
**[Gin](https://github.com/gin-gonic/gin)框架**: 一款轻量级, 高性能, 简单的Golang后端开发框架.

**[Gorm](https://gorm.io/zh_CN/)框架**:
> **ORM**: Object–Relational Mapping, 对象关系映射

对象就是Go中的`struct`, 关系就是MySQL中的各种表

**Viper**: 配置读取绑定

**MySQL**: 一款流行的关系型数据库
> **SQL**: Structure Query Language 结构化查询语言

MySQL 使用 SQL 作为查询语言.

**Git** 一款版本控制软件. 通过Git 可以记录项目代码的变化情况.
> GitHub 是 一个Git的远端管理平台, 可以在上面开源你的代码.

**Apifox** 一个管理API的工具, 可以方便的编写API文档, 进行测试等等.

> API: Application Programming Interface 应用编程接口, 一般也简称接口. 本文后端主要面向的就是RESTful API的编程

**RESTful API**
Representational State Transfer API, 表征状态转移接口

统一资源接口
GET, POST, PUT, DELETE, PATCH

| Method | 用途               |
| ------ | ------------------ |
| GET    | 查询               |
| POST   | 增加               |
| PUT    | 更新(传入全部信息) |
| DELETE | 删除               |
| PATCH  | 更新(传入部分信息) |

# 1. 项目构建
1. 项目的基本构建
```bash
mkdir class-schedule
cd class-schedule
vim main.go
go mod init class-schedule
go mod tidy
```

> go mod 是 golang 的包管理器, 相当用 python 的pip 或者是 node.js 的 npm, yarn...
>
> 使用 go mod 创建项目 可以使用 go mod tidy 自动下载依赖包, 其下载位置在$GOPATH/pkg/mod

2. git 
```bash
git init
```
即可建立一个初始化的仓库,当前这个仓库还没有任何东西.

```bash
vim README.md
```
编写一些东西

```
git add .
git commit -a -m 'init: first commit'
```

3. 项目架构
├── config 配置文件
├── controller 
├── db 数据库相关
├── go.mod
├── go.sum
├── LICENSE
├── main.go 入口文件
├── router 路由
└── utility 工具

# 2. 配置管理
使用 [viper](https://github.com/spf13/viper) 作为配置管理工具
1. config.go

```go
package go
import ...

type config struct {
  // ...
}
var Config config
func InitConfig() {

}
```
> 访问权限控制: Go 中的访问控制为变量名, 函数名等首字母是否大写, 大写则为 Public, 小写则为 Private
> 
> yaml 文件: 相当简单的数据存储文件
# 3. 链接数据库
1. 建立数据库
```bash
sudo mycli -uroot
create database class_schedule;
create user 'class_schedule_admin'@'localhost' identified by 'password';
grant all on class_schedule.* to 'class_schedule_admin'@'localhost';
```
> [mycli](https://github.com/dbcli/mycli): A command line client for MySQL that can do auto-completion and syntax highlighting.
2. 链接

```
dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
	config.Config.DB.UserName,
	config.Config.DB.Password,
	config.Config.DB.Address,
	config.Config.DB.DBName,
	)
```
3. Auto Migrate
> DSN : Data Source Name 
# 4. 编写Utility
1. 密码加密解密
[bcrypt](golang.org/x/crypto/bcrypt)
>bcrypt : 一种非对称加密算法, 加盐hash

2. Jwt生成和解密
[JWT](https://jwt.io/)

通过Jwt鉴权, 一般的鉴权是通过 Session. Session会在服务器端占用存储空间, 返回一个Session ID

而 Jwt 本质上是通过算力换空间.

3. 响应格式化

# 5. 构建模型
```go

type User struct {
	UserID      uint64  `gorm:"primaryKey" json:"user_id"`
	Name        string  `gorm:"not null" json:"name" validate:"required"`
	Email       string  `gorm:"not null" json:"email" validate:"required"`
	Phone       string  `json:"phone"`
	Password    string  `gorm:"not null" json:"password" validate:"required"`
	School      string  `json:"school"`
	StuID       string  `json:"stu_id"`
	ClassAmount uint64  `gorm:"not null; default:12" json:"class_amount"`
	WeekAmount  uint64  `gorm:"not null; default:16" json:"week_amount"`
	Classes     []Class `json:"classes"`
}

type ClassInfo struct {
	ClassID     uint64 `json:"class_id"`
	ClassInfoID uint64 `json:"class_info_id" gorm:"primaryKey"`
	Week        string `json:"week" validate:"required" gorm:"not null"`
	Time        string `json:"time" validate:"required" gorm:"not null"`
	WeekDay     string `json:"week_day" validate:"required" gorm:"not null"`
	Address     string `json:"address" validate:"required" gorm:"not null"`
}

type Class struct {
	UserID     uint64      `json:"-"`
	ClassID    uint64      `gorm:"primaryKey" json:"-"`
	Name       string      `gorm:"not null" json:"name" validate:"required"`
	ClassInfos []ClassInfo `json:"class_infos" validate:"required"`
	Teacher    string      `json:"teacher"`
}
```

三个模型对应三张表(table)

**关系**
1. Belong To
2. Has One
3. Has Many
4. Many to Many

**主键**: 唯一,非空.可以标识这一列. 

**外键**: 用于将两个表链接在一起的键. 

# 6. CURD
Create-Update-Retrieve-Delete

# 7. 路由
1. 路由初始化
2. CORS 中间件

> 同源策略: 只能访问同源的资源. 
> 同源: Scheme, Domain, Port 相同
> 
> Cors: Cross-Origin-Resource-Share 跨域资源共享
# 8. Controller & 中间件
Controller

1. 用户注册 - 增
2. 用户登陆 - 查
3. 用户信息获取 - 查
4. 添加课程 - 增
5. 查询课程 - 查
6. 更改课程 - 改
7. 删除课程 - 删

Middleware:
- 鉴权中间件
