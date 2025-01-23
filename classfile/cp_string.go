package classfile

// ConstantStringInfo 结构体表示字符串常量
type ConstantStringInfo struct {
	cp          ConstantPool // cp 字段表示常量池
	stringIndex uint16       // stringIndex 字段表示字符串在常量池中的索引
}

// readInfo 方法从 ClassReader 中读取字符串常量信息
func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16() // 读取2个字节并转换为uint16，存储字符串索引
}

// String 方法返回字符串常量的值
func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex) // 通过索引从常量池获取对应的UTF-8字符串
}
