package base

type ByteCodeReader struct {
	code []byte
	pc   int
}

func (self *ByteCodeReader) PC() int {
	return self.pc
}

func (self *ByteCodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

func (self *ByteCodeReader) ReadUint8() uint8 {
	val := self.code[self.pc]
	self.pc++
	return val
}

func (self *ByteCodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

func (self *ByteCodeReader) ReadUint16() uint16 {
	u1 := uint16(self.ReadUint8())
	u2 := uint16(self.ReadUint8())
	return (u1 << 8) | u2
}

func (self *ByteCodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}

func (self *ByteCodeReader) ReadUint32() uint32 {
	u1 := uint32(self.ReadUint16())
	u2 := uint32(self.ReadUint16())
	u3 := uint32(self.ReadUint16())
	u4 := uint32(self.ReadUint16())
	return (u1 << 24) | (u2 << 16) | (u3 << 8) | u4
}

func (self *ByteCodeReader) ReadInt32() int32 {
	byte1 := int32(self.ReadUint8())
	byte2 := int32(self.ReadUint8())
	byte3 := int32(self.ReadUint8())
	byte4 := int32(self.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

// used by lookupswitch and tableswitch
func (self *ByteCodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = self.ReadInt32()
	}
	return ints
}

// used by lookupswitch and tableswitch
func (self *ByteCodeReader) SkipPadding() {
	for self.pc%4 != 0 {
		self.ReadUint8()
	}
}
