package socketcore

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"golua/golua/rtos"
	"net"
)
var pool = map[int]*sock{}


func GetValue(L *lua.LState)*lua.LTable{
	ret:=L.NewTable()

	var exports = map[string]lua.LGFunction{
		"sock_conn_ext": sock_conn_ext,
		"sock_send":sock_send,
		"sock_recv":sock_recv,
		"sock_destroy":sock_destroy,
		"sock_close":sock_close,
	}
	L.SetFuncs(ret, exports)
	return ret
}

func sock_send(L *lua.LState)int{
	sock_id := L.CheckInt(1)
	data    := L.CheckString(2)
	if s,ok:=pool[sock_id];ok && s.conn!=nil{
		_,err:=s.conn.Write([]byte(data))
		msg:=L.NewTable()
		L.SetField(msg, "id", rtos.MSG_SOCK_SEND_CNF)
		L.SetField(msg, "socket_index", lua.LNumber(s.id))
		if err!=nil{
			fmt.Println("写错误",err)
			L.SetField(msg, "result", lua.LNumber(1))
		}else{
			L.SetField(msg, "result", lua.LNumber(0))
		}
			rtos.CallReceive(rtos.MSG_SOCK_SEND_CNF,msg,nil)
	}
	return 0
}
func sock_conn_ext(L *lua.LState)int{


	//连接类型
	conn_type := L.CheckInt(1)
	ip:=L.CheckString(2)
	port:=L.CheckAny(3).String()
	/*
	if !IsPublicIP(net.ParseIP(ip)){
		fmt.Println("连接了非公网IP"+ip)
		panic(0)
	}
	*/
	//switch conn_type{
	//case 0:
		s:=newsock()
		go s.start(L,conn_type,ip+":"+port)
		L.Push(lua.LNumber(s.id))
	//}
	return 1
}
func sock_recv(L *lua.LState)int{
	sock_id := L.CheckInt(1)
	if s,ok:=pool[sock_id];ok && s.buf!=nil{
		L.Push(lua.LString(s.buf))
		s.buf=[]byte{}
	}
	return 1
}
func sock_destroy(L *lua.LState)int{
	return 0
}
func sock_close(L *lua.LState)int{
	sock_id := L.CheckInt(1)
	if s,ok:=pool[sock_id];ok && s.conn!=nil{
		s.isclose=true
		s.conn.Close()
		msg:=L.NewTable()
		L.SetField(msg, "id", rtos.MSG_SOCK_CLOSE_CNF)
		L.SetField(msg, "socket_index", lua.LNumber(s.id))
		rtos.CallReceive(rtos.MSG_SOCK_CLOSE_CNF,msg,nil)
	}
	return 0
}


func IsPublicIP(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := IP.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}