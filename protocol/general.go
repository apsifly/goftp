package protocol

import (
	"fmt"
	"io"
	"strings"
)

type Command interface {
	//execute is used on server and accepts current server State
	//and returns new State and result to send to control connection
	Execute(s *State, ch chan *Response)
	//send serializes command on client
	Send(w io.Writer) error
}

func ParseCommand(s string) (Command, *Response) {
	s = strings.TrimRight(s, "\r\n")
	a := strings.SplitN(s, " ", 2)
	if len(a) == 0 {
		return nil, NewResponse(Response500, "", fmt.Errorf("can not parse command in strings.Split"))
	}
	var cmd Command
	var resp *Response
	switch a[0] {
	case "USER":
		cmd, resp = parseUser(a)
	case "PASS":
		cmd, resp = parsePass(a)
	case "QUIT":
		cmd, resp = parseQuit(a)
	case "PORT":
		cmd, resp = parsePort(a)
	case "PASV":
		cmd, resp = parsePassive(a)
	case "TYPE":
		cmd, resp = parseType(a)
	case "MODE":
		cmd, resp = parseMode(a)
	case "STRU":
		cmd, resp = parseStructure(a)
	case "RETR":
		cmd, resp = parseRetrieve(a)
	case "STOR":
		cmd, resp = parseStore(a)
	case "LIST":
		cmd, resp = parseList(a)
	case "NOOP":
		cmd, resp = parseNoop(a)
	case "SYST":
		cmd, resp = parseSystem(a)
	default:
		return nil, NewResponse(Response500, "", fmt.Errorf("unknown command"))

	}
	return cmd, resp
}
