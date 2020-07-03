package vmui

import (
	"flag"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"os"
	"sync"
)
var mw *walk.MainWindow
var useUi = flag.Bool("ui", true, "是否使用UI界面")
func IsUse()bool{
	return *useUi
}
func StartUi(){
	if !*useUi{
		return
	}
	mw = new(walk.MainWindow)
	//Background: SolidColorBrush{Color: walk.RGB(0, 0, 0)},
	//var inTE *walk.TextEdit
	MW:=MainWindow{
		AssignTo: &mw,
		Title:   "LuatVM                 (QQ104978)",
		Size:Size{550,400},
		//横向分裂
		Layout:  HBox{SpacingZero: false,MarginsZero: false,Alignment: AlignHNearVNear},
		//左右切分
		Children: []Widget{
			//垂直分裂
			VSplitter{
				//Layout: Grid{Rows: 3,MarginsZero: true},
				MaxSize: Size{240,600},
				MinSize: Size{240,600},

				Children: []Widget{
					//一个横向分裂
					//控件组
					GroupBox{
						Title: "基本设置",
						Layout: Grid{Rows: 1},
						Children: []Widget{
							setui(),
						},
					},
					//一个横向分裂
					//控件组
					GroupBox{
						MaxSize: Size{0,140},
						Title: "PIN信号",
						Layout: Grid{Columns: 1},
						Children: []Widget{
							pio(),
						},
					},

					//控件组
					GroupBox{
						MaxSize: Size{240,80},
						Title: "UART 设置",
						Layout: Grid{Columns: 1},
						Children: []Widget{
							uart(),
						},
					},
				},
			},
			lcd(),
		},
	}
var mu sync.WaitGroup
var onec sync.Once
go func (){
	mu.Add(1)
	MW.OnBoundsChanged=func(){
		onec.Do(func(){
			mu.Done()
		})
	}
	_,err:=MW.Run()
	if err!=nil{
		panic(err)
	}
	os.Exit(0)
}()
	mu.Wait()
}

/*
	PushButton{
		Text: "SCREAM",
		OnClicked: func() {
			outTE.SetText(strings.ToUpper(inTE.Text()))
		},
	},*/