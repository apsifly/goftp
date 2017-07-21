package protocol

import (
	"fmt"
	"regexp"
	"strings"
)

type TypeCmd struct {
	main string
	sub  string
}

func parseType(a []string) (*TypeCmd, error) {
	if len(a) != 2 {
		return nil, fmt.Errorf("no type argument")
	}
	typearg := a[1]
	if len(typearg) > 1 && typearg[1:2] != " " {
		return nil, fmt.Errorf("wrong type argument")
	}
	cmd := &TypeCmd{}
	switch typearg[0:1] {
	case "A", "a", "E", "e":
		switch typearg[2:3] {
		case "N", "n", "T", "t", "C", "c":
			cmd.main = strings.ToUpper(typearg[0:1])
			cmd.sub = strings.ToUpper(typearg[2:3])
		default:
			return nil, fmt.Errorf("wrong type argument")
		}
	case "I", "i":
		cmd.main = strings.ToUpper(typearg[0:1])

	case "L", "l":
		re1 := regexp.MustCompile("^[0-9]+$")
		if !re1.MatchString(typearg[2:]) {
			return nil, fmt.Errorf("wrong type argument")
		}
		cmd.main = strings.ToUpper(typearg[0:1])
		cmd.sub = strings.ToUpper(typearg[2:])
	default:
		return nil, fmt.Errorf("wrong type argument")
	}
	return cmd, nil
}
