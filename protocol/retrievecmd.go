package protocol

import (
	"fmt"
	"io"
	"os"
)

type RetrieveCmd struct {
	path string
}

func parseRetrieve(a []string) (*RetrieveCmd, *Response) {
	if len(a) != 2 || len(a[1]) == 0 {
		return nil, &Response{
			code:    "501",
			message: "Syntax error in parameters or arguments.",
			err:     fmt.Errorf("no file path provided"),
		}
	}

	return &RetrieveCmd{
		path: a[1],
	}, nil
}

func (c *RetrieveCmd) Execute(s *State, ch chan *Response) {
	s.Lock()

	defer s.Unlock()
	if s.DataConn != nil && s.RetrActive == false {
		f, err := os.Open(c.path)
		if err != nil {
			ch <- &Response{
				code:    "550",
				message: "Requested action not taken.",
				err:     err,
			}
			return
		}
		s.RetrActive = true
		s.Unlock()
		_, err = io.Copy(s.DataConn, f)
		s.Lock()
		s.RetrActive = false
		if err != nil {
			ch <- &Response{
				code:    "450",
				message: "Requested file action not taken.",
				err:     err,
			}
		} else {
			ch <- &Response{
				code:    "250",
				message: "Requested file action okay, completed.",
				err:     err,
			}
			return
		}
	} else {
		ch <- &Response{
			code:    "550",
			message: "Requested action not taken.",
			err:     nil,
		}
	}
}

func (c *RetrieveCmd) Send(w io.Writer) error {
	_, err := io.WriteString(w, "RETR "+c.path+"\r\n")
	return err
}
