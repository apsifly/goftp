package protocol

import "fmt"

type QuitCmd struct {
}

func parseQuit(a []string) (*QuitCmd, *Response) {
	if len(a) != 1 {
		return nil, &Response{
			code:    "501",
			message: "Syntax error in parameters or arguments.",
			err:     fmt.Errorf("additional parameter to quit"),
		}
	}
	return &QuitCmd{}, nil
}
