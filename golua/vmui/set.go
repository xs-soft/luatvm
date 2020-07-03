package vmui

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"golang.org/x/sys/windows/registry"
	"regexp"
)
var imeiCB *walk.ComboBox
//PIO显示台
var imeiset func(string)
//imei选项
var imeis []string
func SetImeis(s []string){
	imeis=s
}
func SetImeiCallBack(f func(string)){
	imeiset=f
}
func setui()Widget{
return Composite{
		Layout: Grid{Columns: 3,MarginsZero: true},
		MaxSize: Size{Width: 250},
		Children: []Widget{
			Label{Text: "IMEI:",MaxSize: Size{Width:30}},
			ComboBox{AssignTo: &imeiCB,Editable: true,
				Model: imeis,
				Persistent: true, Value: imeis[0],MaxSize: Size{120,30}},
			PushButton{Text: "确定",MaxSize: Size{50,30},OnClicked: func(){
				imei := imeiCB.Text()
				b,_:=regexp.MatchString(`\d{10}`,imeiCB.Text())
				if !b{
					walk.MsgBox(mw, "设置IMEI", "IMEI请设置16位数字", walk.MsgBoxIconWarning)
				} else {
					newdata:=[]string{imei}
					//把缓存中的除了数组之外的存储起来
					for _,v:=range imeis{
						find:=false
						for _,v2:=range newdata{
							if v==v2{
								find=true
							}
						}
						if !find && v!="0000000000000000"{
							newdata=append(newdata,v)
						}
					}
					if len(newdata)>10{
						newdata=newdata[0:10]
					}
					key, _, _ := registry.CreateKey(registry.CURRENT_USER, "SOFTWARE\\LuatVm", registry.ALL_ACCESS)
					key.SetStringsValue("imei",newdata)
					if imeiset!=nil{
						imeiset(imei)
					}
				}
			}},
		},
	}
}