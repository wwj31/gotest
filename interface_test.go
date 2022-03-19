package escape

import (
	"testing"
)

// 证明接口函数的参数，如果是指针，会直接导致指针数据逃逸到堆

type iface interface {
	Fn(*int) int
}
type MyStruct struct {
}

func (*MyStruct) Fn(v *int) int {
	return *v
}
func TestInterface(t *testing.T) {
	var (
		i   iface
		st  = &MyStruct{}
		st2 = &MyStruct{}
		i1  = 1
		i2  = 2
	)

	st.Fn(&i1)
	i = st2
	i.Fn(&i2)
}
