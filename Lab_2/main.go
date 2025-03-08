package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader, inFile, err := getReader()
	if err != nil {
		fmt.Println(err)
		return
	}
	writer, outFile, err := getWriter()
	if err != nil {
		fmt.Println(err)
		return
	}
	isFile := false
	if inFile != nil {
		isFile = true
		defer inFile.Close()
	}
	if outFile != nil {
		defer outFile.Close()
	}
	stdin := bufio.NewReader(os.Stdin)
	fmt.Print("Введите 1, если хотите решить уравнение, и 2, если систему линейных уравнений: ")
	ans, err := stdin.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	ans = strings.TrimSpace(ans)
	if ans == "1" {
		fmt.Println("Выберете уравнение:")
		eqMap := GetEqMap()
		for key, _ := range eqMap {
			fmt.Printf("%v - %v\n", key, eqMap[key])
		}
		line, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		number, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			fmt.Println(err)
			return
		}
		err = HandleSimple(reader, writer, isFile, number)
		if err != nil {
			fmt.Println(err)
		}
	} else if ans == "2" {
		fmt.Println("Выберете уравнение:")
		sysEqMap := GetSysEqMap()
		for key, _ := range sysEqMap {
			fmt.Printf("%v - %v\n", key, sysEqMap[key])
		}
		line, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		number, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			fmt.Println(err)
			return
		}
		err = HandleSystem(reader, writer, isFile, number)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Неверный ввод")
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
