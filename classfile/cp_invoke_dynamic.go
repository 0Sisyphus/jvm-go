package classfile

// ConstantInvokeDynamicInfo 表示常量池中的 InvokeDynamic 信息
type ConstantInvokeDynamicInfo struct {
	// bootstrapMethodAttrIndex 引导方法属性的索引
	bootstrapMethodAttrIndex uint16
	// nameAndTypeIndex 名称和类型的索引
	nameAndTypeIndex uint16
}

// readInfo 从 ClassReader 中读取 InvokeDynamic 信息
func (self *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	self.bootstrapMethodAttrIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

// ConstantMethodHandleInfo 表示常量池中的 MethodHandle 信息
type ConstantMethodHandleInfo struct {
	// referenceKind 引用种类
	referenceKind uint8
	// referenceIndex 引用索引
	referenceIndex uint16
}

// readInfo 从 ClassReader 中读取 MethodHandle 信息
func (self *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	self.referenceKind = reader.readUint8()
	self.referenceIndex = reader.readUint16()
}

// ConstantMethodTypeInfo 表示常量池中的 MethodType 信息
type ConstantMethodTypeInfo struct {
	// descriptorIndex 描述符索引
	descriptorIndex uint16
}

// readInfo 从 ClassReader 中读取 MethodType 信息
func (self *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	self.descriptorIndex = reader.readUint16()
}
