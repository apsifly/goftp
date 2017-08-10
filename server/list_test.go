package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/apsifly/goftp/protocol"
)

func TestList(t *testing.T) {
	readConfig("goftp.yml")
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	lowestPort := 1000
	var tryPort int
	var l net.Listener
	var err error
	for i := 0; i < 5; i++ {
		tryPort = random.Int()%(65535-lowestPort) + lowestPort
		l, err = net.Listen("tcp", "localhost:"+strconv.Itoa(tryPort))
		if err == nil {
			break
		}
	}
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		clientTryPort := random.Int()%(65535-lowestPort) + lowestPort
		port := protocol.NewPortCmd("127.0.0.1", clientTryPort)
		list := protocol.NewListCmd("")
		s, err := net.Dial("tcp", "127.0.0.1:"+strconv.Itoa(tryPort))
		if err != nil {
			t.Fatal(err)
		}
		defer s.Close()
		checkResp(s, '2', t)
		if _, err := fmt.Fprintf(s, "USER user1\r\n"); err != nil {
			t.Fatal(err)
		}
		checkResp(s, '3', t)
		if _, err := fmt.Fprintf(s, "PASS 123\r\n"); err != nil {
			t.Fatal(err)
		}
		checkResp(s, '2', t)
		l2, err := net.Listen("tcp", "localhost:"+strconv.Itoa(clientTryPort))
		if err != nil {
			t.Fatal(err)
		}
		port.Send(s)
		checkResp(s, '2', t)
		clientConn, err := l2.Accept()
		list.Send(s)
		checkResp(s, '1', t)
		n, err := io.Copy(os.Stderr, clientConn)
		if n < 50 || err != nil {
			t.Fatal(err)
		}
		checkResp(s, '2', t)
		quit := protocol.NewQuitCmd()
		quit.Send(s)
		checkResp(s, '2', t)

	}()
	if err != nil {
		t.Fatal(err)
	}
	serverConn, err := l.Accept()
	if err != nil {
		t.Fatal(err)
	}
	processConn(serverConn)
}

func checkResp(c net.Conn, expected byte, t *testing.T) {
	buf := make([]byte, 1024)
	n, err := c.Read(buf)
	if err != nil {
		t.Fatal("Unexpected response, ", err)
	}

	if n > 0 && buf[0] == expected {
		log.Println("RECEIVED: ", n, " ", string(buf[:n]))
	} else {
		t.Fatal("Unexpected response, ", string(buf))
	}
}
