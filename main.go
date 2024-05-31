package main

import "fmt"

func main() {
	var op1, op2 int
	var mes, sum string
	fmt.Printf("\nWelcome to the Cypher Tool!\n\nSelect operation (1/2):\n1. Encrypt.\n2. Decrypt.\n")
	fmt.Scan(&op1)

	for pl := "please"; op1 != 1 && op1 != 2; {
		fmt.Printf("\nSelect by entering it's index number %s(1/2):\n1. Encrypt.\n2. Decrypt.\n", pl)
		fmt.Scan(&op1)
		pl += " please"
	}

	fmt.Printf("\nSelect cypher (1/2):\n1. ROT13.\n2. Reverse.\n")
	fmt.Scan(&op2)

	for pl := "please"; op2 != 1 && op2 != 2; {
		fmt.Printf("\nSelect by entering it's index number %s(1/2):\n1. Encrypt.\n2. Decrypt.\n", pl)
		fmt.Scan(&op2)
		pl += " please"
	}

	fmt.Printf("\nEnter the message:\n")
	fmt.Scan(&mes)

	switch op1 {
	case 1:
		switch op2 {
		case 1:
			sum = survivalists.encrypt_rot13(mes)
		case 2:
			sum = survivalists.encrypt_reverse(mes)
		}
	case 2:
		switch op2 {
		case 1:
			sum = survivalists.decrypt_rot13(mes)
		case 2:
			sum = survivalists.decrypt_reverse(mes)
		}
	}

	fmt.Printf("\nDecrypted message using %s:\n")

	fmt.Printf("%s\n", sum)
}
