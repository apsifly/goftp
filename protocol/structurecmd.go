package protocol

import (
	"fmt"
	"strings"
)

type StructureCmd struct {
	s string
}

func parseStructure(a []string) (*StructureCmd, error) {
	if len(a) != 2 || len(a[1]) != 1 {
		return nil, fmt.Errorf("wrong file structure")
	}
	switch a[1][0:1] {
	case "F", "f", "R", "r", "P", "p":
		return &StructureCmd{
			s: strings.ToUpper(a[1][0:1]),
		}, nil
	default:
		return nil, fmt.Errorf("wrong file structure")
	}
}
