package classfile

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct {
	AttributeInfo
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
