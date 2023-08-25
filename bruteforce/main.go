package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

const (
	savedHash = "lHZK77WacKCo1yAylcMrGi1ZActe6w5lf6OCfwirHdc="
)

func main() {
	file, err := os.Open("rockyou.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	readerStrings := bufio.NewScanner(file)

	for readerStrings.Scan() {

		password := readerStrings.Text()

		gotHash, err := hashString(password)
		if err != nil {
			log.Println(err)
			continue
		}

		if gotHash != savedHash {
			continue
		}
		fmt.Println("Пароль найден:", password)
		break
	}

	if err := readerStrings.Err(); err != nil {
		log.Fatal(err)
	}

}

func hashString(str string) (string, error) {

	inputBytes := []byte(str)

	hashSum := sha256.Sum256(inputBytes)
	return base64.StdEncoding.EncodeToString(hashSum[:]), nil
}
