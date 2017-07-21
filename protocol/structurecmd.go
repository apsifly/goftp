package protocol

import (
	"fmt"
	"strings"
)

type StructureCmd struct {
	s string
}

func parseStructure(a []string) (*StructureCmd, *Response) {
	if len(a) != 2 || len(a[1]) != 1 {
		return nil, &Response{
			code:    "501",
			message: "Syntax error in parameters or arguments.",
			err:     fmt.Errorf("not enough arguments"),
		}
	}
	switch a[1][0:1] {
	case "F", "f", "R", "r", "P", "p":
		return &StructureCmd{
			s: strings.ToUpper(a[1][0:1]),
		}, nil
	default:
		return nil, &Response{
			code:    "504",
			message: "Command not implemented for that parameter.",
			err:     fmt.Errorf("wrong structure marker"),
		}
	}
}
