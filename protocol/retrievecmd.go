package protocol

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
)

type RetrieveCmd struct {
	Path string
}

func parseRetrieve(a []string) (*RetrieveCmd, *Response) {
	if len(a) != 2 || len(a[1]) == 0 {
		return nil, NewResponse(Response501, "", fmt.Errorf("no file path provided"))
	}

	return &RetrieveCmd{
		Path: path.Clean(a[1]),
	}, nil
}

func (c *RetrieveCmd) Execute(s *State, ch chan *Response) {
	s.Lock()

	defer s.Unlock()
	if s.DataConn != nil && s.RetrActive == false {
		f, err := os.Open(c.Path)
		if err != nil {
			ch <- NewResponse(Response550, "", err)
			return
		}
		s.RetrActive = true
		ch <- NewResponse(Response150, "", nil)
		s.Unlock()
		err = copyWithTransform(s.DataConn, f, s.TransferType)
		s.DataConn.Close()
		s.Lock()
		s.RetrActive = false
		if err != nil {
			ch <- NewResponse(Response450, "", err)
		} else {
			ch <- NewResponse(Response250, "", err)
			return
		}
	} else {
		ch <- NewResponse(Response450, "", nil)
	}
}

func (c *RetrieveCmd) Send(w io.Writer) error {
	_, err := io.WriteString(w, "RETR "+c.Path+"\r\n")
	return err
}

func copyWithTransform(dst io.ReadWriteCloser, src io.Reader, mode string) error {
	buf := make([]byte, 32*1024)
	var err error
	var dbuf []byte
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			switch mode[0] {
			case 'A':
				tmpbuf := bytes.Replace(buf[:nr], []byte("\r\n"), []byte("\n"), -1) //keep \r\n
				dbuf = bytes.Replace(tmpbuf, []byte("\n"), []byte("\r\n"), -1)      //and replace \n with \r\n
			default:
				dbuf = buf[:nr]
			}
			_, ew := dst.Write(dbuf)
			if ew != nil {
				err = ew
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	return err

}
