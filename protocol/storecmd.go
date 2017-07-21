package protocol

import "fmt"

type StoreCmd struct {
	path string
}

func parseStore(a []string) (*StoreCmd, *Response) {
	if len(a) != 2 || len(a[1]) == 0 {
		return nil, &Response{
			code:    "501",
			message: "Syntax error in parameters or arguments.",
			err:     fmt.Errorf("no file path provided"),
		}
	}

	return &StoreCmd{
		path: a[1],
	}, nil
}
