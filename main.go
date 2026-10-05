package main

import "fmt"

func main() {
	urlMap := map[string]string{
		"PurpleSchool": "https://puprleschool.ru",
		"Yandex":       "https://yandex.ru",
		"Google":       "https://www.google.com",
	}
	urlMap["PurpleSchool"] = "https://www.purple-school.ru"
	delete(urlMap, "PurpleSchool")
	value, ok := urlMap["Yandex"]

	fmt.Println(value, ok)

	_, ok = urlMap["NotExist"]

	if ok == false {
		fmt.Println("Kye 'NotExist' doesnt exist")
	}

	fmt.Printf("Url map len is %d\n", len(urlMap))

	for key, value := range urlMap {
		fmt.Println(key, value)
	}
}
