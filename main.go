package main

import (
	"fmt"
	"math/rand"
	"os"
)

func generateRandomNumber() string {
	formats := []string{
		"380%07d",
		"(380) %07d",
		"380-%07d",
		"380 %07d",
		"380.%07d",
		"+380 %07d",
	}

	format := formats[rand.Intn(len(formats))]

	return fmt.Sprintf(format, rand.Intn(10000000))
}

func createFile(filename string) error {
	file, err := os.Create(filename)

	if err != nil {
		return fmt.Errorf("невдалося створити файл: %v", err)
	}

	defer file.Close()

	randQuantity := rand.Intn(100)

	for i := 0; i < randQuantity; i++ {
		randNumber := generateRandomNumber()
		_, err := file.WriteString(randNumber + "\n")
		if err != nil {
			return fmt.Errorf("помилка запису рядка у файл: %v", err)
		}
	}
	return nil
}

func readFile(filename string) {
	fileContent, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Помилка читання файлу", err)
	}

	fmt.Printf(string(fileContent))
}

func main() {
	const filename = "a.txt"
	err := createFile(filename)
	if err != nil {
		fmt.Printf("помилка", err)
	}
	readFile(filename)
}
