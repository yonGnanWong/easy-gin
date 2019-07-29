# gin类laravel封装

## 使用须知
### 文档类
> gin文档地址 https://github.com/gin-gonic/gin

> xorm文档地址 https://www.kancloud.cn/xormplus/xorm/167077

### 配置类
> 本项目包管理使用go modules .请将项目放在除gopath以外的目录地址并且设置export GO111MODULE=on
意思为开启go modules模式.默认是auto,如果想取消这个设置可输入export GO111MODULE=auto

> 第一次使用该项目时,如需加载依赖包.请输入命令.go mod tidy.由于某些包在国外所以被墙只能科学上网或者找github相应包.或者配置代理,操作入下:设置
export GOPROXY=https://goproxy.io 然后在terminal中输入 echo $GOPROXY 查看是否设置成功

## 目录结构
```
.
├── App                     //项目初始化目录
│   ├── app.go              //初始化所有配置并且start httpserver
│   ├── config.go           //初始化配置文件
│   ├── database.go         //初始化db
│   ├── log.go              //初始化log配置
│   ├── redis.go            //初始化redis
│   └── server.go           //server start
├── Config                  //配置文件
│   ├── dev.yaml            //dev配置文件
│   ├── product.yaml        
│   └── qa.yaml
├── Console                 //console目录.暂代开发
├── Controllers             //控制器目录
│   └── Test.go
├── Helpers                 //助手包
│   ├── JsonTrait.go        
│   ├── errorcode.go        //错误状态码
│   ├── readme.md           
│   └── request.go          //http请求
├── Logs                    //log日志
│   └── 2019-07-28.log
├── Middleware              //中间件
│   ├── header.go
│   └── readme.md
├── Models                  //模型包
├── Routers                 //路由包
│   └── router.go
├── Service                 //service服务包,编写项目服务
│   └── server.go           //服务启动.包括热重载
├── go.mod                  //go的包依赖文件
├── go.sum                  //相当于composer.lock 提交时可以igonre
├── main.go                 //项目入口
├── readme.md               
└── tree.text
```
## 操作类
    热加载
    // kill -USR2 pid 重启服务
    // kill -INT pid 关闭服务
