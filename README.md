# gin框架

gin框架入门到实战

原文：https://www.liwenzhou.com/posts/Go/Gin_framework/#autoid-0-2-2

> 只要API程序遵循了REST风格，那就可以称其为RESTful API。目前在前后端分离的架构中，前后端基本都是通过RESTful API来进行交互。

- GET用来获取资源	
- POST用来新建资源		
- PUT用来更新资源
- DELETE用来删除资源。

## 项目结构

```go
├─.idea
├─gin-demo-01-hello
│  └─gin_demo
├─gin-demo-02-template
├─gin-demo-03-template-grammar
├─gin-demo-04-template-nesting
├─gin-demo-05-template-tmpl
│  └─templates
├─gin-demo-06-template-supplement
├─gin-demo-07-template-Rendering
│  ├─statics
│  │  ├─css
│  │  ├─fonts
│  │  ├─images
│  │  ├─js
│  │  └─picture
│  └─templates
│      ├─posts
│      └─users
├─gin-demo-08-json
├─gin-demo-09-parameter
├─gin-demo-10-form
├─gin-demo-11-url
├─gin-demo-12-parameter-binding
├─gin-demo-13-file-upload
├─gin-demo-14-redirect
├─gin-demo-15-router
├─gin-demo-16-middleware
├─gin-demo-17-gorm-mysql
├─gin-demo-18-gorm-model
├─gin-demo-19-gorm-create
├─gin-demo-20-gorm-select
├─gin-demo-21-gorm-update
├─gin-demo-22-gorm-delete
├─gin-demo-23-bubble-fromtend
│  ├─bubble
│  │  ├─static
│  │  │  ├─css
│  │  │  ├─fonts
│  │  │  └─js
│  │  └─templates
│  └─dist
│      └─dist
│          └─static
│              ├─css
│              ├─fonts
│              └─js
├─gin-demo-24-bubble-v1
│  └─bubble
│      ├─static
│      │  ├─css
│      │  ├─fonts
│      │  └─js
│      └─templates
└─gin-demo-25-bubble-v2
    └─bubble
        ├─controller
        ├─dao
        ├─models
        ├─routers
        ├─static
        │  ├─css
        │  ├─fonts
        │  └─js
        └─templates
```

| 序号 |          project_name           |           简介           |
| :--: | :-----------------------------: | :----------------------: |
|  1   |        gin-demo-01-hello        |       gin框架初识        |
|  2   |      gin-demo-02-template       |       template初识       |
|  3   |  gin-demo-03-template-grammar   |      go模板语法详解      |
|  4   |  gin-demo-04-template-nesting   |         模板嵌套         |
|  5   |    gin-demo-05-template-tmpl    |         模板继承         |
|  6   | gin-demo-06-template-supplement |         模板补充         |
|  7   | gin-demo-07-template-Rendering  |     gin框架模板渲染      |
|  8   |        gin-demo-08-json         |     gin框架返回json      |
|  9   |      gin-demo-09-parameter      |  gin获取querystring参数  |
|  10  |        gin-demo-10-form         |     gin获取form参数      |
|  11  |         gin-demo-11-url         |    gin获取URI路径参数    |
|  12  |  gin-demo-12-parameter-binding  |       gin参数绑定        |
|  13  |     gin-demo-13-file-upload     |       gin文件上传        |
|  14  |      gin-demo-14-redirect       |      gin请求重定向       |
|  15  |       gin-demo-15-router        |     gin路由和路由组      |
|  16  |     gin-demo-16-middleware      |          中间件          |
|  17  |     gin-demo-17-gorm-mysql      |  GORM连接MySQL基本示例   |
|  18  |     gin-demo-18-gorm-model      |   GORM模型定义那些事儿   |
|  19  |     gin-demo-19-gorm-create     |   GORM创建记录及字段默   |
|  20  |     gin-demo-20-gorm-select     | P22lesson22_GORM查询操作 |
|  21  |     gin-demo-21-gorm-update     |       GORM更新操作       |
|  22  |     gin-demo-22-gorm-delete     |       GORM删除操作       |
|  23  |   gin-demo-23-bubble-fromtend   |      小清单项目启动      |
|  24  |      gin-demo-24-bubble-v1      |      小清单功能实现      |
|  25  |      gin-demo-25-bubble-v2      |    企业级项目结构拆分    |
|      |                                 |                          |
