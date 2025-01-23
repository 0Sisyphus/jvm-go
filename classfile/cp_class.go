package classfile

// ConstantClassInfo 结构体表示类常量信息
type ConstantClassInfo struct {
	cp        ConstantPool // cp 字段表示常量池
	nameIndex uint16       // nameIndex 字段表示类名在常量池中的索引
}

// readInfo 方法从 ClassReader 中读取类常量信息
func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16() // 读取2个字节并转换为uint16，存储类名索引
}

// Name 方法返回类常量的名称
func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex) // 通过索引从常量池获取对应的UTF-8字符串
}
