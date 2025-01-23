package classfile

import "encoding/binary"

// ClassReader 是一个用于读取Java字节码文件的结构体
type ClassReader struct {
	// data 是一个字节数组，包含要读取的字节码数据
	data []byte
}

// readUint8 读取一个无符号8位整数（uint8）
func (self *ClassReader) readUint8() uint8 {
	// 读取第一个字节并将其转换为uint8类型
	val := self.data[0]
	// 更新data字段，去掉已读取的字节
	self.data = self.data[1:]
	return val
}

// readUint16 读取一个无符号16位整数（uint16）
func (self *ClassReader) readUint16() uint16 {
	// 使用大端序读取两个字节并将其转换为uint16类型
	val := binary.BigEndian.Uint16(self.data)
	// 更新data字段，去掉已读取的两个字节
	self.data = self.data[2:]
	return val
}

// readUint32 读取一个无符号32位整数（uint32）
func (self *ClassReader) readUint32() uint32 {
	// 使用大端序读取四个字节并将其转换为uint32类型
	val := binary.BigEndian.Uint32(self.data)
	// 更新data字段，去掉已读取的四个字节
	self.data = self.data[4:]
	return val
}

// readUint64 读取一个无符号64位整数（uint64）
func (self *ClassReader) readUint64() uint64 {
	// 使用大端序读取八个字节并将其转换为uint64类型
	val := binary.BigEndian.Uint64(self.data)
	// 更新data字段，去掉已读取的八个字节
	self.data = self.data[8:]
	return val
}

// readUint16s 读取一个uint16表
func (self *ClassReader) readUint16s() []uint16 {
	// 先读取表的长度
	n := self.readUint16()
	// 创建一个长度为n的uint16切片
	s := make([]uint16, n)
	// 循环读取n个uint16值并存入切片
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

// readBytes 读取指定长度的字节
func (self *ClassReader) readBytes(length uint32) []byte {
	// 读取指定长度的字节并返回
	bytes := self.data[:length]
	// 更新data字段，去掉已读取的字节
	self.data = self.data[length:]
	return bytes
}
