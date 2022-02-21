# ipio

[🇨🇳 中文](README_ZH.md)

[🗣 News](https://t.me/txthinking_news)
[💬 Chat](https://join.txthinking.com)
[🩸 Youtube](https://www.youtube.com/txthinking) 
[❤️ Sponsor](https://github.com/sponsors/txthinking)

tun2brooklink, tun2socks5, tun2brookserver, tun2brookwsserver, tun2brookwssserver. IPv4 and IPv6, TCP and UDP.

❤️ A project by [txthinking.com](https://www.txthinking.com)

### Pure Go

ipio implements the TCP/IP stack with part of gopacket lib, No C, No CGO, can be used in any OS, such as Linux, macOS, Windows, Android, iOS, etc...

### Install via [nami](https://github.com/txthinking/nami)

```
nami install ipio
```

or download binary from [releases](https://github.com/txthinking/ipio/releases)

### Usage

```
NAME:
   ipio - ipio can convert Network/IPv4/IPv6 layer packets to Transport/TCP/UDP layer packets, such as tun to socks5, tun to brook server, tun to brook wsserver, tun to brook wssserver, ...

USAGE:
   ipio [global options] command [command options] [arguments...]

AUTHOR:
   Cloud <cloud@txthinking.com>

COMMANDS:
   tun2brooklink      Tun to brook link
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

### tun2brooklink

> ROOT/Admin privileges required`

```
ipio tun2brooklink -l 'brook://...'
```

Custom DNS, bypass domain and IP, custom route: `ipio tun2brooklink -h`

### tun2brookserver

> Assume your brook server is `1.2.3.4:9999`, password is `hello`<br/>
> ROOT/Admin privileges required, more `ipio tun2brookserver -h`

```
ipio tun2brookserver -s 1.2.3.4:9999 -p hello
```

Custom DNS, bypass domain and IP, custom route: `ipio tun2brookserver -h`

### tun2brookwsserver

> Assume your brook wsserver is `ws://1.2.3.4:9999`, password is `hello`<br/>
> ROOT/Admin privileges required, more `ipio tun2brookwsserver -h`

```
ipio tun2brookwsserver -s ws://1.2.3.4:9999 -p hello
```

Custom DNS, bypass domain and IP, custom route: `ipio tun2brookwsserver -h`

### tun2brookwssserver

> Assume your brook wssserver is `wss://domain.com:443`, password is `hello`<br/>
> ROOT/Admin privileges required`

```
ipio tun2brookwssserver -s wss://domain.com:443 -p hello
```

Custom DNS, bypass domain and IP, custom route: `ipio tun2brookwssserver -h`

### tun2socks5

> Assume your socks5 server is `1.2.3.4:1080`<br/>
> ROOT/Admin privileges required`

```
ipio tun2socks5 -s 1.2.3.4:1080
```

Custom DNS, bypass domain and IP, custom route: `ipio tun2socks5 -h`

## Developer

Coming soon

## License

Licensed under The GPLv3 License
