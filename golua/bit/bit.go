package bit

import lua "github.com/yuin/gopher-lua"
const bitsPerWord = 32

// BitsPerWord is the implementation-specific size of int and uint in bits.
const BitsPerWord = bitsPerWord // either 32 or 64

// Implementation-specific integer limit values.
const (
	MaxInt  = 1<<(BitsPerWord-1) - 1 // either 1<<31 - 1 or 1<<63 - 1
	MinInt  = -MaxInt - 1            // either -1 << 31 or -1 << 63
	MaxUint = 1<<BitsPerWord - 1     // either 1<<32 - 1 or 1<<64 - 1
)


func GetValue(L *lua.LState)*lua.LTable{
	ret:=L.NewTable()
	var exports = map[string]lua.LGFunction{
		"bit": bit,
		"isset": isset,
		"isclear": isclear,
		"set": set,
		"bnot": bnot,
		"band": band,
		"bor": bor,
		"bxor": bxor,
		"lshift": lshift,
		"rshift":rshift,
		"arshift":arshift,
	}
	L.SetFuncs(ret, exports)
	return ret
}


//左移位计算
func bit(L *lua.LState)int{
	ret:=L.CheckInt(1) << 1
	L.Push(lua.LNumber(ret))
	return 1
}
//测试位数是否被置1
func isset(L *lua.LState)int{
	val:=L.CheckInt(1)
	pos:=L.CheckInt(2)
	//b:=(val & 1<<pos)>>pos
	L.Push(lua.LBool((val & (1<<uint(pos)))>0))
	return 1
}

func isclear(L *lua.LState)int{
	val:=L.CheckInt(1)
	pos:=L.CheckInt(2)
	//b:=(val & 1<<pos)>>pos
	L.Push(lua.LBool((val & (1<<uint(pos)))==0))
	return 1
}
func set(L *lua.LState)int{
	ret:=L.CheckInt(1)
	for i:=2;i<=L.GetTop();i++{
		ret=ret | (1<<uint(L.CheckInt(i)))
	}
	L.Push(lua.LNumber(ret))
	return 1
}
func clear(L *lua.LState){
	ret:=L.CheckInt(1)
	for i:=2;i<=L.GetTop();i++{
		ret=ret &^ (1<<uint(L.CheckInt(i)))
	}
	L.Push(lua.LNumber(ret))
}
func bnot(L *lua.LState)int{
	ret:=L.CheckInt(1)
	ret = ^ret
	L.Push(lua.LNumber(ret))
	return 1
}
func band(L *lua.LState)int{
	ret:=L.CheckInt(1)
	for i:=2;i<=L.GetTop();i++{
		ret=ret & L.CheckInt(i)
	}
	L.Push(lua.LNumber(ret))
	return 1
}
func bor(L *lua.LState)int{
	ret:=L.CheckInt(1)
	for i:=2;i<=L.GetTop();i++{
		ret=ret | L.CheckInt(i)
	}
	L.Push(lua.LNumber(ret))
	return 1
}
func bxor(L *lua.LState)int{
	ret:=L.CheckInt(1)
	for i:=2;i<=L.GetTop();i++{
		ret=ret ^ L.CheckInt(i)
	}
	L.Push(lua.LNumber(ret))
	return 1
}
func lshift(L *lua.LState)int{
	ret:=L.CheckInt(1)
	ret = ret << uint(ret)
	L.Push(lua.LNumber(ret))
	return 1
}
//逻辑右移
func rshift(L *lua.LState)int{
	ret:=L.CheckInt(1)
	pos:=L.CheckInt(2)

	L.Push(lua.LNumber(_rshift(ret,pos)))
	return 1
}
//算数右移
func arshift(L *lua.LState)int{
	ret:=L.CheckInt(1)
	ret = ret >> uint(ret)
	L.Push(lua.LNumber(ret))
	return 1
}
func _rshift(a,b int)int{
	if a>MaxUint{

	}
	//去符号位
	v:=(a & MaxUint)
	v=v>>uint(b)
	return v
}
func dumpbit(val int)string{
	s:=""
	for i:=31;i>=0;i--{
		if (val & (1<<uint(i)))>0{
			s+="1"
		}else{
			s+="0"
		}
	}
	return s
}