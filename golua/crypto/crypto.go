package crypto

import (
	"crypto/hmac"
	md52 "crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"io"
	"strings"
)

func GetValue(L *lua.LState)*lua.LTable{
	ret:=L.NewTable()
	var exports = map[string]lua.LGFunction{
		"md5": md5,
		"hmac_md5":hmac_md5,
		"base64_encode":base64_encode,
		"hmac_sha1":hmac_sha1,
	}
	L.SetFuncs(ret, exports)
	return ret
}
func md5(L *lua.LState)int{
	ss:=L.CheckString(1)

	L.Push(lua.LString(md5V3(ss)))
	return 1
}
func base64_encode(L *lua.LState)int{
	t:=L.CheckString(1)
	l:=L.CheckInt(2)
	if len(t)>l {
		t=t[:l]
	}
	s:=base64.StdEncoding.EncodeToString([]byte(t))
	L.Push(lua.LString(s))
	return 1
}
func hmac_md5(L *lua.LState)int{
	originstr:=L.CheckString(1)
	len_str:=L.ToInt(2)
	if len(originstr)>len_str{
		originstr=originstr[0:len_str]
	}
	signkey:=L.CheckString(3)
	len_key:=L.ToInt(4)
	if len(signkey)>len_key{
		signkey=signkey[0:len_key]
	}

	h:=hmac.New(md52.New,[]byte(signkey))//.Sum()
	h.Write([]byte(originstr))
	b:=h.Sum(nil)
	L.Push(lua.LString(strings.ToUpper(hex.EncodeToString(b))))
	return 1
}


func hmac_sha1(L *lua.LState)int{
	originstr:=L.CheckString(1)
	len_str:=L.ToInt(2)
	if len(originstr)>len_str{
		originstr=originstr[0:len_str]
	}
	signkey:=L.CheckString(3)
	len_key:=L.ToInt(4)
	if len(signkey)>len_key{
		signkey=signkey[0:len_key]
	}

	h:=hmac.New(sha1.New,[]byte(signkey))//.Sum()
	h.Write([]byte(originstr))
	b:=h.Sum(nil)
	L.Push(lua.LString(strings.ToUpper(hex.EncodeToString(b))))
	return 1
}

func md5V3(str string) string {
	w := md52.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}