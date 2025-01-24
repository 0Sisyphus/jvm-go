package constants

import (
	"jvm-go/instructions/base"
	"jvm-go/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// do nothing
}
