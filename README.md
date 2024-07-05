# 基于公司规范的http server的封装

## 使用

```shell
go get github.com/penitence1992/go-server-v1
```

## 规范内容

1. 全局的统一响应

```json
{
  "bizCode": "B0001",
  "code": 404,
  "data": null,
  "msg": "test error"
}
```

2. 统一的异常处理中间件和异常类 `BaseCwError` 和接口`CwError`

3. 目前golang有尝试使用go-swagger来生成swagger文件具体可以查看 [go-swagger](https://goswagger.io/generate/spec.html)

## 功能说明

### 对Sql数据库的支持

相关代码存放在`pkg/storage`下, 目前实现了`pg`数据库的连接, 相关的orm使用的是gorm v2版本

### 添加Rabbitmq的相关基础模块

### 添加了Eureka的注册服务模块

### 添加了常用的actuator模块

添加了`/actuator/info`, `/actuator/health`, `/swagger` 这几个接口

其中swagger接口需要把文件放置到和执行文件相同目录下,并且命名为`swagger.yml`
