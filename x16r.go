package x16r

// #cgo LDFLAGS: libx16r_hash.a
// #include "x16r.h"
import "C"
import "fmt"

func Sum(data []byte) []byte {
	var res [32]C.char
	C.x16r_hash(C.CString(string(data)), &res[0])
	return []byte(C.GoStringN(&res[0], 32))
}

func main() {
	fmt.Println(Sum([]byte("test")))
}
