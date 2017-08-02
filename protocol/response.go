package protocol

import (
	"io"
	"log"
	"strings"
)

type Response struct {
	Base          baseResponse
	CustomMessage string
	err           error //for internal usage
}

func (r *Response) Send(w io.Writer) {

	var m string
	if r.CustomMessage != "" {
		m = r.CustomMessage
	} else {
		m = r.Base.message
	}
	if m[len(m)-1:] != "\n" {
		m += "\n"
	}
	if r.err != nil {
		log.Println("sent response ", r.Base.code+" "+m, " with error ", r.err.Error())
	} else {
		log.Println("sent response ", r.Base.code+" "+m)
	}
	io.WriteString(w, r.Base.code+" "+strings.Replace(m, "\r", "\r\n", -1))

}

func NewResponse(b baseResponse, customMessage string, e error) *Response {
	return &Response{
		Base:          b,
		CustomMessage: customMessage,
		err:           e,
	}
}
