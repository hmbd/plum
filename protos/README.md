# gRPC

## 实践
* 1.通过一个 protocol buffers 模式，定义一个简单的带有 Hello World 方法的 RPC 服务。
* 2.用你最喜欢的语言(如果可用的话)来创建一个实现了这个接口的服务端。
* 3.用你最喜欢的(或者其他你愿意的)语言来访问你的服务端。

## 目标
* 1.go实现服务端
* 2.python实现客户端
* 3.客户端端可以正常访问服务端

## 安装依赖

* 脚本安装

> curl -fsSL https://goo.gl/getgrpc | bash -s -- --with-plugins

* 或者直接使用brew安装：

> brew tap grpc/grpc

> brew install --with-plugins grpc

* 检查
```text
➜ ls /usr/local/bin/ | grep -E "protoc|grpc_"
grpc_cli
grpc_cpp_plugin
grpc_csharp_plugin
grpc_node_plugin
grpc_objective_c_plugin
grpc_php_plugin
grpc_python_plugin
grpc_ruby_plugin
protoc
```

* 只安装protobuf

> brew install protobuf

## 生成gRPC代码
在当前目录下执行

* 生成go代码

> protoc -I helloworld helloworld/helloworld.proto --go_out=plugins=grpc:helloworld

* 生成python代码

> protoc -I helloworld --python_out=helloworld --grpc_out=helloworld --plugin=protoc-gen-grpc=`which grpc_python_plugin` helloworld/helloworld.proto

* -I: 指定import路径，可以指定多个-I参数，编译时按顺序查找，不指定时默认查找当前目录
* --go_out: golang编译支持，支持以下参数
    - plugins=plugin1+plugin2 - 指定插件，目前只支持grpc，即：plugins=grpc
    - M 参数 - 指定导入的.proto文件路径编译后对应的golang包名(不指定本参数默认就是.proto文件中import语句的路径)
    - import_prefix=xxx - 为所有import路径添加前缀，主要用于编译子目录内的多个proto文件，这个参数按理说很有用，尤其适用替代一些情况时的M参数，但是实际使用时有个蛋疼的问题导致并不能达到我们预想的效果，自己尝试看看吧
    - import_path=foo/bar - 用于指定未声明package或go_package的文件的包名，最右面的斜线前的字符会被忽略
* 末尾: 编译文件路径 .proto文件路径(支持通配符)


## 使用
* 启动服务端

> cd greeter_server && go run main.go

* 服务端测试

> export PYTHONPATH=`pwd` && python greeter_client/main.py
