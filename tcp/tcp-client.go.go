package tcp

import (
	"fmt"
	"log"
	"net"
)

func main() {

	connection, err := net.Dial("tcp", "127.0.0.1")
	if err != nil {
		log.Fatalln("cant dial the given address", err)
	}

	fmt.Println("local address:", connection.LocalAddr())
}
