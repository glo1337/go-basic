package main

import (
	"binstruct/bins"
	"binstruct/storage"
	"encoding/json"
	"fmt"
)

func main() {
	fileStorage := storage.NewFileStorage("./bin.json")

	binList, err := fileStorage.Read()
	if err == nil {
		fmt.Println(binList)
	} else {
		fmt.Println("ошибка при чтении:", err)
	}

	bin1 := bins.NewBin()
	bin2 := bins.NewBin()
	bin3 := bins.NewBin()

	binList.List = append(binList.List, *bin1, *bin2, *bin3)

	content, err := json.Marshal(binList)
	if err != nil {
		fmt.Println("ошибка при формировании json bin:", err)
		return
	}

	err = fileStorage.Write(content)
	if err != nil {
		fmt.Println("ошибка при записи файла:", err)
	}
}
