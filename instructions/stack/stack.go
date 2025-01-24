package stack

import (
	"jvm-go/instructions/base"
	"jvm-go/rtda"
)

type POP struct {
	base.NoOperandsInstruction
}

func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	stack.PopSlot()
}

type POP2 struct {
	base.NoOperandsInstruction
}

func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	stack.PopSlot()
	stack.PopSlot()
}
