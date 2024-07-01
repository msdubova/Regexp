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
		"380%s",
		"(380) %s",
		"380-%s",
		"380 %s",
		"380.%s",
		"+380 %s",
	}

	randomNumberFormats := []string{
		"%02d-%02d-%02d-%03d",
		"%02d %02d %02d %03d",
		"%02d.%02d.%02d.%03d",
		"%09d",
	}

	format := formats[rand.Intn(len(formats))]
	randomNumberFormat := randomNumberFormats[rand.Intn(len(randomNumberFormats))]

	var randomNumber string

	if randomNumberFormat == "%09d" {
		randomNumber = fmt.Sprintf(randomNumberFormat, rand.Intn(1000000000))
	} else {
		part1 := rand.Intn(100)
		part2 := rand.Intn(100)
		part3 := rand.Intn(100)
		part4 := rand.Intn(1000)
		randomNumber = fmt.Sprintf(randomNumberFormat, part1, part2, part3, part4)
	}

	return fmt.Sprintf(format, randomNumber)
}

func createFile(filename string) error {
	file, err := os.Create(filename)

	if err != nil {
		return fmt.Errorf("не вдалося створити файл: %v", err)
	}

	defer file.Close()

	randQuantity := rand.Intn(10)

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

	pattern := regexp.MustCompile(`\+?380\s?(\(?380\)?\s?)?(?:\s|-|\.?)?\d{2,3}[\s.-]?\d{2}[\s.-]?\d{2}[\s.-]?\d{2,3}`)

	for i, line := range lines {

		matches := pattern.FindAllString(line, -1)
		for _, match := range matches {
			fmt.Printf("%d) Знайдено номер : %s \n", i+1, match)
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
