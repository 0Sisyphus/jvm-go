package classfile

import (
	"fmt"
	"unicode/utf16"
)

// ConstantUtf8Info 结构体表示 UTF-8 编码的字符串常量
type ConstantUtf8Info struct {
	// str 字段存储解码后的字符串
	str string
}

// readInfo 方法从 ClassReader 中读取 UTF-8 编码的字符串信息
func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	// 读取字符串长度
	length := uint32(reader.readUint16())
	// 读取指定长度的字节数组
	bytes := reader.readBytes(length)
	// 使用 decodeMUTF8 函数将字节数组解码为字符串
	self.str = decodeMUTF8(bytes)
}

// decodeMUTF8 函数将 UTF-8 编码的字节数组解码为字符串
func decodeMUTF8(bytearr []byte) string {
	// utflen 存储字节数组的长度
	utflen := len(bytearr)
	// chararr 用于存储解码后的 UTF-16 编码字符
	chararr := make([]uint16, utflen)

	var c, char2, char3 uint16
	// count 用于遍历字节数组
	count := 0
	// chararr_count 用于记录解码后的字符数量
	chararr_count := 0

	// 遍历字节数组，处理单字节字符
	for count < utflen {
		c = uint16(bytearr[count])
		if c > 127 {
			break
		}
		count++
		chararr[chararr_count] = c
		chararr_count++
	}

	// 处理多字节字符
	for count < utflen {
		c = uint16(bytearr[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			chararr[chararr_count] = c
			chararr_count++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			chararr[chararr_count] = c&0x1F<<6 | char2&0x3F
			chararr_count++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-2])
			char3 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count-1))
			}
			chararr[chararr_count] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			chararr_count++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	// The number of chars produced may be less than utflen
	chararr = chararr[0:chararr_count]
	// 将 UTF-16 编码字符转换为 runes
	runes := utf16.Decode(chararr)
	// 返回解码后的字符串
	return string(runes)
}
