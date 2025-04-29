package file

import (
	"errors"
	"os"
	"path"
)

func ReadFromFile(fileName string) ([]byte, error) {
	info, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("файл не существует")
		}
		return nil, errors.New("ошибка получения информации о файле")
	}

	if info.IsDir() {
		return nil, errors.New("ожидался файл, но это папка")
	}

	if !IsJson(fileName) {
		return nil, errors.New("ожидается json")
	}

	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, errors.New("ошибка чтения файла")
	}
	return content, nil
}

func IsJson(fileName string) bool {
	fileExt := path.Ext(fileName)
	return fileExt == ".json"
}
