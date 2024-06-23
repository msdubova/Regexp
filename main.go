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
		randNumber := rand.Int63n(9000000000) + 1000000000
		_, err := file.WriteString(fmt.Sprintf("%d\n", randNumber))
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
