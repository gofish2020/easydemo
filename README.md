# 学习 gRPC

本项目的目标

1. 学习如何编写 `proto`文件
2. 一个 `proto`如何 `import`另外一个 `proto`,需要结合编译指令 `-I`参数，解读
3. `option go_package = "easydemo/proto/hellopb;hellopb";` 含义 `easydemo/proto/hellopb`表示输出目录，还需要结合`-go_out`编译参数，才能知道真实的输出目录，`hellopb`表示 `*.go`文件的包名（当然可以省略）自动用目录名作为包名
4. `protoc -I . --go_out=module=easydemo:.  --go-grpc_out=module=easydemo:.  proto/*.proto` 编译指令和 `*.proto`文件有着内在的关系

```go
-I，--proto_path: 指定进行搜索依赖包的目录，针对的是 *.proto文件，可以指定多个，例如 -I . -I .. 与 *.proto 文件中的 import 一起组成绝对的搜索路径

--go_out: 指定输出 Go 代码的目录，默认为当前目录。 和 *.proto 文件中的 go_package 一起组成绝对路径 （如果 存在--go_out=module 选项，会将输出路径中的【匹配的名字去掉】）； go_package还有一层的意义，当作为被其他 proto文件引入import的时候，编译后生成的 *.go的引用路径

--go_opt: 参数为 module 或 paths。
```
5. 实现 grpc client/server 
6. 理解 `stream` 在 `client/server` 之间的双通道流，可以实现客户端和服务端之间的数据相互推送，底层是通过复用`tcp`连接
7. 利用 `stream`实现文件发送`sendfile`到服务端

