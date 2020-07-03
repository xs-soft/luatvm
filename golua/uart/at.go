package uart

import (
	lua "github.com/yuin/gopher-lua"
	"golang.org/x/sys/windows/registry"
	"golua/golua/rtos"
	"golua/golua/vmui"
	"regexp"
	"strings"
	"time"
)
var Imei = "866262047153787"
var inputcache=[]byte{}
var inputfunc  func(string)
var inputlen   int
//取得当前IMEI
func init(){

	key, _, _ := registry.CreateKey(registry.CURRENT_USER, "SOFTWARE\\LuatVm", registry.ALL_ACCESS)
	imeis,_,err:=key.GetStringsValue("imei")
	if err!=nil{
		imeis=[]string{}
	}
	if len(imeis)==0 {
		imeis=append(imeis,"0000000000000000")
	}else{
		Imei=imeis[0]
	}
	vmui.SetImeis(imeis)

}
type EXP struct{
	exp string
	do []interface{}
}
var ATEXP []EXP

func regfunc(name string,f func([]string)){
	e:=EXP{
		exp:"AT[\\+\\*]"+name,
		do:[]interface{}{f},
	}
	ATEXP = append(ATEXP,e)
}

var AtMap=map[string][]interface{}{
	"AT+CGSN":{ATCGSN},
}
var SIMST = false
//获取IMEI指令响应
func ATCGSN(){
	sendat(Imei)
	sendat("OK")
}
func readat(DATA string,L *lua.LState){

	if inputfunc!=nil{
		//存在等待输入
		inputcache=append(inputcache,[]byte(DATA)...)
		if len(inputcache)>=inputlen{
			//输入完成
			inputfunc(string(inputcache))
		}
		inputcache=[]byte{}
		inputfunc=nil
		return
	}
	DATA=strings.ReplaceAll(DATA,"\n","")
	DATA=strings.ReplaceAll(DATA,"\r","")
	for _,exp:=range ATEXP{
		ok,err:=regexp.MatchString(exp.exp,DATA)
		if err!=nil{
			panic(err)
		}
		if !ok{
			continue
		}
		for _,v:=range exp.do{
			switch v.(type) {
			case string:
				sendat(v.(string))
			case func():
				v.(func())()
			case func(i []string):
				v.(func(i []string))(regexp.MustCompile(exp.exp).FindStringSubmatch(DATA))
			}
		}
		return
	}
	//设置IMEI特别处理
	if  strings.Contains(DATA,"AT+WIMEI="){
		sendat("OK")
		return
	}
	//fmt.Println("串口",UART_ID,DATA)
	//fmt.Println(DATA)
	if s,ok:=AtMap[DATA];ok{
		for _,v:=range s{
			switch v.(type){
			case string:
				sendat(v.(string))
			case func():
				v.(func())()
			}
		}
	}else{
		//处理不了的AT送给lua
		if err := L.CallByParam(lua.P{
			Fn: rilatfunc,
			NRet: 1,
			Protect: true,
		}, lua.LString(DATA)); err != nil {
			panic(err)
		}
	}
	//L.Get(-1) // returned value
	L.Pop(1)  // remove received value
}
func sendatwait(s string,len int,f func (string)){
	rtos.CallReceive(rtos.MSG_UART_RXDATA,lua.LNumber(umap[0].Id),func(){
		umap[0].cache = []byte(s+"\r\n")
		inputfunc=f
	})
}
func sendat(s string){
	rtos.CallReceive(rtos.MSG_UART_RXDATA,lua.LNumber(umap[0].Id),func(){
		umap[0].cache = []byte(s+"\r\n")
	})
}

func init(){
	go func(){
	time.Sleep(time.Second)
		sendat("RDY")
		sendat("RDY")
	}()
}