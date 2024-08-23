package tcp

import (
	"fmt"
	"log"
	"net"
)

func main() {

	const (
		network = "tcp"
		address = "127.0.0.1"
	)

	// create new listener
	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatalln("cant listen on given address", address, err)
	}

	fmt.Println("listener address is : ", listener.Addr())

	for {
		// listen for new connection
		connection, aErr := listener.Accept()
		if err != nil {
			log.Fatalln("cant listen to new connection", aErr)
		}

		fmt.Println("client address is : ", connection.RemoteAddr(), connection.LocalAddr())

		// process the request
		var data = make([]byte, 1024)
		numberOfBytes, rErr := connection.Read(data)

		if rErr != nil {
			log.Println("cant read data from connection", rErr)

			continue
		}

		fmt.Println("number Of Bytes: ", numberOfBytes)
		fmt.Println("data: ", data)
	}

}
