package main

import (
	"encoding/json"
	"fmt"
	"struct/bins"
	"struct/file"
	"struct/storage"
)

func main() {
	binList := bins.NewBinList()
	content, err := file.ReadFromFile("./bin.json")
	if err != nil {
		fmt.Println("произошла ошибка во время чтения:", err)
		return
	}

	fmt.Println(json.Unmarshal(content, binList))

	bin1 := bins.NewBin()
	bin2 := bins.NewBin()
	bin3 := bins.NewBin()

	binList.List = append(binList.List, *bin1)
	binList.List = append(binList.List, *bin2)
	binList.List = append(binList.List, *bin3)

	content, err = json.Marshal(binList)
	if err != nil {
		fmt.Println("ошибка при формировании json bin")
	}
	storage.SaveAsJson(content, "bin")
}
