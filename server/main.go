package main

import (
	"bufio"
	"ftp/protocol"
	"log"
	"net"
	"regexp"
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
	var st protocol.State
	var readerCh = make(chan string)
	go reader(readerCh, c)
	responseCh := make(chan *protocol.Response)

	s := protocol.State{}
mainloop:
	for {
		select {
		case recvStr, ok := <-readerCh:
			if !ok {
				log.Println("control connection closed")
				break mainloop
			}
			log.Println(recvStr)
			comm, r := protocol.ParseCommand(recvStr)
			if r != nil {
				r.Send(c)
			} else {
				log.Printf("executing command %t, %v\n", comm, comm)
				go comm.Execute(&s, responseCh)
			}
			c.Write([]byte(recvStr))
		case r := <-responseCh:
			r.Send(c)
		}

	}

	// Shut down the connection.
	c.Close()
	if st.DataConn != nil {
		st.DataConn.Close()
	}
}

func reader(ch chan string, c net.Conn) {
	re1 := regexp.MustCompile("^[^\r]*\r\n")
	bufreader := bufio.NewReader(c) //.ReadString('\n')
	fullstr := ""
	for {
		str, err := bufreader.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}
		if re1.FindString(fullstr+str) != "" {
			ch <- fullstr + str
			fullstr = ""
		} else {
			fullstr += str
		}
	}
	close(ch)
}
