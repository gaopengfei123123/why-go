package main

/* 下面是 C 代码 */

// int add(int a, int b) {
//     return a + b;
// }
import "C"

import "fmt"

func main() {
	a := C.int(1)
	b := C.int(2)
	value := C.add(a, b)
	fmt.Printf("%v\n", value)
}
