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
		return nil, &Response{
			code:    "501",
			message: "Syntax error in parameters or arguments.",
			err:     fmt.Errorf("wrong number of arguments"),
		}
	}
	re1 := regexp.MustCompile("^([0-9]+,){5}[0-9]+$")

	if !re1.MatchString(a[1]) {
		return nil, &Response{
			code:    "501",
			message: "Syntax error in parameters or arguments.",
			err:     fmt.Errorf("wrong host-port syntax"),
		}
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
	log.Println("connecting " + c.host + ":" + strconv.Itoa(c.port))
	dc, err := net.Dial("tcp", c.host+":"+strconv.Itoa(c.port))
	if err != nil {
		log.Println("error connecting ", err)
		ch <- &Response{
			code:    "425",
			message: "Can't open data connection.",
			err:     err,
		}
	} else {
		log.Println("success")
		s.DataConn = dc
		ch <- &Response{
			code:    "225",
			message: "Data connection open; no transfer in progress.",
			err:     err,
		}
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
