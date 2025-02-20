package main

import (
	matrix2 "Lab_1/matrix"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var size, maxIter int
	var accuracy float64
	var isFile bool
	reader, file, readerErr := getReader()
	if file != nil {
		defer file.Close()
		isFile = true
	} else {
		isFile = false
	}
	if readerErr != nil {
		fmt.Println(readerErr.Error())
		return
	}
	var line string
	var splitline []string
	if !isFile {
		fmt.Print("Введите размер матрицы: ")
	}
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	size, err = strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		fmt.Println(err.Error())
	}
	if !isFile {
		fmt.Println("Введите матрицу:")
	}
	data := make([][]float64, size)
	for i := 0; i < size; i++ {
		data[i] = make([]float64, size)
		line, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		splitline = strings.Split(strings.TrimSpace(line), " ")
		if len(splitline) != size {
			fmt.Printf("Длина %d строки матрицы некорректна\n", i+1)
		}
		for j := 0; j < size; j++ {
			data[i][j], err = strconv.ParseFloat(splitline[j], 64)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}
	if !isFile {
		fmt.Println("Введите вектор b: ")
	}
	line, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	splitline = strings.Split(strings.TrimSpace(line), " ")
	if len(splitline) != size {
		fmt.Println("Некорректное количество элементов для вектора b")
		return
	}
	vecB := make([]float64, size)
	for i := 0; i < size; i++ {
		vecB[i], err = strconv.ParseFloat(splitline[i], 64)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	if !isFile {
		fmt.Print("Введите максимальное количество итераций: ")
	}
	line, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	maxIter, err = strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if !isFile {
		fmt.Print("Введите точность: ")
	}
	line, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	accuracy, err = strconv.ParseFloat(strings.TrimSpace(line), 64)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	matrix := matrix2.BuildMatrix(size, maxIter, data, vecB, accuracy)
	if matrix.MakeDiagonallyDominant() {
		fmt.Println("Диагональное преобразование найдено:")
		fmt.Println(matrix.StringMatrix())
	} else {
		fmt.Println("Диагональное преобразование невозможно")
	}
	fmt.Printf("Норма матрицы: %f\n", matrix.Norm())
	x, iter, errors, err := matrix.GaussSeidel()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Отклонения: ")
		for i := 0; i < size; i++ {
			fmt.Println(errors[i])
		}
		return
	}
	fmt.Printf("Решено за %d итераций\n", iter)
	fmt.Println("Отклонения: ")
	for i := 0; i < size; i++ {
		fmt.Println(errors[i])
	}
	fmt.Println("Решение: ")
	for i := 0; i < size; i++ {
		fmt.Println(x[i])
	}
}

func getReader() (*bufio.Reader, *os.File, error) {
	var line string
	stdin := bufio.NewReader(os.Stdin)
	fmt.Print("Введите любой символ, если хотите ввести данные из файла: ")
	line, err := stdin.ReadString('\n')
	line = line[:len(line)-1]
	if err != nil {
		return nil, nil, err
	}
	if line == "" {
		return bufio.NewReader(stdin), nil, nil
	}
	fmt.Print("Введите название файла: ")
	line, err = stdin.ReadString('\n')
	line = line[:len(line)-1]
	if err != nil {
		return nil, nil, err
	}
	file, err := os.Open(line)
	if err != nil {
		return nil, nil, err
	}
	return bufio.NewReader(file), file, nil
}
