package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func HandleSimple(reader *bufio.Reader, writer *bufio.Writer, isFile bool, number int) error {
	eq, err := GetEquation(number)
	if err != nil {
		return err
	}
	cancel := make(chan struct{})
	errChan := make(chan error)
	go func() {
		defer func() {
			close(errChan)
			close(cancel)
		}()
		if !isFile {
			fmt.Println("Выберете каким методом решить уравнение:")
			fmt.Println("1 - метод простой итерации")
			fmt.Println("2 - метод Ньютона")
			fmt.Println("3 - метод половинного деления")
		}
		ans, err := reader.ReadString('\n')
		if err != nil {
			cancel <- struct{}{}
			errChan <- err
			return
		}
		ans = strings.TrimSpace(ans)
		if ans == "1" {
			err := handleSimpleIter(reader, writer, eq, isFile)
			if err != nil {
				cancel <- struct{}{}
				errChan <- err
				return
			}
		} else if ans == "2" {
			err := handleNewton(reader, writer, eq, isFile)
			if err != nil {
				cancel <- struct{}{}
				errChan <- err
				return
			}
		} else if ans == "3" {
			err := handleBisection(reader, writer, eq, isFile)
			if err != nil {
				cancel <- struct{}{}
				errChan <- err
				return
			}
		} else {
			cancel <- struct{}{}
			errChan <- errors.New("Неверно выбран метод для решения")
			return
		}
		cancel <- struct{}{}
		errChan <- nil
	}()
	DrawGraphic(cancel, eq)
	return <-errChan
}

func handleSimpleIter(reader *bufio.Reader, writer *bufio.Writer, eq *Equation, isFile bool) error {
	for i := 0; i < len(eq.roots); i++ {
		if !isFile {
			fmt.Printf("Введите данные для корня %v:\n", i+1)
		}
		eps, a, b, err := handleInterval(reader, writer, eq, isFile)
		if err != nil {
			return err
		}
		if !eq.checkConvergence(a, b) {
			return errors.New("Функция на сходится на данном промежутке")
		}
		res, iter := SimpleIteration(eq, (a+b)/2, eps)
		_, err = writer.WriteString(fmt.Sprintf("Корень: %v\n", res))
		if err != nil {
			return err
		}
		err = writer.Flush()
		if err != nil {
			return err
		}
		_, err = writer.WriteString(fmt.Sprintf("Количество итераций: %v\n", iter))
		if err != nil {
			return err
		}
		err = writer.Flush()
		if err != nil {
			return err
		}
	}
	return nil
}

func handleBisection(reader *bufio.Reader, writer *bufio.Writer, eq *Equation, isFile bool) error {
	for i := 0; i < len(eq.roots); i++ {
		if !isFile {
			fmt.Printf("Введите данные для корня %v:\n", i+1)
		}
		eps, a, b, err := handleInterval(reader, writer, eq, isFile)
		if err != nil {
			return err
		}
		res, iter, err := Bisection(eq, a, b, eps)
		if err != nil {
			return err
		}
		_, err = writer.WriteString(fmt.Sprintf("Корень: %v\n", res))
		if err != nil {
			return err
		}
		err = writer.Flush()
		if err != nil {
			return err
		}
		_, err = writer.WriteString(fmt.Sprintf("Количество итераций: %v\n", iter))
		if err != nil {
			return err
		}
		err = writer.Flush()
		if err != nil {
			return err
		}
	}
	return nil
}

func handleNewton(reader *bufio.Reader, writer *bufio.Writer, eq *Equation, isFile bool) error {
	for i := 0; i < len(eq.roots); i++ {
		if !isFile {
			fmt.Printf("Введите данные для корня %v:\n", i+1)
		}
		eps, a, b, err := handleInterval(reader, writer, eq, isFile)
		if err != nil {
			return err
		}
		res, iter, err := Bisection(eq, a, b, eps)
		if err != nil {
			return err
		}
		_, err = writer.WriteString(fmt.Sprintf("Корень: %v\n", res))
		if err != nil {
			return err
		}
		err = writer.Flush()
		if err != nil {
			return err
		}
		_, err = writer.WriteString(fmt.Sprintf("Количество итераций: %v\n", iter))
		if err != nil {
			return err
		}
		err = writer.Flush()
		if err != nil {
			return err
		}
	}
	return nil
}

func handleInterval(reader *bufio.Reader, writer *bufio.Writer, eq *Equation, isFile bool) (float64, float64, float64, error) {
	if !isFile {
		fmt.Print("Введите точность: ")
	}
	line, err := reader.ReadString('\n')
	if err != nil {
		return 0, 0, 0, err
	}
	eps, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
	if err != nil {
		return 0, 0, 0, err
	}
	if eps <= 0 {
		return 0, 0, 0, errors.New("Неверно указана точность")
	}
	if !isFile {
		fmt.Print("Введите начало промежутка: ")
	}
	line, err = reader.ReadString('\n')
	if err != nil {
		return 0, 0, 0, err
	}
	a, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
	if !isFile {
		fmt.Print("Введите конец промежутка: ")
	}
	line, err = reader.ReadString('\n')
	if err != nil {
		return 0, 0, 0, err
	}
	b, err := strconv.ParseFloat(strings.TrimSpace(line), 64)

	if !eq.checkRoots(a, b) {
		return 0, 0, 0, errors.New("На промежутке имеется количество корней не равное 1")
	}
	return eps, a, b, nil
}
