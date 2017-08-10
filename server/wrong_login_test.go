package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"testing"
	"time"

	"github.com/apsifly/goftp/protocol"
)

func TestWrongUser(t *testing.T) {
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
		s, err := net.Dial("tcp", "127.0.0.1:"+strconv.Itoa(tryPort))
		if err != nil {
			t.Fatal(err)
		}
		defer s.Close()
		checkResp(s, '2', t)
		if _, err := fmt.Fprintf(s, "USER user2\r\n"); err != nil {
			t.Fatal(err)
		}
		checkResp(s, '3', t)
		if _, err := fmt.Fprintf(s, "PASS 123\r\n"); err != nil {
			t.Fatal(err)
		}
		checkResp(s, '5', t)
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
