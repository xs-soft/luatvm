package vmui

import lua "github.com/yuin/gopher-lua"

func GetValue(L *lua.LState)*lua.LTable {
	ret:=L.NewTable()
	var exports = map[string]lua.LGFunction{
		"PioSet": pioSet,
	}
	L.SetFuncs(ret, exports)
	return nil
}

func pioSet(L *lua.LState)int{
	val:=L.CheckInt(1)
	for i:=2;i<=L.GetTop();i++{
		PioSet(L.CheckInt(i),val)
	}
	return 0
}