# ipio

[🇨🇳 中文](README_ZH.md)

[🗣 News](https://t.me/txthinking_news)
[💬 Chat](https://join.txthinking.com)
[🩸 Youtube](https://www.youtube.com/txthinking) 
[❤️ Sponsor](https://github.com/sponsors/txthinking)

👉 **Proxy All Traffic** 👈 : tun2brooklink, tun2socks5, tun2brookserver, tun2brookwsserver, tun2brookwssserver. IPv4 and IPv6, TCP and UDP.

A project by [txthinking.com](https://www.txthinking.com)

### Install via [nami](https://github.com/txthinking/nami)

```
nami install ipio
```

### Usage

**ROOT/Admin privileges required**

### tun2brooklink

```
ipio tun2brooklink -l 'brook://...'
```

Custom DNS, bypass domain and IP, custom route, more: `ipio tun2brooklink -h`

### tun2brookserver

Assume your brook server is `1.2.3.4:9999`, password is `hello`<br/>

```
ipio tun2brookserver -s 1.2.3.4:9999 -p hello
```

Custom DNS, bypass domain and IP, custom route, more: `ipio tun2brookserver -h`

### tun2brookwsserver

Assume your brook wsserver is `ws://1.2.3.4:9999`, password is `hello`<br/>

```
ipio tun2brookwsserver -s ws://1.2.3.4:9999 -p hello
```

Custom DNS, bypass domain and IP, custom route, more: `ipio tun2brookwsserver -h`

### tun2brookwssserver

Assume your brook wssserver is `wss://domain.com:443`, password is `hello`<br/>

```
ipio tun2brookwssserver -s wss://domain.com:443 -p hello
```

Custom DNS, bypass domain and IP, custom route, more: `ipio tun2brookwssserver -h`

### tun2socks5

Assume your socks5 server is `1.2.3.4:1080`. **Your socks5 must support standard UDP, recommand [brook](https://github.com/txthinking/brook) socks5**

```
ipio tun2socks5 -s 1.2.3.4:1080
```

Custom DNS, bypass domain and IP, custom route, more: `ipio tun2socks5 -h`

## Developer

Coming soon

## License

Licensed under The GPLv3 License
