package main

import (
	"fgsoftware1/fgcrypto/helper"
	"fmt"
)

func main() {
	Message := ""
	Secret := ""

	fmt.Print("\x1b[38;5;4m")
	fmt.Println("message: ")
	var messageToDec string
	fmt.Scanln(&messageToDec)
	Message = string(messageToDec)

	fmt.Print("\x1b[38;5;4m")
	fmt.Println("key: ")
	var key string
	fmt.Scanln(&key)
	Secret = string(key)

	decText, err := helper.Decrypt(Message, Secret)

	if err != nil {
		fmt.Println("error decrypting your encrypted text: ", err)
	}
	fmt.Println(decText)
}
