package protocol

import (
	"fmt"
	"ftp/protocol/osdependent"
	"io"
)

type SystemCmd struct {
}

func parseSystem(a []string) (*SystemCmd, *Response) {
	if len(a) != 1 {
		return nil, NewResponse(Response501, "", fmt.Errorf("got additional parameter to syst"))
	}
	return &SystemCmd{}, nil
}

func (c *SystemCmd) Execute(s *State, ch chan *Response) {
	ch <- NewResponse(Response215, osdependent.FtpSystem+" Type: "+ //right type for transfer mode detection
		osdependent.FtpType, nil)
}
func (c *SystemCmd) Send(w io.Writer) error {
	io.WriteString(w, "SYST\r\n")
	return nil
}
func NewSystemCmd() *SystemCmd {
	return &SystemCmd{}
}
