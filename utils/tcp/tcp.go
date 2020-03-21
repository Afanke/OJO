package tcp

import (
	"net"
	"time"
)

type Conn struct {
	conn net.Conn
}

type Listener struct {
	listener net.Listener
}

func Listen(addr string) (Listener, error) {
	res, err := net.Listen("tcp", addr)
	return Listener{listener: res}, err
}

func Dial(addr string) (Conn, error) {
	conn, err := net.Dial("tcp", addr)
	return Conn{conn: conn}, err
}

func (l Listener) Accept() (Conn, error) {
	conn, err := l.listener.Accept()
	return Conn{conn: conn}, err
}

func (l Listener) Addr() net.Addr {
	return l.listener.Addr()
}

func (l Listener) Close() error {
	return l.Close()
}

func (c Conn) Send(s []byte) (int, error) {
	var n [4]byte
	x := len(s)
	n[0] = uint8(x)
	n[1] = uint8(x >> 8)
	n[2] = uint8(x >> 16)
	n[3] = uint8(x >> 24)
	_, err := c.conn.Write(n[0:])
	if err != nil {
		return 0, err
	}
	lens, err := c.conn.Write(s)
	return lens, err
}

func (c Conn) Recv() (uint32, []byte, error) {
	var n [4]byte
	for {
		_, err := c.conn.Read(n[0:])
		if err != nil {
			return 0, nil, err
		}
		var total uint32
		total |= uint32(n[0])
		total |= uint32(n[1]) << 8
		total |= uint32(n[2]) << 16
		total |= uint32(n[3]) << 24
		buf := make([]byte, total)
		for read := uint32(0); read < total; {
			lens, err := c.conn.Read(buf[read:])
			//fmt.Println("lens",lens)
			if err != nil {
				return 0, nil, err
			}
			read += uint32(lens)
			//fmt.Println("read",read)
		}
		return total, buf, nil
	}
}

func (c Conn) Close() error {
	return c.conn.Close()
}

func (c Conn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c Conn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c Conn) SetDeadline(t time.Time) error {
	return c.conn.SetDeadline(t)
}

func (c Conn) SetReadDeadline(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}

func (c Conn) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}
