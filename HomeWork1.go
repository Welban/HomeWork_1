package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	uniq(true)
}

func getFile() (*bufio.Scanner, *os.File, error) {

	// Открытие файла для чтения
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return nil, nil, err
	}

	// Создание нового сканера для чтения из файла
	scanner := bufio.NewScanner(file)

	return scanner, file, nil
}

// Функция для подсчета строк в файле
func countByString(c bool) ([]string, map[string]int, error) {
	var (
		line          string
		inputLines    []string
		countByString = make(map[string]int)
	)

	scanner, file, err := getFile()
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	// Чтение файла по одной строке
	for scanner.Scan() {
		currentLine := scanner.Text()
		if c {
			count, ok := countByString[currentLine]
			if ok {
				countByString[currentLine] = count + 1
			} else {
				countByString[currentLine] = 1
			}
		}

		if line != currentLine {
			line = currentLine
			inputLines = append(inputLines, line)
			continue
		}
	}

	// Проверка ошибок после завершения сканирования
	err = scanner.Err()
	if err != nil {
		return nil, nil, err
	}

	return inputLines, countByString, nil
}

// Функция для подсчета строк в файле
func repeatByLines(c bool) ([]string, error) {
	var (
		line          string
		inputLines    []string
		repeatByLines = make(map[string]int)
	)

	scanner, file, err := getFile()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Чтение файла по одной строке
	for scanner.Scan() {
		currentLine := scanner.Text()
		if c {
			count, ok := repeatByLines[currentLine]
			if ok {
				repeatByLines[currentLine] = count + 1
			} else {
				repeatByLines[currentLine] = 1
			}
		}

		if line != currentLine && repeatByLines[currentLine] > 1 {
			line = currentLine
			inputLines = append(inputLines, line)
			continue
		}
	}

	// Проверка ошибок после завершения сканирования
	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return inputLines, nil
}

func uniq(c bool) {

	inputLinesTwo, err := repeatByLines(c)
	if err != nil {
		fmt.Println("Ошибка при подсчете строк:", err)
		return
	}

	inputLines, countByString, err := countByString(c)
	if err != nil {
		fmt.Println("Ошибка при подсчете строк:", err)
		return
	}

	// Создание нового файла для записи
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Ошибка при создании файла для записи:", err)
		return
	}
	defer outputFile.Close()

	// Создание нового писателя для записи в файл
	writer := bufio.NewWriter(outputFile)

	// Запись строк из box в файл
	for _, line := range inputLines {
		if c {
			line = fmt.Sprintf("%d %s", countByString[line], line)
		}
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
	}
	_, err = writer.WriteString("\n")

	for _, line := range inputLinesTwo {
		if c {
			line = fmt.Sprintf("%s ", line)
		}
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
	}
	_, err = writer.WriteString("\n")

	// Сброс буфера, чтобы убедиться, что все данные записаны в файл
	err = writer.Flush()
	if err != nil {
		fmt.Println("Ошибка при сбросе буфера:", err)
		return
	}

}

/*
читаем файл
выписываем первую строку
    сравниваем со следующей
    пока а = б, пропуск строки
    в момент, когда а != б, выводим новую строку
*/

/*
	countByString := make(map[string]int)

	{
	    "I love music.": 3,
	    "": 1,
	}

	Если строки нет в мапе:
	    Добавить строку в мапу, count = 1
	Если строка есть в мапе
	    count = count + 1

*/
