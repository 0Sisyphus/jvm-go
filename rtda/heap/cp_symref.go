package heap

type SymRef struct {
	// cp字段存放符号引用所在的运行时常量池指针，
	// 这样就可以通过符号引用访问到运行时常量池，进一步又可以访问到类数据
	cp        *ConstantPool
	className string
	class     *Class
}
