package main

import (
	"binstruct/bins"
	"binstruct/storage"
	"encoding/json"
	"fmt"
)

func main() {
	binList, err := storage.ReadBins()
	if err != nil {
		fmt.Println("при чтении bins произошла ошибка:", err)
		return
	}

	bin1 := bins.NewBin()
	bin2 := bins.NewBin()
	bin3 := bins.NewBin()

	binList.List = append(binList.List, *bin1)
	binList.List = append(binList.List, *bin2)
	binList.List = append(binList.List, *bin3)

	content, err := json.Marshal(binList)
	if err != nil {
		fmt.Println("ошибка при формировании json bin")
	}
	storage.SaveBins(content, "bin")
}
