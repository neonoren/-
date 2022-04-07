//Урок посвящен повтору информации по Git
//код взят рандомный

package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["Андрей"] = 38
	ages["Ирина"] = 35
	ages["Максим"] = 7
	ages["Василиса"] = 3

	for key, value := range ages {
		fmt.Printf("%s - %d\n", key, value)

	}
	fmt.Printf("Максиму сейчас %d лет", ages["Максим"])
}
