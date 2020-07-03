package uart

import (
	"golua/golua/at/CIP"
	"strconv"
)

func init(){
	CIP.ConnFailFunc=func(i string){
		sendat(i+", CONNECT FAIL")
	}
	CIP.ConnOkFunc=func(i string){
		sendat(i+", CONNECT OK")
	}
	CIP.SendOkFunc=func(i string){
		sendat(i+", SEND OK")
	}
	CIP.SendFailFunc=func(i string){
		sendat(i+", UDP ERROR")
	}
	CIP.ReadOkFunc=func(i string,b []byte){
		if len(b)==0 {
			return
		}
		sendat("+RECEIVE,"+i+","+strconv.Itoa(len(b))+":")
		sendat(string(b))
	}
	CIP.ReadFailFunc=func(i string){
		sendat(i+", UDP ERROR")
	}
	CIP.CloseOk=func(i string){
		sendat(i+", CLOSE OK")
	}
	regfunc(`CIPSTART=(\d+),"(.*)","(.*)",(\d+)`, func(strs []string) {
		sendat("OK")
		CIP.Start(strs[1],strs[2],strs[3],strs[4])
	})
	regfunc(`CIPSEND=(\d+),(\d+)`, func(strs []string) {
		l,_:=strconv.Atoi(strs[2])
		sendatwait("> ",l,func(s string){
			CIP.Send(strs[1],s)
		})
	})
	regfunc(`CIPCLOSE=(\d+)`, func(strs []string) {
		CIP.Close(strs[1])
		go func() {
			sendat(strs[1]+", CLOSE OK")
		}()
	})
}
