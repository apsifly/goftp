package protocol

import "fmt"

type StoreCmd struct {
	path string
}

func parseStore(a []string) (*StoreCmd, error) {
	if len(a) != 2 || len(a[1]) == 0 {
		return nil, fmt.Errorf("no file path provided")
	}

	return &StoreCmd{
		path: a[1],
	}, nil
}
