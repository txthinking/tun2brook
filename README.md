# ipio
ipio can convert Network/IPv4/IPv6 layer packets to Transport/TCP/UDP layer packets, such as tun to socks5, tun to brook server, tun to brook wsserver.

### Pure Go

No C, No CGO, written in Pure Go, can be used in any OS, such as Linux, macOS, Android, iOS, etc...

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
    c, _ := socks5.NewClient(server, username, password, 60, 0, 60)
    rc, _ := s.Client.Dial("tcp", conn.RemoteAddr().String()) 
    go io.Copy(conn, rc)
    io.Copy(rc, conn)
}

func (t *Socks5) UDP(conn net.Conn) {
    c, _ := socks5.NewClient(server, username, password, 60, 0, 60)
    rc, _ := s.Client.Dial("udp", conn.RemoteAddr().String()) 
    go io.Copy(conn, rc)
    io.Copy(rc, conn)
}

// func NewIPIO(t Interface, h Handler) *IPIO
ii := ipio.NewIPIO(tun, &Socks5{})

ii.Run()

ii.Close()
```
