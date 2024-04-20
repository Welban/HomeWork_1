package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

// Функция для подсчита количества встречаний строки во входных данных
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

// Функция для вывода только тех строк, которые повторились во входных данных
func repeatByLinesD(c bool) ([]string, error) {
	var (
		uniqueLines   []string
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
	}

	// Формирование списка повторяющихся строк
	for line, count := range repeatByLines {
		if count > 1 {
			uniqueLines = append(uniqueLines, line)
		}
	}

	// Проверка ошибок после завершения сканирования
	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return uniqueLines, nil
}

// Функция для вывода только тех строк, которые не повторились во входных данных
func repeatByLinesU(c bool) ([]string, error) {
	var (
		uniqueLines   []string
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
		//inputLines = append(inputLines, currentLine)

		if c {
			count, ok := repeatByLines[currentLine]
			if ok {
				repeatByLines[currentLine] = count + 1
			} else {
				repeatByLines[currentLine] = 1
			}
		}
	}

	// Формирование списка неповторяющихся строк
	for line, count := range repeatByLines {
		if count == 1 {
			uniqueLines = append(uniqueLines, line)
		}
	}

	// Проверка ошибок после завершения сканирования
	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return uniqueLines, nil
}

// Функция для вывода только тех строк, которые не повторились во входных данных
func repeatByLinesF(c bool, numFields int) ([]string, error) {
	var (
		countEntered  int
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
		fields := strings.Fields(currentLine)
		memoryLine := currentLine

		if len(fields) > numFields {

			currentLine = strings.Join(fields[numFields:], " ")

		}

		if c {
			count, ok := repeatByLines[currentLine]
			if ok {
				repeatByLines[currentLine] = count + 1
			} else {
				repeatByLines[currentLine] = 1
			}
		}

		// Формирование списка неповторяющихся строк
		if line != currentLine {
			countEntered++
			if countEntered == 1 {
				line = memoryLine
				inputLines = append(inputLines, line)
				line = currentLine
				countEntered = 0
				continue
			}
			line = currentLine
			inputLines = append(inputLines, line)
		}
		countEntered = 0

	}

	// Проверка ошибок после завершения сканирования
	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return inputLines, nil
}

func uniq(c bool) {

	linesOfC, countByString, err := countByString(c)
	if err != nil {
		fmt.Println("Ошибка при подсчете строк:", err)
		return
	}

	linesOfD, err := repeatByLinesD(c)
	if err != nil {
		fmt.Println("Ошибка при подсчете строк:", err)
		return
	}

	linesOfU, err := repeatByLinesU(c)
	if err != nil {
		fmt.Println("Ошибка при подсчете строк:", err)
		return
	}

	linesOfF, err := repeatByLinesF(c, 1)
	if err != nil {
		fmt.Println("Ошибка при подсчете строк:", err)
		return
	}

	//////////////////////////////////////////////////////

	// Создание нового файла для записи
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Ошибка при создании файла для записи:", err)
		return
	}
	defer outputFile.Close()

	// Создание нового писателя для записи в файл
	writer := bufio.NewWriter(outputFile)

	//////////////////////////////////////////////////////
	_, err = writer.WriteString("parameter -c \n")
	// Запись строк из inputLines в файл
	for _, line := range linesOfC {
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
	_, err = writer.WriteString("parameter -d \n")

	for _, line := range linesOfD {
		if c {
			line = fmt.Sprintf("%s", line)
		}
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
	}
	_, err = writer.WriteString("\n")
	_, err = writer.WriteString("parameter -u \n")

	for _, line := range linesOfU {
		if c {
			line = fmt.Sprintf("%s", line)
		}
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
	}

	_, err = writer.WriteString("\n")
	_, err = writer.WriteString("parameter -f \n")

	for _, line := range linesOfF {
		if c {
			line = fmt.Sprintf("%s", line)
		}
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
	}

	//////////////////////////////////////////////////////

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
