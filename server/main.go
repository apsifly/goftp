package main

import (
	"bufio"
	"flag"
	"fmt"
	"ftp/protocol"
	"ftp/protocol/osdependent"
	"io"
	"log"
	"net"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	readConfig(*configPath)

	// Listen on TCP port 2000 on all interfaces.
	l, err := net.Listen("tcp", config.Server.Host+":"+
		strconv.Itoa(config.Server.Port))
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
	s := protocol.State{
		TransferType: osdependent.FtpType,
		LocalIP:      config.Server.Host,
	}

	io.WriteString(c, "220 (goFTPD 0.1)\r\n")
mainloop:
	for {
		log.Println("starting select")
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
				log.Printf("executing command %T, %v\n", comm, comm)
				go allowAndExecute(comm, &s, responseCh)
			}
			//c.Write([]byte(recvStr))
		case r := <-responseCh:
			switch r.Base {
			case protocol.Response221:
				r.Send(c)
				break mainloop
			default:
				r.Send(c)
			}
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
	bufreader := bufio.NewReader(c)
	fullstr := ""
	for {
		str, err := bufreader.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}
		if re1.FindString(fullstr+str) != "" { //protection from strings that do not contain "\r"
			ch <- fullstr + str
			fullstr = ""
		} else {
			fullstr += str
		}
	}
	close(ch)
}

func allowAndExecute(c protocol.Command, s *protocol.State, ch chan *protocol.Response) {
	switch c := c.(type) {
	case *protocol.PassCmd:
		allowAndExecutePassCmd(c, s, ch)
	case *protocol.RetrieveCmd:
		user, ok := config.Users[s.User]
		matched := false
		if ok {
			for _, readPath := range user.ReadPath {
				if readPath == c.Path || strings.HasPrefix(c.Path, path.Clean(readPath+"/")) {
					matched = true
				}
			}
		}
		if s.Logged && matched { //TODO: better access denied message
			c.Execute(s, ch)
		} else {
			ch <- protocol.NewResponse(protocol.Response530, "", fmt.Errorf("user not logged in"))
		}
	case *protocol.ListCmd: //TODO: remove duplicate code
		user, ok := config.Users[s.User]
		matched := false
		if ok {
			for _, readPath := range user.ReadPath {
				if readPath == c.Path || strings.HasPrefix(c.Path, path.Clean(readPath+"/")) {
					matched = true
				}
			}
		}
		if s.Logged && matched {
			c.Execute(s, ch)
		} else {
			ch <- protocol.NewResponse(protocol.Response530, "", fmt.Errorf("user not logged in"))
		}
	case *protocol.StoreCmd: //TODO: remove duplicate code
		user, ok := config.Users[s.User]
		matched := false
		if ok {
			for _, writePath := range user.WritePath {
				if writePath == c.Path || strings.HasPrefix(c.Path, path.Clean(writePath+"/")) {
					matched = true
				}
			}
		}
		if s.Logged && matched {
			c.Execute(s, ch)
		} else {
			ch <- protocol.NewResponse(protocol.Response530, "", fmt.Errorf("user not logged in"))
		}
	case *protocol.UserCmd:
		if !s.Logged {
			c.Execute(s, ch)
		} else {
			ch <- protocol.NewResponse(protocol.Response503, "", fmt.Errorf("login reqiest from already logged user, state: %v", s))
		}
	default:
		if s.Logged {
			c.Execute(s, ch)
		} else {
			ch <- protocol.NewResponse(protocol.Response530, "", fmt.Errorf("user not logged in"))
		}
	}

}

func allowAndExecutePassCmd(c protocol.Command, s *protocol.State, ch chan *protocol.Response) {
	c.Execute(s, ch)
	log.Println("pass", s)
	user, ok := config.Users[s.User]

	if ok && user.Password == s.Pass && !s.Logged {
		s.Logged = true
		s.Pass = ""
		ch <- protocol.NewResponse(protocol.Response230, "", fmt.Errorf("login succesful"))
	} else if !s.Logged {
		ch <- protocol.NewResponse(protocol.Response530, "", fmt.Errorf("login error"))
	} else {
		s.Pass = ""
		ch <- protocol.NewResponse(protocol.Response503, "", fmt.Errorf("login reqiest from already logged user, state: %v", s))
	}
}
