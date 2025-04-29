package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

func SaveAsJson(content []byte, fileName string) {
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
