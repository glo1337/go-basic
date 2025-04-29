package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"struct/bins"
	"struct/file"
)

func SaveBins(content []byte, fileName string) {
	if filepath.Ext(fileName) == "" {
		fileName = fileName + ".json"
	}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		fmt.Println("ошибка при записи файла:", err)
		return
	}
	fmt.Println("запись успешна:", fileName)
}

func ReadBins() (*bins.BinList, error) {
	binList := bins.NewBinList()
	content, err := file.ReadFromFile("./bin.json")
	if err != nil {
		return nil, errors.New("произошла ошибка во время чтения")
	}

	err = json.Unmarshal(content, binList)
	if err != nil {
		return nil, err
	}

	return binList, nil
}
