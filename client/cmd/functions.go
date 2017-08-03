package cmd

import (
	"bufio"
	"ftp/protocol"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func ftpQuit(conn net.Conn, bufReader *bufio.Reader) {
	var command protocol.Command
	command = protocol.NewQuitCmd()
	command.Send(conn)
	log.Println("sent cmd quit")
	getResponse(bufReader)
}
func ftpGetFile(conn net.Conn, bufReader *bufio.Reader, arg string) bool {
	arg = path.Clean(arg)
	dataConn := openDataConn(conn, bufReader)
	var command protocol.Command
	//get file
	command = protocol.NewRetrieveCmd(arg)
	command.Send(conn)
	log.Println("sent cmd retr")

	getPrevResponse(bufReader) //before starting data transfer
	f, err := os.Create(arg)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(f, dataConn)
	if err != nil {
		log.Fatal(err)
	}
	dataConn.Close()
	log.Println("conn closed")

	_, ok := getResponse(bufReader) //need 2 responses for each get
	return ok
}

func ftpList(conn net.Conn, bufReader *bufio.Reader, arg string) bool {
	arg = path.Clean(arg)
	dataConn := openDataConn(conn, bufReader)
	var command protocol.Command
	//list
	command = protocol.NewListCmd(arg)
	command.Send(conn)

	getPrevResponse(bufReader) //before starting data transfer

	_, err := io.Copy(os.Stdout, dataConn)
	if err != nil {
		log.Fatal(err)
	}
	dataConn.Close()
	log.Println("conn closed")

	_, ok := getResponse(bufReader) //need 2 responses for each get
	return ok
}

func ftpPutFile(conn net.Conn, bufReader *bufio.Reader, arg string) bool {
	arg = path.Clean(arg)
	dataConn := openDataConn(conn, bufReader)
	var command protocol.Command
	//get file
	command = protocol.NewStoreCmd(arg)
	command.Send(conn)
	log.Println("sent cmd stor")

	getPrevResponse(bufReader) //before starting data transfer
	f, err := os.Open(arg)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(dataConn, f)
	if err != nil {
		log.Fatal(err)
	}
	dataConn.Close()
	log.Println("conn closed")

	_, ok := getResponse(bufReader) //need 2 responses for each get
	return ok
}

func ftpLogin(conn net.Conn, bufReader *bufio.Reader) bool {
	var command protocol.Command
	command = protocol.NewUserCmd(ftpUser)
	command.Send(conn)
	getResponse(bufReader)
	command = protocol.NewPassCmd(ftpPassword)
	command.Send(conn)
	_, ok := getResponse(bufReader)
	return ok
}
func ftpSetFlags(conn net.Conn, bufReader *bufio.Reader) bool {
	var command protocol.Command
	failed := false
	command = protocol.NewSystemCmd()
	command.Send(conn)
	_, ok := getResponse(bufReader)
	failed = failed || !ok
	command = protocol.NewTypeCmd(ftpType, "")
	command.Send(conn)
	_, ok = getResponse(bufReader)
	failed = failed || !ok
	return !failed
}
func openDataConn(cmdConn net.Conn, bufReader *bufio.Reader) net.Conn {
	var dataConn net.Conn
	var sock net.Listener
	var err error
	var command protocol.Command
	if ftpPassive {
		command = protocol.NewPassiveCmd()
		command.Send(cmdConn)
		response, ok := getResponse(bufReader)
		re1 := regexp.MustCompile("([0-9]+,){5}[0-9]+")
		if ok && re1.MatchString(response.message) {
			addr := re1.FindString(response.message)
			addrArr := strings.Split(addr, ",")
			remoteHost := strings.Join(addrArr[:4], ".")
			port1, _ := strconv.Atoi(addrArr[4])
			port2, _ := strconv.Atoi(addrArr[5])
			remotePort := port1*256 + port2
			dataConn, err = net.Dial("tcp", remoteHost+":"+strconv.Itoa(remotePort))
			if err == nil {
				log.Fatal(err)
			} else {
				return dataConn
			}
		} else {
			log.Fatal(response)
		}
		//net.Dial(dataConn)
	} else {
		localIP := cmdConn.LocalAddr().String()
		localIP = localIP[:strings.LastIndex(localIP, ":")] //remove port
		lowestPort := 1000
		var tryPort int
		for i := 0; i < 15; i++ {
			log.Println("start listen")
			tryPort = rand.Int()%(65535-lowestPort) + lowestPort
			log.Println("data port ", localIP, tryPort) //TODO: remove
			sock, err = net.Listen("tcp", localIP+":"+strconv.Itoa(tryPort))
			if err == nil {
				break
			}
		}

		if err != nil {
			log.Fatal("Error creating client port ", err)
		} else {
			command = protocol.NewPortCmd(localIP, tryPort)
			command.Send(cmdConn)

			log.Println("sent cmd")
			dataConn, err = sock.Accept()
			if err != nil {
				log.Fatal(err)
			}
			getResponse(bufReader)
		}
	}
	return dataConn
}
