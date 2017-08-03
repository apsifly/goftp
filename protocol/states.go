package protocol

import (
	"net"
	"sync"
)

type State struct {
	sync.Mutex
	User         string
	Pass         string
	Logged       bool
	DataConn     net.Conn
	RetrActive   bool
	StorActive   bool
	TransferType string
	LocalAddr    string
	CmdConn      net.Conn
}
