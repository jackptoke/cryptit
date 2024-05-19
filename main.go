package main

import (
	"fmt"
	"github.com/jackptoke/cryptit/decrypt"
	"github.com/jackptoke/cryptit/encrypt"
)

func main() {
	text := "Something"
	encrypted := encrypt.Nimbus(text)
	fmt.Println("Encrypted: ", encrypted)
	fmt.Println("Decrypted: ", decrypt.Nimbus(encrypted))
}
