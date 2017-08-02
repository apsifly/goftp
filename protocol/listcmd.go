package protocol

import (
	"fmt"
	"io"
	"io/ioutil"
	"path"
)

type ListCmd struct {
	Path string
}

func parseList(a []string) (*ListCmd, *Response) {
	if len(a) != 2 || len(a[1]) == 0 {
		return &ListCmd{
			Path: ".",
		}, nil
	}

	return &ListCmd{
		Path: path.Clean(a[1]),
	}, nil
}

func (c *ListCmd) Execute(s *State, ch chan *Response) {
	if s.DataConn != nil {
		ch <- NewResponse(Response150, "", nil)
		files, err := ioutil.ReadDir(c.Path)
		if err != nil {
			ch <- NewResponse(Response451, "Error while reading specified directory", err)
		} else {
			var lasterr error = nil
			for _, file := range files {
				_, err := io.WriteString(s.DataConn, fmt.Sprintf("%s\t%d\t%s\t%s\r\n", file.Mode(), file.Size(), file.ModTime(), file.Name()))
				if err != nil {
					lasterr = err
				}
			}
			if lasterr != nil {
				ch <- NewResponse(Response425, "", lasterr)
			} else {

				ch <- NewResponse(Response226, "", nil)
				s.DataConn.Close() //expected by client
				s.DataConn = nil
			}
		}

	} else {
		ch <- NewResponse(Response425, "", nil)
	}
}

func (c *ListCmd) Send(w io.Writer) error {
	_, err := io.WriteString(w, "LIST "+c.Path+"\r\n")
	return err
}
