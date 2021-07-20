# go-user-agent
Gets the "OSName" or "BrowserName" for USER_AGENT

[![License](https://img.shields.io/github/license/wenlng/go-user-agent.svg)](https://github.com/wenlng/go-user-agent/blob/master/LICENSE)
[![Version](https://img.shields.io/github/tag/wenlng/go-user-agent.svg)](https://github.com/wenlng/go-user-agent/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/wenlng/go-user-agent)](https://goreportcard.com/report/github.com/wenlng/go-user-agent)
[![GoDoc](https://godoc.org/github.com/wenlng/go-user-agent?status.svg)](https://godoc.org/github.com/wenlng/go-user-agent)

- 安装项目依赖
>加载依赖管理包 (解决国内下载依赖太慢问题)
>使用国内七牛云的 go module 镜像。
>
>>参考 https://github.com/goproxy/goproxy.cn。
>
>>阿里： https://mirrors.aliyun.com/goproxy/
>
>>官方： https://goproxy.io/
>
>>中国：https://goproxy.cn
>
>>其他：https://gocenter.io

##### golang 1.13+ 可以直接执行：
```shell script
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

- 安装模块
```
go get github.com/wenlng/go-user-agent
```

- 获取系统名称
```
import "github.com/wenlng/go-user-agent"

func main(){
    userAgent := "Mozilla/5.0 {Macintosh; Intel Mac OS X 10.6.8; U; en) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60"
    name := useragent.GetOsName(userAgent)
    fmt.Println(name)   // "Windows 10"
}

```

- 获取浏览器名称
```
import "github.com/wenlng/go-user-agent"

func main(){
    userAgent := "Mozilla/5.0 {Macintosh; Intel Mac OS X 10.6.8; U; en) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60"
    name := useragent.GetBrowserName(userAgent)
    fmt.Println(name)   // "Chrome/39.0.2171.95"
}

```

IT技术网站: [witkeycode.com](witkeycode.com)

<div align="center">
    <img src="http://47.104.180.148/reward-support.png" alt="Reward Support">
</div>

## LICENSE
    MIT
