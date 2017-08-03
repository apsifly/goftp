package cmd

import (
	"bufio"
	"log"
	"net"
	"regexp"
	"strconv"
	"strings"
)

type ftpResponse struct {
	code    string
	message string
}

var reResponse = regexp.MustCompile("^[0-9]{3}")

func ftpConnect() net.Conn {
	tc, err := net.Dial("tcp", ftpServer+":"+strconv.Itoa(ftpPort))
	if err != nil {
		log.Fatal(err)
	}
	return tc
}

func parseResponse(s string) *ftpResponse {
	if len(s) < 3 || !reResponse.MatchString(s) {
		return nil
	} else if len(s) >= 5 && s[3:4] == " " {
		return &ftpResponse{
			code:    s[:3],
			message: s[4:],
		}
	} else {
		return &ftpResponse{
			code: s[:3],
		}
	}
}

func getResponse(r *bufio.Reader) (ftpResponse, bool) {

	str, err := r.ReadString('\n')
	response := parseResponse(str)
	if err != nil || response == nil {
		log.Println(err, response)
		return ftpResponse{}, false
	}
	log.Print(response.code, "\t", response.message, "\n")
	if len(str) > 3 {
		switch response.code {
		case "215":
			typeS := "Type: "
			index := strings.Index(str, typeS)
			if len(str) > index+len(typeS) && index > 0 && ftpType == "server" {
				log.Println("set local ftp type")
				ftpLocalMode := str[index+len(typeS) : index+len(typeS)+1]
				switch ftpLocalMode {
				case "A", "a":
					ftpType = "A"
				case "L", "l":
					ftpType = "L"
				default:
					log.Fatal("unrecognized transfer mode ", str)
				}

			}
		}
	}
	if response.code[:1] != "2" {
		return *response, false
	}
	return *response, true
}
func getPrevResponse(bufReader *bufio.Reader) {
	resp, _ := getResponse(bufReader)
	if resp.code == "" || resp.code[0:1] != "1" {
		log.Fatal("error")
	}
}
