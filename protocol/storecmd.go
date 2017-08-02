package protocol

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
)

type StoreCmd struct {
	Path string
}

func parseStore(a []string) (*StoreCmd, *Response) {
	if len(a) != 2 || len(a[1]) == 0 {
		return nil, NewResponse(Response501, "", fmt.Errorf("no file path provided"))
	}

	return &StoreCmd{
		Path: path.Clean(a[1]),
	}, nil
}

func (c *StoreCmd) Execute(s *State, ch chan *Response) {
	s.Lock()

	defer s.Unlock()
	if s.DataConn != nil && s.StorActive == false {
		f, err := os.Create(c.Path)
		if err != nil {
			ch <- NewResponse(Response550, "", err)
			return
		}
		ch <- NewResponse(Response150, "", nil)
		s.StorActive = true
		s.Unlock()
		//err = storeWithTransform(f, s.DataConn, s.TransferType)
		_, err = io.Copy(f, s.DataConn)
		s.Lock()
		s.StorActive = false
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

func (c *StoreCmd) Send(w io.Writer) error {
	_, err := io.WriteString(w, "STOR "+c.Path+"\r\n")
	return err
}

func storeWithTransform(dst io.ReadWriteCloser, src io.Reader, mode string) error {
	buf := make([]byte, 32*1024)
	var err error
	var dbuf []byte
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			switch mode[0] {
			case 'A':
				dbuf = bytes.Replace(buf[:nr], []byte("\r\n"), []byte("\n"), -1)
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
