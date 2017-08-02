package protocol

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
)

type PassiveCmd struct {
}

func parsePassive(a []string) (*PassiveCmd, *Response) {
	if len(a) != 1 {
		return nil, NewResponse(Response501, "", fmt.Errorf("got additional parameter to PASV"))
	}

	return &PassiveCmd{}, nil
}

func (c *PassiveCmd) Execute(s *State, ch chan *Response) {
	s.Lock()
	defer s.Unlock()
	if s.DataConn != nil {
		s.DataConn.Close()
	}
	lowestPort := 1000
	tryPort := rand.Int()%(65535-lowestPort) + lowestPort
	var sock net.Listener
	var err error
	for i := 0; i < 5; i++ {
		sock, err = net.Listen("tcp", s.LocalIP+":"+strconv.Itoa(tryPort))
		if err == nil {
			break
		} else {
			tryPort = rand.Int()*(65535-lowestPort) + lowestPort
		}
	}

	if err != nil {
		log.Println("Error accepting passive connection ", err)
		ch <- NewResponse(Response425, "", err)
	} else {
		respString := "Entering Passive Mode ("
		respString += strings.Replace(s.LocalIP, ".", ",", -1)
		respString += "," + strconv.Itoa(tryPort/256) + "," +
			strconv.Itoa(tryPort%256) + ")."
		ch <- NewResponse(Response227, respString, err)
		log.Println("success")

		dc, _ := sock.Accept()
		s.DataConn = dc
		log.Println("accepted passive")
	}
}

func (c *PassiveCmd) Send(w io.Writer) error {
	_, err := io.WriteString(w, "PASV\r\n")
	return err
}
