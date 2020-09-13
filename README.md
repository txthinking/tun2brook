# ipio
ipio can convert Network/IPv4/IPv6 layer packets to Transport/TCP/UDP layer packets, such as tun to socks5, tun to brook server, tun to brook wsserver.

### Pure Go

No C, No CGO, written in Pure Go, can be used in any OS, such as Linux, macOS, Android, iOS, etc...

### Install via [nami](https://github.com/txthinking/nami)

```
$ nami install github.com/txthinking/ipio
```

### Usage

```
NAME:
   ipio - ipio can convert Network/IPv4/IPv6 layer packets to Transport/TCP/UDP layer packets, such as tun to socks5, tun to brook server, tun to brook wsserver, ...

USAGE:
   ipio [global options] command [command options] [arguments...]

VERSION:
   20200909

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

> Assume your brook server is `1.2.3.4:9999`, password is `hello`

```
$ ipio tun2brookserver -s 1.2.3.4:9999 -p hello
```

> ROOT privileges required, more `$ ipio tun2brookserver -h`

#### tun2brookwsserver

> Assume your brook wsserver is `ws://1.2.3.4:9999`, password is `hello`

```
$ ipio tun2brookwsserver -s ws://1.2.3.4:9999 -p hello
```

> ROOT privileges required, more `$ ipio tun2brookwsserver -h`

#### tun2socks

> Assume your socks5 server is `1.2.3.4:1080`

```
$ ipio tun2socks -s 1.2.3.4:1080
```

> ROOT privileges required, more `$ ipio tun2socks -h`

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
