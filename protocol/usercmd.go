package protocol

import "fmt"

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

//func (c *UserCmd) Execute(w io.Writer, s State) (*Response, State)
