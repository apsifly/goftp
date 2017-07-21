package main

import (
	"log"
	"net"
)

func main() {
	// Listen on TCP port 2000 on all interfaces.
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go processConn(conn)
	}
}

func processConn(c net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := c.Read(buf)
		if err != nil {
			log.Println(err)
			break
		}
		c.Write(buf[:n])
	}

	// Shut down the connection.
	c.Close()
}
