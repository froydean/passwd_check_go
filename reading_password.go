// This program generate simple serial key "KEY$...$" if password in password file is correct
package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

// ReadingPassword() reads whole data from the file with
// its given name and returns that whole data
func ReadingPassword(fileName string) []byte {
	//WARNING! in this code ReadFile has been used
	//only because file isn't meant to be large
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
func passwdCorrect(userPasswd []byte) bool {
	res_compare := strings.Compare(string(userPasswd), "passw0rd")
	if res_compare == 0 {
		return true
	}
	return false
}
func generateSerialKey() string {
	var serialNum [10]byte
	for i := 0; i < 10; i++ {
		serialNum[i] = byte(rand.Intn(128))
	}
	tmp2 := string(serialNum[:])
	tmp := "KEY$" + tmp2 + "$"
	return tmp
}

func printSerialKey(fileName string) {
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.WriteString(generateSerialKey()); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	if passwdCorrect(ReadingPassword("password.txt")) {
		fmt.Println("Correct password. Access allowed.")
		printSerialKey("serial.txt")
	} else {
		fmt.Println("Password is incorrect. Access denied.")
	}
}
