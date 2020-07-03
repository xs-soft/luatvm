package socketcore

import (
	"crypto/tls"
	lua "github.com/yuin/gopher-lua"
	"golua/golua/rtos"
	"io"
	"net"
	"time"
)

func newsock()*sock{
	ret:=new(sock)
	i:=0
	for{
		i++
		if _,ok:=pool[i];!ok{
			ret.id=i
			pool[i]=ret
			return ret
		}
	}
	return ret
}
type sock struct{
	id int
	network string
	addr string
	conn net.Conn
	typ int
	//数据缓冲
	buf []byte
	err error
	isclose bool
}

func (s *sock)start(L *lua.LState,conn_type int,addr string){
	s.isclose=false
	if conn_type == 0{
		s.conn,s.err = net.DialTimeout("tcp",addr,time.Second*10)
	}
	if conn_type == 2{
		s.conn,s.err =tls.Dial("tcp",addr,&tls.Config{})
	}
	if conn_type == 1 {
		s.conn,s.err = net.DialTimeout("udp",addr,time.Second*10)
	}
	msg:=L.NewTable()
	L.SetField(msg, "id", rtos.MSG_SOCK_CONN_CNF)
	L.SetField(msg, "socket_index", lua.LNumber(s.id))
	if s.err == nil{
		go s.read(L)
		L.SetField(msg, "result", lua.LNumber(0))
	}else{
		L.SetField(msg, "result", lua.LNumber(1))
	}
	rtos.CallReceive(rtos.MSG_SOCK_CONN_CNF,msg,nil)
}
func (s *sock)read(L *lua.LState){
	for {
		b := make([]byte, 8000)
		//s.conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		i,err:=s.conn.Read(b)
		if err==io.EOF{
			continue
		}
		if err!=nil {
			if !s.isclose{
				msg:=L.NewTable()
				L.SetField(msg, "id", rtos.MSG_SOCK_CLOSE_IND)
				L.SetField(msg, "socket_index", lua.LNumber(s.id))
				rtos.CallReceive(rtos.MSG_SOCK_CLOSE_IND,msg,nil)
			}
			return
		}
		b=b[:i]
		msg:=L.NewTable()
		L.SetField(msg, "id", rtos.MSG_SOCK_RECV_IND)
		L.SetField(msg, "socket_index", lua.LNumber(s.id))
		L.SetField(msg, "recv_len"    , lua.LNumber(len(b)))
		rtos.CallReceive(rtos.MSG_SOCK_RECV_IND,msg,func(){
			s.buf=b
		})
	}
}