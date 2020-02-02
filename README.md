# gin_web

#### 介绍
负责linkbook的api服务

#### 包管理
使用gomod来管理go的包依赖

#### 安装
*  1. Redis          (v. 4.0.9+)
*  2. Mysql       (v. 5.7.27)

#### 使用说明

1. git clone git@github.com:haimait/gin_web.git
2. go run main.go



#### 程序目录结构介绍
    router　路由条目

    App

        controller 路由参数处理
        model      数据操作层
        service    业务逻辑层
                    --base.go 初始化数据层model
        middleware 中间件层

    config  配置文件

    flags   一些常量和变量

    main.go 主函数

#### 分支操作
1. Checkout master 分支　
2. 新建 feature/name_xx 分支
3. 提交代码
4. 新建 Pull Request,如果和主分支有冲突要解决所有冲突,未解决冲突的代码,不可强行合入主分支


