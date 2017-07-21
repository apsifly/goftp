package protocol

import "fmt"

type RetrieveCmd struct {
	path string
}

func parseRetrieve(a []string) (*RetrieveCmd, *Response) {
	if len(a) != 2 || len(a[1]) == 0 {
		return nil, &Response{
			code:    "501",
			message: "Syntax error in parameters or arguments.",
			err:     fmt.Errorf("no file path provided"),
		}
	}

	return &RetrieveCmd{
		path: a[1],
	}, nil
}
