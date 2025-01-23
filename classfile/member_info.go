package classfile

// MemberInfo 结构体表示类或接口的字段或方法的信息
type MemberInfo struct {
	cp              ConstantPool    // 常量池
	accessFlags     uint16          // 访问标志
	nameIndex       uint16          // 名称索引
	descriptorIndex uint16          // 描述符索引
	attributes      []AttributeInfo // 属性表
}

// readMembers 函数读取多个成员信息
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

// readMember 函数读取单个成员信息
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

// AccessFlags 方法返回访问标志
func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}

// Name 方法返回成员名称
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

// Descriptor 方法返回成员描述符
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}
