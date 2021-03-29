# DSF - 极简文件服务

[![Go Reference](https://pkg.go.dev/badge/github.com/nanmu42/dsf.svg)](https://pkg.go.dev/github.com/nanmu42/dsf)
[![Go Report Card](https://goreportcard.com/badge/github.com/nanmu42/dsf)](https://goreportcard.com/report/github.com/nanmu42/dsf)

[English](https://github.com/nanmu42/dsf) | **中文**

一个非常简单的HTTP文件服务，可以把本地的文件暴露到局域网上。

虽然`python3 -m http.server`也能用，但是有时我没法接受它单线程一次只能服务一个文件的特性。

## 用法

```
$ dsf -h
dsf 的用法:
  -port int
    	监听端口 (默认 8080)
  -root string
    	文件根目录 (默认 "./")
```

分享当前工作目录：

```bash
$ dsf
listening on: http://127.0.0.1:8080
listening on: http://::1:8080
```

在特定端口分享特定目录：

```bash
$ dsf -port 3000 -root ~/images
listening on: http://127.0.0.1:3000
listening on: http://::1:3000
```

## 安装

适用于Linux, Mac 和 Windows预编译二进制可以在这里找到：https://github.com/nanmu42/dsf/releases/

Gophers 可以这样安装：

```bash
go install github.com/nanmu42/dsf
```

使用愉快！

## License

MIT