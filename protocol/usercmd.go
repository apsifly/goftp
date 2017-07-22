package protocol

import (
	"fmt"
	"io"
)

type UserCmd struct {
	user string
}

func parseUser(a []string) (*UserCmd, *Response) {
	if len(a) != 2 {
		return nil, &Response{
			code:    "501",
			message: "Syntax error in parameters or arguments.",
			err:     fmt.Errorf("username not provided"),
		}
	}
	return &UserCmd{
		user: a[1],
	}, nil
}

func (c *UserCmd) Execute(s *State, ch chan *Response) {
	//switch s.(type) {
	if s.Logged == false {
		s.Lock()
		s.User = c.user
		s.Unlock()
		ch <- &Response{
			code:    "331",
			message: "User name okay, need password.",
			err:     nil,
		}

	} else {
		ch <- &Response{
			code:    "503",
			message: "Bad sequence of commands.",
			err:     fmt.Errorf("state: %v, tried to log user, %s", s, c.user),
		}
	}
}
func (c *UserCmd) Send(w io.Writer) error {
	_, err := io.WriteString(w, "USER "+c.user+"\r\n")
	return err
}
