package typelearning

//const定义的常量不可改变
//iota
const (
	x = iota //从零开始累加
	y
	z
)

const (
	x100 = 100 + iota
	y101
	z102
)

const (
	c = iota << 1 //移位累加
	d
	e
)
