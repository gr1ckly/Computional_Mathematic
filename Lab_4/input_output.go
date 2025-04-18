package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ReadIntervals(reader *bufio.Reader, writer *bufio.Writer, isOutFile bool) (int, error) {
	if !isOutFile {
		_, err := writer.WriteString("Введите количество точек от 8 до 12: ")
		err = writer.Flush()
		if err != nil {
			return -1, err
		}
	}
	data, err := reader.ReadString('\n')
	if err != nil {
		return -1, err
	}
	number, err := strconv.Atoi(strings.TrimSpace(data))
	if err != nil {
		return -1, err
	}
	if number > 12 || number < 8 {
		return -1, fmt.Errorf("Число не входит в указанный интервал")
	}
	return number, nil
}

func ReadPoints(reader *bufio.Reader, writer *bufio.Writer, isOutFile bool, intervalsNumber int) ([]Point, error) {
	if !isOutFile {
		_, err := writer.WriteString("Введите x и y через пробел:\n")
		err = writer.Flush()
		if err != nil {
			return nil, err
		}
	}
	var line string
	var splitline []string
	var err error
	var x, y float64
	answer := make([]Point, intervalsNumber)
	for i := 0; i < intervalsNumber; i++ {
		line, err = reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		splitline = strings.Split(strings.TrimSpace(line), " ")
		if len(splitline) != 2 {
			return nil, fmt.Errorf("Неверное количество аргументов")
		}
		x, err = strconv.ParseFloat(strings.TrimSpace(splitline[0]), 64)
		if err != nil {
			return nil, err
		}
		y, err = strconv.ParseFloat(strings.TrimSpace(splitline[1]), 64)
		if err != nil {
			return nil, err
		}
		answer[i] = Point{x, y}
	}
	return answer, nil
}

func PrintFuncInfo(writer *bufio.Writer, points []Point, approximate *ApproxFunc) error {
	_, err := writer.WriteString(approximate.Name + ":\n")
	err = writer.Flush()
	if err != nil {
		return err
	}
	_, err = writer.WriteString("Коэффициенты:\n")
	err = writer.Flush()
	if err != nil {
		return err
	}
	for _, coef := range approximate.Coefficients {
		_, err = writer.WriteString(fmt.Sprintf("%v\n", coef))
		if err != nil {
			return err
		}
	}
	_, err = writer.WriteString(fmt.Sprintf("СКО: %v\n", CalcStdDeviation(points, approximate.Func, approximate.Coefficients)))
	err = writer.Flush()
	if err != nil {
		return err
	}
	_, err = writer.WriteString("x_i, y_i, fi(x_i), eps_i\n")
	err = writer.Flush()
	if err != nil {
		return err
	}
	for _, point := range points {
		_, err = writer.WriteString(fmt.Sprintf("%v, %v, %v, %v\n", point.X, point.Y, approximate.Func(point.X, approximate.Coefficients), approximate.Func(point.X, approximate.Coefficients)-point.Y))
		err = writer.Flush()
		if err != nil {
			return err
		}
	}
	if approximate.Name == "Линейная" {
		_, err = writer.WriteString(fmt.Sprintf("Коэффициент Пирсона: %v\n", PearsonCorrelation(points)))
		err = writer.Flush()
		if err != nil {
			return err
		}
	}
	det := CalcDetermination(points, approximate.Func, approximate.Coefficients)
	_, err = writer.WriteString(fmt.Sprintf("Коэффицент детерминации: %v\n", det))
	err = writer.Flush()
	if err != nil {
		return err
	}
	if det >= 0.95 {
		_, err = writer.WriteString("Высокая точность аппроксимации\n")
		err = writer.Flush()
		if err != nil {
			return err
		}
	} else if det >= 0.75 {
		_, err = writer.WriteString("Удовлетворительная точность аппроксимации\n")
		err = writer.Flush()
		if err != nil {
			return err
		}
	} else if det >= 0.5 {
		_, err = writer.WriteString("Слабая точность аппроксимации\n")
		err = writer.Flush()
		if err != nil {
			return err
		}
	} else {
		_, err = writer.WriteString("Недостаточная точность аппроксимации\n")
		err = writer.Flush()
		if err != nil {
			return err
		}
	}
	return nil
}
