package disp

import (
	"bytes"
	"fmt"
	"github.com/lxn/walk"
	"github.com/nfnt/resize"
	lua "github.com/yuin/gopher-lua"
	"golua/golua/vmui"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"path"
)
var RGB888_RED   =   0x00ff0000
var RGB888_GREEN =   0x0000ff00
var RGB888_BLUE  =   0x000000ff
var RGB565_RED   =  0xf800
var RGB565_GREEN =  0x07e0
var RGB565_BLUE  =  0x001f
var colorset     =color.RGBA{0, 0, 0, 255}
var bgcolorset   =color.RGBA{255, 255, 255, 255}
func GetValue(L *lua.LState)*lua.LTable{

	ret:=L.NewTable()
	var exports = map[string]lua.LGFunction{
		"init": _init,
		"puttext":puttext,
		"clear":clear,
		"update":update,
		"putimage":putimage,
		"drawrect":drawrect,
		"setcolor":setcolor,
		"setbkcolor":setbkcolor,
		"loadfont":loadfont,
		"setfont":setfont,
		"sleep":sleep,
		"getlcdinfo":getlcdinfo,
		"putqrcode":putqrcode,
	}

	L.SetFuncs(ret, exports)
	L.SetField(ret,"BUS_SPI4LINE",lua.LNumber(1))
	return ret
}
var img *image.RGBA
var bpp,width,height int
func GetTableInt(t *lua.LTable,name string,dval ...int)int{
	v:=t.RawGetString(name)
	if i,ok:=v.(lua.LNumber);ok{
		return int(i)
	}
	if len(dval)>0 {
		return dval[0]
	}
	panic("未找到参数："+name)
}


func _init(L *lua.LState)int{
	param:=L.CheckTable(1)
	width  =GetTableInt(param,"width")
	height =GetTableInt(param,"height")

	bpp    =GetTableInt(param,"bpp")

	//xoffset:=GetTableInt(param,"xoffset",0) //:=lua.LVAsNumber(param.RawGet(lua.LString("xoffset")))
	//yoffset:=GetTableInt(param,"yoffset") //:=lua.LVAsNumber(param.RawGet(lua.LString("xoffset")))
	if !vmui.IsUse(){
		return 0
	}
	//pinrst :=GetTableInt(param,"pinrst")
	//pincs  :=GetTableInt(param,"pincs")
	//initcmd:=param.RawGetString("initcmd")
	//hwfillcolor:=GetTableInt(param,"hwfillcolor") //填充色
	img = image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img,img.Bounds(),&image.Uniform{color.RGBA{255, 255, 255, 255}}, image.Point{}, draw.Src)

	return 0
}
func puttext(L *lua.LState)int{
	if !vmui.IsUse(){
		return 0
	}

	str:=L.CheckString(1)
	x  :=L.CheckInt(2)
	y  :=L.CheckInt(3)
	addLabel(img,x,y,colorset,str)
	return 0
}
//清空缓冲
func clear(L *lua.LState)int{
	if !vmui.IsUse(){
		return 0
	}

	draw.Draw(img,img.Bounds(),&image.Uniform{bgcolorset}, image.Point{}, draw.Src)
	return 0
}
func update(L *lua.LState)int{
	if !vmui.IsUse(){
		return 0
	}
	b,err:=walk.NewBitmapFromImageForDPI(img,96)
	if err!=nil{
		panic(err)
	}
	if vmui.LcdView!=nil{
		vmui.LcdView.SetImage(b)
	}
	return 0
}


