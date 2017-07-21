package protocol

import "fmt"

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
