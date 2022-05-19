package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("leyendo ficheros linea linea")
	open, err := os.Open("./eje12-ficheros/stand_by_me.txt")
	defer open.Close()
	if err != nil {
		fmt.Println("Hubo un error" + err.Error())
	} else {
		scanner := bufio.NewScanner(open)
		contador := 1
		for scanner.Scan() {
			fmt.Println("Linea:" + string(contador) + scanner.Text())
			contador++
		}

	}

}
