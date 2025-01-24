package base

import "jvm-go/rtda"

type Instruction interface {
	FetchOperands(reader *ByteCodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {
}

func (self *NoOperandsInstruction) FetchOperands(reader *ByteCodeReader) {
	// nothing to do
}

type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *ByteCodeReader) {
	self.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *ByteCodeReader) {
	self.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *ByteCodeReader) {
	self.Index = uint(reader.ReadUint16())
}
