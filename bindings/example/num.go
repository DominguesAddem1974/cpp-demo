package num

// #cgo windows CFLAGS: -DCGO_OS_WINDOWS=1
// #cgo darwin CFLAGS: -DCGO_OS_DARWIN=1
// #cgo linux CFLAGS: -DCGO_OS_LINUX=1
// #cgo LDFLAGS: -L${SRCDIR} -lstdc++
// #cgo CXXFLAGS: -std=c++17 -I${SRCDIR}
// #include "num.h"
import "C"
import "unsafe"

type GoNum struct {
	num C.Num
}

func New() GoNum {
	var ret GoNum
	ret.num = C.NumInit()
	return ret
}
func (n GoNum) Free() {
	C.NumFree((C.Num)(unsafe.Pointer(n.num)))
}
func (n GoNum) Inc() {
	C.NumIncrement((C.Num)(unsafe.Pointer(n.num)))
}
func (n GoNum) GetValue() int {
	return int(C.NumGetValue((C.Num)(unsafe.Pointer(n.num))))
}
