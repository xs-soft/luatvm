package iconv

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"testing"
)
func TestHelloWorld(t *testing.T) {
	L := lua.NewState()
	err:=L.DoString(`
d2=7
offset = d2 or 0
print(type(offset))
`)
	fmt.Println(err)
	/*
	fmt.Println("utf8",len("你好LUA"))
	v,_:=GetEncoding("ucs2").NewDecoder().String("你好")
	fmt.Println("utf8",len("你好"))
	v,_=GetEncoding("gb2312").NewDecoder().String(v)

	fmt.Println(v)
*/
	/*转换 utf8 ucs2 欢迎使用Luat ��"kΏO(uL u a t
	ucs2 欢迎使用Luat
	转换 ucs2 gb2312 ��"kΏO(uL u a t  ��ӭʹ��Luat


	 */
	}
