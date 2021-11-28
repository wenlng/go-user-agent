# go-user-agent - Parses the "User-Agent" info

[![Go Report Card](https://goreportcard.com/badge/github.com/wenlng/go-user-agent)](https://goreportcard.com/report/github.com/wenlng/go-user-agent)
[![Version](https://img.shields.io/github/tag/wenlng/go-user-agent.svg)](https://github.com/wenlng/go-user-agent/releases)
[![GoDoc](https://godoc.org/github.com/wenlng/go-user-agent?status.svg)](https://godoc.org/github.com/wenlng/go-user-agent)
[![License](https://img.shields.io/github/license/wenlng/go-user-agent.svg)](https://github.com/wenlng/go-user-agent/blob/master/LICENSE)

This library is Gets the "OSName" or "BrowserName" for USER_AGENT

- Github：[https://github.com/wenlng/go-user-agent](https://github.com/wenlng/go-user-agent)
- Author Website: [http://witkeycode.com](http://witkeycode.com)

## Installation of proxy go module in China
- GoProxy https://github.com/goproxy/goproxy.cn
- AliProxy： https://mirrors.aliyun.com/goproxy/
- OfficialProxy： https://goproxy.io/
- ChinaProxy：https://goproxy.cn
- Other：https://gocenter.io

#### Set Proxy of go module 
- Window
```shell script
$ set GO111MODULE=on
$ set GOPROXY=https://goproxy.io,direct

### The Golang 1.13+ can be executed directly
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.io,direct
```
- Linux or Mac
```shell script
$ export GO111MODULE=on
$ export GOPROXY=https://goproxy.io,direct

### or
$ echo "export GO111MODULE=on" >> ~/.profile
$ echo "export GOPROXY=https://goproxy.cn,direct" >> ~/.profile
$ source ~/.profile
```

#### Installation
```
$ go get -u github.com/wenlng/go-user-agent
```

#### Get OS Name
```go
import useragent "github.com/wenlng/go-user-agent"

func main(){
    userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60"
    name := useragent.GetOsName(userAgent)
    fmt.Println(name)   // "Windows 10"
}

```

#### Get Browser Name
```go
import useragent "github.com/wenlng/go-user-agent"

func main(){
    userAgent := "Mozilla/5.0 {Macintosh; Intel Mac OS X 10.6.8; U; en) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60"
    name := useragent.GetBrowserName(userAgent)
    fmt.Println(name)   // "Chrome/39.0.2171.95"
}
```

<br/>

<div align="center">
    <img src="http://47.104.180.148/reward-support.png?v=1" alt="Reward Support">
</div>

## LICENSE
    MIT
