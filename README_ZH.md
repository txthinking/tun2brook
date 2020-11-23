# ipio

[English](README.md)

[![捐赠](https://img.shields.io/badge/%E6%94%AF%E6%8C%81-%E6%8D%90%E8%B5%A0-ff69b4.svg)](https://www.txthinking.com/opensource-support.html)
[![交流群](https://img.shields.io/badge/%E7%94%B3%E8%AF%B7%E5%8A%A0%E5%85%A5-%E4%BA%A4%E6%B5%81%E7%BE%A4-ff69b4.svg)](https://docs.google.com/forms/d/e/1FAIpQLSdzMwPtDue3QoezXSKfhW88BXp57wkbDXnLaqokJqLeSWP9vQ/viewform)

ipio 实现了TCP/IP栈, 可以将网络层/IPv4/IPv6数据包转化为传输层/TCP/UDP数据包, 比如tun to socks5, tun to brook server, tun to brook wsserver.

### Pure Go

没有C, 没有 CGO, 使用纯Go编写, 可以被用在Linux, macOS, Windows, Android, iOS, 等...

> Windows, 需要先安装 [tap-windows](http://swupdate.openvpn.net/community/releases/tap-windows-9.21.2.exe), Copyright OpenVPN Technologies, Inc.

### 通过 [nami](https://github.com/txthinking/nami) 安装

```
$ nami install github.com/txthinking/ipio
```

或直接下载二进制命令文件 [releases](https://github.com/txthinking/ipio/releases)

### 使用

```
NAME:
   ipio - ipio can convert Network/IPv4/IPv6 layer packets to Transport/TCP/UDP layer packets, such as tun to socks5, tun to brook server, tun to brook wsserver, ...

USAGE:
   ipio [global options] command [command options] [arguments...]

VERSION:
   20201111

AUTHOR:
   Cloud <cloud@txthinking.com>

COMMANDS:
   tun2brookserver    Tun to brook server
   tun2brookwsserver  Tun to brook wsserver
   tun2socks          Tun to socks5 server
   help, h            Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)

COPYRIGHT:
   https://github.com/txthinking/ipio
```

#### tun2brookserver

> 假设你的 brook server 是 `1.2.3.4:9999`, password 是 `hello`

```
$ ipio tun2brookserver -s 1.2.3.4:9999 -p hello
```

> ROOT 权限需要, 更多 `$ ipio tun2brookserver -h`

#### tun2brookwsserver

> 假设你的 brook wsserver 是 `ws://1.2.3.4:9999`, password 是 `hello`

```
$ ipio tun2brookwsserver -s ws://1.2.3.4:9999 -p hello
```

> ROOT 权限需要, 更多 `$ ipio tun2brookwsserver -h`

#### tun2socks

> 假设你的 socks5 server is `1.2.3.4:1080`

```
$ ipio tun2socks -s 1.2.3.4:1080
```

> ROOT 权限需要, 更多 `$ ipio tun2socks -h`

## 开发者

即将到来

### Interface

```
# Interface can be tun device, fd
type Interface interface {
    io.ReadWriteCloser
}
```

```
# Handler you need to implement
type Handler interface {
    TCP(conn net.Conn)
    UDP(conn net.Conn)
}
```

### Tun2socks Example

```
import "github.com/txthinking/ipio"
import "github.com/txthinking/socks5"

type Socks5 struct{
}

func (t *Socks5) TCP(conn net.Conn) {
    c, _ := socks5.NewClient(server, username, password, 0, 60)
    rc, _ := c.Dial("tcp", conn.RemoteAddr().String()) 
    go io.Copy(conn, rc)
    io.Copy(rc, conn)
}

func (t *Socks5) UDP(conn net.Conn) {
    c, _ := socks5.NewClient(server, username, password, 0, 60)
    rc, _ := c.Dial("udp", conn.RemoteAddr().String()) 
    go io.Copy(conn, rc)
    io.Copy(rc, conn)
}

// func NewIPIO(t Interface, h Handler) *IPIO
ii := ipio.NewIPIO(tun, &Socks5{})

// func (t *IPIO) Run() error
ii.Run()

// func (t *IPIO) Close() error
ii.Close()
```
