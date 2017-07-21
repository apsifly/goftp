package protocol

import "fmt"

type UserCmd struct {
	user string
}

func parseUser(a []string) (*UserCmd, error) {
	if len(a) != 2 {
		return nil, fmt.Errorf("username not provided")
	}
	return &UserCmd{
		user: a[1],
	}, nil
}
