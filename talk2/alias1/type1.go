package main

import (
	"fmt"
)

type farenheit float64

func (f farenheit) String() string {
	return fmt.Sprintf("%.2f F", f)
}

type temp farenheit

func main() {
	f := farenheit(10)
	fmt.Println(f)

	//same type but methods are not passed
	t := temp(20)
	fmt.Println(t)
}
