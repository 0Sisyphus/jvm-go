package stack

import (
	"jvm-go/instructions/base"
	"jvm-go/rtda"
)

type Swap struct {
	base.NoOperandsInstruction
}

func (self *Swap) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
