package main

import "fmt"

func printMenu() {
	fmt.Println(`Сервис работы с закладками:
	1. Показать закладку
	2. Добавить закладку
	3. Удалить закладку
	4. Выход
	`)
}
func showBookmarks(bookmarks map[string]string) {
	if len(bookmarks) == 0 {
		fmt.Println("Список закладок пуст")
	} else {
		for key, value := range bookmarks {
			fmt.Println(key, value)
		}
	}
}

func addBookmark(bookmarks map[string]string) {
	var name string
	var url string
	fmt.Println("Введите название ресурса и урл через пробел")
	_, err := fmt.Scanln(&name, &url)
	if err != nil {
		fmt.Println(err)
		return
	}

	bookmarks[name] = url
}

func deleteBookmark(bookmarks map[string]string) {
	var name string
	fmt.Println("Введите название ресурса для удаления")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, ok := bookmarks[name]
	if !ok {
		fmt.Printf("Закладка с именем %s не существует", name)
	} else {
		delete(bookmarks, name)
	}
}

func main() {
	bookmarks := map[string]string{}
	var choiceNum int
	for {
		printMenu()
		_, err := fmt.Scanln(&choiceNum)

		if err != nil {
			fmt.Println("Выходим из программы")
			break
		}

		switch choiceNum {
		case 1:
			showBookmarks(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 3:
			deleteBookmark(bookmarks)
		case 4:
			fmt.Println("Выход")
			break
		default:
			fmt.Printf("Значение %d не поддерживается\n", choiceNum)
		}
	}
}
