package classfile

import "fmt"

type UnparsedAttribute struct {
	attrName string
	attrLen  uint32
	info     []byte
}

func (self *UnparsedAttribute) readInfo(reader *ClassReader) {
	// 无法解析方法体为空
	fmt.Printf("UnparsedAttribute attrName: %v\n", self.attrName)
	fmt.Printf("UnparsedAttribute attrLen: %v\n", self.attrLen)
	// 读取属性值
	self.info = reader.readBytes(self.attrLen)
	fmt.Printf("UnparsedAttribute info: %v\n", self.info)
}
