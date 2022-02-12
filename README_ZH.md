# ipio

[🇬🇧 English](README.md)

[🗣 News](https://t.me/txthinking_news)
[💬 Chat](https://join.txthinking.com)
[🩸 Youtube](https://www.youtube.com/txthinking) 
[❤️ Sponsor](https://github.com/sponsors/txthinking)

tun2socks5, tun2brookserver, tun2brookwsserver, tun2brookwssserver. IPv4 and IPv6, TCP and UDP.

❤️ A project by [txthinking.com](https://www.txthinking.com)

### Pure Go

ipio implements the TCP/IP stack with part of gopacket lib, No C, No CGO, can be used in any OS, such as Linux, macOS, Windows, Android, iOS, etc...

### 通过 [nami](https://github.com/txthinking/nami) 安装

```
$ nami install ipio
```

或直接下载二进制命令文件 [releases](https://github.com/txthinking/ipio/releases)

### 使用

```
NAME:
   ipio - ipio can convert Network/IPv4/IPv6 layer packets to Transport/TCP/UDP layer packets, such as tun to socks5, tun to brook server, tun to brook wsserver, tun to brook wssserver, ...

USAGE:
   ipio [global options] command [command options] [arguments...]

VERSION:
   20210701

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

> 假设你的 brook server 是 `1.2.3.4:9999`, password 是 `hello`<br/>
> ROOT 权限需要, 更多 `$ ipio tun2brookserver -h`

```
$ ipio tun2brookserver -s 1.2.3.4:9999 -p hello
```
Bypass規則
```
$ ipio tun2brookserver -s 1.2.3.4:9999 -p hello --bypassDomainList https://txthinking.github.io/bypass/china_domain.txt --bypassCIDR4List https://txthinking.github.io/bypass/china_cidr4.txt --bypassCIDR6List https://txthinking.github.io/bypass/china_cidr6.txt
```
不修改路由表和DNS, 自己修改
```
$ ipio tun2brookserver -s 1.2.3.4:9999 -p hello --doNotChangeRouteAndDNS
```

### tun2brookwsserver

> 假设你的 brook wsserver 是 `ws://1.2.3.4:9999`, password 是 `hello`<br/>
> ROOT 权限需要, 更多 `$ ipio tun2brookwsserver -h`

```
$ ipio tun2brookwsserver -s ws://1.2.3.4:9999 -p hello
```
Bypass規則
```
$ ipio tun2brookwsserver -s ws://1.2.3.4:9999 -p hello --bypassDomainList https://txthinking.github.io/bypass/china_domain.txt --bypassCIDR4List https://txthinking.github.io/bypass/china_cidr4.txt --bypassCIDR6List https://txthinking.github.io/bypass/china_cidr6.txt
```
不修改路由表和DNS, 自己修改
```
$ ipio tun2brookwsserver -s ws://1.2.3.4:9999 -p hello --doNotChangeRouteAndDNS
```

### tun2brookwssserver

> 假设你的 brook wssserver 是 `wss://domain.com:443`, password 是 `hello`<br/>
> ROOT 权限需要, 更多 `$ ipio tun2brookwssserver -h`

```
$ ipio tun2brookwssserver -s wss://domain.com:443 -p hello
```
Bypass規則
```
$ ipio tun2brookwssserver -s wss://domain.com:443 -p hello --bypassDomainList https://txthinking.github.io/bypass/china_domain.txt --bypassCIDR4List https://txthinking.github.io/bypass/china_cidr4.txt --bypassCIDR6List https://txthinking.github.io/bypass/china_cidr6.txt
```
不修改路由表和DNS, 自己修改
```
$ ipio tun2brookwssserver -s wss://domain.com:443 -p hello --doNotChangeRouteAndDNS
```

### tun2socks5

> 假设你的 socks5 server is `1.2.3.4:1080`<br/>
> ROOT 权限需要, 更多 `$ ipio tun2socks5 -h`

```
$ ipio tun2socks5 -s 1.2.3.4:1080
```
Bypass規則
```
$ ipio tun2socks5 -s 1.2.3.4:1080 --bypassDomainList https://txthinking.github.io/bypass/china_domain.txt --bypassCIDR4List https://txthinking.github.io/bypass/china_cidr4.txt --bypassCIDR6List https://txthinking.github.io/bypass/china_cidr6.txt
```
不修改路由表和DNS, 自己修改
```
$ ipio tun2socks5 -s 1.2.3.4:1080 --doNotChangeRouteAndDNS
```

## 开发者

即将到来

## 开源协议

基于 GPLv3 协议开源
