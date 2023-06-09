package main

import (
	"fmt"
	_ "github.com/go-redis/redis"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//lon := "London"
	// Получение данных с weatherapi.com
	resp, err := http.Get("http://api.weatherapi.com/v1/future.json?key=34b637d2858a40d9b1542532233105&q=London&dt=2023-07-09")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Чтение ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//err = ioutil.WriteFile("data1.txt", body, 0644)
	//if err != nil {
	//	fmt.Println("Ошибка при сохранении данных в файл:", err)
	//	return
	//}
	// Запуск локального сервера и передача данных
	http.HandleFunc("/weather/London", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", body)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
