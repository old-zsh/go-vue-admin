目的

本项目旨在打造一个运行即用的全栈脚手架 
后端采用了一系列Golang中比较流行的组件，可以以本项目为基础快速搭建Restful Web API特色
前端采用一系列Vue+ElementUi快速构建后台项目

#### 后台

后台已经整合了许多开发API所必要的组件：

    Gin: 轻量级Web框架，自称路由速度是golang最快的
    GORM: ORM工具。本项目需要配合Mysql使用
    Gin-Session: Gin框架提供的Session操作工具
    Go-Redis: Golang Redis客户端
    godotenv: 开发环境下的环境变量工具，方便使用环境变量
    Gin-Cors: Gin框架提供的跨域中间件
    自行实现了国际化i18n的一些基本功能
    本项目是使用基于cookie实现的session来保存登录状态的，如果需要可以自行修改为token验证

本项目已经预先实现了一些常用的代码方便参考和复用:

    创建了用户模型
    实现了/api/v1/user/register用户注册接口
    实现了/api/v1/user/login用户登录接口
    实现了/api/v1/user/me用户资料接口(需要登录后获取session)
    实现了/api/v1/user/logout用户登出接口(需要登录后获取session)

本项目已经预先创建了一系列文件夹划分出下列模块:

    api文件夹就是MVC框架的controller，负责协调各部件完成任务
    model文件夹负责存储数据库模型和数据库操作相关的代码
    service负责处理比较复杂的业务，把业务代码模型化可以有效提高业务代码的质量（比如用户注册，充值，下单等）
    serializer储存通用的json模型，把model得到的数据库模型转换成api需要的json对象
    cache负责redis缓存相关的代码
    auth权限控制文件夹
    util一些通用的小工具
    conf放一些静态存放的配置文件，其中locales内放置翻译相关的配置文件

Godotenv

项目在启动的时候依赖以下环境变量，但是在也可以在项目根目录创建.env文件设置环境变量便于使用(建议开发环境使用)

MYSQL_DSN="db_user:db_password@/db_name?charset=utf8&parseTime=True&loc=Local" # Mysql连接地址
REDIS_ADDR="127.0.0.1:6379" # Redis端口和地址
REDIS_PW="" # Redis连接密码
REDIS_DB="" # Redis库从0到10
SESSION_SECRET="setOnProducation" # Seesion密钥，必须设置而且不要泄露
GIN_MODE="debug"

Go Mod

本项目使用Go Mod管理依赖。

go mod init go-crud
export GOPROXY=http://mirrors.aliyun.com/goproxy/
go run main.go // 自动安装

运行

go run main.go

项目运行后启动在3000端口（可以修改，参考gin文档)



#### 前台

To start
This is a project template for vue-cli

# install dependencies
npm install

# serve with hot reload at localhost:8081
npm run dev

# build for production with minification
npm run build
Folder structure
build - webpack config files
config - webpack config files
dist - build
src -your app
api
assets
common
components - your vue components
mock
styles
views - your pages
vuex
App.vue
main.js - main file
routes.js
static - static assets
Theme
You can change theme by

Generate theme packages by https://elementui.github.io/theme-preview/#/
Put theme packages in src/assets/theme/
Edit src/main.js
   import 'element-ui/lib/theme-default/index.css'
   to
   import './assets/theme/your-theme/index.css'
Edit src/styles/vars.scss
theme-blue theme-green

Browser support
Modern browsers and IE 10+.

License
MIT
