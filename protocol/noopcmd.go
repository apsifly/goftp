package protocol

import (
	"fmt"
	"io"
)

type NoopCmd struct {
}

func parseNoop(a []string) (*NoopCmd, *Response) {
	if len(a) != 1 {
		return nil, &Response{
			code:    "501",
			message: "Syntax error in parameters or arguments.",
			err:     fmt.Errorf("additional parameter to noop"),
		}
	}

	return &NoopCmd{}, nil
}

func (c *NoopCmd) Execute(s *State, ch chan *Response) {
	ch <- &Response{
		code:    "200",
		message: "Command okay.",
		err:     nil,
	}
}

func (c *NoopCmd) Send(w io.Writer) error {
	io.WriteString(w, "NOOP\r\n")
	return nil
}
