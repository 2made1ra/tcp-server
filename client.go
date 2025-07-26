package main

import (
	"bufio"
	"log"
	"net"
)

const (
	Protocol = "tcp"
	Port     = ":8080"
)

func main() {
	//Подключаюсь по TCP
	conn, err := net.Dial(Protocol, Port)

	//Проверка соединения
	if err != nil {
		log.Fatal("Connection error:", err)
	}

	//defer нужен, чтобы в конце выполнения main корреткно закрыть соединение
	defer conn.Close()

	//Читаю и проверяю ответ
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("Reading error:", err)
	}

	if response == "OK\n" {
		log.Println("Correct response:", response)
	}
}
