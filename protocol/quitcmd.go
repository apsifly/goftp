package protocol

import "fmt"

type QuitCmd struct {
}

func parseQuit(a []string) (*QuitCmd, error) {
	if len(a) != 1 {
		return nil, fmt.Errorf("syntax error")
	}
	return &QuitCmd{}, nil
}
