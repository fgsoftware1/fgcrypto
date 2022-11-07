package main

import (
	"fgsoftware1/fgcrypto/helper"
	"fmt"
)

var (
	Message string
	Secret  string
)

func main() {
	fmt.Print("\x1b[38;5;4m")
	fmt.Println("message: ")
	var messageToEnc string
	fmt.Scanln(&messageToEnc)
	Message = string(messageToEnc)

	fmt.Print("\x1b[38;5;4m")
	fmt.Println("key: ")
	var key string
	fmt.Scanln(&key)
	Secret = string(key)

	encText, err := helper.Encrypt(Message, Secret)
	if err != nil {
		fmt.Println("error encrypting your classified text: ", err)
	}

	fmt.Println("\x1b[38;5;4mKEY: ", "\x1b[38;5;13m", Secret)
	fmt.Println("\x1b[38;5;4mMessage: ", "\x1b[38;5;13m", string(encText))
}
