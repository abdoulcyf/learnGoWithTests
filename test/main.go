package main

import (
	"fmt"
	"github.com/ediallocyf/learnGoWithTest/structs_methods_interfaces"
)

func main() {
	area1 := structs_methods_interfaces.Circle{5.5}
	area2 := structs_methods_interfaces.Rectangle{5.5, 10.0}

	shapes := []structs_methods_interfaces.Shape{area1, area2}

	for _, shape := range shapes {
		fmt.Println(shape)
	}

}
