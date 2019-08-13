# gin类laravel封装

## 使用须知
### 文档类
> gin文档地址 https://github.com/gin-gonic/gin  

> gin看云文档地址 https://www.kancloud.cn/shuangdeyu/gin_book/949414 

> xorm文档地址 https://www.kancloud.cn/xormplus/xorm/167077


### 配置类
> 本项目包管理使用go modules .请将项目放在除gopath以外的目录地址并且设置export GO111MODULE=on
意思为开启go modules模式.默认是auto,如果想取消这个设置可输入export GO111MODULE=auto

> 第一次使用该项目时,如需加载依赖包.请输入命令.go mod tidy.由于某些包在国外所以被墙只能科学上网或者找github相应包.或者配置代理,操作入下:设置
export GOPROXY=https://goproxy.io或者阿里云代理https://mirrors.aliyun.com/goproxy/` 然后在terminal中输入 echo $GOPROXY 查看是否设置成功


## 目录结构
```
.
├── Api
│   └── v1
│       └── test.go
├── App
│   ├── app.go
│   ├── config.go
│   ├── cron.go
│   ├── log.go
│   └── server.go
├── Config
│   ├── dev.yaml
│   ├── product.yaml
│   └── qa.yaml
├── Console
│   └── cron1.go
├── Dockerfile
├── Helpers
│   ├── array.go
│   ├── readme.md
│   └── request.go
├── Logs
│   ├── 2019-08-11.log
│   └── 2019-08-12.log
├── Middleware
│   ├── Header
│   │   └── header.go
│   ├── Log
│   │   └── log.go
│   └── readme.md
├── Models
│   ├── db
│   │   ├── course.go
│   │   └── user.go
│   └── test.go
├── Routers
│   └── router.go
├── Service
│   ├── Bean
│   │   └── factory.go
│   ├── Database
│   │   ├── database.go
│   │   └── redis.go
│   └── Server
│       └── server.go
├── build.sh
├── go.mod
├── go.sum
├── main.go
├── readme.md
└── tree.text
```
## 操作类
    热加载
    // kill -USR2 pid 重启服务
    // kill -INT pid 关闭服务
    
## 工具项
xrom cmd工具
> go get github.com/go-xorm/cmd/xorm
> * reverse 反转一个数据库结构，生成代码
> * shell 通用的数据库操作客户端，可对数据库结构和数据操作
> * dump Dump数据库中所有结构和数据到标准输出
> * source 从标注输入中执行SQL文件
> * driver 列出所有支持的数据库驱动
> 执行一次mysql表到struct的装换命令 
xorm reverse mysql root:123456@\(127.0.0.1:3306\)/test\?charset=utf8 $GOPATH/pkg/mod/github.com/go-xorm/cmd/xorm@v0.0.0-20190426080617-f87981e709a1/templates/goxorm /Users/yongnan_wong/Documents/obj/uuzu/gin/Models/db
    
      
