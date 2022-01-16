package demo

/*
#cgo LDFLAGS: -L${SRCDIR}/release/lib -lstdc++ -ldemo_static
#cgo CXXFLAGS: -std=c++14 -I${SRCDIR}
#include "include/function.h"
*/
import "C"

func CAdd(num1 int, num2 int) int {
	num1C := C.int(num1)
	num2C := C.int(num2)
	r := C.Add(num1C, num2C)

	return int(r)
}

func Add(num1 int, num2 int) int {
	return num1 + num2
}
