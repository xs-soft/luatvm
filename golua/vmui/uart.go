package vmui

import (
	"fmt"
	"github.com/lxn/walk"
	"github.com/tarm/serial"
	"golang.org/x/sys/windows/registry"
	"log"
	"strings"
)
import . "github.com/lxn/walk/declarative"

var uartSelect=[]*walk.ComboBox{
	{},
	{},
}
var uartLabel=[]*walk.TextLabel{
	{},
	{},
}
var comList []string
var UartCallback func (int,string)
//设置UART回调
func UartCallFunc(f func (int,string)){
	UartCallback=f
}
//
func UartState(id int,msg string){
	uartLabel[id-1].SetText(msg)
}
func init(){
	comList=append(comList,"不使用")
	for i:=1;i<=100;i++{
		name:="COM"+fmt.Sprint(i)
		con:=&serial.Config{Name: name, Baud: 115200}
		conn,err:=serial.OpenPort(con)
		//未找到串口,找下一个
		if err!=nil && strings.Contains(err.Error(),"cannot find the file"){
			continue
		}
		//断开连接
		if err==nil{
			conn.Close()
		}
		comList=append(comList,name)
	}
	//fmt.Println()

	//NewUart(1)
	//NewUart(2)
}
func NewUart(id int)HSplitter{
	key, exists, _ := registry.CreateKey(registry.CURRENT_USER, "SOFTWARE\\LuatVm", registry.ALL_ACCESS)
	value := "不使用"
	if exists{
		value,_,_=key.GetStringValue("UART"+fmt.Sprint(id))
	}
	return HSplitter{
		Children: []Widget{
			TextLabel{Text: "UART"+fmt.Sprint(id)+":",MaxSize: Size{30,20}},
			ComboBox{AssignTo: &uartSelect[id-1],Enabled:true,Model: comList,
				Editable: false,
				//+fmt.Sprint(id)
				Value:    value,
				OnCurrentIndexChanged: func() {
					if UartCallback!=nil{
						key, _, err := registry.CreateKey(registry.CURRENT_USER, "SOFTWARE\\LuatVm", registry.ALL_ACCESS)
						if err != nil {
							log.Fatal(err)
						}
						defer key.Close()

						key.SetStringValue("UART"+fmt.Sprint(id),uartSelect[id-1].Text())


						UartCallback(id,uartSelect[id-1].Text())
					}
					//var url string
					/*
					if index := tableView.CurrentIndex(); index > -1 {
						name := tableModel.items[index].Name
						dir := treeView.CurrentItem().(*Directory)
						url = filepath.Join(dir.Path(), name)
					}
					webView.SetURL(url)
					 */
				},
			},
			TextLabel{Text: "未调用",MinSize: Size{100,20},AssignTo: &uartLabel[id-1]},
		},
	}
}
type Species struct {
	Id   int
	Name string
}
//PIO显示台
func uart()Widget {
	return VSplitter{
		StretchFactor: 0,
		MaxSize:       Size{200, 80},
		Children: []Widget{
			NewUart(1),
			NewUart(2),
		},
	}
}