package protocol

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
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
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	s.Lock()
	defer s.Unlock()
	if s.DataConn != nil {
		s.DataConn.Close()
	}
	lowestPort := 1000
	var tryPort int
	var sock net.Listener
	var err error
	for i := 0; i < 5; i++ {
		tryPort = random.Int()%(65535-lowestPort) + lowestPort
		sock, err = net.Listen("tcp", s.LocalAddr+":"+strconv.Itoa(tryPort))
		if err == nil {
			break
		}
	}

	if err != nil {
		log.Println("Error accepting passive connection ", err)
		ch <- NewResponse(Response425, "", err)
	} else {
		//server can accept connections at multiple addresses
		//we should identify current one
		localip := s.CmdConn.LocalAddr().String()
		localip = localip[:strings.LastIndex(localip, ":")] //remove port
		respString := "Entering Passive Mode ("
		respString += strings.Replace(localip, ".", ",", -1)
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
func NewPassiveCmd() *PassiveCmd {
	return &PassiveCmd{}
}
