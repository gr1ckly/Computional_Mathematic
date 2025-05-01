package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	isFile := false
	reader, inFile, err := getReader()
	if err != nil {
		fmt.Println(err)
		return
	}
	if inFile != nil {
		defer inFile.Close()
		isFile = true
	}
	points, err := GetPoints(reader, isFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	table := FiniteDifferences(points)
	fmt.Println("\nТаблица конечных разностей:")
	fmt.Printf("%10s  %10s", "x", "y")
	for i := 1; i < len(table); i++ {
		fmt.Printf("  %12s", fmt.Sprintf("Δ^%d", i))
	}
	fmt.Println()
	for i := 0; i < len(table); i++ {
		fmt.Printf("%10.4f  %10.4f", points[i].X, points[i].Y)
		for j := 1; j < len(table)-i; j++ {
			fmt.Printf("  %12.6f", table[i][j])
		}
		fmt.Println()
	}
	x, err := GetX(reader, isFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	lagrange := LagrangeInterpolation(points, x)
	newton := NewtonInterpolation(points, x)
	gauss := GaussInterpolation(points, x)
	fmt.Printf("Результаты интерполяции для x = %v\n", x)
	fmt.Printf("Интерполяция по Лагранжу: %v\n", lagrange)
	fmt.Printf("Интерполяция по Ньютону с разельными суммами: %v\n", newton)
	fmt.Printf("Интерполяция по Гауссу: %v\n", gauss)
	err = DrawGraphic(points, x, lagrange, newton, gauss, "out.png")
	if err != nil {
		fmt.Println(err)
		return
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
