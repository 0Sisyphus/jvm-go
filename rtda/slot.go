package rtda

import "jvm-go/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
