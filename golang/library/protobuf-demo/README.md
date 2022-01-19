
## 安装步骤

### 一、安装 protobuf (protoc)
#### 下载地址：https://github.com/google/protobuf 下载对平台的二进制文件就行，并配置环境变量

### 二、安装 protoc-gen-go 用于生成 golang 代码的插件
> go get google.golang.org/protobuf
>
> go get github.com/golang/protobuf
>
> go get google.golang.org/protobuf/cmd/protoc-gen-go
>
> go get github.com/golang/protobuf/protoc-gen-go

### 三、生成代码
> protoc --go_out=. xxx.proto
