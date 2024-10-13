package main

import (
	"fmt"
	"log"
	"net/http"
	"smpk/service"
)

func main() {

	//маршруты
	http.HandleFunc("/", service.MainPage)
	http.HandleFunc("/jsonUser", service.GetUsers)
	http.HandleFunc("/jsonUser/", service.GetUserById)

	//запуск сервера
	fmt.Println("Stated...")
	log.Fatal(http.ListenAndServe(":8181", nil))
}
