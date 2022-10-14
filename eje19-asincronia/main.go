package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	var i int = 1
	inicio := time.Now()
	canal := make(chan string)
	for ; i < 30; i++ { //lanzar gorutine
		go getPokemonById(i, canal)
	}

	i = 1
	for ; i < 30; i++ { //Leer los mensajes cada que respondan
		fmt.Println(<-canal)
	}

	fmt.Printf("Duracion: %v", time.Since(inicio))
}

func getPokemonById(id int, canal chan string) {
	response, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon-form/%d", id))
	if responseData, err2 := ioutil.ReadAll(response.Body); err != nil || err2 != nil {
		canal <- ""
	} else {
		canal <- string(responseData)
	}
}
