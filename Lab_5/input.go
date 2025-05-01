package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func GetPoints(reader *bufio.Reader, isFile bool) ([]Point, error) {
	if !isFile {
		fmt.Print("Введите 1, если хотите задать функцию таблицей и 2, если хотите выбрать из предложенных: ")
	}
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("Ошибка при чтении данных: ", err)
	}
	line = strings.TrimSpace(line)
	switch line {
	case "1":
		return ReadTable(reader, isFile)
	case "2":
		return ReadReadyFunction(reader, isFile)
	default:
		return nil, fmt.Errorf("Некорректный ввод.")
	}
}

func GetX(reader *bufio.Reader, isFile bool) (float64, error) {
	if !isFile {
		fmt.Print("Введите значение X для интерполяции: ")
	}
	line, err := reader.ReadString('\n')
	if err != nil {
		return -1.0, fmt.Errorf("Ошибка при чтении данных: ", err)
	}
	ans, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
	if err != nil {
		return -1.0, fmt.Errorf("Ошибка при чтении кординат точки: ", err)
	}
	return ans, nil
}

func ReadTable(reader *bufio.Reader, isFile bool) ([]Point, error) {
	if !isFile {
		fmt.Print("Введите количество точек: ")
	}
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("Ошибка при чтении данных: ", err)
	}
	number, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		return nil, fmt.Errorf("Ошибка при вводе количества точек: ", err)
	}
	if !isFile {
		fmt.Println("Введите значения x и y через пробел в каждой отдельной строке: ")
	}
	ans := make([]Point, number)
	for i := 0; i < number; i++ {
		line, err = reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("Ошибка при чтении данных: ", err)
		}
		splitline := strings.Split(line, " ")
		if len(splitline) != 2 {
			return nil, fmt.Errorf("Неверное количество аргументов при вводе точки")
		}
		ans[i].X, err = strconv.ParseFloat(strings.TrimSpace(splitline[0]), 64)
		if err != nil {
			return nil, fmt.Errorf("Ошибка при чтении кординат точки: ", err)
		}
		ans[i].Y, err = strconv.ParseFloat(strings.TrimSpace(splitline[1]), 64)
		if err != nil {
			return nil, fmt.Errorf("Ошибка при чтении кординат точки: ", err)
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i].X < ans[j].X
	})
	return ans, nil
}

func ReadReadyFunction(reader *bufio.Reader, isFile bool) ([]Point, error) {
	if !isFile {
		fmt.Println("Выберите функцию: ")
		descMap := GetFuncsDescription()
		for key, _ := range descMap {
			fmt.Printf("%v - %v\n", key, descMap[key])
		}
	}
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("Ошибка при выборе функции: ", err)
	}
	line = strings.TrimSpace(line)
	fn, ok := GetFuncs()[line]
	if !ok {
		return nil, fmt.Errorf("Некорректно указана функция")
	}
	if !isFile {
		fmt.Println("Укажите начало и конец интервала через пробел: ")
	}
	line, err = reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("Ошибка при чтении данных: ", err)
	}
	splitline := strings.Split(line, " ")
	if len(splitline) != 2 {
		return nil, fmt.Errorf("Неверное количество аргументов при вводе интервала")
	}
	start, err := strconv.ParseFloat(strings.TrimSpace(splitline[0]), 64)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при чтении начала интервала: ", err)
	}
	end, err := strconv.ParseFloat(strings.TrimSpace(splitline[1]), 64)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при чтении конца интервала: ", err)
	}
	if end <= start {
		return nil, fmt.Errorf("Интервал введен некорректно")
	}
	if !isFile {
		fmt.Print("Введите количество точек: ")
	}
	line, err = reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("Ошибка при чтении данных: ", err)
	}
	number, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		return nil, fmt.Errorf("Ошибка при чтении количества точек: ", err)
	}
	if number < 2 {
		return nil, fmt.Errorf("Некорректное значение для количества точек (должно быть минимум 2).")
	}
	ans := make([]Point, number)
	for i := 0; i < number-1; i++ {
		ans[i].X = float64(start) + float64(end-start)/float64(number-1)*float64(i)
		ans[i].Y = fn(ans[i].X)
	}
	ans[number-1].X = end
	ans[number-1].Y = fn(ans[number-1].X)
	sort.Slice(ans, func(i, j int) bool {
		return ans[i].X < ans[j].X
	})
	return ans, nil
}
