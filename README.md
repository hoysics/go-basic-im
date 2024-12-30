

# Basic IM

基于[Plato](https://github.com/hardcore-os/plato)，结合个人理解改写，用于学习和试验IM相关的各种技术

## 目录

- [开发前的配置要求](#开发前的配置要求)
- [安装步骤](#安装步骤)
- [架构设计](#架构设计)
- [文件目录说明](#文件目录说明)
- [部署](#部署)
- [使用到的框架](#使用到的框架)

### 开发前的配置要求

1. golang > 1.22

### **安装步骤**

1. Clone the repo

```sh
git clone https://github.com/hoysics/basic-im.git
```

### 架构设计

整体遵守Plato的架构设计，主要变动为：

1. 服务内引入Kratos作为微服务框架
2. 面向客户端的接口协议改为gRPC，TCP版协议暂未实现

### 文件目录说明

```
.
├── README.md
├── api
│         ├── common
│         ├── gateway
│         ├── ipconf
│         └── state
├── app
│         ├── client
│         ├── cmd
│         ├── gateway
│         ├── ipconf
│         └── state
├── doc
│         └── test
├── go.mod
├── go.sum
├── openapi.yaml
├── pkg
│         ├── cache
│         ├── net
│         ├── registry
│         ├── router
│         ├── sdk
│         └── timingwheel
└── third_party
    ├── README.md
    ├── errors
    ├── google
    ├── openapi
    └── validate

```

### 部署

暂无

### 使用到的框架

- [kratos](https://go-kratos.dev/)

### 作者

Github: hoysics
