package constants

import (
	"jvm-go/instructions/base"
	"jvm-go/rtda"
)

// BIPUSH bipush指令从操作数中获取一个byte型整数，扩展成int型，然 后推入栈顶
type BIPUSH struct {
	val int8
}

// SIPUSH sipush指令从操作数中获取一个short型整数，扩展成 int型，然后推入栈顶
type SIPUSH struct {
	val int16
}

func (self *BIPUSH) FetchOperands(reader *base.ByteCodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack.PushInt(i)
}

func (self *SIPUSH) FetchOperands(reader *base.ByteCodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack.PushInt(i)
}
