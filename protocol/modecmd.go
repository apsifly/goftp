package protocol

import (
	"fmt"
	"io"
	"strings"
)

type ModeCmd struct {
	mode string
}

func parseMode(a []string) (*ModeCmd, *Response) {
	if len(a) != 2 || len(a[1]) != 1 {
		return nil, NewResponse(Response501, "", fmt.Errorf("wrong mode"))
	}
	switch a[1][0:1] {
	case "S", "s", "B", "b", "C", "c":
		return &ModeCmd{
			mode: strings.ToUpper(a[1][0:1]),
		}, nil
	default:
		return nil, NewResponse(Response501, "", fmt.Errorf("wrong mode"))
	}
}

func (c *ModeCmd) Execute(s *State, ch chan *Response) {
	if c.mode == "S" {
		ch <- NewResponse(Response200, "", nil)
	} else {
		ch <- NewResponse(Response504, "", nil)
	}
}

func (c *ModeCmd) Send(w io.Writer) error {
	_, err := io.WriteString(w, "MODE "+c.mode+"\r\n")
	return err
}
