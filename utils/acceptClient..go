package utils

import (
	"fmt"
	"net"
	"os"
	"sync"
)

// Récupère un nom auprès d'un client connecté, le stocke, envoie un message de bienvenue et diffuse l'historique du chat au client
func acceptClients(conn net.Conn, mutex *sync.Mutex) string {
	mutex.Lock()
	txt, err := os.ReadFile("penguin.txt")
	if err != nil {
		fmt.Println("welcom I'm a penguin but there is and error")
	}
	mutex.Unlock()
	conn.Write(txt)

ReName:
	buffer := make([]byte, 1024)
	bufLen, err := conn.Read(buffer)
	if err != nil {
		NbrClents--
		conn.Close()
		fmt.Println(err)
		return "" // Interrompt l'étiquette de saut de renommage lorsque le client se déconnecte
	}
	msgBuf := string(buffer)
	name := string(msgBuf[:bufLen-1])
	if name == "" || len(name) > 20 {
		conn.Write([]byte("Name invali. Enter a valid name .\n[ENTER YOUR NAME]:"))
		goto ReName
	}
	mutex.Lock()
	for val := range Clients {
		if name == val {
			conn.Write([]byte("Name already used, Cheack another name:"))
			mutex.Unlock()
			goto ReName
		}
	}
	Clients[name] = conn
	SendMeessageHistory(SendMsgHistory, conn)
	mutex.Unlock()
	return name
}
