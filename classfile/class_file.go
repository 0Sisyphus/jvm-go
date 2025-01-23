package classfile

import "fmt"

// ClassFile 结构体表示一个Java类文件的完整结构
type ClassFile struct {
	// minorVersion 表示类文件的次版本号
	minorVersion uint16
	// majorVersion 表示类文件的主要版本号
	majorVersion uint16
	// constantPool 是常量池，存储类中用到的各种常量
	constantPool ConstantPool
	// accessFlags 表示类或接口的访问标志
	accessFlags uint16
	// thisClass 表示当前类的常量池索引
	thisClass uint16
	// superClass 表示父类的常量池索引
	superClass uint16
	// interfaces 存储接口的常量池索引列表
	interfaces []uint16
	// fields 存储字段信息列表
	fields []*MemberInfo
	// methods 存储方法信息列表
	methods []*MemberInfo
	// attributes 存储类的属性信息列表
	attributes []AttributeInfo
}

// Parse 函数解析字节码数据并返回一个ClassFile对象
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	reader := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(reader)
	return
}

// read 方法读取并解析ClassFile的所有字段
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

// MajorVersion 返回类文件的主要版本号
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

// MinorVersion 返回类文件的次版本号
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

// ConstantPool 返回类文件的常量池
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

// AccessFlags 返回类或接口的访问标志
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

// Fields 返回类的字段信息列表
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

// Methods 返回类的方法信息列表
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

// ClassName 返回类的全限定名
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

// SuperClassName 返回父类的全限定名
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}

// InterfaceNames 返回接口的全限定名列表
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

// readAndCheckMagic 读取并检查魔数是否正确
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

// readAndCheckVersion 读取并检查版本号是否正确
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}
