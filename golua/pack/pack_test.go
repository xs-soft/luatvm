package pack

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/zhuangsirui/binpacker"
	"testing"
)
func TestRunRabbitConsumer(t *testing.T) {
	buffer := new(bytes.Buffer)
	packer := binpacker.NewPacker(binary.BigEndian,buffer)
	packer.PushUint16(0x3234)

	//A 字符串型     =packer.PushString()
	//H 无符号短型.指 =PushUint16
	//H 无符号短型.指 =PushUint16
	//H 无符号短型.指 =PushUint16
	//b 无符号字符型  =packer.PushByte()
	//P 长字符优先    =packer.PushString()
	//P 长字符优先    =packer.PushString()
	//I 无符号整形    =packer.PushUint32()
	//i 整形
	//packer.PushFloat64()
	fmt.Println(hex.EncodeToString(buffer.Bytes()))
	//fmt.Println(fmt.Sprint("%b",0x3234))

}