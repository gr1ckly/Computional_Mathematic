package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader, inFile, err := getReader()
	if err != nil {
		fmt.Println(err)
		return
	}
	if inFile != nil {
		defer inFile.Close()
	}
	writer, outFile, err := getWriter()
	if err != nil {
		fmt.Println(err)
		return
	}
	if outFile != nil {
		defer outFile.Close()
	}
	defer writer.Flush()
	intNumber, err := ReadIntervals(reader, writer, outFile != nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	points, err := ReadPoints(reader, writer, outFile != nil, intNumber)
	if err != nil {
		fmt.Println(err.Error() + "\n")
		return
	}
	approximations := []*ApproxFunc{}
	exp, err := NewExp(points)
	if err != nil {
		_, err = writer.WriteString(err.Error() + "\n")
		if err != nil {
			fmt.Println(err)
		}
	} else {
		approximations = append(approximations, exp)
	}
	log, err := NewLog(points)
	if err != nil {
		_, err = writer.WriteString(err.Error() + "\n")
		if err != nil {
			fmt.Println(err)
		}
	} else {
		approximations = append(approximations, log)
	}
	lin, err := NewLinear(points)
	if err != nil {
		_, err = writer.WriteString(err.Error() + "\n")
		if err != nil {
			fmt.Println(err)
		}
	} else {
		approximations = append(approximations, lin)
	}
	poly2, err := NewPoly2(points)
	if err != nil {
		_, err = writer.WriteString(err.Error() + "\n")
		if err != nil {
			fmt.Println(err)
		}
	} else {
		approximations = append(approximations, poly2)
	}
	power, err := NewPower(points)
	if err != nil {
		_, err = writer.WriteString(err.Error() + "\n")
		if err != nil {
			fmt.Println(err)
		}
	} else {
		approximations = append(approximations, power)
	}
	poly3, err := NewPoly3(points)
	if err != nil {
		_, err = writer.WriteString(err.Error() + "\n")
		if err != nil {
			fmt.Println(err)
		}
	} else {
		approximations = append(approximations, poly3)
	}
	err = DrawApproximations(points, approximations, "out.png", writer)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, appr := range approximations {
		err = PrintFuncInfo(writer, points, appr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	least, appr := GetBestApproximation(points, approximations)
	_, err = writer.WriteString(fmt.Sprintf("Наименьшей S является: %v\n", least))
	if err != nil {
		fmt.Println(err)
		return
	}
	writer.Flush()
	_, err = writer.WriteString("Ею обладает: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	writer.Flush()
	for _, approx := range appr {
		_, err = writer.WriteString(approx.Name + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
		writer.Flush()
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

func getWriter() (*bufio.Writer, *os.File, error) {
	var line string
	stdin := bufio.NewReader(os.Stdin)
	fmt.Print("Введите любой символ, если хотите перенаправить вывод в файл: ")
	line, err := stdin.ReadString('\n')
	line = line[:len(line)-1]
	if err != nil {
		return nil, nil, err
	}
	if line == "" {
		return bufio.NewWriter(os.Stdout), nil, nil
	}
	fmt.Print("Введите название файла: ")
	line, err = stdin.ReadString('\n')
	line = line[:len(line)-1]
	if err != nil {
		return nil, nil, err
	}
	file, err := os.Create(line)
	if err != nil {
		return nil, nil, err
	}
	return bufio.NewWriter(file), file, nil
}
