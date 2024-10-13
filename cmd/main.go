package main

import (
	"fmt"
	"log"
	"net/http"
	"smpk/service"
)

func main() {

	//стили html файла

	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	//маршруты
	http.HandleFunc("/", service.MainPage)
	http.HandleFunc("/jsonUser", service.GetUsers)

	//запуск сервера
	fmt.Println("Stated...")
	log.Fatal(http.ListenAndServe(":8181", nil))
}
