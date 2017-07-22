package protocol

import (
	"io"
	"sync"
)

type State struct {
	sync.Mutex
	User       string
	Pass       string
	Logged     bool
	DataConn   io.ReadWriteCloser
	RetrActive bool
	StorActive bool
	CmdActive  map[int]*Command
}
