package main

import (
	"errors"
	"fmt"
	"io"
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

func readFileBuff(filename string) {

	buff := make([]byte, 5)

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Помилка відкриття файлу", err)
	}

	for {
		n, err := file.Read(buff)
		if err != nil {

			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Println("помилка відкриття буферу та читання файла")
		}

		fmt.Println(n, "\"", string(buff), "\"")
	}
}

func main() {
	const filename = "a.txt"
	err := createFile(filename)
	if err != nil {
		fmt.Printf("помилка", err)
	}

	readFileBuff(filename)

}
