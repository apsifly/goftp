package protocol

import (
	"fmt"
	"io"
)

type PassCmd struct {
	pass string
}

func parsePass(a []string) (*PassCmd, *Response) {
	if len(a) != 2 {
		return nil, NewResponse(Response501, "", fmt.Errorf("Password not provided"))
	}
	return &PassCmd{
		pass: a[1],
	}, nil
}

func (c *PassCmd) Execute(s *State, ch chan *Response) {

	s.Lock()
	//stores password for server verification
	//other logic is at server side
	s.Pass = c.pass
	s.Unlock()

}

func (c *PassCmd) Send(w io.Writer) error {
	_, err := io.WriteString(w, "PASS "+c.pass+"\r\n")
	return err
}
func NewPassCmd(p string) *PassCmd {
	return &PassCmd{
		pass: p,
	}
}
