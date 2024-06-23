package main

import (
	"fmt"
	"math/rand"
	"os"
)

func createFile(filename string) error {
	file, err := os.Create(filename)

	if err != nil {
		return fmt.Errorf("невдалося створити файл: %v", err)
	}
	defer file.Close()

	randQuantity := rand.Intn(100)

	for i := 0; i < randQuantity; i++ {
		randNumber := fmt.Sprintf("380%07d", rand.Intn(1000000000))
		_, err := file.WriteString(randNumber + "\n")
		if err != nil {
			return fmt.Errorf("помилка запису рядка у файл: %v", err)
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
