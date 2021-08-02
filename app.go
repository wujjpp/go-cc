package main

import "fmt"
// #cgo CFLAGS: -I${SRCDIR}/lib
// #cgo LDFLAGS: -lstdc++ -L${SRCDIR}/lib -lhash64
// #include "bridge.h"
import "C"

func main() {
	var s string = "f1bc319830aa4bd525b6b211fc128e5e"
	// [
  // 	"-6331077761356739522",
  // 	"6750279366282124699"
	// ]
	res1 := C.hash64(C.CString(s), C.int(len(s)))
	fmt.Println(res1)

	s = "5608c68f0f84a28a5c8c0479ebbd69c1"
	// [
  // 	"8004344832703457518",
  // 	"4306029215693896848"
	// ]
	res1 = C.hash64(C.CString(s), C.int(len(s)))
	fmt.Println(res1)
}
