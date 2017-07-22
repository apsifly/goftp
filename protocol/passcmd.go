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
		return nil, &Response{
			code:    "501",
			message: "Syntax error in parameters or arguments.",
			err:     fmt.Errorf("Password not provided"),
		}
	}
	return &PassCmd{
		pass: a[1],
	}, nil
}

func (c *PassCmd) Execute(s *State, ch chan *Response) {
	//switch s := s.(type) {
	if s.Logged == false && s.User != "" {
		s.Lock()
		s.Pass = c.pass
		s.Unlock()
		ch <- &Response{
			code:    "530", //pass stored for authentification
			message: "Not logged in.",
			err:     nil,
		}

	} else if s.Logged == false {
		ch <- &Response{
			code:    "332",
			message: "Need account for login.",
			err:     nil,
		}

	} else {
		ch <- &Response{
			code:    "503",
			message: "Bad sequence of commands.",
			err:     fmt.Errorf("state: %v, tried to send password, %s", s, c.pass),
		}
	}
}

func (c *PassCmd) Send(w io.Writer) error {
	_, err := io.WriteString(w, "PASS "+c.pass+"\r\n")
	return err
}
