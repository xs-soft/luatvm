package CIP

import (
	"crypto/tls"
	"io"
	"net"
	"strings"
	"sync"
	"time"
)
//连接失败
var ConnFailFunc func(string)
//连接成功
var ConnOkFunc func(string)
//发送失败
var SendFailFunc func(string)
var SendOkFunc func(string)
var ReadFailFunc func(string)
var ReadOkFunc func(string,[]byte)
var CloseOk func(string)

var SSLConnOkFunc func(string)
var SSLConnFailFunc func(string)
//发送失败
var SSLSendFailFunc func(string)
var SSLSendOkFunc func(string)
var SSLReadFailFunc func(string)
var SSLReadOkFunc func(string,[]byte)
var SSLCloseOk func(string)


var pool=map[string]net.Conn{}
var mu sync.Mutex
func Start(id string,protocol string,host string,port string){
	go func (){
		var conn net.Conn
		var err error
		conn,err=net.Dial(strings.ToLower(protocol),host+":"+port)
		if err!=nil{
			ConnFailFunc(id)
			return
		}
		mu.Lock()
		pool[id]=conn
		mu.Unlock()
		//读线程
		go func(){
			buf := make([]byte, 4096)    //创建一个切片，存储客户端发送的数据

			for {
				conn.SetReadDeadline(time.Now().Add(time.Second*5))
				//读取用户数据
				n, err := conn.Read(buf)
				if err != nil && err!=io.EOF {

					//fmt.Println("网络通讯错误",err)
					//if err!=err
					//ReadFailFunc(id)
					return
				}

				ReadOkFunc(id,buf[:n])
				//服务器处理数据：把客户端数据转大写，再写回给client
				//conn.Write([]byte(strings.ToUpper(string())))
			}
		}()
		ConnOkFunc(id)
	}()
}
func Send(id string,s string){
	go func(){
	mu.Lock()
	conn:=pool[id]
	mu.Unlock()
	_,err:=conn.Write([]byte(s))
	if err!=nil{
		SendFailFunc(id)
		return
	}
	SendOkFunc(id)

	}()
}

func Close(id string){
	mu.Lock()
	pool[id].Close()
	pool[id]=nil
	mu.Unlock()
	CloseOk(id)
}
func SslInit()string{
	return "0"
}

func SSLStart(id string,host string){
	oldid:=id
	id="SSL&"+id
	go func (){
		var conn net.Conn
		var err error
		conn,err=tls.Dial("tcp",host,nil)
		if err!=nil{
			SSLConnFailFunc(id)
			return
		}
		mu.Lock()
		pool[id]=conn
		mu.Unlock()
		SSLConnOkFunc(id)
		go func (){
			for {
			//conn.SetReadDeadline(time.Now().Add(time.Second*5))
			b:=make([]byte,1024,1024)
			i,err:=conn.Read(b)
			if err!=nil{
				SSLReadFailFunc(id)
				return
			}
			b=b[0:i]
			SSLReadOkFunc(oldid,b)
			}
		}()
	}()
}

func SSLSend(id string,s string){
	//oldid:=id
	id="SSL&"+id
	go func(){
		mu.Lock()
		conn:=pool[id]
		mu.Unlock()
		_,err:=conn.Write([]byte(s))
		if err!=nil{
			SSLSendFailFunc(id)
			return
		}
		SSLSendOkFunc(id)
	}()
}
func SSLClose(id string){
	id="SSL&"+id
	mu.Lock()
	pool[id].Close()
	pool[id]=nil
	mu.Unlock()
	SSLCloseOk(id)
}