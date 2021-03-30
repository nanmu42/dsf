# DSF - Dead Simple Fileserver

[![Go Reference](https://pkg.go.dev/badge/github.com/nanmu42/dsf.svg)](https://pkg.go.dev/github.com/nanmu42/dsf)
[![Go Report Card](https://goreportcard.com/badge/github.com/nanmu42/dsf)](https://goreportcard.com/report/github.com/nanmu42/dsf)

**English** | [中文](https://github.com/nanmu42/dsf/blob/main/README.zh-cn.md)

A dead simple HTTP fileserver to share your files across LAN.

## Motivation

Yes, `python3 -m http.server` works fine, but...

* sometimes I got bitten by its single-threaded nature(serving one downloading at a time);
* `dsf` prints copy-paste-friendly host IP and port, so I don't need to figure out by myself.

## Usage

```
$ dsf -h
Usage of dsf:
  -port int
    	listening port (default 8080)
  -root string
    	root for files (default "./")
```

Serve current working directory:

```bash
$ dsf
listening on: http://127.0.0.1:8080
listening on: http://::1:8080
listening on: http://192.168.0.6:8080
listening on: http://fe80::1234:1e11:abec:5678:8080
```

Serve a specific directory and listen on a desired port:

```bash
$ dsf -port 3000 -root ~/images
listening on: http://127.0.0.1:3000
listening on: http://::1:3000
listening on: http://192.168.0.6:3000
listening on: http://fe80::1234:1e11:abec:5678:3000
```

## Installation

Precompiled binaries are available for Linux, Mac and Windows: https://github.com/nanmu42/dsf/releases/

For gophers:

```bash
go install github.com/nanmu42/dsf@latest
```

## License

MIT