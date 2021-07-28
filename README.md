# ipio

[🇨🇳 中文](README_ZH.md)

[![Donate](https://img.shields.io/badge/Support-Donate-ff69b4.svg)](https://github.com/sponsors/txthinking)

tun2socks5, tun2brookserver, tun2brookwsserver, tun2brookwssserver. IPv4 and IPv6, TCP and UDP.

❤️ A project by [txthinking.com](https://www.txthinking.com)

### Pure Go

ipio implements the TCP/IP stack with part of gopacket lib, No C, No CGO, can be used in any OS, such as Linux, macOS, Windows, Android, iOS, etc...

### Install via [nami](https://github.com/txthinking/nami)

```
$ nami install github.com/txthinking/ipio
```

or download binary from [releases](https://github.com/txthinking/ipio/releases)

### Usage

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

> Assume your brook server is `1.2.3.4:9999`, password is `hello`<br/>
> ROOT privileges required, more `$ ipio tun2brookserver -h`

```
$ ipio tun2brookserver -s 1.2.3.4:9999 -p hello
```
With bypass rules
```
$ ipio tun2brookserver -s 1.2.3.4:9999 -p hello --bypassDomainList https://txthinking.github.io/bypass/china_domain.txt --bypassCIDR4List https://txthinking.github.io/bypass/china_cidr4.txt --bypassCIDR6List https://txthinking.github.io/bypass/china_cidr6.txt
```
Do not change route(include bypass CIDR/IP) and DNS when you want to change by yourself
```
$ ipio tun2brookserver -s 1.2.3.4:9999 -p hello --doNotChangeRouteAndDNS
```

### tun2brookwsserver

> Assume your brook wsserver is `ws://1.2.3.4:9999`, password is `hello`<br/>
> ROOT privileges required, more `$ ipio tun2brookwsserver -h`

```
$ ipio tun2brookwsserver -s ws://1.2.3.4:9999 -p hello
```
With bypass rules
```
$ ipio tun2brookwsserver -s ws://1.2.3.4:9999 -p hello --bypassDomainList https://txthinking.github.io/bypass/china_domain.txt --bypassCIDR4List https://txthinking.github.io/bypass/china_cidr4.txt --bypassCIDR6List https://txthinking.github.io/bypass/china_cidr6.txt
```
Do not change route(include bypass CIDR/IP) and DNS when you want to change by yourself
```
$ ipio tun2brookwsserver -s ws://1.2.3.4:9999 -p hello --doNotChangeRouteAndDNS
```

### tun2brookwssserver

> Assume your brook wssserver is `wss://domain.com:443`, password is `hello`<br/>
> ROOT privileges required, more `$ ipio tun2brookwssserver -h`

```
$ ipio tun2brookwssserver -s wss://domain.com:443 -p hello
```
With bypass rules
```
$ ipio tun2brookwssserver -s wss://domain.com:443 -p hello --bypassDomainList https://txthinking.github.io/bypass/china_domain.txt --bypassCIDR4List https://txthinking.github.io/bypass/china_cidr4.txt --bypassCIDR6List https://txthinking.github.io/bypass/china_cidr6.txt
```
Do not change route(include bypass CIDR/IP) and DNS when you want to change by yourself
```
$ ipio tun2brookwssserver -s wss://domain.com:443 -p hello --doNotChangeRouteAndDNS
```

### tun2socks5

> Assume your socks5 server is `1.2.3.4:1080`<br/>
> ROOT privileges required, more `$ ipio tun2socks5 -h`

```
$ ipio tun2socks5 -s 1.2.3.4:1080
```
With bypass rules
```
$ ipio tun2socks5 -s 1.2.3.4:1080 --bypassDomainList https://txthinking.github.io/bypass/china_domain.txt --bypassCIDR4List https://txthinking.github.io/bypass/china_cidr4.txt --bypassCIDR6List https://txthinking.github.io/bypass/china_cidr6.txt
```
Do not change route(include bypass CIDR/IP) and DNS when you want to change by yourself
```
$ ipio tun2socks5 -s 1.2.3.4:1080 --doNotChangeRouteAndDNS
```

## Developer

Coming soon

## License

Licensed under The GPLv3 License
