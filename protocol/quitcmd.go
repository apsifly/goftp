package protocol

import (
	"fmt"
	"io"
)

type QuitCmd struct {
}

func parseQuit(a []string) (*QuitCmd, *Response) {
	if len(a) != 1 {
		return nil, NewResponse(Response501, "", fmt.Errorf("got additional parameter to quit"))
	}
	return &QuitCmd{}, nil
}

func (c *QuitCmd) Execute(s *State, ch chan *Response) {
	ch <- NewResponse(Response221, "", nil)
}
func (c *QuitCmd) Send(w io.Writer) error {
	io.WriteString(w, "QUIT\r\n")
	return nil
}
func NewQuitCmd() *QuitCmd {
	return &QuitCmd{}
}
