package type_conversion

/*
 #include <stdlib.h>
#include <stdio.h>
#include <string.h>

int val_int = 1;
int *p_v_int = &val_int;
*/
import "C"
import (
	"fmt"
	"math"
	"unsafe"
)

// 指针类型的转化
func GetSetPointerTypeVariables() {
	// ------ get value ------
	var pVInt = uintptr(unsafe.Pointer(C.p_v_int))

	// print: 0x525c64 == 0x525c64 == 0x525c64
	fmt.Println(fmt.Sprintf("0x%x", pVInt),
		"==", unsafe.Pointer(C.p_v_int), "==",
		C.p_v_int)

	// ------ set value ------
	// go way :
	*(*int32)(unsafe.Pointer(pVInt)) = 10
	// print: 10
	fmt.Println(*(*int32)(unsafe.Pointer(pVInt)))

	// cgo way :
	*C.p_v_int = math.MaxInt32
	// print: MaxInt32
	fmt.Println(*C.p_v_int)
}
