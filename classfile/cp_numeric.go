package classfile

import "math"

// ConstantIntegerInfo 结构体表示整数常量
type ConstantIntegerInfo struct {
	// val 字段存储整数值
	val int32
}

// ConstantFloatInfo 结构体表示浮点数常量
type ConstantFloatInfo struct {
	// val 字段存储浮点数值
	val float32
}

// ConstantLongInfo 结构体表示长整数常量
type ConstantLongInfo struct {
	// val 字段存储长整数值
	val int64
}

// ConstantDoubleInfo 结构体表示双精度浮点数常量
type ConstantDoubleInfo struct {
	// val 字段存储双精度浮点数值
	val float64
}

// readInfo 方法从 ClassReader 中读取整数信息
func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	// 读取4个字节并转换为int32
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

// readInfo 方法从 ClassReader 中读取浮点数信息
func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	// 读取4个字节并转换为float32
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}

// readInfo 方法从 ClassReader 中读取长整数信息
func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	// 读取8个字节并转换为int64
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

// readInfo 方法从 ClassReader 中读取双精度浮点数信息
func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	// 读取8个字节并转换为float64
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}
