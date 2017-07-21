package protocol

import (
	"fmt"
	"io"
	"log"
	"strings"
)

type Command interface {
	//Execute()
	//Send()
}

func ParseCommand(s string) (Command, *Response) {
	s = strings.TrimRight(s, "\r\n")
	a := strings.SplitN(s, " ", 2)
	if len(a) == 0 {
		return nil, &Response{
			code:    "500",
			message: "Syntax error, command unrecognized.",
			err:     fmt.Errorf("can not parse command in strings.Split"),
		}
	}
	var cmd Command
	var resp *Response
	switch a[0] {
	case "USER":
		cmd, resp = parseUser(a)
	case "QUIT":
		cmd, resp = parseQuit(a)
	case "PORT":
		cmd, resp = parsePort(a)
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
	case "NOOP":
		cmd, resp = parseNoop(a)
	default:
		return nil, &Response{
			code:    "500",
			message: "Syntax error, command unrecognized.",
			err:     fmt.Errorf("unknown command"),
		}

	}
	return cmd, resp
}

type Response struct {
	code    string
	message string
	err     error //for internal usage
}

func (r *Response) Send(w io.Writer) {
	if r.err != nil {
		log.Println("sent response ", r.code+" "+r.message, " with error ", r.err.Error())
	}
	io.WriteString(w, r.code+" "+r.message+"\r\n")

}
