package protocol

import "fmt"

type NoopCmd struct {
}

func parseNoop(a []string) (*NoopCmd, error) {
	if len(a) != 1 {
		return nil, fmt.Errorf("wrong argument")
	}

	return &NoopCmd{}, nil
}
