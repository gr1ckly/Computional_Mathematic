package main

import (
	matrix2 "Lab_1/matrix"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var size, maxIter int
	var accuracy float64
	var isFile bool
	var isRandom bool
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
	data := make([][]float64, size)
	vecB := make([]float64, size)
	if !isFile {
		fmt.Print("Введите любой символ если хотите заполнить матрицу случайными числами: ")
		line, err = reader.ReadString('\n')
		line = line[:len(line)-1]
		if line != "" {
			rand.Seed(time.Now().UnixNano())
			for i := 0; i < size; i++ {
				data[i] = make([]float64, size)
				for j := 0; j < size; j++ {
					data[i][j] = rand.Float64() * 100
				}
			}
			for i := 0; i < size; i++ {
				vecB[i] = rand.Float64()
			}
			isRandom = true
		} else {
			isRandom = false
		}
	}
	if !isRandom {
		if !isFile {
			fmt.Println("Введите матрицу:")
		}
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
		for i := 0; i < size; i++ {
			vecB[i], err = strconv.ParseFloat(splitline[i], 64)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
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
	if isRandom {
		fmt.Println("Сгенерированная матрица: ")
		fmt.Println(matrix.StringMatrix())
	}
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
