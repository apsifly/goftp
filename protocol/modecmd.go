package protocol

import (
	"fmt"
	"strings"
)

type ModeCmd struct {
	mode string
}

func parseMode(a []string) (*ModeCmd, *Response) {
	if len(a) != 2 || len(a[1]) != 1 {
		return nil, &Response{
			code:    "501",
			message: "Syntax error in parameters or arguments.",
			err:     fmt.Errorf("wrong mode"),
		}
	}
	switch a[1][0:1] {
	case "S", "s", "B", "b", "C", "c":
		return &ModeCmd{
			mode: strings.ToUpper(a[1][0:1]),
		}, nil
	default:
		return nil, &Response{
			code:    "501",
			message: "Syntax error in parameters or arguments.",
			err:     fmt.Errorf("wrong mode"),
		}
	}
}
