package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/apsifly/goftp/protocol"
)

func TestGetBinary(t *testing.T) {
	getMessage, err := base64.StdEncoding.DecodeString(base64Message)
	ioutil.WriteFile("testfile.bin", getMessage, os.FileMode(0777))
	if err != nil {
		t.Fatal("bad test")
	}
	readConfig("goftp.yml")
	config.Users["user1"].WritePath[0] = "."
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	lowestPort := 1000
	var tryPort int
	var l net.Listener
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
		list := protocol.NewRetrieveCmd("testfile.bin")
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
		finalFile, err := ioutil.ReadAll(clientConn)
		if err != nil {
			t.Fatal(err)
		}

		checkResp(s, '2', t)

		if !testEq(getMessage, finalFile) {
			t.Fatal("file changed while in transfer")
		}
		os.Remove("testfile.bin")
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
