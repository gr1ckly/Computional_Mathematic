package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func HandleSystem(reader *bufio.Reader, writer *bufio.Writer, isFile bool, number int) error {
	sysEq, err := GetSystemEquation(number)
	if err != nil {
		return err
	}
	cancel := make(chan struct{})
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(errChan)
			close(cancel)
		}()
		if !isFile {
			fmt.Print("Введите точность: ")
		}
		line, err := reader.ReadString('\n')
		if err != nil {
			cancel <- struct{}{}
			errChan <- err
			return
		}
		eps, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
		if eps <= 0 {
			cancel <- struct{}{}
			errChan <- errors.New("Неверно указана точность")
			return
		}
		if !isFile {
			fmt.Print("Введите максимальное количество итераций: ")
		}
		line, err = reader.ReadString('\n')
		if err != nil {
			cancel <- struct{}{}
			errChan <- err
			return
		}
		maxIter, err := strconv.Atoi(strings.TrimSpace(line))
		if maxIter <= 0 {
			cancel <- struct{}{}
			errChan <- errors.New("Неверно указано максимальное количество итераций")
			return
		}
		if !isFile {
			fmt.Println("Введите начальные приближения: ")
			fmt.Print("x0: ")
		}
		line, err = reader.ReadString('\n')
		if err != nil {
			cancel <- struct{}{}
			errChan <- err
			return
		}
		x0, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
		if !isFile {
			fmt.Print("y0: ")
		}
		line, err = reader.ReadString('\n')
		if err != nil {
			cancel <- struct{}{}
			errChan <- err
			return
		}
		y0, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
		x, y, iter, maxError, err := NewtonMethod(sysEq, x0, y0, eps, maxIter)
		if err != nil {
			cancel <- struct{}{}
			errChan <- err
			return
		}
		_, err = writer.WriteString(fmt.Sprintf("X: %v\n", x))
		if err != nil {
			cancel <- struct{}{}
			errChan <- err
			return
		}
		err = writer.Flush()
		if err != nil {
			cancel <- struct{}{}
			errChan <- err
			return
		}
		_, err = writer.WriteString(fmt.Sprintf("Y: %v\n", y))
		if err != nil {
			cancel <- struct{}{}
			errChan <- err
			return
		}
		err = writer.Flush()
		if err != nil {
			cancel <- struct{}{}
			errChan <- err
			return
		}
		_, err = writer.WriteString(fmt.Sprintf("Итерация: %v\n", iter))
		if err != nil {
			cancel <- struct{}{}
			errChan <- err
			return
		}
		err = writer.Flush()
		if err != nil {
			cancel <- struct{}{}
			errChan <- err
			return
		}
		_, err = writer.WriteString(fmt.Sprintf("Максимальная ошибка: %v\n", maxError))
		if err != nil {
			cancel <- struct{}{}
			errChan <- err
			return
		}
		err = writer.Flush()
		if err != nil {
			cancel <- struct{}{}
			errChan <- err
			return
		}
		cancel <- struct{}{}
		errChan <- nil
	}()
	DrawSystemGraphic(cancel, sysEq)
	return <-errChan
}
