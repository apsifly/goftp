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
		return nil, NewResponse(Response501, "", fmt.Errorf("username not provided"))
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
		ch <- NewResponse(Response331, "", nil)

	} else {
		ch <- NewResponse(Response331, "", fmt.Errorf("state: %v, tried to log user, %s", s, c.user))
	}
}
func (c *UserCmd) Send(w io.Writer) error {
	_, err := io.WriteString(w, "USER "+c.user+"\r\n")
	return err
}

func NewUserCmd(u string) *UserCmd {
	return &UserCmd{
		user: u,
	}
}
