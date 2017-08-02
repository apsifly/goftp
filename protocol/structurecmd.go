package protocol

import (
	"fmt"
	"io"
	"strings"
)

type StructureCmd struct {
	s string
}

func parseStructure(a []string) (*StructureCmd, *Response) {
	if len(a) != 2 || len(a[1]) != 1 {
		return nil, NewResponse(Response501, "", fmt.Errorf("not enough arguments"))
	}
	switch a[1][0:1] {
	case "F", "f", "R", "r", "P", "p":
		return &StructureCmd{
			s: strings.ToUpper(a[1][0:1]),
		}, nil
	default:
		return nil, NewResponse(Response501, "", fmt.Errorf("wrong structure marker"))
	}
}

func (c *StructureCmd) Execute(s *State, ch chan *Response) {
	switch c.s {

	case "F":
		ch <- NewResponse(Response200, "", nil)
	default:
		ch <- NewResponse(Response504, "", fmt.Errorf("unsupported structure option %s", c.s))
		return
	}
}

func (c *StructureCmd) Send(w io.Writer) error {
	_, err := io.WriteString(w, "STRU "+c.s+"\r\n")
	return err
}
