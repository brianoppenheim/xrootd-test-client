package main
import (
	"flag"
	"fmt"
	"net"
)
func main() {
	//Extract the specified flag (stored in addr pointer)
	addr := flag.String("addr", "0.0.0.0:9001", "Address to dial over TCP")
	flag.Parse()
	fmt.Println(*addr)
	println("xrootd-test-client")
	//Dial over address over tcp
	conn, err := net.Dial("tcp", *addr)
	fmt.Println(conn)
	fmt.Println(err)
}