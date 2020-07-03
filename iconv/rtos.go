package rtos

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"time"
)
var MSG_UART_RXDATA =lua.LNumber(1)
var MSG_UART_TX_DONE=lua.LNumber(2)
var MSG_SOCK_CONN_CNF=lua.LNumber(3)
var MSG_SOCK_CLOSE_CNF=lua.LNumber(4)
var MSG_SOCK_CLOSE_IND=lua.LNumber(5)
var MSG_SOCK_SEND_CNF=lua.LNumber(6)
var MSG_SOCK_RECV_IND=lua.LNumber(7)
var MSG_INT          =lua.LNumber(8)
var MSG_TIMER        =lua.LNumber(9)
var MSG_PDP_DEACT_IND=lua.LNumber(10)
var INF_TIMEOUT      =lua.LNumber(11)
func GetValue(L *lua.LState)*lua.LTable{
	ret:=L.NewTable()
	var exports = map[string]lua.LGFunction{
		"open": open,
	}
	L.SetFuncs(ret, exports)
	L.SetField(ret, "MSG_UART_RXDATA", MSG_UART_RXDATA)
	L.SetField(ret, "MSG_UART_TX_DONE", MSG_UART_TX_DONE)
	L.SetField(ret, "MSG_SOCK_CONN_CNF", MSG_SOCK_CONN_CNF)
	L.SetField(ret, "MSG_SOCK_CLOSE_CNF",MSG_SOCK_CLOSE_CNF)
	L.SetField(ret, "MSG_SOCK_CLOSE_IND", MSG_SOCK_CLOSE_IND)
	L.SetField(ret, "MSG_SOCK_SEND_CNF", MSG_SOCK_SEND_CNF)

	L.SetField(ret, "MSG_SOCK_RECV_IND", MSG_SOCK_RECV_IND)
	L.SetField(ret, "MSG_INT", MSG_INT)
	L.SetField(ret, "MSG_TIMER", MSG_TIMER)

	L.SetField(ret, "MSG_PDP_DEACT_IND", MSG_PDP_DEACT_IND)
	L.SetField(ret, "INF_TIMEOUT", INF_TIMEOUT)
	L.Push(ret)
	return ret
}
