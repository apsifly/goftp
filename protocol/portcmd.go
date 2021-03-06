package protocol

import (
	"fmt"
	"io"
	"log"
	"net"
	"regexp"
	"strconv"
	"strings"
)

type PortCmd struct {
	host string
	port int
}

func parsePort(a []string) (*PortCmd, *Response) {
	if len(a) != 2 {
		return nil, NewResponse(Response501, "", fmt.Errorf("wrong number of arguments"))
	}
	re1 := regexp.MustCompile("^([0-9]+,){5}[0-9]+$")

	if !re1.MatchString(a[1]) {
		return nil, NewResponse(Response501, "", fmt.Errorf("wrong host-port syntax"))
	}
	hostport := strings.Split(a[1], ",")
	mustAtoi := func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}
	cmd := &PortCmd{
		host: strings.Join(hostport[:4], "."),
		port: mustAtoi(hostport[4])*256 + mustAtoi(hostport[5]),
	}
	return cmd, nil
}

func (c *PortCmd) Execute(s *State, ch chan *Response) {
	s.Lock()
	defer s.Unlock()
	if s.DataConn != nil {
		s.DataConn.Close()
	}
	log.Println("connecting " + c.host + ":" + strconv.Itoa(c.port))
	dc, err := net.Dial("tcp", c.host+":"+strconv.Itoa(c.port))
	if err != nil {
		log.Println("error connecting ", err)
		ch <- NewResponse(Response425, "", err)
	} else {
		log.Println("success")
		s.DataConn = dc
		ch <- NewResponse(Response200, "", err)
	}
}
func (c *PortCmd) Send(w io.Writer) error {
	message := "PORT " +
		strings.Replace(c.host, ".", ",", 3) + "," +
		strconv.Itoa(c.port/256) + "," +
		strconv.Itoa(c.port%256) + "\r\n"

	_, err := io.WriteString(w, message)
	return err
}

func NewPortCmd(host string, port int) *PortCmd {
	return &PortCmd{
		host: host,
		port: port,
	}
}
