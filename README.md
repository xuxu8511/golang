# golang
简单封装的网络库

windows 和 linux 环境都可以编译

windows编译使用Complie.bat

linux编译使用Makefile


protobuf编译步骤:

1.https://github.com/google/protobuf/releases链接下载protoc

2.go get github.com/golang/protobuf/protoc-gen-go命令下载protoc-gen-go

3.protoc和protoc-gen-go放到go\bin路径下

4.使用tools下的build脚本（linux使用sh，windows使用bat）生成go文件
