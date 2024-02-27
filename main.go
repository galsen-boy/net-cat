package main

import (
	"fmt"
	"os"
	"utilities/utils"
)

func main() {

	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	if len(args) == 1 {
		port := (args[0])

		if port < "1024" || port > "65535" {
			fmt.Println("Only ports 1024 a 65535 allowed")
			return
		}

		utils.Port = args[0]

	}

	utils.Server()

}
