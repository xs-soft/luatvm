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
var MSG_AUDIO      =lua.LNumber(11)

var starttime =time.Now().UnixNano()
func GetValue(L *lua.LState)*lua.LTable{
	ret:=L.NewTable()
	var exports = map[string]lua.LGFunction{
		"on": On,
		"poweron_reason": poweron_reason,
		"remove_fatal_info": remove_fatal_info,
		"timer_start": timer_start,
		"timer_stop": timer_stop,
		"get_version": get_version,
		"poweron"    : poweron,
		"fota_start" :fota_start,
		"receive"    :receive,
		"restart"    :restart,
		"get_fs_free_size":get_fs_free_size,
		"tick"       :tick,
		"set_trace_port":set_trace_port,
		"setvol"     :setvol,
		"sys32k_clk_out":sys32k_clk_out,
		"gotest":gotest,
		"fota_process": fota_process,
		"fota_end":fota_end,
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
	L.SetField(ret, "MSG_AUDIO", MSG_AUDIO)

	L.SetField(ret, "INF_TIMEOUT", INF_TIMEOUT)

	L.Push(ret)
	return ret
}

var TimeMap =map[int64]int{}
var TimeId  =0
var msg_chan =make(chan func (L *lua.LState)int,1000)
func On(L *lua.LState)int{
	L.Push(lua.LString("aaa"))
	return 1
}
//设置调试端口???
func set_trace_port(L *lua.LState)int{
	return 0
}
func poweron_reason(L *lua.LState)int {
	L.Push(lua.LString(""))
	return 1
}

func remove_fatal_info(L *lua.LState)int{
	return 0
}
func setvol(L *lua.LState)int{
	return 0
}

func timer_start(L *lua.LState)int{
	id  :=L.CheckAny(1)
	val:=L.CheckInt(2)
	go func(){
		time.Sleep(time.Duration(val)*time.Millisecond)
		msg_chan<-func(L *lua.LState)int{
			L.Push(MSG_TIMER)
			L.Push(id)
			return 2
		}
	}()
	L.Push(lua.LNumber(1))
	return 1
}
func get_fs_free_size(L *lua.LState)int{
	L.Push(lua.LNumber(0))
	return 1
}

func timer_stop(L *lua.LState)int{
	return 0
}
func get_version(L *lua.LState)int{
	L.Push(lua.LString("Luat_V0025_ASR1802_FLOAT_720H SSL"))
	return 1
}
func poweron(L *lua.LState)int{
	return 0
}
func fota_start(L *lua.LState)int{
	L.Push(lua.LNumber(1))
	return 1
}
func receive(L *lua.LState)int{
	newfun := <-msg_chan
	return newfun(L)
}
func restart(L *lua.LState)int{
	fmt.Println("系统重启")
	return 0
}
func tick(L *lua.LState)int{
	ms:=(time.Now().UnixNano()-starttime)/61
	L.Push(lua.LNumber(ms))
	return 1
}
func sys32k_clk_out(L *lua.LState)int{
	return 0
}
func gotest(L *lua.LState)int{
	//L.Push(lua.LTrue)
	L.Push(lua.LFalse)
	return 1
}
func fota_process(L *lua.LState)int{
	L.Push(lua.LNumber(0))
	return 1
}
func fota_end(L *lua.LState)int{
	L.Push(lua.LNumber(0))
	return 1
}


func CallReceive(msgid lua.LValue,value lua.LValue,run func()){
	//fmt.Println("插入函数")
	msg_chan<-func(L *lua.LState)int{
		if run!=nil{
			run()
		}
		L.Push(msgid)
		L.Push(value)
		return 2
	}
}
