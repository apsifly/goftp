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

func (c *PassCmd) Execute(rw io.ReadWriter, s State) (*Response, State) {
	switch s := s.(type) {
	case SUserProvided:
		return &Response{
				code:    "530", //pass stored for authentification
				message: "Not logged in.",
				err:     nil,
			},
			SUserPassProvided{
				user: s.user,
				pass: c.pass,
			}

	case SConnected:
		return &Response{
				code:    "332",
				message: "Need account for login.",
				err:     nil,
			},
			SConnected{}

	default:
		return &Response{
			code:    "503",
			message: "Bad sequence of commands.",
			err:     fmt.Errorf("state: %t, %v, tried to send password, %s", s, s, c.pass),
		}, s
	}
}

func (c *PassCmd) Send(w io.Writer) error {
	_, err := io.WriteString(w, "PASS "+c.pass+"\r\n")
	return err
}
