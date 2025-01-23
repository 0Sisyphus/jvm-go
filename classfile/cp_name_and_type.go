package classfile

// ConstantNameAndTypeInfo 结构体表示名称和类型常量信息
// 名称和类型常量用于表示字段和方法的名称和描述符。
// 例如，对于一个方法 int add(int, int)，其名称为 "add"，描述符为 "(II)I"。
// 描述符的格式如下：
// - 基本类型：byte(B)、short(S)、char(C)、int(I)、long(J)、float(F) 和 double(D)
// - 引用类型：类或接口类型使用 L<classname>; 表示，例如 Ljava/lang/String;
// - 数组类型：使用 [ 表示，例如 [I 表示 int[]，[[I 表示 int[][]
// - 方法描述符：使用 (参数描述符)返回值描述符 表示，例如 (II)I 表示接受两个 int 参数并返回一个 int 的方法; 其中void返回值由单个字母V表示
type ConstantNameAndTypeInfo struct {
	// nameIndex 字段表示名称在常量池中的索引
	nameIndex uint16
	// descriptorIndex 字段表示类型描述符在常量池中的索引
	descriptorIndex uint16
}

// readInfo 方法从 ClassReader 中读取名称和类型常量信息
// 读取两个 uint16 值，分别表示名称索引和类型描述符索引
func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	// 读取2个字节并转换为uint16，存储名称索引
	self.nameIndex = reader.readUint16()
	// 读取2个字节并转换为uint16，存储类型描述符索引
	self.descriptorIndex = reader.readUint16()
}
