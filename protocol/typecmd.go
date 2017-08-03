package protocol

import (
	"fmt"
	"io"
	"strings"
)

type TypeCmd struct {
	main string
	sub  string
}

func parseType(a []string) (*TypeCmd, *Response) {
	if len(a) != 2 {
		return nil, NewResponse(Response501, "", fmt.Errorf("not enough arguments"))
	}
	typearg := a[1]
	if len(typearg) > 1 && typearg[1:2] != " " {
		return nil, NewResponse(Response501, "", fmt.Errorf("wrong argument: %s", typearg))
	}
	cmd := &TypeCmd{}
	switch typearg[0:1] {
	case "A", "a", "E", "e":
		cmd.main = strings.ToUpper(typearg[0:1])
		if len(typearg) > 1 {
			switch typearg[2:3] {
			case "N", "n", "T", "t", "C", "c":
				cmd.sub = strings.ToUpper(typearg[2:3])
			default:
				return nil, NewResponse(Response504, "", fmt.Errorf("wrong secondary type"))
			}
		}
	case "I", "i":
		cmd.main = strings.ToUpper(typearg[0:1])

	case "L", "l":

		cmd.main = strings.ToUpper(typearg[0:1])
		if len(typearg) > 2 {
			cmd.sub = strings.ToUpper(typearg[2:])
		}
	default:
		return nil, NewResponse(Response504, "", fmt.Errorf("wrong primary type"))
	}
	return cmd, nil
}

func (c *TypeCmd) Execute(s *State, ch chan *Response) {
	s.TransferType = c.main + c.sub
	ch <- NewResponse(Response200, "", nil)
}
func (c *TypeCmd) Send(w io.Writer) error {
	m := "TYPE " + c.main
	if len(c.sub) != 0 {
		m += " " + c.sub
	}
	m += "\r\n"
	_, err := io.WriteString(w, m)
	return err
}
func NewTypeCmd(main, sub string) *TypeCmd {
	return &TypeCmd{
		main: main,
		sub:  sub,
	}
}
