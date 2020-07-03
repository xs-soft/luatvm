package uart

import (
	"golua/golua/at/CIP"
	"strconv"
)

func init(){
	CIP.SSLConnFailFunc=func(i string){
		sendat(i+", CONNECT FAIL")
	}
	CIP.SSLConnOkFunc=func(i string){
		sendat(i+", CONNECT OK")
	}
	CIP.SSLSendOkFunc=func(i string){
		sendat(i+", SEND OK")
	}
	CIP.SSLSendFailFunc=func(i string){
		sendat(i+", UDP ERROR")
	}
	CIP.SSLReadOkFunc=func(i string,b []byte){
		sendat("+SSL RECEIVE,"+i+", "+strconv.Itoa(len(b))+":")
		sendat(string(b))
	}
	CIP.SSLReadFailFunc=func(i string){
		sendat(i+", UDP ERROR")
	}
	CIP.SSLCloseOk=func(i string){
		sendat(i+", CLOSE OK")
	}
	regfunc(`SSLINIT`, func(strs []string) {
		//sendat(strs[1]+", CLOSE OK")
		sendat("SSL&"+CIP.SslInit()+",INIT OK")
	})
	regfunc(`SSLCONNECT=(\d+)`, func(strs []string) {
		sendat("OK")
	})
	regfunc(`SSLCREATE=(\d+),"(.+)",(\d+)`, func(strs []string) {
		CIP.SSLStart(strs[1],strs[2])
		sendat("SSL&"+strs[1]+", CREATE OK")
	})
	regfunc(`SSLSEND=(\d+),(\d+)`, func(strs []string) {
		l,_:=strconv.Atoi(strs[2])
		sendatwait("> ",l,func(s string){
			CIP.SSLSend(strs[1],s)
		})
	})
	regfunc(`SSLDESTROY=(\d+)`, func(strs []string) {
		CIP.SSLClose(strs[1])
	})
}
