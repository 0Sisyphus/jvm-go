package classfile

// 常量类型定义
const (
	CONSTANT_Class              = 7  // 类常量
	CONSTANT_Fieldref           = 9  // 字段引用常量
	CONSTANT_Methodref          = 10 // 方法引用常量
	CONSTANT_InterfaceMethodref = 11 // 接口方法引用常量
	CONSTANT_String             = 8  // 字符串常量
	CONSTANT_Integer            = 3  // 整数常量
	CONSTANT_Float              = 4  // 浮点数常量
	CONSTANT_Long               = 5  // 长整数常量
	CONSTANT_Double             = 6  // 双精度浮点数常量
	CONSTANT_NameAndType        = 12 // 名称和类型常量
	CONSTANT_Utf8               = 1  // UTF-8 编码的字符串常量
	CONSTANT_MethodHandle       = 15 // 方法句柄常量
	CONSTANT_MethodType         = 16 // 方法类型常量
	CONSTANT_InvokeDynamic      = 18 // 动态调用常量
)

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

// 读取常量信息
func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

// 根据标签创建新的常量信息
func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_String:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantMemberRefInfo{cp: cp}
	case CONSTANT_Methodref:
		return &ConstantMemberRefInfo{cp: cp}
	case CONSTANT_InterfaceMethodref:
		return &ConstantMemberRefInfo{cp: cp}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}
