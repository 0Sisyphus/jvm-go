package rtda

type Frame struct {
	lower        *Frame
	LocalVars    LocalVars
	OperandStack *OperandStack
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		LocalVars:    newLocalVars(maxLocals),
		OperandStack: newOperandStack(maxStack),
	}
}
