package main

import (
	"fmt"
	"log"
	"net"
)

func handleMessage(conn *net.UDPConn) {
	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buffer)
	fmt.Println("received:", string(buffer[:n]))

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	config := newConfig()

	udpAddr, err := net.ResolveUDPAddr("udp4", config.Address())

	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("UDP server listening on", config.Address())
	defer ln.Close()

	for {
		handleMessage(ln)
	}

}
