package main
//#include <stdio.h>
// #cgo CFLAGS: -I${SRCDIR}
// #cgo LDFLAGS: -lstdc++ -L${SRCDIR} -lbuffer
import "C"
import "unsafe"
func main() {
    buf := NewMyBuffer(1024)
    defer buf.Delete()
    copy(buf.Data(), []byte("hello\x00"))
    C.puts((*C.char)(unsafe.Pointer(&(buf.Data()[0]))))
}
