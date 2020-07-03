package qrencode

import (
	"bytes"
	"github.com/skip2/go-qrcode"
	lua "github.com/yuin/gopher-lua"
	"image/png"
)

func GetValue(L *lua.LState)*lua.LTable{

	ret:=L.NewTable()
	var exports = map[string]lua.LGFunction{
		"encode": encode,
	}

	L.SetFuncs(ret, exports)
	L.SetField(ret,"BUS_SPI4LINE",lua.LNumber(1))
	return ret
}

//清空缓冲
func encode(L *lua.LState)int{
	url:=L.CheckString(1)
	b,_:=qrcode.Encode(url,qrcode.Low, 256)
	image,_:=png.Decode(bytes.NewReader(b))
	//draw.Draw(img,img.Bounds(),&image.Uniform{bgcolorset}, image.Point{}, draw.Src)
	L.Push(lua.LNumber(image.Bounds().Size().Y))
	L.Push(lua.LString(b))
	return 2
}