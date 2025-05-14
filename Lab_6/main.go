package main

import "fmt"

func main() {
	diffur, err := ReadDiffurInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	euler, err := Euler(diffur)
	if err != nil {
		fmt.Println(err)
		return
	}
	accEuler, err := estimateRungeError(euler, Euler, 1.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	runge, err := RungeKutta4(diffur)
	if err != nil {
		fmt.Println(err)
		return
	}
	accRunge, err := estimateRungeError(runge, RungeKutta4, 15.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	milne, accMilne, err := Milne(diffur)
	if err != nil {
		fmt.Println(err)
		return
	}
	PrintTable(euler, runge, milne)
	fmt.Printf("Точность для метода Эйлера: %v\n", accEuler)
	fmt.Printf("Точность для метода Рунге-Кутта 4-го порядка: %v\n", accRunge)
	fmt.Printf("Точность для метода Милне: %v\n", accMilne)
	err = DrawGraphic(euler, runge, milne, "output.png")
	if err != nil {
		fmt.Println(err)
		return
	}
}
