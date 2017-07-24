package protocol

import (
	"fmt"
	"io"
	"os"
)

type StoreCmd struct {
	path string
}

func parseStore(a []string) (*StoreCmd, *Response) {
	if len(a) != 2 || len(a[1]) == 0 {
		return nil, &Response{
			code:    "501",
			message: "Syntax error in parameters or arguments.",
			err:     fmt.Errorf("no file path provided"),
		}
	}

	return &StoreCmd{
		path: a[1],
	}, nil
}

func (c *StoreCmd) Execute(s *State, ch chan *Response) {
	s.Lock()

	defer s.Unlock()
	if s.DataConn != nil && s.StorActive == false {
		f, err := os.Create(c.path)
		if err != nil {
			ch <- &Response{
				code:    "550",
				message: "Requested action not taken.",
				err:     err,
			}
			return
		}
		s.StorActive = true
		s.Unlock()
		_, err = io.Copy(f, s.DataConn)
		s.Lock()
		s.StorActive = false
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

func (c *StoreCmd) Send(w io.Writer) error {
	_, err := io.WriteString(w, "STOR "+c.path+"\r\n")
	return err
}
