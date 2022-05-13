package main

import "fmt"

func main() {

	counterPersonas := map[string]int{
		"Luis":   1,
		"Maria":  10,
		"Ana":    3,
		"Carlos": 60,
	}
	counterPersonas["Dairo"] = 1
	fmt.Println(counterPersonas)
	for s, _ := range counterPersonas {
		fmt.Println(s)
	}
	delete(counterPersonas, "Ana")
	fmt.Println(counterPersonas)

}
