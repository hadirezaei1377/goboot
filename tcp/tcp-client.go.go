package tcp

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("command", os.Args[1])

	message := "default message"

	if len(os.Args) > 1 {
		message = os.Args[1]
	}

	connection, err := net.Dial("tcp", "127.0.0.1")
	if err != nil {
		log.Fatalln("cant dial the given address", err)
	}

	fmt.Println("local address:", connection.LocalAddr())

	numberOfWrittenBytes, wErr := connection.Write([]byte(message))

	if wErr != nil {
		log.Fatalln("cant write data to connection: ", wErr)
	}

	fmt.Println("number Of Written Bytes: ", numberOfWrittenBytes)

	// write data
	var data = make([]byte, 1024)
	_, rErr := connection.Read(data)

	if rErr != nil {
		log.Fatalln("cant read data from connection", rErr)

	}

	fmt.Println("server response:", string(data))

}
