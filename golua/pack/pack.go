package pack

import (
	"bytes"
	"encoding/binary"
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"github.com/zhuangsirui/binpacker"
	"strings"
)

func GetValue(L *lua.LState)*lua.LTable{

	ret:=L.NewTable()
	var exports = map[string]lua.LGFunction{
		"pack": pack,
		"unpack": unpack,
	}
	L.SetFuncs(ret, exports)
	L.Push(ret)
	return ret
}
func pack(L *lua.LState)int{
	f:=L.CheckString(1)
	var p *binpacker.Packer
	buffer := new(bytes.Buffer)
	if f[0:1]==">" || f[0:1]!="<"{
		p=binpacker.NewPacker(binary.BigEndian,buffer)
	}
	if f[0:1]=="<"{
		p=binpacker.NewPacker(binary.LittleEndian,buffer)
	}
	//裁剪前缀
	f=strings.ReplaceAll(f,">","")
	f=strings.ReplaceAll(f,"<","")
	if p == nil {
		panic("处理失败")
	}
	//fmt.Println("参数数量",L.)
	for i:=0;i<len(f);i++{
		switch f[i:i+1] {
			case "A":
				//如果参数为nil则直接忽略
				//if L.CheckAny(i+2).Type()==lua.LTNil{
					//fmt.Println("忽略")
					//	continue
				//}

				p.PushString(L.ToString(i+2))
			case "H":

				p.PushUint16(uint16(L.CheckNumber(i+2)))
			case "b":

				p.PushByte(byte(L.CheckNumber(i+2)))
			case "P"://有争议

				p.PushInt16(int16(len((L.CheckString(i+2)))))
				p.PushString(L.CheckString(i+2))
		default:
			fmt.Println("编码未知类型",f[i:i+1],L.CheckString(1))
			panic("!!!!!!!!")
		}
	}
	L.Push(lua.LString(buffer.Bytes()))
	//fmt.Println("处理PACK",hex.EncodeToString(buffer.Bytes()))
	return 1
}

func unpack(L *lua.LState)int{

	f:=L.CheckString(2)
	pos:=L.CheckInt(3)


	var p *binpacker.Unpacker
	buffer := new(bytes.Buffer)
	byte:=[]byte(L.CheckString(1))
	//fmt.Println("原始数据信息",hex.EncodeToString(byte))
	//if pos>1{
		byte=byte[pos-1:]
	//}
	buffer.Write(byte)

	if f[0:1]==">" || f[0:1]!="<"{
		p=binpacker.NewUnpacker(binary.BigEndian,buffer)
	}
	if f[0:1]=="<"{
		p=binpacker.NewUnpacker(binary.LittleEndian,buffer)
	}
	f=strings.ReplaceAll(f,">","")
	f=strings.ReplaceAll(f,"<","")
	nextpos:=pos
	retlist:=[]lua.LValue{}
	//fmt.Println("参数数量",L.)
	for i:=0;i<len(f);i++{
		switch f[i:i+1] {
		case "H":
			h,err:=p.ShiftInt16()
			if err!=nil{
				panic(err)
			}
			retlist = append(retlist,lua.LNumber(h))
			//2字节
			nextpos+=2
		case "b":
			b,err:=p.ShiftByte()
			if err!=nil{
				panic(err)
			}
			retlist = append(retlist,lua.LNumber(b))
			nextpos+=1
		case "P"://有争议
			l,err:=p.ShiftInt16()
			if err!=nil{
				panic(err)
			}
			s,err:=p.ShiftString(uint64(l))
			if err!=nil{
				panic(err)
			}
			nextpos+=(2+int(l))
			retlist = append(retlist,lua.LString(s))
		default:
			fmt.Println("解码未知类型",f[i:i+1],L.CheckString(2))
			panic("!!!!!!!!")
		}
	}
	L.Push(lua.LNumber(nextpos))
	//fmt.Println("下一字节",lua.LNumber(nextpos))
	for _,v:=range retlist{
		L.Push(v)
		//fmt.Println("值",v.String())
	}
	return len(retlist)+1
	//binpacker.NewUnpacker()
}