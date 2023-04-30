package main

import "C"

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

func main() {
	b := C.CString("hello world.")
	C.puts(b)
	C.free(unsafe.Pointer(b))
}
