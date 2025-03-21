package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	integral, err := ParseIntegral(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	acc, err := ParseAccuracy(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	method, err := ParseMethod(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	ans, iter := SolveIntegral(integral, acc, method)
	fmt.Printf("Значение интеграла: %v\n", ans)
	fmt.Printf("Итерация: %v\n", iter)
}
