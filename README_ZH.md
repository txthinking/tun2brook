# ipio

[🇬🇧 English](README.md)

[🗣 News](https://t.me/txthinking_news)
[💬 Chat](https://join.txthinking.com)
[🩸 Youtube](https://www.youtube.com/txthinking) 
[❤️ Sponsor](https://github.com/sponsors/txthinking)

全局代理: tun2brooklink, tun2socks5, tun2brookserver, tun2brookwsserver, tun2brookwssserver. IPv4 and IPv6, TCP and UDP.

❤️ A project by [txthinking.com](https://www.txthinking.com)

### 通过 [nami](https://github.com/txthinking/nami) 安装

```
nami install ipio
```

或直接下载二进制命令文件 [releases](https://github.com/txthinking/ipio/releases)

### 使用

**需要ROOT/Admin权限**

### tun2brooklink

```
ipio tun2brooklink -l 'brook://...'
```

自定义DNS，分流，路由，等等：`ipio tun2brooklink -h`

### tun2brookserver

假设你的 brook server 是 `1.2.3.4:9999`, password 是 `hello`

```
ipio tun2brookserver -s 1.2.3.4:9999 -p hello
```

自定义DNS，分流，路由，等等：`ipio tun2brookserver -h`

### tun2brookwsserver

假设你的 brook wsserver 是 `ws://1.2.3.4:9999`, password 是 `hello`

```
ipio tun2brookwsserver -s ws://1.2.3.4:9999 -p hello
```

自定义DNS，分流，路由，等等：`ipio tun2brookwsserver -h`

### tun2brookwssserver

假设你的 brook wssserver 是 `wss://domain.com:443`, password 是 `hello`

```
ipio tun2brookwssserver -s wss://domain.com:443 -p hello
```

自定义DNS，分流，路由，等等：`ipio tun2brookwssserver -h`

### tun2socks5

假设你的 socks5 server is `1.2.3.4:1080`. **你的socks5必须支持标准 UDP 协议, 推荐 [brook](https://github.com/txthinking/brook) socks5**

```
ipio tun2socks5 -s 1.2.3.4:1080
```

自定义DNS，分流，路由，等等：`ipio tun2socks5 -h`

## 开发者

即将到来

## 协议

基于 GPLv3 协议
