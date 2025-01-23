package classfile

// ConstantMemberRefInfo 结构体表示成员引用常量信息，包括字段引用和方法引用
type ConstantMemberRefInfo struct {
	cp               ConstantPool // cp 字段表示常量池
	classIndex       uint16       // classIndex 字段表示类在常量池中的索引
	nameAndTypeIndex uint16       // nameAndTypeIndex 字段表示名称和类型在常量池中的索引
}

// ConstantInterfaceMethodRefInfo 结构体表示接口方法引用常量信息
type ConstantInterfaceMethodRefInfo struct {
	ConstantMemberRefInfo // 嵌入 ConstantMemberRefInfo 结构体
}

// ConstantMethodRefInfo 结构体表示方法引用常量信息
type ConstantMethodRefInfo struct {
	ConstantMemberRefInfo // 嵌入 ConstantMemberRefInfo 结构体
}

// ConstantFieldRefInfo 结构体表示字段引用常量信息
type ConstantFieldRefInfo struct {
	ConstantMemberRefInfo // 嵌入 ConstantMemberRefInfo 结构体
}

// readInfo 方法从 ClassReader 中读取成员引用常量信息
func (self *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()       // 读取2个字节并转换为uint16，存储类索引
	self.nameAndTypeIndex = reader.readUint16() // 读取2个字节并转换为uint16，存储名称和类型索引
}

// ClassName 方法返回成员引用常量的类名
func (self *ConstantMemberRefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex) // 通过索引从常量池获取类名
}

// NameAndDescriptor 方法返回成员引用常量的名称和描述符
func (self *ConstantMemberRefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex) // 通过索引从常量池获取名称和类型
}
