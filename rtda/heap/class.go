package heap

import "jvm-go/classfile"

type Class struct {
	accessFlags       uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint  // 实例变量占据的空间大小
	staticSlotCount   uint  // 类变量占据空间大小
	staticVars        Slots // 存放静态变量
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	//class.methods = newMethods(class, cf.Methods())
	return class
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}

func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}

func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}

func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}

func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}

func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}

func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}

func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}
