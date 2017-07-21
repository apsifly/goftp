package protocol

type State interface {
}

type SConnected struct {
}

type SUserProvided struct {
	user string
}
type SUserPassProvided struct {
	user string
	pass string
}
type SLogged struct {
	user string
}
