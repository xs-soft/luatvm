package uart

import (
	"github.com/jacobsa/go-serial/serial"
	lua "github.com/yuin/gopher-lua"
	"golua/golua/rtos"
	"golua/golua/vmui"
	"io"
	"strings"
	"sync"
	"time"
)
var umap=map[int]*uart{}
func init(){
	umap[0]=newuart(0,"COM0")
	umap[1]=newuart(1,"")
	umap[2]=newuart(2,"")
	vmui.UartCallFunc(OnCOMEVENT)
}

//设置绑定端口号
func OnCOMEVENT(uart_id int ,name string){
	umap[uart_id].setname(name)
}
func start(){

//	umap[1]=newuart(2,"COM6")
//	umap[2]=newuart(2,"COM6")
//	umap[2].start()
}
//
func SetUART(uart_id int,comname string){

}
func newuart(id int64,name string)*uart{
	ret:=new(uart)
	ret.Id=id
	ret.options=serial.OpenOptions{
		PortName:name,
	}
	ret.cache =[]byte{}
	if id>0 {
		ret.start()
	}

	return ret
}

type uart struct{
	Id      int64
	options serial.OpenOptions
	conn    io.ReadWriteCloser
	cache   []byte
	mu      sync.RWMutex
	isset   bool
}
func (u *uart)start(){
	go func(){
		for {
			//没有指令发出，或者未设置端口信息
			if u.options.PortName=="" || !u.isset{
				time.Sleep(time.Millisecond*10)
				continue
			}
			var err error
			u.conn, err = serial.Open(u.options)
			if err != nil {
				if strings.Contains(err.Error(),"Access is denied"){
					vmui.UartState(int(u.Id),"被占用")
				}
				if strings.Contains(err.Error(),"The parameter is incorrect.") {
					vmui.UartState(int(u.Id),"参数异常")
				}
				time.Sleep(time.Second)
				continue
			}
			//log.Print("连接成功: %v")
			vmui.UartState(int(u.Id),"连接成功")
			for{
				but   :=make([]byte,1024,1024)
				n,err := u.conn.Read(but)
				if err!=nil{
					//log.Print("读错误",err)
					break
				}
				but     = but[0:n]
				u.cache = append(u.cache,but...)
				rtos.CallReceive(rtos.MSG_UART_RXDATA,lua.LNumber(u.Id),nil)
			}
		}
	}()
}
func (u *uart)setname(name string){
	u.options.PortName=name
	if u.conn!=nil{
		u.conn.Close()
	}
}
func (u *uart)setup(BaudRate uint,DataBits uint, StopBits uint){
	u.isset=true
	u.options.BaudRate=BaudRate
	u.options.DataBits=DataBits
	u.options.StopBits=StopBits
	if u.conn!=nil{
		u.conn.Close()
	}
}
func (u *uart)read(l int)[]byte{
	u.mu.Lock()
	defer u.mu.Unlock()
	if l>len(u.cache){
		l=len(u.cache)
	}
	ret:=u.cache[:l]
	u.cache=u.cache[l:]
	return ret
}