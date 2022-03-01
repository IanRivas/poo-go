package main

import (
	"fmt"

	"github.com/IanRivas/poo-go/course"
)

func main(){
	Go := &course.Course{
		"Go desde Cero",
		 12.34,
		 false,
		 []uint{12,56,89},
		 map[uint]string{
			1: "introduccion",
			2: "Estructuras",
			3: "Maps",
		},
	} 
	Go.PrintClasses()
	Go.ChangePrice(67.12)
	fmt.Println(Go.Price)
}