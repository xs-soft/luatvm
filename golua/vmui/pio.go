package vmui

import (
	"github.com/lxn/walk"
)
import . "github.com/lxn/walk/declarative"

func PioSet(id int,val int){
	if id>4{
		return
	}
	var cval byte
	if val ==1{
		cval=255
	}
	b,_:=SolidColorBrush{Color:  walk.RGB(cval, 0, 0)}.Create()
	pios[id-1].SetBackground(b)
}
var pios []*walk.CustomWidget
func init(){
	for i:=0;i<=3;i++{
		pios=append(pios,&walk.CustomWidget{})
	}
}
func NewLed(name string,id int)HSplitter{
	return HSplitter{
		MaxSize: Size{Width: 100},
		Children: []Widget{
			TextLabel{Text: name,MaxSize: Size{50,20}},
			CustomWidget{AssignTo: &pios[id],Enabled:true,MaxSize: Size{20,20},Background:
			SolidColorBrush{Color:  walk.RGB(0, 0, 0)}},
		},
	}
}
//PIO显示台
func pio()Widget{
return VSplitter{
	StretchFactor:0,
	MaxSize: Size{80,40},
	Children: []Widget{
		NewLed("P0_1:",0),
		NewLed("P0_2:",1),
		NewLed("P0_3:",2),
		NewLed("P0_4:",3),
	},
}
}