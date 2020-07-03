package iconv

import (
	lua "github.com/yuin/gopher-lua"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
)
func GetValue(L *lua.LState)*lua.LTable{

	ret:=L.NewTable()
	var exports = map[string]lua.LGFunction{
		"open": open,

	}
	L.SetFuncs(ret, exports)
	L.SetField(ret, "LDO_VLCD", lua.LNumber(55))
	return ret
}
func GetEncoding(name string)encoding.Encoding{
	switch name{
	case "utf8":
		return unicode.UTF8
	case "ucs2be":
		return unicode.UTF16(unicode.BigEndian, unicode.ExpectBOM)
	case "ucs2":
		return unicode.UTF16(unicode.LittleEndian, unicode.ExpectBOM)
	case "gb2312":
		return simplifiedchinese.GBK
	}
	return nil
}
func open(L *lua.LState)int {

	to:=L.CheckString(1)
	from:=L.CheckString(2)
	f:=func(L *lua.LState)int{
		s:=L.CheckString(2)
		v,_:=GetEncoding(from).NewDecoder().String(s)
		v,_=GetEncoding(to).NewEncoder().String(v)
		L.Push(lua.LString(v))
		return 1
	}

	ret:=L.NewTable()
	L.SetFuncs(ret,map[string]lua.LGFunction{
		"iconv": f,
	})
	L.Push(ret)
	return 1
}