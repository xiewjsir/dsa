package main

import "net"

func main() {
	listen,err := net.Listen("tcp",":8083")
	if err != nil{
	
	}
	
	for  {
		conn,_ := listen.Accept()
		go func(conn net.Conn) {
		
		}(conn)
	}
	
}
