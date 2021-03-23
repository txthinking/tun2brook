# ipio

[中文](README_ZH.md)

[![Donate](https://img.shields.io/badge/Support-Donate-ff69b4.svg)](https://www.txthinking.com/opensource-support.html)

ipio has implemented the TCP/IP stack, can convert Network/IPv4/IPv6 layer packets to Transport/TCP/UDP layer packets, such as tun to socks5, tun to brook server, tun to brook wsserver, tun to brook wssserver. **Let all traffic go through socks5, brook server, brook wsserver, brook wssserver**

### Pure Go

No C, No CGO, written in Pure Go, can be used in any OS, such as Linux, macOS, Windows, Android, iOS, etc...

> Windows, need to install [tap-windows](http://swupdate.openvpn.net/community/releases/tap-windows-9.21.2.exe) first, Copyright OpenVPN Technologies, Inc.

### Install via [nami](https://github.com/txthinking/nami)

```
$ nami install github.com/txthinking/ipio
```

or download binary from [releases](https://github.com/txthinking/ipio/releases)

### Usage

```
NAME:
   ipio - ipio can convert Network/IPv4/IPv6 layer packets to Transport/TCP/UDP layer packets, such as tun to socks5, tun to brook server, tun to brook wsserver, ...

USAGE:
   ipio [global options] command [command options] [arguments...]

VERSION:
   20210401

AUTHOR:
   Cloud <cloud@txthinking.com>

COMMANDS:
   tun2brookserver    Tun to brook server
   tun2brookwsserver  Tun to brook wsserver
   tun2brookwssserver Tun to brook wssserver
   tun2socks5         Tun to socks5 server
   help, h            Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)

COPYRIGHT:
   https://github.com/txthinking/ipio
```

### tun2brookserver

> Assume your brook server is `1.2.3.4:9999`, password is `hello`<br/>
> ROOT privileges required, more `$ ipio tun2brookserver -h`

```
$ ipio tun2brookserver -s 1.2.3.4:9999 -p hello
```
With bypass rules
```
$ ipio tun2brookserver -s 1.2.3.4:9999 -p hello --bypassDomainList https://txthinking.github.io/bypass/chinadomain.txt --bypassCIDR4List https://txthinking.github.io/bypass/chinacidr4.txt --bypassCIDR6List https://txthinking.github.io/bypass/chinacidr6.txt
```
Do not change route(include bypass CIDR/IP) and DNS when you want to change by yourself
```
$ ipio tun2brookserver -s 1.2.3.4:9999 -p hello --doNotChangeAnything
```

### tun2brookwsserver

> Assume your brook wsserver is `ws://1.2.3.4:9999`, password is `hello`<br/>
> ROOT privileges required, more `$ ipio tun2brookwsserver -h`

```
$ ipio tun2brookwsserver -s ws://1.2.3.4:9999 -p hello
```
With bypass rules
```
$ ipio tun2brookwsserver -s ws://1.2.3.4:9999 -p hello --bypassDomainList https://txthinking.github.io/bypass/chinadomain.txt --bypassCIDR4List https://txthinking.github.io/bypass/chinacidr4.txt --bypassCIDR6List https://txthinking.github.io/bypass/chinacidr6.txt
```
Do not change route(include bypass CIDR/IP) and DNS when you want to change by yourself
```
$ ipio tun2brookwsserver -s ws://1.2.3.4:9999 -p hello --doNotChangeAnything
```

### tun2brookwssserver

> Assume your brook wssserver is `wss://domain.com:443`, password is `hello`<br/>
> ROOT privileges required, more `$ ipio tun2brookwssserver -h`

```
$ ipio tun2brookwssserver -s wss://domain.com:443 -p hello
```
With bypass rules
```
$ ipio tun2brookwssserver -s wss://domain.com:443 -p hello --bypassDomainList https://txthinking.github.io/bypass/chinadomain.txt --bypassCIDR4List https://txthinking.github.io/bypass/chinacidr4.txt --bypassCIDR6List https://txthinking.github.io/bypass/chinacidr6.txt
```
Do not change route(include bypass CIDR/IP) and DNS when you want to change by yourself
```
$ ipio tun2brookwssserver -s wss://domain.com:443 -p hello --doNotChangeAnything
```

### tun2socks5

> Assume your socks5 server is `1.2.3.4:1080`<br/>
> ROOT privileges required, more `$ ipio tun2socks5 -h`

```
$ ipio tun2socks5 -s 1.2.3.4:1080
```
With bypass rules
```
$ ipio tun2socks5 -s 1.2.3.4:1080 --bypassDomainList https://txthinking.github.io/bypass/chinadomain.txt --bypassCIDR4List https://txthinking.github.io/bypass/chinacidr4.txt --bypassCIDR6List https://txthinking.github.io/bypass/chinacidr6.txt
```
Do not change route(include bypass CIDR/IP) and DNS when you want to change by yourself
```
$ ipio tun2socks5 -s 1.2.3.4:1080 --doNotChangeAnything
```

## Developer

Coming soon

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
