package utils

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

var Clients = make(map[string]net.Conn)

type MessageStruct struct {
	ClientNom string
	Time      string
	Message   string
}

var SendMsgHistory []string
var NbrClents int

// Gérer les noms de connexion, les entrées et les déconnexions
func ClientHandler(conn net.Conn, mutex *sync.Mutex) {
	name := acceptClients(conn, mutex)

	BroadcastMessage(name, conn)
	for {
		if !WriteMessage(conn, mutex, name) { // arrêter de lire le tampon si le client se déconnecte
			break
		}
	}
}

// Écrit l'historique du chat précédent pour les nouveaux clients qui se connectent
func SendMeessageHistory(history []string, conn net.Conn) {
	for _, v := range history {
		conn.Write([]byte(v))
	}
}

// Utilise un tampon pour attendre le message d'entrée du client
func WriteMessage(conn net.Conn, mutex *sync.Mutex, name string) bool {
	if name == "" {
		return false
	}

	buffer := make([]byte, 1024)
	bufLen, err := conn.Read(buffer)
	if err != nil {
		mutex.Lock()
		fmt.Println(err)
		conn.Close()
		delete(Clients, name)
		for _, c := range Clients {
			c.Write([]byte(name + " has left our chat" + "\n"))
		}

		NbrClents--
		mutex.Unlock()
		return false
	}

	msgBuf := string(buffer)
	msgtxt := string(msgBuf[:bufLen-1])

	// Vérifie si le message est vide
	if isEmptyMessage([]byte(msgtxt)) {
		conn.Write([]byte("No empty msg in the chat. Write something.\n"))
		return true // Continuer à lire les messages du client
	}

	Clients[name] = conn

	msg := MessageStruct{
		Message:   msgtxt,
		ClientNom: name,
		Time:      time.Now().Format("01-02-2006 15:04:05"),
	}
	msgbroadcast := "[" + msg.Time + "][" + msg.ClientNom + "]:" + msg.Message + "\n"
	for key, c := range Clients {
		if key == name {
			conn.Write([]byte("\033[A"))
		}
		c.Write([]byte(msgbroadcast))
	}
	SendMsgHistory = append(SendMsgHistory, msgbroadcast)
	return true
}

// Diffuse le nouveau client à la liste actuelle des clients
func BroadcastMessage(s string, excludeClient net.Conn) {
	if s != "" {
		for _, c := range Clients {
			// Exclure le nouveau client de l'envoi du message de bienvenue
			if c != excludeClient {
				c.Write([]byte(s + " has join our chat" + "\n"))
			}
		}
	}
}

func isEmptyMessage(message []byte) bool {
	messageStr := string(message)
	// Supprime les espaces de début et de fin du message
	trimMessage := strings.TrimSpace(messageStr)
	return len(trimMessage) == 0
}
