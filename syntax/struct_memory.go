package main

import (
	"fmt"
	"unsafe"
)

/*
	struct的内存布局
	————《go语言学习笔记》雨痕 P126 ~ 129

	不管结构体包含多少字段，其内存总是一次性分配的，各字段在相邻的地址空间按定义顺序排列。
	当然，对于引用类型、字符串和指针，结构内存中只包含其基本（头部）数据。还有，所有匿名字段成员也被包含在内。
	借助unsafe包中的相关函数，可输出所有字段的偏移量和长度。
	unsafe.Sizeof返回值的单位是字节（byte）
*/

type point struct {
	x, y int
}

type value struct {
	id   int
	name string
	data []byte
	next *value
	point
}

func main() {
	v := value{
		id:    1,
		name:  "xinmin",
		data:  []byte{1, 2, 3, 4},
		point: point{x: 100, y: 200},
	}

	s := `
	    v: %p ~ %x, size: %d, align: %d

		field    address             offset       size
		-------+---------------+---------------+-----------
		id       %p             %d       %d
		name     %p             %d       %d
		data     %p             %d       %d
		next     %p             %d       %d
		x        %p             %d       %d
		y        %p             %d       %d
	`
	fmt.Printf(s,
		&v, uintptr(unsafe.Pointer(&v))+unsafe.Sizeof(v), unsafe.Sizeof(v), unsafe.Alignof(v),
		&v.id, unsafe.Offsetof(v.id), unsafe.Sizeof(v.id),
		&v.name, unsafe.Offsetof(v.name), unsafe.Sizeof(v.name),
		&v.data, unsafe.Offsetof(v.data), unsafe.Sizeof(v.data),
		&v.next, unsafe.Offsetof(v.next), unsafe.Sizeof(v.next),
		&v.x, unsafe.Offsetof(v.x), unsafe.Sizeof(v.x),
		&v.y, unsafe.Offsetof(v.y), unsafe.Sizeof(v.y),
	)

	/*
		field    address             offset       size
		-------+---------------+---------------+-----------
		id       0x14000100050             0       8
		name     0x14000100058             8       16
		data     0x14000100068             24       24
		next     0x14000100080             48       8
		x        0x14000100088             56       8
		y        0x14000100090             64       8
	*/

	// 以上内存地址对应的10进制数为：
	// 1374390583376
	// 1374390583384
	// 1374390583400
	// 1374390583424
	// 1374390583432
	// 1374390583440

	/* 字节占用解释：
	+----+----------+----------+----------+----------+----------+------+---------+---------+
	| id | name.ptr | name.len | data.ptr | data.len | data.cap | next | point.x | point.y |
	+----+----------+----------+----------+----------+----------+------+---------+---------+
	0    8          16         24         32         40         48     56        64        72
	*/

	// 在分配内存时，字段须做对齐处理，通常以所有字段中最长的基础类型宽度为标准。
	// 为什么要做字段对齐处理？
	// 是为了提高内存访问效率。计算机内存是以字节为基本单位的，但是许多计算机体系结构要求数据在内存中的地址必须是特定大小的倍数，例如4字节或8字节。这种对齐要求是为了提高内存访问的速度和效率，因为许多计算机架构在访问未对齐的数据时会导致性能下降，甚至在某些体系结构上会触发硬件异常。
	// 当分配内存给结构体时，编译器会自动进行字段对齐，以确保结构体的每个字段都按照正确的字节对齐要求排列在内存中，从而最大化内存访问效率。
	// 总之，字段对齐是为了优化内存访问，提高计算机程序的性能。
	v1 := struct {
		a byte
		b byte
		c int32 // 对齐宽度4
	}{}
	v2 := struct {
		a byte
		b byte // 对齐宽度1
	}{}
	v3 := struct {
		a byte
		b []int // 基础类型int，对齐宽度8
		c byte
	}{}
	fmt.Printf("v1: %d, %d\n", unsafe.Alignof(v1), unsafe.Sizeof(v1)) // v1: 4, 8
	fmt.Printf("v2: %d, %d\n", unsafe.Alignof(v2), unsafe.Sizeof(v2)) // v2: 1, 2
	fmt.Printf("v3: %d, %d\n", unsafe.Alignof(v3), unsafe.Sizeof(v3)) // v3: 8, 40

	// 比较特殊的是空结构体类型字段。如果它是最后一个字段，那么编译器将其当作长度为1的类型做对齐处理，以便其地址不会越界，避免引发垃圾回收错误。
	v4 := struct {
		a struct{}
		b int
		c struct{}
	}{}
	s4 := `
		v: %p ~ %x, size: %d, align: %d
		field   address         offset    size
		------+---------------+---------+---------
		a       %p        %d        %d
		b       %p        %d        %d
		c       %p        %d        %d
	`
	fmt.Printf(s4,
		&v4, uintptr(unsafe.Pointer(&v4))+unsafe.Sizeof(v4), unsafe.Sizeof(v4), unsafe.Alignof(v4),
		&v4.a, unsafe.Offsetof(v4.a), unsafe.Sizeof(v4.a),
		&v4.b, unsafe.Offsetof(v4.b), unsafe.Sizeof(v4.b),
		&v4.c, unsafe.Offsetof(v4.c), unsafe.Sizeof(v4.c),
	)
	/*
		v: 0x14000124020 ~ 14000124030, size: 16, align: 8
		field   address         offset    size
		------+---------------+---------+---------
		a       0x14000124020        0        0
		b       0x14000124020        0        8
		c       0x14000124028        8        0
	*/

	// 如果仅有一个空结构字段，那么同样按1对齐，只不过长度为0，且指向runtime.zerobase变量。
	v5 := struct {
		a struct{} // 对齐宽度1
	}{}
	fmt.Printf("%p, %d, %d\n", &v5, unsafe.Sizeof(v5), unsafe.Alignof(v5)) // 0x104c7f098, 0, 1

}
