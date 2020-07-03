package vmui
import (
	"github.com/lxn/walk"
)
import . "github.com/lxn/walk/declarative"
func lcd()Widget{
	return ImageView{
		AssignTo:            &LcdView,
	}
}
func init(){

}
var PaintWidget *walk.CustomWidget
var LcdView=&walk.ImageView{}
