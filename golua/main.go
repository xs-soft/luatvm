package main

import (
	"flag"
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"golua/golua/bit"
	"golua/golua/crypto"
	"golua/golua/disp"
	"golua/golua/iconv"
	"golua/golua/pack"
	"golua/golua/qrencode"
	"golua/golua/rtos"
	"golua/golua/socketcore"
	"golua/golua/uart"
	"golua/golua/vmui"
	_ "golua/golua/vmui"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)
var L *lua.LState
var syslog =flag.Bool("syslog",true,"是否显示系统级日志,可查看core/print.lua")
func main() {
	flag.Parse()
	os.Remove("lib_err.txt")
	go func (){http.ListenAndServe(":8080", nil)}()
	vmui.StartUi()
	time.Sleep(time.Millisecond*500)
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println("错误信息:",err) // 这里的err其实就是panic传入的内容
		}
		fmt.Println("程序已经停止：请输入任意内容退出")
		input:=""
		fmt.Scanln(&input)
	}()
	_,err:=os.Stat("lib")
	if err != nil{
		fmt.Println("未找到lib目录,请复制luat的lib目录到工作目录")
	}
	//imei更换绑定

	vmui.SetImeiCallBack( func(s string){
		uart.Imei=s
		L.DoString(`misc.setImei("`+s+`")`)
	})
	//fmt.Println(0xfff87654)
	//fmt.Println(bit.MaxInt)
	//fmt.Println(dumpbit(bit.MinInt))
	//fmt.Println(ashift(0x87654321,8))
	//fmt.Println(dumpbit(int(0x87654321>>12)))
	//fmt.Println(dumpbit(int(0xfff87654)))
	//fmt.Println(ashift(-255,8))
	//fmt.Println(ashift(0x87654321,12))
	//fmt.Println(0x87654321&^(1<<32))
	//fmt.Println(dumpbit(16777215))
	//fmt.Println(dumpbit(-256))
	//fmt.Println(dumpbit((-256 & bit.MaxUint)>>8))
	//fmt.Println((-256 & bit.MaxUint)>>8)
	//fmt.Println(dumpbit(bit.MaxUint))
	//fmt.Println(dumpbit())
	//fmt.Println(dumpbit(^4))
	//fmt.Println(dumpbit(-1))
	time.Sleep(time.Millisecond*100)
	os.Remove("/lib_err.txt")
	L = lua.NewState()
	defer L.Close()
	AddMod("uart",uart.GetValue(L))
	AddMod("rtos",rtos.GetValue(L))
	//AddMod("json",ljson.GetValue(L))
	AddMod("crypto",crypto.GetValue(L))
	//AddMod("pio",pio.GetValue(L))
	AddMod("socketcore",socketcore.GetValue(L))
	AddMod("bit",bit.GetValue(L))
	AddMod("disp",disp.GetValue(L))
	AddMod("iconv",iconv.GetValue(L))
	AddMod("pack",pack.GetValue(L))
	AddMod("qrencode",qrencode.GetValue(L))
	AddMod("vmui",vmui.GetValue(L))

	L.SetGlobal("ISVM", lua.LBool(true))
	L.SetGlobal("DISPSYSLOG", lua.LBool(*syslog))
	corelist,err:=ioutil.ReadDir("core")
	if err!=nil{
		panic("未找到core目录")
	}
	for _,v:=range corelist{
		if !v.IsDir() {
			if err := L.DoFile("core/"+v.Name()); err != nil {
				panic(err)
			}
		}
	}



	if err := L.DoFile("main.lua"); err != nil {
		panic(err)
	}
}
func AddMod(name string,table *lua.LTable){
	L.SetGlobal(name, table)
	L.PreloadModule(name, func(L *lua.LState) int {
		L.Push(table)
		return 1
	})
}

type tfile struct {
	*os.File
}
func (t *tfile)Write(b []byte) (n int, err error){
	return n,nil
}