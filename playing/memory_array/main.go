package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// ints
	array := []uint8{1, 2, 3, 4, 5, 6}
	fmt.Println(
		uintptr(unsafe.Pointer(&array[0])), uintptr(unsafe.Pointer(&array[1])), 
		// uintptr(unsafe.Pointer(&array[2]))- 
		// uintptr(unsafe.Pointer(&array[3])),
	)

	// strings
	arrayString := []string{"1", "2", "3", "4", "5", "6"}
	fmt.Println(
		uintptr(unsafe.Pointer(&arrayString[0])), 
		uintptr(unsafe.Pointer(&arrayString[1])), 
		uintptr(unsafe.Pointer(&arrayString[2])), 
		uintptr(unsafe.Pointer(&arrayString[3])),
	)
}
