package decrypt

func Nimbus(s string) string {
	decrypted := ""
	for _, c := range s {
		ascii := int(c)
		character := string(ascii - 3)
		decrypted += character
	}
	return decrypted
}
