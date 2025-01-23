package classfile

type ConstantPool []ConstantInfo

// readConstantPool 读取类文件中的常量池
func readConstantPool(reader *ClassReader) ConstantPool {
	// 读取常量池计数
	cpCount := int(reader.readUint16())
	// 创建常量池切片
	cp := make([]ConstantInfo, cpCount)
	// 遍历常量池，读取每个常量信息
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		// 如果常量是 long 或 double 类型(占两个位置)，则跳过下一个索引
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

// getConstantInfo 根据索引获取常量信息
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	// 检查索引对应的常量信息是否存在
	if info := self[index]; info != nil {
		return info
	}
	// 如果索引无效，抛出 panic
	panic("Invalid constant pool index!")
}

// getNameAndType 根据索引获取名称和类型信息
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	// 获取 ConstantNameAndTypeInfo 类型的常量信息
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	// 获取名称和类型字符串
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

// getClassName 根据索引获取类名
func (self ConstantPool) getClassName(index uint16) string {
	// 获取 ConstantClassInfo 类型的常量信息
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	// 获取类名字符串
	return self.getUtf8(classInfo.nameIndex)
}

// getUtf8 根据索引获取 UTF-8 字符串
func (self ConstantPool) getUtf8(index uint16) string {
	// 获取 ConstantUtf8Info 类型的常量信息
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	// 返回 UTF-8 字符串
	return utf8Info.str
}
