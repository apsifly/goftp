package protocol

import "fmt"

type RetrieveCmd struct {
	path string
}

func parseRetrieve(a []string) (*RetrieveCmd, error) {
	if len(a) != 2 || len(a[1]) == 0 {
		return nil, fmt.Errorf("no file path provided")
	}

	return &RetrieveCmd{
		path: a[1],
	}, nil
}
