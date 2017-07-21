package protocol

import (
	"fmt"
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
