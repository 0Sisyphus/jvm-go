package math

import (
	"jvm-go/instructions/base"
	"jvm-go/rtda"
)

type IINC struct {
	Index uint
	Const int32
}

func (self *IINC) FetchOperands(reader *base.ByteCodeReader) {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	val += self.Const
	localVars.SetInt(self.Index, val)
}
