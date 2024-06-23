package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
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

func readFile(filename string) ([]string, error) {
	fileContent, err := os.ReadFile(filename)

	if err != nil {
		return nil, fmt.Errorf("Помилка читання файлу", err)
	}
	lines := strings.Split(string(fileContent), "\n")
	return lines, nil
}

func findPhoneNumbers(filename string) error {
	lines, err := readFile(filename)

	if err != nil {
		return fmt.Errorf("помилка читання файлу: %v", err)
	}

	pattern := regexp.MustCompile(`380\d{7}`)

	for _, line := range lines {
		line = strings.ReplaceAll(line, " ", "")
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")
		line = strings.ReplaceAll(line, "-", "")
		line = strings.ReplaceAll(line, "+", "")
		matches := pattern.FindAllString(line, -1)
		for _, match := range matches {
			fmt.Println("Знайдено номер : ", match)
		}
	}

	return nil
}

func main() {
	const filename = "a.txt"
	err := createFile(filename)
	if err != nil {
		fmt.Printf("помилка створення файлу  з номерами", err)
		return
	}

	err = findPhoneNumbers(filename)
	if err != nil {
		fmt.Printf("помилка пошуку номерів : %v\n", err)
		return
	}
}

// приклади з лекціі розбір
// func someExample() {

// 	matched, err := regexp.MatchString(`\w+`, ".")
// 	if err != nil {
// 		fmt.Println("ПОмилка", err)
// 	}

// 	fmt.Println("тут ", matched)

// 	pattern, err := regexp.Compile(`C\w+`)
// 	if err != nil {
// 		fmt.Println("ПОмилка", err)
// 	}

// 	matches := pattern.FindAllString("Cat Cuc i meows meows", -1)

// 	for i, v := range matches {
// 		fmt.Println(i, v)
// 	}
// }
// func readFileBuff(filename string) {

// 	buff := make([]byte, 5)

// 	file, err := os.Open(filename)

// 	if err != nil {
// 		fmt.Println("Помилка відкриття файлу", err)
// 	}

// 	for {
// 		n, err := file.Read(buff)
// 		if err != nil {

// 			if errors.Is(err, io.EOF) {
// 				break
// 			}
// 			fmt.Println("помилка відкриття буферу та читання файла")
// 		}

// 		fmt.Println(n, "\"", string(buff), "\"")
// 	}
// }