func putimage(L *lua.LState)int{
	if !vmui.IsUse(){
		return 0
	}
	file  :=L.CheckString(1)
		x     :=L.CheckInt(2)
		y     :=L.CheckInt(3)
		left  :=L.ToInt(5)
		top   :=L.ToInt(6)
		right :=L.ToInt(7)
		bottom:=L.ToInt(8)
	fp, err := os.Open(path.Base(file))
	if err!=nil{
		panic("文件加载错误1")
	}
	src,err:=png.Decode(fp)
	if err!=nil{
		panic("文件加载错误2"+err.Error())
	}
	if right==0{
		right=src.Bounds().Size().X
	}
	if bottom==0{
		bottom=src.Bounds().Size().Y
	}

	var subImg image.Image
	//fmt.Println("加载文件错误",err)
	if rgbImg, ok := src.(*image.YCbCr); ok {
		subImg = rgbImg.SubImage(image.Rect(left, top,right, bottom)).(*image.YCbCr) //图片裁剪x0 y0 x1 y1
	} else if rgbImg, ok := src.(*image.RGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(left,top, right, bottom)).(*image.RGBA) //图片裁剪x0 y0 x1 y1
	} else if rgbImg, ok := src.(*image.NRGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(left, top, right, bottom)).(*image.NRGBA) //图片裁剪x0 y0 x1 y1
	} else {
		panic("图片解码失败")
		//return subImg, errors.New("图片解码失败")
	}
	dp:=image.Point{X:x,Y:y }
	r := image.Rectangle{dp, dp.Add(subImg.Bounds().Size())}
	r=r
	//img=subImg.Bounds()
	draw.Draw(img,r,subImg,subImg.Bounds().Min,draw.Src)
	return 0
}

func drawrect(L *lua.LState)int{
	if !vmui.IsUse(){
		return 0
	}

	fmt.Println("绘制方形")
	return 0
}

func setcolor(L *lua.LState)int{
	if !vmui.IsUse(){
		return 0
	}

	L.Push(L.CheckNumber(1))
	c565:=L.CheckInt(1)
	r:= (c565 & RGB565_RED)   >> 8
	g:= (c565 & RGB565_GREEN) >> 3
	b:= (c565 & RGB565_BLUE)  << 3
	colorset     =color.RGBA{uint8(r), uint8(g), uint8(b), 255}
	return 1
}

func setbkcolor(L *lua.LState)int{
	if !vmui.IsUse(){
		return 0
	}

	L.Push(L.CheckNumber(1))
	c565:=L.CheckInt(1)
	r:= (c565 & RGB565_RED)   >> 8
	g:= (c565 & RGB565_GREEN) >> 3
	b:= (c565 & RGB565_BLUE)  << 3
	bgcolorset     =color.RGBA{uint8(r), uint8(g), uint8(b), 255}
	return 0
}
func loadfont(L*lua.LState)int{
	if !vmui.IsUse(){
		return 0
	}

	fmt.Println("读取字体")
	return 0
}
func setfont(L *lua.LState)int{
	if !vmui.IsUse(){
		return 0
	}

	fmt.Println("设置字体")
	return 0
}
//休眠。干啥用的
func sleep(L *lua.LState)int{
	if !vmui.IsUse(){
		return 0
	}

	return 0
}
func getlcdinfo(L *lua.LState)int {
	L.Push(lua.LNumber(width))
	L.Push(lua.LNumber(height))
	L.Push(lua.LNumber(bpp))
	return 3
}
func putqrcode(L *lua.LState)int {
	if !vmui.IsUse(){
		return 0
	}

	//数据
	data:=L.CheckString(1)
	//width:=L.check
	display_width:=L.CheckInt(3)
	x:=L.CheckInt(4)
	y:=L.CheckInt(5)
	qrimg,_:=png.Decode(bytes.NewReader([]byte(data)))
	qrimg=resize.Resize(uint(display_width),uint(display_width),qrimg,resize.Lanczos3)
	//draw2.Scaler()
	//dc:=gg.NewContext(x,y)
	//dc.DrawImage(qrimg,0,0)
	//dc:=gg.NewContextForImage(img)
	//dc.Scale(float64(display_width),float64(display_width))
	//qrimg=dc.Image()
	//dc.SavePNG("1111.png")
	dp:=image.Point{X:x,Y:y }
	r := image.Rectangle{dp, dp.Add(img.Bounds().Size())}
	r=r
	draw.Draw(img,r,qrimg,qrimg.Bounds().Min,draw.Src)
	return 0
}