package main

import (
	"fmt"
	"github.com/jackptoke/cryptit.git/decrypt"
	"github.com/jackptoke/cryptit.git/encrypt"
)

func main() {
	text := "Something"
	encrypted := encrypt.Nimbus(text)
	fmt.Println("Encrypted: ", encrypted)
	fmt.Println("Decrypted: ", decrypt.Nimbus(encrypted))
}
