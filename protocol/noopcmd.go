package protocol

import (
	"fmt"
	"io"
)

type NoopCmd struct {
}

func parseNoop(a []string) (*NoopCmd, *Response) {
	if len(a) != 1 {
		return nil, NewResponse(Response501, "", fmt.Errorf("got additional parameter to noop"))
	}

	return &NoopCmd{}, nil
}

func (c *NoopCmd) Execute(s *State, ch chan *Response) {
	ch <- NewResponse(Response200, "", nil)
}

func (c *NoopCmd) Send(w io.Writer) error {
	_, err := io.WriteString(w, "NOOP\r\n")
	return err
}
