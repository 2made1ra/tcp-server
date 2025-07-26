package main

import (
	"log"
	"net"
)

const (
	Protocol = "tcp"
	Port     = ":8080"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	log.Printf("Client have been connected: %s\n", conn.RemoteAddr().String())

	_, err := conn.Write([]byte("OK\n"))

	if err != nil {
		log.Println("Error with sending answer", err)
	}

	log.Printf("Message was sending to the client: %s\n", conn.RemoteAddr().String())
}

func main() {
	//Создание TCP адреса
	tcpAddr, _ := net.ResolveTCPAddr(Protocol, Port)

	//Открытие сокета для прослушивания
	listener, err := net.ListenTCP(Protocol, tcpAddr)

	if err != nil {
		//Вызываю метод Fatal, чтобы корретно вызвать defer на закрытие
		log.Fatal("Connection error:", err)
	}

	defer listener.Close()

	log.Printf("Listening on port %v\\n", Port)

	//Обработка множественных соединений
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error with accept connection", err)
			continue
		}

		go handleConnection(conn)
	}

}
