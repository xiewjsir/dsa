// hello.go
package main

//void SayHello(const char* s);
import "C"

func main() {
	//C.puts(C.CString("Hello, World\n"))
	C.SayHello(C.CString("Hello, World\n"))
}
