package main

import (
	"fmt"
	"math/rand"
	"os"
)

func createFile(filename string) error {
	file, err := os.Create(filename)

	if err != nil {
		return fmt.Errorf("недавла спрба створити файл", err)
	}

	for i := 0; i < rand.Intn(100); i++ {
		_, err := file.WriteString("abc\n")
		if err != nil {
			return fmt.Errorf("помика запису рядка у файл", err)
		}
	}
	return nil

}
func main() {

	err := createFile("a.txt")
	if err != nil {
		fmt.Printf("помилка")
	}

}
