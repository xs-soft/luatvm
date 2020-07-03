package disp

import (
	"github.com/golang/freetype"
	"golang.org/x/image/font"
	"golua/golua/iconv"
	"image"
	"image/color"
	"io/ioutil"
)

func addLabel(img *image.RGBA, x, y int,color color.RGBA, label string) {

	label,_=iconv.GetEncoding("gb2312").NewDecoder().String(label)
	//col := color.RGBA{200, 100, 0, 255}
	//point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}
	fontBytes, _ := ioutil.ReadFile(`C:\WINDOWS\FONTS\SIMFANG.TTF`)
	f, err := freetype.ParseFont(fontBytes)
	if err!=nil{
		panic(err)
	}
	c := freetype.NewContext()
	c.SetFont(f)
	c.SetDPI(72)
	c.SetFontSize(16)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	//NewUniform
	c.SetSrc( image.NewUniform(color))
	c.SetHinting(font.HintingVertical)
	//fmt.Println(label,x,y)
	//字体高度从底部算起。所以要加字体高度。
	pt := freetype.Pt(x,y+16)
	c.DrawString(label, pt)
}