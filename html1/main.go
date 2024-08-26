package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	var response = []byte(`HTTP/1.1 500 OK   
Server: Go Casts TCP Server
Accept-Ranges: bytes
Content-Length: 4
Content-Type: text/html

TEST`) // write a response manually in http format on tcp server

	listener, _ := net.Listen("tcp", ":8080")
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("err", err)

		return
	}

	var data = make([]byte, 1024)
	_, rErr := conn.Read(data)
	if rErr != nil {
		fmt.Println("rErr", rErr)
		return
	}

	method, path := parseHTTPRequest(string(data))
	if method == "GET" && path != "/" { // customize get to our html file
		var fData = make([]byte, 200480)
		f, _ := os.OpenFile(strings.Replace(path, "/", "", 1)+".html", os.O_RDONLY, 0777)
		fData, _ = io.ReadAll(f)
		contentLength := len(fData)
		contentType := "text/html"
		responseStatusCode := 200
		responseStatusMessage := "OK"

		response = createHTTPResponse(fData, contentType, responseStatusMessage, contentLength, responseStatusCode)

	}

	_, wErr := conn.Write(response)
	if wErr != nil {
		fmt.Println("wErr", wErr)

		return
	}
	conn.Close()
}

// parse request.txt
func parseHTTPRequest(data string) (string, string) {
	lines := strings.Split(data, "\n")

	req := strings.Split(lines[0], " ")
	method, path, httpVersion := req[0], req[1], req[2]

	var host string
	for _, line := range lines {
		if strings.Contains(line, "Host:") {
			host = strings.Replace(line, "Host: ", "", 1)
		}
	}

	fmt.Println("-----")
	fmt.Println("Method", method)
	fmt.Println("Path", path)
	fmt.Println("HTTP Version", httpVersion)
	fmt.Println("Host", host)
	fmt.Println("-----")

	return method, path
}

func createHTTPResponse(fData []byte, contentType, responseStatusMessage string, contentLength, responseStatusCode int) []byte {
	return []byte(fmt.Sprintf(`HTTP/1.1 %d %s
Content-Length: %d
Content-Type: %s

%s
`, responseStatusCode, responseStatusMessage, contentLength, contentType, fData))
}
