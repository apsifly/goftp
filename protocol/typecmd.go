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

func parseType(a []string) (*TypeCmd, *Response) {
	if len(a) != 2 {
		return nil, &Response{
			code:    "501",
			message: "Syntax error in parameters or arguments.",
			err:     fmt.Errorf("not enough arguments"),
		}
	}
	typearg := a[1]
	if len(typearg) > 1 && typearg[1:2] != " " {
		return nil, &Response{
			code:    "501",
			message: "Syntax error in parameters or arguments.",
			err:     fmt.Errorf("wrong argument: %s", typearg),
		}
	}
	cmd := &TypeCmd{}
	switch typearg[0:1] {
	case "A", "a", "E", "e":
		switch typearg[2:3] {
		case "N", "n", "T", "t", "C", "c":
			cmd.main = strings.ToUpper(typearg[0:1])
			cmd.sub = strings.ToUpper(typearg[2:3])
		default:
			return nil, &Response{
				code:    "504",
				message: "Command not implemented for that parameter.",
				err:     fmt.Errorf("wrong secondary type"),
			}
		}
	case "I", "i":
		cmd.main = strings.ToUpper(typearg[0:1])

	case "L", "l":
		re1 := regexp.MustCompile("^[0-9]+$")
		if !re1.MatchString(typearg[2:]) {
			return nil, &Response{
				code:    "501",
				message: "Syntax error in parameters or arguments.",
				err:     fmt.Errorf("byte size is not a number"),
			}
		}
		cmd.main = strings.ToUpper(typearg[0:1])
		cmd.sub = strings.ToUpper(typearg[2:])
	default:
		return nil, &Response{
			code:    "504",
			message: "Command not implemented for that parameter.",
			err:     fmt.Errorf("wrong primary type"),
		}
	}
	return cmd, nil
}
