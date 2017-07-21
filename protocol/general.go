package protocol

import (
	"fmt"
	"strings"
)

type Command interface {
	Execute()
	Send()
}
type Response interface {
	Send()
}

func ParseCommand(s string) (Command, error) {
	s = strings.TrimRight(s, "\r\n")
	a := strings.SplitN(s, " ", 2)
	if len(a) == 0 {
		return nil, fmt.Errorf("can not parse command")
	}
	var cmd Command
	var err error
	switch a[0] {
	case "USER":
		cmd, err = parseUser(a)
	case "QUIT":
		cmd, err = parseQuit(a)
	case "PORT":
		cmd, err = parsePort(a)
	case "TYPE":
		cmd, err = parseType(a)

	}
}
