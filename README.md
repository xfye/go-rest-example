# rest-example

#### 项目介绍

一个Go语言写的restful api service例子。后端使用MySQL数据库作为存储。


#### 安装依赖

1. 安装golang dep https://github.com/golang/dep/releases
2. 使用dep安装依赖包
```
export GOPATH=`pwd`
cd src/rest-example; dep init
```

#### 编译
在Makefile所在目录执行:
```
make
```

#### 运行

1. 创建db和table
```
mysql -h127.0.0.1 -uroot -proot < sql/db.sql
```
2. 设置DB_URL环境变量，例如：
```
export DB_URL='root:root@tcp(127.0.0.1:3306)/rest-example?charset=utf8'
```
3. 运行
```
build/bin/main
```
4. 测试
```
curl -H "content-type:application/json" http://127.0.0.1:8080/users/1
```
