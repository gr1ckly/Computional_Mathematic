package main

import (
	"fmt"
	"strings"
)

func ReadDiffurInput() (Diffur, error) {
	fmt.Println("Выберите ОДУ:")
	for key, _ := range fnsDescription {
		fmt.Printf("%v - %v\n", key, fnsDescription[key])
	}
	var fnChoice string
	_, err := fmt.Scan(&fnChoice)
	if err != nil {
		return Diffur{}, err
	}
	fn, ok := fns[strings.TrimSpace(fnChoice)]
	if !ok {
		return Diffur{}, fmt.Errorf("Некорректно указан номер ОДУ")
	}
	fmt.Print("Введите y0: ")
	var y0 float64
	_, err = fmt.Scan(&y0)
	if err != nil {
		return Diffur{}, err
	}
	fmt.Print("Введите интервал дифференцирования через пробел: ")
	var x0, xn float64
	_, err = fmt.Scan(&x0)
	if err != nil {
		return Diffur{}, err
	}
	_, err = fmt.Scan(&xn)
	if err != nil {
		return Diffur{}, err
	}
	if x0 >= xn {
		return Diffur{}, fmt.Errorf("Интервал указан некорректно")
	}
	exactY, err := GetExactSolutionFunction(strings.TrimSpace(fnChoice), x0, y0)
	if err != nil {
		return Diffur{}, err
	}
	fmt.Print("Введите шаг: ")
	var step float64
	_, err = fmt.Scan(&step)
	if err != nil {
		return Diffur{}, err
	}
	if step > xn-x0 || step <= 0 {
		return Diffur{}, fmt.Errorf("Шаг указан некорректно")
	}
	fmt.Print("Введите точность: ")
	var accuracy float64
	_, err = fmt.Scan(&accuracy)
	if err != nil {
		return Diffur{}, err
	}
	if accuracy <= 0 {
		return Diffur{}, fmt.Errorf("Точность указана некорректно")
	}
	return BuildDiffur(fn, exactY, getXArray(x0, xn, step), y0, accuracy), nil
}

func getXArray(x0, xn, step float64) []float64 {
	xArr := make([]float64, int((xn-x0)/step)+1)
	pointer := 0
	for i := x0; i < xn; i += step {
		xArr[pointer] = i
		pointer++
	}
	xArr[len(xArr)-1] = xn
	return xArr
}

func PrintTable(euler, rungeKutta, milne Diffur) {
	fmt.Printf("%-10s %-15s %-15s %-15s\n", "x", "Эйлер", "Рунге-Кутта", "Милн")
	for i := 0; i < len(euler.x); i++ {
		x := euler.x[i]
		yEuler := euler.y[i]
		yRK := rungeKutta.y[i]
		yMilne := milne.y[i]
		fmt.Printf("%-10.4f %-15.8f %-15.8f %-15.8f\n", x, yEuler, yRK, yMilne)
	}
}
