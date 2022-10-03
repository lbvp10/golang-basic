package main

import "net/http"

func main() {

	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "./eje15-servidor_web/index.html")
}
