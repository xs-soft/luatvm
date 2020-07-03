package uart

import (
	lua "github.com/yuin/gopher-lua"
)
var rilatfunc *lua.LFunction
func GetValue(L *lua.LState)*lua.LTable{
	ret:=L.NewTable()
	var exports = map[string]lua.LGFunction{
		"setup": setup,
		"write": write,
		"read" : read,
		"on"   : on,
		"set_rs485_oe":set_rs485_oe,
		"close":uartclose,
		"rilAT":rilAT,
		"rilATSend":rilATSend,
	}

	L.SetFuncs(ret, exports)
	L.SetField(ret, "ATC", lua.LNumber(0))
	L.SetField(ret, "PAR_NONE", lua.LNumber(0))
	L.SetField(ret, "STOP_1", lua.LNumber(1))

	L.Push(ret)
	return ret
}
func set_rs485_oe(L *lua.LState)int{
	return 0
}
func uartclose(L *lua.LState)int{
	return 0
}
func setup(L *lua.LState)int{
	UART_ID  :=L.CheckInt(1)
	if UART_ID==0{
		return 0
	}
	BaudRate :=L.ToInt(2)
	DataBits :=L.CheckInt(3)
	StopBits :=L.CheckInt(5)
	//fmt.Println(UART_ID)
	umap[UART_ID].setup(uint(BaudRate),uint(DataBits),uint(StopBits))
	return 0
}
func write(L *lua.LState)int{
	UART_ID:=L.CheckInt(1)
	DATA   :=L.CheckString(2)
	if u,ok:=umap[UART_ID];ok && u.conn!=nil{
		u.conn.Write([]byte(DATA))
	}else{
		if UART_ID==0{
			readat(DATA,L)
		}
	}
	return 0
}
func read(L *lua.LState)int{
	UART_ID:=L.CheckInt(1)
	//Len    :=L.CheckInt(2)
	if u,ok:=umap[UART_ID];ok && (u.conn!=nil || UART_ID==0){
		L.Push(lua.LString(u.read(10000)))
	}else{
		L.Push(lua.LString(""))
	}
	return 1
}

func on(L *lua.LState)int{
	return 0
}
func rilAT(L *lua.LState)int{
	f:=L.CheckFunction(1)
	rilatfunc=f
	return 1
}
func rilATSend(L *lua.LState)int{
	f:=L.CheckString(1)
	//--go func(){
	sendat(f)
	//--}()
	return 0
}
