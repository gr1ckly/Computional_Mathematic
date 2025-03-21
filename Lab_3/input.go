package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ParseIntegral(reader *bufio.Reader) (*Integral, error) {
	fmt.Println("Выберите интеграл для вычисления: ")
	descMap := GetIntegralDescription()
	for key, _ := range descMap {
		fmt.Printf("%v - %v\n", key, descMap[key])
	}
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	intNum, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		return nil, err
	}
	integral, err := GetIntegral(intNum)
	if err != nil {
		return nil, err
	}
	fmt.Println("Укажите начало промежутка: ")
	line, err = reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	a, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
	if err != nil {
		return nil, err
	}
	fmt.Println("Укажите конец промежутка: ")
	line, err = reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	b, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
	if err != nil {
		return nil, err
	}
	if a >= b {
		return nil, errors.New("Некорректно указан промежуток")
	}
	integral.start = a
	integral.end = b
	return integral, nil
}

func ParseAccuracy(reader *bufio.Reader) (float64, error) {
	fmt.Println("Укажите точность: ")
	line, err := reader.ReadString('\n')
	if err != nil {
		return -1.0, err
	}
	acc, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
	if err != nil {
		return -1.0, err
	}
	if acc <= 0.0 {
		return -1.0, errors.New("Некорректно указана точность")
	}
	return acc, nil
}

func ParseMethod(reader *bufio.Reader) (func(*Integral, int) float64, error) {
	fmt.Println("Выберите метод для вычисления интеграла: ")
	methodMap := GetMethodDescription()
	for key, _ := range methodMap {
		fmt.Printf("%v - %v\n", key, methodMap[key])
	}
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	methodNum, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		return nil, err
	}
	method, err := GetMethod(methodNum)
	return method, err
}
